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
	"unsafe"
)

type (
	ClearFlags uint
	BlendMode  uint
	BlendFn    uint
	DepthFn    uint
	Enum       uint
)

type (
	Capability   uint
	WindingOrder uint
	CullMode     uint
)

const (
	Float32 Type = 0x1406
	Float64 Type = 0x140A
	Int8    Type = 0x1400
	Int16   Type = 0x1402
	Int32   Type = 0x1404
	Uint8   Type = 0x1401
	Uint16  Type = 0x1403
	Uint32  Type = 0x1405
)

const (
	VENDOR                   Enum = 0x1F00
	RENDERER                 Enum = 0x1F01
	VERSION                  Enum = 0x1F02
	SHADING_LANGUAGE_VERSION Enum = 0x8B8C
)

const (
	BACK           CullMode = 0x0405
	FRONT          CullMode = 0x0404
	FRONT_AND_BACK CullMode = 0x0408
)

const (
	CW  WindingOrder = 0x0900 // Clockwise
	CCW WindingOrder = 0x0901 // Counter-clockwise
)

const (
	BLEND                     Capability = 0x0BE2
	CLIP_DISTANCE0            Capability = 0x3000
	CLIP_DISTANCE1            Capability = 0x3001
	CLIP_DISTANCE2            Capability = 0x3002
	CLIP_DISTANCE3            Capability = 0x3003
	CLIP_DISTANCE4            Capability = 0x3004
	CLIP_DISTANCE5            Capability = 0x3005
	CLIP_DISTANCE6            Capability = 0x3006
	CLIP_DISTANCE7            Capability = 0x3007
	COLOR_LOGIC_OP            Capability = 0x0BF2
	CULL_FACE                 Capability = 0x0B44
	DEPTH_CLAMP               Capability = 0x864F
	DEPTH_TEST                Capability = 0x0B71
	DITHER                    Capability = 0x0BD0
	FRAMEBUFFER_SRGB          Capability = 0x8DB9
	LINE_SMOOTH               Capability = 0x0B20
	MULTISAMPLE               Capability = 0x809D
	POLYGON_OFFSET_FILL       Capability = 0x8037
	POLYGON_OFFSET_LINE       Capability = 0x2A02
	POLYGON_OFFSET_POINT      Capability = 0x2A01
	POLYGON_SMOOTH            Capability = 0x0B41
	PRIMITIVE_RESTART         Capability = 0x8F9D
	PROGRAM_POINT_SIZE        Capability = 0x8642
	SAMPLE_ALPHA_TO_COVERAGE  Capability = 0x809E
	SAMPLE_ALPHA_TO_ONE       Capability = 0x809F
	SAMPLE_COVERAGE           Capability = 0x80A0
	SCISSOR_TEST              Capability = 0x0C11
	TEXTURE_CUBE_MAP_SEAMLESS Capability = 0x884F
)

const (
	COLOR_BUFFER_BIT   ClearFlags = 0x00004000
	DEPTH_BUFFER_BIT   ClearFlags = 0x00000100
	STENCIL_BUFFER_BIT ClearFlags = 0x00000400
)

const (
	FUNC_ADD              BlendMode = 0x8006
	FUNC_SUBTRACT         BlendMode = 0x800B
	FUNC_REVERSE_SUBTRACT BlendMode = 0x800A
	MIN                   BlendMode = 0x8007
	MAX                   BlendMode = 0x8008
)

const (
	ZERO                     BlendFn = 0x0000
	ONE                      BlendFn = 0x0001
	SRC_COLOR                BlendFn = 0x0300
	ONE_MINUS_SRC_COLOR      BlendFn = 0x0301
	DST_COLOR                BlendFn = 0x0306
	ONE_MINUS_DST_COLOR      BlendFn = 0x0307
	SRC_ALPHA                BlendFn = 0x0302
	ONE_MINUS_SRC_ALPHA      BlendFn = 0x0303
	DST_ALPHA                BlendFn = 0x0304
	ONE_MINUS_DST_ALPHA      BlendFn = 0x0305
	CONSTANT_COLOR           BlendFn = 0x8002
	ONE_MINUS_CONSTANT_COLOR BlendFn = 0x8002
	CONSTANT_ALPHA           BlendFn = 0x8003
	ONE_MINUS_CONSTANT_ALPHA BlendFn = 0x8004
)

const (
	NEVER    DepthFn = 0x0200
	LESS     DepthFn = 0x0201
	EQUAL    DepthFn = 0x0202
	LEQUAL   DepthFn = 0x0203
	GREATER  DepthFn = 0x0204
	NOTEQUAL DepthFn = 0x0205
	GEQUAL   DepthFn = 0x0206
	ALWAYS   DepthFn = 0x0207
)

// GetString can be used to retrieve information about the
// loaded OpenGL implementation, such as the supported shader
// version.
func GetString(key Enum) string {
	val := unsafe.Pointer(C.goglGetString(C.GLenum(key)))
	return C.GoString((*C.char)(val))
}

