package main

import (
	"log"
	"fmt"
	"flag"
	"time"
	"errors"
	"math/rand"
	"aqwari.net/exp/gl"
	"aqwari.net/exp/display"
)

const (
	Alive = 1
	Dead = 0
)

var vertexGLSL = []byte(
`#version 150 core

in vec2 position;
uniform vec2 offset;

void main() {
	gl_Position = vec4(position + offset, 0, 1);
}
`)

var fragmentGLSL = []byte(
`#version 150

out vec4 color;

void main(void) {
	color = vec4(1,1,1,1);
}
`)

var (
	columns = 256
	rows = 256
	width = flag.Int("w", 600, "Screen width")
	height = flag.Int("h", 600, "Screen height")
	config = display.Config{
		"OpenGL Version": "3.2",
		"Title": "Life",
		"Resize": "False",
		"Swap Interval": "1",
	}
)

func initGL(win *display.Buffer, nrow, ncol int) (gl.Program, error) {
	if err := gl.Init(config["OpenGL Version"]); err != nil {
		return 0, err
	}
	gl.ClearColor(0, 0, 0, 0)
	
	prog := gl.CreateProgram()
	
	frag := gl.CreateShader(gl.FRAGMENT_SHADER)
	defer gl.DeleteShader(frag)
	
	vert := gl.CreateShader(gl.VERTEX_SHADER)
	defer gl.DeleteShader(vert)
	
	gl.ShaderSource(vert, vertexGLSL)
	gl.ShaderSource(frag, fragmentGLSL)
	
	if err := gl.CompileShader(vert); err != nil {
		return 0, errors.New("Vertex Shader " + err.Error())
	}
	if err := gl.CompileShader(frag); err != nil {
		return 0, errors.New("Fragment Shader " + err.Error())
	}
	gl.AttachShader(prog, vert)
	defer gl.DetachShader(prog, vert)
	
	gl.AttachShader(prog, frag)
	defer gl.DetachShader(prog, frag)
	
	if err := gl.LinkProgram(prog); err != nil {
		return 0, err
	}
	if err := gl.UseProgram(prog); err != nil {
		return 0, err
	}
	
	cells := make([]float32, 0, 2 * nrow * ncol)
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			x := float32(j) * 2 / float32(ncol)
			y := float32(i) * 2 / float32(nrow)
			cells = append(cells, x, y)
		}
	}
	
	buffers := gl.GenBuffers(1)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffers[0])
	gl.BufferData(gl.ARRAY_BUFFER, cells, gl.STATIC_DRAW)

	vao := gl.GenVertexArrays(1)[0]
	gl.BindVertexArray(vao)
		
	pos, err := gl.GetAttribLocation(prog, "position")
	if err != nil {
		return 0, err
	}
	gl.EnableVertexAttribArray(pos)
	gl.VertexAttribPointer(pos, 2, gl.Float32, false, 0, 0)
	
	return prog, nil
}

func Life(win *display.Buffer, prog gl.Program, nrow, ncol int) {
	clickCells := make(chan [2]int, 5)
	drawElems := make(chan []uint16, 1)
	
	go Simulate(nrow, ncol, clickCells, drawElems)
	Draw(win, nrow, ncol, prog, drawElems, clickCells)
}

