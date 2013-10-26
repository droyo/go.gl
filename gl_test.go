package gl

import (
	"aqwari.net/exp/display"
	"image/color"
	"runtime"
	"testing"
	"time"
)

func init() {
	runtime.LockOSThread()
}

var (
	vertShader = `#version 120
		attribute vec2 position;
		void main() {
			gl_Position = vec4(position, 0, 1);
		}`
	fragShader = `#version 120
		varying vec4 outColor;
		void main() {
			outColor = vec4(1,1,1,1);
		}`
	config = display.Config{
		"OpenGL Version": "2.1",
		"Geometry":       "600x600",
		"Title":          "Laser Grid",
		"Resize":         "False",
		"Swap Interval":  "1",
	}
)

func TestWindow(t *testing.T) {
	win, err := display.Open(config)
	if err != nil {
		t.Fatal(err)
	}
	defer win.Close()
	
	if err := Init("3.2"); err != nil {
		t.Fatal(err)
	}
	ClearColor(color.Black)
	Clear(ColorBufferBit)
	time.Sleep(time.Second * 10)
}
