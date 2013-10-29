package main

import (
	"aqwari.net/exp/gl"
	"aqwari.net/exp/display"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	vertShader = []byte(`#version 150
		in vec2 position;
		void main() {
			gl_Position = vec4(position, 0, 1);
		}`)
	fragShader = []byte(`#version 150
		out vec4 color;
		void main() {
			color = vec4(1,1,1,1);
		}`)
	config = display.Config{
		"OpenGL Version": "3.2",
		"Geometry":       "600x600",
		"Title":          "Laser Grid",
		"Resize":         "False",
		"Swap Interval":  "1",
	}
)

func Display(win *display.Buffer, rows, cols int) {
	if err := gl.Init(config["OpenGL Version"]); err != nil {
		log.Fatalln(err)
	}
	clock := time.Tick(time.Second)
	log.Print("GL Version ", gl.GetString(gl.VERSION))
	log.Print("GLSL Version ", gl.GetString(gl.SHADING_LANGUAGE_VERSION))
	log.Print("Renderer ", gl.GetString(gl.RENDERER))
	log.Print("Vendor ", gl.GetString(gl.VENDOR))

	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Create our Shader program
	pg := gl.CreateProgram()
	defer gl.DeleteProgram(pg)

	vs := gl.CreateShader(gl.VERTEX_SHADER)
	defer gl.DeleteShader(vs)
	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	defer gl.DeleteShader(fs)

	if err := gl.ShaderSource(vs, vertShader); err != nil {
		log.Fatal(err)
	}
	if err := gl.CompileShader(vs); err != nil {
		log.Fatal("Vert ", err)
	}
	if err := gl.ShaderSource(fs, fragShader); err != nil {
		log.Fatal(err)
	}
	if err := gl.CompileShader(fs); err != nil {
		log.Fatal("Frag ", err)
	}
	if err := gl.AttachShader(pg, vs); err != nil {
		log.Fatal(err)
	}
	defer gl.DetachShader(pg, vs)
	if err := gl.AttachShader(pg, fs); err != nil {
		log.Fatal(err)
	}
	defer gl.DetachShader(pg, fs)

	if err := gl.LinkProgram(pg); err != nil {
		log.Fatal(err)
	}
	if err := gl.UseProgram(pg); err != nil {
		log.Fatal(err)
	}

	lines := make([]float32, 0, (rows-1)*4+(cols-1)*4)
	for i, sz := 1, 2/float32(rows); i < rows; i++ {
		lines = append(lines,
			-1.0, 1.0-(sz*float32(i)),
			1.0, 1.0-(sz*float32(i)),
		)
	}
	for i, sz := 1, 2/float32(cols); i < cols; i++ {
		lines = append(lines,
			1.0-(sz*float32(i)), 1.0,
			1.0-(sz*float32(i)), -1.0,
		)
	}

	buf := gl.GenBuffers(1)[0]
	if err := gl.BindBuffer(gl.ARRAY_BUFFER, buf); err != nil {
		log.Fatal(err)
	}


	// Inform GL of our vertex structure.
	vao := gl.GenVertexArrays(1)[0]
	if err := gl.BindVertexArray(vao); err != nil {
		log.Fatal(err)
	}
	attr, err := gl.GetAttribLocation(pg, "position")
	if err != nil {
		log.Fatal(err)
	}
	
	gl.EnableVertexAttribArray(attr)
	gl.VertexAttribPointer(attr, 2, gl.Float, false, 0, 0)
	
	if err := gl.BufferData(gl.ARRAY_BUFFER, lines, gl.STATIC_DRAW); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case e, ok := <-win.Event:
			if !ok {
				log.Print("Closed window")
				return
			}
			switch e := e.(type) {
			case display.KeyPress:
				if e.Code == display.KeyEscape {
					win.Close()
					return
				}
			}
		case <-clock:
			gl.Clear(gl.COLOR_BUFFER_BIT)
			gl.DrawArrays(gl.LINES, 0, len(lines)/2)
			win.Flip()
			win.CheckEvent()
		}
	}
}

func main() {
	var nrow, ncol int
	var err error
	var win *display.Buffer

	log.SetFlags(0)
	log.SetPrefix("")
	if len(os.Args) != 3 {
		log.Fatal("Usage: draw_grid nrows ncols")
	}
	nrow, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ncol, err = strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if nrow < 1 || ncol < 1 {
		log.Fatal("Must have positive number of rows & columns")
	}
	win, err = display.Open(config)
	if err != nil {
		log.Fatal(err)
	}
	Display(win, nrow, ncol)
}