func Draw(win *display.Buffer, nrow, ncol int, prog gl.Program, living chan []uint16, toggle chan [2]int) {
	var down bool
	var cells []uint16
	
	width, height := *width, *height
	
	padding := float32(2)
	pointSize := float32(width) / float32(nrow) - padding
	gl.PointSize(pointSize)
	offset,  _ := gl.GetUniformLocation(prog, "offset")
	halfPoint := 2/float32(width) * pointSize/2
	gl.Uniformf(offset, halfPoint - 1, -halfPoint - 1)
	
	drag := make(chan display.Cursor, 5)
	go func () {
		row, col := -1, -1
		for m := range drag {
			newRow := m.Y / (height / nrow)
			newCol := m.X / (width / ncol)
			log.Printf("(%d,%d) -> [%d][%d]", m.X, m.Y, newRow, newCol)
			if row != newRow || col != newCol {
				toggle <- [2]int{newRow, newCol}
				row, col = newRow, newCol
			}
		}
	}()
	
	frameClock := time.Tick(time.Second / 60)

	// Adjust with mouse wheel? TODO
	genClock := time.Tick(time.Millisecond * 20)
	
	// Initialize element array buffer
	buffers := gl.GenBuffers(1)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers[0])
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, make([]uint16, nrow * ncol), gl.STREAM_DRAW)
	
	for {
		select {
		case <-genClock:
			select {
			default:
				break
			case cells = <-living:
				gl.Clear(gl.COLOR_BUFFER_BIT)
				if len(cells) > 0 {
					gl.BufferSubData(gl.ELEMENT_ARRAY_BUFFER, 0, cells)
					gl.DrawElements(gl.POINTS, len(cells), gl.Uint16, 0)
				}
				win.Flip()
			}
		case event, ok := <-win.Event:
			if !ok {
				return
			}
			switch event := event.(type) {
			case display.Resize:
				width, height = event.Width, event.Height
				gl.Viewport(0, 0, width, height)

				size := width
				if width > height {
					size = height
				}
				
				pointSize = float32(size) / float32(nrow) - padding
				gl.PointSize(pointSize)
				halfPoint = 2/float32(size) * pointSize/2
				gl.Uniformf(offset, halfPoint - 1, -halfPoint - 1)
				
			case display.KeyPress:
				if event.Code == display.KeyEscape {
					close(drag)
					return
				}
			case display.Cursor:
				if down {
					drag <- event
				}
			case display.Mouse:
				down = event.Btn & display.Mouse1 > 0
			}
		case <-frameClock:
			win.CheckEvent()
		}
	}
}

func Simulate(nrow, ncol int, toggle chan [2]int, out chan []uint16) {
	living := make([]uint16, 0, nrow * ncol)
	grid := make([][]uint8, nrow)
	next := make([][]uint8, nrow)
	
	for i := range grid {
		grid[i] = make([]uint8, ncol)
		next[i] = make([]uint8, ncol)
	}
	
	for _, row := range(grid) {
		for col := range row {
			if rand.Int() % 14 == 0 {
				row[col] = 1
			}
		}
	}

Loop:
	for {
		select {
		case cell := <-toggle:
			row, col := cell[0], cell[1]
			if row > 0 && row < nrow && col > 0 && col < ncol {
				grid[row][col] ^= 1
			}
			continue Loop
		case out <-living:
			living = living[0:0]
			
			for y := range grid {
				for x, status := range grid[y] {
					neighbors := 0
					for i := -1; i < 2; i++ {
						for j := -1; j < 2; j++ {
							row := (y + i) % nrow
							col := (x + j) % ncol
							if row < 0 {
								row += nrow
							}
							if col < 0 {
								col += ncol
							}
							neighbors += int(grid[row][col])
						}
					}
					switch status {
					case Alive:
						neighbors-- // Don't count yourself
						if neighbors < 2 || neighbors > 3 {
							next[y][x] = Dead
						} else {
							next[y][x] = Alive
						}
					case Dead:
						if neighbors == 3 {
							next[y][x] = Alive
						} else {
							next[y][x] = Dead
						}
					default:
						panic("Never reached")
					}
					if next[y][x] == 1 {
						living = append(living, uint16(y * ncol + x))
					}
				}
			}
		}
		next, grid = grid, next
	}
}

func main() {
	flag.Parse()
	config["Geometry"] = fmt.Sprintf("%dx%d", *width, *height)
	
	log.SetFlags(0)
	log.SetPrefix("")
	
	win, err := display.Open(config)
	if err != nil {
		log.Fatal(err)
	}
	if prog, err := initGL(win, rows, columns); err != nil {
		log.Fatal(err)
	} else {
		Life(win, prog, rows, columns)
	}
}