// Init loads the OpenGL functions using system-specific
// library loaders. The string hint is currently unused
// but reserved for future use.
func Init(hint string) error {
	if err := C.goglInit(); err != nil {
		defer C.free(unsafe.Pointer(err))
		return errors.New("Could not initialize function " + C.GoString(err))
	}
	return nil
}

func Enable(cap Capability) {
	C.goglEnable(C.GLenum(cap))
}

func Disable(cap Capability) {
	C.goglDisable(C.GLenum(cap))
}

func IsEnabled(cap Capability) bool {
	return C.goglIsEnabled(C.GLenum(cap)) == 1
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

// https://www.opengl.org/sdk/docs/man3/xhtml/glDepthRange.xml
func DepthRange(near, far float64) {
	C.goglDepthRange(clamp64(near), clamp64(far))
}

// DepthMask enables or disables writing to the
// depth buffer.
// https://www.opengl.org/sdk/docs/man3/xhtml/glDepthMask.xml
func DepthMask(on bool) {
	if on {
		C.goglDepthMask(1)
	} else {
		C.goglDepthMask(0)
	}
}

// ClearColor defines the preset value to set the
// color buffer to when using Clear(COLOR_BUFFER_BIT)
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearDepth.xml
func ClearColor(r, g, b, a float32) {
	C.goglClearColor(clamp(r), clamp(g), clamp(b), clamp(a))
}

// ClearDepth defines the preset value to set the
// depth buffer to when using Clear(DEPTH_BUFFER_BIT)
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearDepth.xml
func ClearDepth(depth float64) {
	C.goglClearDepth(clamp64(depth))
}

// ClearStencil defines the preset value to set the
// stencil buffer to when using Clear(STENCIL_BUFFER_BIT).
// https://www.opengl.org/sdk/docs/man3/xhtml/glClearStencil.xml
func ClearStencil(index int) {
	C.goglClearStencil(C.GLint(index))
}

// Clear clears buffers to preset values.
// https://www.opengl.org/sdk/docs/man3/xhtml/glClear.xml
func Clear(bit ClearFlags) {
	C.goglClear(C.GLbitfield(bit))
}

// DrawArrays renders primitives from vertex data in ARRAY_BUFFER.
// A buffer should be bound to ARRAY_BUFFER with BindBuffer, and
// initialized with vertex data using BufferData before calling
// DrawArrays.
//
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawArrays.xml
func DrawArrays(mode DrawMode, first, count int) {
	C.goglDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

// DrawElements draws vertices in the buffer object bound to 
// ARRAY_BUFFER as selected by the indices in the buffer object bound to 
// ELEMENT_ARRAY_BUFFER.
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElements.xml
func DrawElements(mode DrawMode, count int, typ Type, offset uintptr) {
	C.goglDrawElements(
		C.GLenum(mode),
		C.GLsizei(count),
		C.GLenum(typ),
		unsafe.Pointer(offset))
}

// DrawElementsBaseVertex draws in the same manner as DrawElements, but
// adds a fixed offset to each index in ELEMENT_ARRAY_BUFFER before 
// selecting vertices.
// https://www.opengl.org/sdk/docs/man3/xhtml/glDrawElementsBaseVertex.xml
func DrawElementsBaseVertex(mode DrawMode, count int, typ Type, offset uintptr, base int) {
	C.goglDrawElementsBaseVertex(
		C.GLenum(mode),
		C.GLsizei(count),
		C.GLenum(typ),
		unsafe.Pointer(offset),
		C.GLint(base))
}

// Viewport sets the size and position of the screen within the window.
// https://www.opengl.org/sdk/docs/man3/xhtml/glViewport.xml
func Viewport(x, y, w, h int) {
	C.goglViewport(C.GLint(x), C.GLint(y), C.GLsizei(w), C.GLsizei(h))
	if C.goglGetError() == gl_INVALID_VALUE {
		panic("negative width or height")
	}
}

// BufferData uploads vertex data from the system into server memory.
// The argument data must be a slice or array of numeric values. Passing
// a nil or forbidden type will result in a run-time panic.
func BufferData(t Target, data interface{}, hint UsageHint) error {
	_, size, count, ptr := sliceInfo(data)
	C.goglBufferData(C.GLenum(t), C.GLsizeiptr(size * uintptr(count)), ptr, C.GLenum(hint))
	return getError()
}

// BufferSubData replaces a portion of existing data in a buffer with new
// data from client memory. The target buffer must have already been
// initialized with BufferData, and the data to upload must be an array or
// slice of numeric values.
func BufferSubData(t Target, offset uintptr, data interface{}) error {
	_, size, count, ptr := sliceInfo(data)
	C.goglBufferSubData(C.GLenum(t), C.GLintptr(offset), C.GLsizeiptr(size * uintptr(count)), ptr)
	return getError()
}

// CullFace specifies whether front- or back-facing facets are culled.
// https://www.opengl.org/sdk/docs/man3/xhtml/glCullFace.xml
func CullFace(mode CullMode) {
	C.goglCullFace(C.GLenum(mode))
}

// FrontFace specifies the winding order when using facet culling.
// https://www.opengl.org/sdk/docs/man3/xhtml/glFrontFace.xml
func FrontFace(order WindingOrder) {
	C.goglFrontFace(C.GLenum(order))
}
