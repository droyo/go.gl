package gl


/*
#include "gogl.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type (
	Attrib       uint32
	// A Vertex Array object stores the state that describes
	// arrays of vertex data that will be used in rendering commands.
	VertexArray  uint32
	// Buffers hold data in server memory. They can be bound
	// to ARRAY_BUFFER with BindBuffer to hold vertex and color
	// data uploaded with BufferData, or they can be bound to
	// ELEMENT_ARRAY_BUFFER to reference vertices in another
	// buffer.
	Buffer       uint32
	// Framebuffers are off-screen drawing destinations that can
	// be used to render to Textures or Renderbuffers with the
	// same GLSL drawing routines used to render to the display.
	Framebuffer  uint32
	// Renderbuffers provide storage for Framebuffer objects. They
	// must be initialized with RenderBufferStorage and attached
	// to a Framebuffer with FrameBufferRenderbuffer
	Renderbuffer uint32
)

type(
	// A Target specifies the slot for a bound OpenGL object, such
	// as a vertex buffer bound to ARRAY_BUFFER
	Target     uint
	// A UsageHint informs OpenGL of the intended use case for a 
	// buffer object. The OpenGL implementation is not required to
	// act on a UsageHint, but often will.
	UsageHint  uint
	// DrawMode informs OpenGL how to assemble vertex data into
	// shapes.
	DrawMode   uint
)

const (
	POINTS         DrawMode = 0x0000 // Independent points
	LINE_STRIP     DrawMode = 0x0003 // Points connected with lines
	LINE_LOOP      DrawMode = 0x0002 // Line strip with first and last points joined
	LINES          DrawMode = 0x0001 // Series of two-point straigth lines
	TRIANGLE_STRIP DrawMode = 0x0005 // Triangle vertex N has vertices N-1, N-2
	TRIANGLE_FAN   DrawMode = 0x0006 // Triangles sharing same central vertex
	TRIANGLES      DrawMode = 0x0004 // Independent triangles
)

const (
	ARRAY_BUFFER                 Target = 0x8892
	ELEMENT_ARRAY_BUFFER         Target = 0x8893
	FRAMEBUFFER                  Target = 0x8D40
	RENDERBUFFER                 Target = 0x8D41
	TEXTURE_1D                   Target = 0x0DE0
	TEXTURE_2D                   Target = 0x0DE1
	TEXTURE_3D                   Target = 0x806F
	TEXTURE_BUFFER               Target = 0x8C2A
	TEXTURE_CUBE_MAP             Target = 0x8513
	TEXTURE_RECTANGLE            Target = 0x84F5
	UNIFORM_BUFFER               Target = 0x8A11
)

const (
	STREAM_DRAW  UsageHint = 0x88E0 // Modify once, use seldom.
	STATIC_DRAW  UsageHint = 0x88E4 // Modify once, use often.
	DYNAMIC_DRAW UsageHint = 0x88E8 // Modify often, use often.
)

// GenBuffers returns a slice of n buffer object names. 
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenBuffers.xml
func GenBuffers(n int) []Buffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Buffer, n)
	C.goglGenBuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Buffer(buf)
}

// Bind makes a Buffer the current active Buffer.
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func BindBuffer(t Target, b Buffer) error {
	C.goglBindBuffer(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %v", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%x is not a Buffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

// DeleteBuffers deletes several buffers at once.
func DeleteBuffers(b []Buffer) {
	if b != nil {
		C.goglDeleteBuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

// GenFramebuffers returns a slice of n framebuffer object names.
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenFramebuffers.xml
func GenFramebuffers(n int) []Framebuffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Framebuffer, n)
	C.goglGenFramebuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Framebuffer(buf)
}

func BindFramebuffer(t Target, b Framebuffer) error {
	C.goglBindFramebuffer(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %v", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%x is not a Framebuffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func FramebufferRenderbuffer(tgt Target, a Attachment, bufTgt Target, buf Renderbuffer) error {
	C.goglFramebufferRenderbuffer(
		C.GLenum(tgt),
		C.GLenum(a),
		C.GLenum(bufTgt),
		C.GLuint(buf))
	return getError()
}

func FramebufferTexture2D(tgt Target, a Attachment, texTgt Target, tex Texture, lvl int) error {
	C.goglFramebufferTexture2D(
		C.GLenum(tgt),
		C.GLenum(a),
		C.GLenum(texTgt),
		C.GLuint(tex),
		C.GLint(lvl))
	return getError()
}

func DeleteFramebuffers(b []Framebuffer) {
	if b != nil {
		C.goglDeleteFramebuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

// GenRenderbuffers returns a slice of n renderbuffer object names.
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenRenderbuffers.xml
func GenRenderbuffers(n int) []Renderbuffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Renderbuffer, n)
	C.goglGenRenderbuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Renderbuffer(buf)
}

func BindRenderbuffer(t Target, b Renderbuffer) error {
	C.goglBindRenderbuffer(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %v", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%x is not a Renderbuffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func DeleteRenderbuffers(b []Renderbuffer) {
	if b != nil {
		C.goglDeleteRenderbuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

// GenVertexArrays returns a slice of n vertex array object names.
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenVertexArrays.xml
func GenVertexArrays(n int) []VertexArray {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]VertexArray, n)
	C.goglGenVertexArrays(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []VertexArray(buf)
}

func BindVertexArray(vao VertexArray) error {
	C.goglBindVertexArray(C.GLuint(vao))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_OPERATION:
		return fmt.Errorf("%x is not a Vertex Array", uint(vao))
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func EnableVertexAttribArray(a Attrib) {
	C.goglEnableVertexAttribArray(C.GLuint(a))
}

func DisableVertexAttribArray(a Attrib) {
	C.goglDisableVertexAttribArray(C.GLuint(a))
}

func VertexAttribPointer(a Attrib, size int, typ Type, normalize bool, stride, offset uintptr) error {
	norm := 0
	if normalize {
		norm = 1
	}
	C.goglVertexAttribPointer(
		C.GLuint(a), C.GLint(size), C.GLenum(typ),
		C.GLboolean(norm), C.GLsizei(stride),
		unsafe.Pointer(offset),
	)
	return getError()
}

func DeleteVertexArrays(b []VertexArray) {
	if b != nil {
		C.goglDeleteVertexArrays(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}
