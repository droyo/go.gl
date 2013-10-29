package gl

/*
#cgo darwin  LDFLAGS: -framework OpenGL
#cgo linux   LDFLAGS: -lGL
#cgo windows LDFLAGS: -lopengl32

#include "gogl.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type (
	ClearFlags uint
	BlendMode  uint
	BlendFn    uint
	DepthFn    uint
	Enum       uint
)


const (
	VENDOR                   Enum = gl_VENDOR
	RENDERER                 Enum = gl_RENDERER
	VERSION                  Enum = gl_VERSION
	SHADING_LANGUAGE_VERSION Enum = gl_SHADING_LANGUAGE_VERSION
)

const (
	COLOR_BUFFER_BIT   ClearFlags = gl_COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT   ClearFlags = gl_DEPTH_BUFFER_BIT
	STENCIL_BUFFER_BIT ClearFlags = gl_STENCIL_BUFFER_BIT
)

const (
	FUNC_ADD              BlendMode = gl_FUNC_ADD
	FUNC_SUBTRACT         BlendMode = gl_FUNC_SUBTRACT
	FUNC_REVERSE_SUBTRACT BlendMode = gl_FUNC_REVERSE_SUBTRACT
	MIN                   BlendMode = gl_MIN
	MAX                   BlendMode = gl_MAX
)

const (
	ZERO                     BlendFn = gl_ZERO
	ONE                      BlendFn = gl_ONE
	SRC_COLOR                BlendFn = gl_SRC_COLOR
	ONE_MINUS_SRC_COLOR      BlendFn = gl_ONE_MINUS_SRC_COLOR
	DST_COLOR                BlendFn = gl_DST_COLOR
	ONE_MINUS_DST_COLOR      BlendFn = gl_ONE_MINUS_DST_COLOR
	SRC_ALPHA                BlendFn = gl_SRC_ALPHA
	ONE_MINUS_SRC_ALPHA      BlendFn = gl_ONE_MINUS_SRC_ALPHA
	DST_ALPHA                BlendFn = gl_DST_ALPHA
	ONE_MINUS_DST_ALPHA      BlendFn = gl_ONE_MINUS_DST_ALPHA
	CONSTANT_COLOR           BlendFn = gl_CONSTANT_COLOR
	ONE_MINUS_CONSTANT_COLOR BlendFn = gl_ONE_MINUS_CONSTANT_COLOR
	CONSTANT_ALPHA           BlendFn = gl_CONSTANT_ALPHA
	ONE_MINUS_CONSTANT_ALPHA BlendFn = gl_ONE_MINUS_CONSTANT_ALPHA
)

const (
	NEVER    DepthFn = gl_NEVER
	LESS     DepthFn = gl_LESS
	EQUAL    DepthFn = gl_EQUAL
	LEQUAL   DepthFn = gl_LEQUAL
	GREATER  DepthFn = gl_GREATER
	NOTEQUAL DepthFn = gl_NOTEQUAL
	GEQUAL   DepthFn = gl_GEQUAL
	ALWAYS   DepthFn = gl_ALWAYS
)

func GetString(key Enum) string {
	val := unsafe.Pointer(C.goglGetString(C.GLenum(key)))
	return C.GoString((*C.char)(val))
}

func Init(version string) error {
	if err := C.goglInit(); err != nil {
		defer C.free(unsafe.Pointer(err))
		return errors.New("Could not initialize function " + C.GoString(err))
	}
	return nil
}

func clamp(x float32) C.GLfloat {
	switch {
	case x > 1:
		return 1
	case x < 0:
		return 0
	default:
		return C.GLfloat(x)
	}
}

func clamp64(x float64) C.GLdouble {
	switch {
	case x > 1:
		return 1
	case x < 0:
		return 0
	default:
		return C.GLdouble(x)
	}
}

func BlendColor(r, g, b, a float32) {
	C.goglBlendColor(clamp(r), clamp(g), clamp(b), clamp(a))
}

func BlendEquation(mode BlendMode) {
	C.goglBlendEquation(C.GLenum(mode))
}

func BlendEquationSeparate(rgb, alpha BlendMode) {
	C.goglBlendEquationSeparate(C.GLenum(rgb), C.GLenum(alpha))
}

func BlendFunc(src, dst BlendFn) {
	C.goglBlendFunc(C.GLenum(src), C.GLenum(dst))
}

func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha BlendFn) {
	C.goglBlendFuncSeparate(C.GLenum(srcRGB), C.GLenum(dstRGB), C.GLenum(srcAlpha), C.GLenum(dstAlpha))
}

func DepthFunc(fn DepthFn) {
	C.goglDepthFunc(C.GLenum(fn))
}

func DepthRange(near, far float64) {
	C.goglDepthRange(clamp64(near), clamp64(far))
}

func DepthMask(on bool) {
	if on {
		C.goglDepthMask(gl_TRUE)
	} else {
		C.goglDepthMask(gl_FALSE)
	}
}

func ClearColor(r, g, b, a float32) {
	C.goglClearColor(clamp(r), clamp(g), clamp(b), clamp(a))
}

func ClearDepth(depth float64) {
	C.goglClearDepth(clamp64(depth))
}

func ClearStencil(index int) {
	C.goglClearStencil(C.GLint(index))
}

func Clear(bit ClearFlags) {
	C.goglClear(C.GLbitfield(bit))
}

func DrawArrays(draw DrawMode, first, count int) {
	C.goglDrawArrays(C.GLenum(draw), C.GLint(first), C.GLsizei(count))
}

func Viewport(x, y, w, h int) {
	C.goglViewport(C.GLint(x), C.GLint(y), C.GLsizei(w), C.GLsizei(h))
	if C.goglGetError() == gl_INVALID_VALUE {
		panic("negative width or height")
	}
}

// BufferData uploads vertex data from the system into server memory.
// The argument data must be a slice or array of numeric values. The
// data must be a non-nil slice or array of numeric elements. Passing
// a nil or forbidden type will result in a run-time panic.
func BufferData(t Target, data interface{}, hint UsageHint) error {
	_, size, count, ptr := sliceInfo(data)
	C.goglBufferData(C.GLenum(t), C.GLsizeiptr(size * uintptr(count)), ptr, C.GLenum(hint))
	return getError()
}

func BufferSubData(t Target, offset uintptr, data interface{}, hint UsageHint) error {
	_, size, count, ptr := sliceInfo(data)
	C.goglBufferSubData(C.GLenum(t), C.GLintptr(offset), C.GLsizeiptr(size * uintptr(count)), ptr)
	return getError()
}

func sliceInfo(v interface{}) (fmt Type, size uintptr, count int, ptr unsafe.Pointer) {
        val := reflect.ValueOf(v)
        typ := reflect.TypeOf(v)

		switch typ.Kind() {
		case reflect.Array, reflect.Slice:
			switch typ.Elem().Kind() {
			case reflect.Float32:
				fmt = Float
			default:
				panic("Invalid element type " + typ.Elem().String())
			}
			size = typ.Elem().Size()
			count = val.Len()
			if val.IsNil() {
				ptr = nil
			} else {
				ptr = unsafe.Pointer(val.Index(0).UnsafeAddr())
			}
		default:
			panic("value must be a slice or array")
		}
		return
}
