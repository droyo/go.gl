package main

import (
	"aqwari.net/exp/display"
	"aqwari.net/exp/gl"
	"log"
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

func Display(win *display.Buffer) {
	if err := gl.Init(config["OpenGL Version"]); err != nil {
		log.Fatalln(err)
	}
	clock := time.Tick(time.Second / 30)

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
	gl.ShaderSource(vs, vertShader)
	gl.ShaderSource(fs, fragShader)
	
	if err := gl.CompileShader(vs); err != nil {
		log.Fatal("Vert ", err)
	}
	if err := gl.CompileShader(fs); err != nil {
		log.Fatal("Frag ", err)
	}
	gl.AttachShader(pg, vs)
	defer gl.DetachShader(pg, vs)
	
	gl.AttachShader(pg, fs)
	defer gl.DetachShader(pg, fs)

	if err := gl.LinkProgram(pg); err != nil {
		log.Fatal(err)
	}
	if err := gl.UseProgram(pg); err != nil {
		log.Fatal(err)
	}

	// Upload our Vertices
	triPoints := []float32{
		0.75, 0.75,
		0.75, -0.75,
		-0.75, -0.75,
	}

	buffers := gl.GenBuffers(1)
	buf := buffers[0]
	defer gl.DeleteBuffers(buffers)
	gl.BindBuffer(gl.ARRAY_BUFFER, buf)

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
	if err := gl.VertexAttribPointer(attr, 2, gl.Float, false, 0, 0); err != nil {
		log.Fatal(err)
	}
	if err := gl.BufferData(gl.ARRAY_BUFFER, triPoints, gl.STATIC_DRAW); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case e, ok := <-win.Event:
			if !ok {
				return
			}
			switch e := e.(type) {
			case display.KeyPress:
				if e.Code == display.KeyEscape {
					win.Close()
					return
				}
				log.Print(e)
			}
		case <-clock:
			gl.Clear(gl.COLOR_BUFFER_BIT)
			gl.DrawArrays(gl.TRIANGLES, 0, 3)
			win.Flip()
			win.CheckEvent()
		}
	}
}

func main() {
	win, err := display.Open(config)
	if err != nil {
		log.Fatalln(err)
	}
	Display(win)
}
