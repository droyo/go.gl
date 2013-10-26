package gl

/*
#include "gogl.h"
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Type int

const (
	Float  Type = gl_FLOAT
	Int    Type = gl_INT
	Int16  Type = gl_SHORT
	Int8   Type = gl_BYTE
	Uint   Type = gl_UNSIGNED_INT
	Uint16 Type = gl_UNSIGNED_SHORT
	Uint8  Type = gl_UNSIGNED_BYTE
)

type Target int

const (
	ArrayBuffer        Target = gl_ARRAY_BUFFER
	ElementArrayBuffer Target = gl_ELEMENT_ARRAY_BUFFER
)

type (
	Buffer       uint32
	Framebuffer  uint32
	UsageHint    uint32
	Renderbuffer uint32
)

const (
	StreamDraw  UsageHint = 0x88E0 // Modify once, use seldom.
	StaticDraw  UsageHint = 0x88E4 // Modify once, use often.
	DynamicDraw UsageHint = 0x88E8 // Modify often, use often.
)

// GenBuffers creates a slice of N Buffer objects.
func GenBuffers(n int) []Buffer {
	buf := make([]Buffer, n)
	C.goglGenBuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Negative value (%d) to GenBuffers", n))
	}
	return buf
}

func GenFramebuffers(n int) []Framebuffer {
	buf := make([]Framebuffer, n)
	C.goglGenFramebuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Negative value (%d) to GenFramebuffers", n))
	}
	return buf
}

func GenRenderbuffers(n int) []Renderbuffer {
	buf := make([]Renderbuffer, n)
	C.goglGenRenderbuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Negative value (%d) to GenRenderbuffers", n))
	}
	return buf
}

// DeleteBuffers deletes several buffers at once.
func DeleteBuffers(b []Buffer) {
	if b == nil {
		panic("nil Buffer slice")
	}
	C.goglDeleteBuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

func DeleteFramebuffers(b []Framebuffer) {
	if b == nil {
		panic("nil Buffer slice")
	}
	C.goglDeleteFramebuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

func DeleteRenderbuffers(b []Renderbuffer) {
	if b == nil {
		panic("nil Buffer slice")
	}
	C.goglDeleteRenderbuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

// Bind makes a Buffer the current active Buffer.
// http://www.opengl.org/sdk/docs/man3/xhtml/glBindBuffer.xml
func (b Buffer) Bind(t Target) error {
	C.goglBindBuffer(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %d", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%d is not a Buffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func (b Framebuffer) Bind(t Target) error {
	C.goglBindFramebuffer(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %d", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%d is not a Framebuffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

// BufferData uploads vertex data from the system into server memory.
// The argument data must be a non-nil slice of numeric values. The
// data must be a non-nil slice or array of numeric elements. Passing
// a nil or forbidden type will result in a run-time panic.
func BufferData(t Target, data interface{}, hint UsageHint) error {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		if val.IsNil() {
			panic("Nil data argument")
		}
		switch typ.Elem().Kind() {
		case reflect.Int,reflect.Int32,reflect.Int64,reflect.Uint,reflect.Uint32,reflect.Uint64,reflect.Float32,reflect.Float64:
			C.goglBufferData(
				C.GLenum(t),
				C.GLsizeiptr(typ.Elem().Size()) * C.GLsizeiptr(val.Len()),
				unsafe.Pointer(val.Index(0).UnsafeAddr()),
				C.GLenum(hint))
		default:
			panic("Non-numeric data type " + typ.Elem().String())
		}
	default:
		panic("Non-slice or array data type " + typ.String())
	}
	return getError()
}

type Mode uint32
const (
        Points Mode = gl_POINTS
        LineStrip Mode = gl_LINE_STRIP
        LineLoop Mode = gl_LINE_LOOP
        Lines Mode = gl_LINES
        TriangleStrip Mode = gl_TRIANGLE_STRIP
        TriangleFan Mode = gl_TRIANGLE_FAN
        Triangles Mode = gl_TRIANGLES
)

func DrawArrays(draw Mode, first, count int) {
	C.goglDrawArrays(C.GLenum(draw), C.GLint(first), C.GLsizei(count))
}
