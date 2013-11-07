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
	// A texture is an OpenGL Object that contains one or more
	// images that all have the same image format. Textures can
	// be accessed from a shader, or used as a render target for
	// a Framebuffer.
	Texture      uint32
)

type (
	// A Type is an enumeration of an OpenGL type. The gl package
	// names types after their Go counterparts, rather than the
	// identifiers used in the OpenGL standard.
	Type       uint
)

type(
	Attachment uint
	// A TexFormat describes the representation of a texture in memory.
	TexFormat  uint
	TexParam   uint
)

const (
	DEPTH_ATTACHMENT   Attachment = 0x8D00
	STENCIL_ATTACHMENT Attachment = 0x8D20
	COLOR_ATTACHMENT0  Attachment = 0x8CE0
)

const (
	DEPTH_COMPONENT     TexFormat = 0x1902
	DEPTH_COMPONENT16   TexFormat = 0x81A5
	DEPTH_COMPONENT24   TexFormat = 0x81A6
	DEPTH_COMPONENT32   TexFormat = 0x81A7
	RED                 TexFormat = 0x1903
	RGB                 TexFormat = 0x1907
	RGBA                TexFormat = 0x1908
	RGBA12              TexFormat = 0x805A
	RGBA16              TexFormat = 0x805B
	RGBA2               TexFormat = 0x8055
	RGBA4               TexFormat = 0x8056
	RGBA8               TexFormat = 0x8058
	RGB4                TexFormat = 0x804F
	RGB5                TexFormat = 0x8050
	RGB8                TexFormat = 0x8051
	RGB565              TexFormat = 0x8D62
	RGB10               TexFormat = 0x8052
	RGB12               TexFormat = 0x8053
	RGB16               TexFormat = 0x8054
	INTENSITY           TexFormat = 0x8049
	INTENSITY12         TexFormat = 0x804C
	INTENSITY16         TexFormat = 0x804D
	INTENSITY4          TexFormat = 0x804A
	INTENSITY8          TexFormat = 0x804B
	LUMINANCE           TexFormat = 0x1909
	LUMINANCE12         TexFormat = 0x8041
	LUMINANCE12_ALPHA12 TexFormat = 0x8047
	LUMINANCE12_ALPHA4  TexFormat = 0x8046
	LUMINANCE16         TexFormat = 0x8042
	LUMINANCE16_ALPHA16 TexFormat = 0x8048
	LUMINANCE4          TexFormat = 0x803F
	LUMINANCE4_ALPHA4   TexFormat = 0x8043
	LUMINANCE6_ALPHA2   TexFormat = 0x8044
	LUMINANCE8          TexFormat = 0x8040
	LUMINANCE8_ALPHA8   TexFormat = 0x8045
	LUMINANCE_ALPHA     TexFormat = 0x190A
)

const (
	TEXTURE0  Texture = 0x84C0
	TEXTURE1  Texture = 0x84C1
	TEXTURE2  Texture = 0x84C2
	TEXTURE3  Texture = 0x84C3
	TEXTURE4  Texture = 0x84C4
	TEXTURE5  Texture = 0x84C5
	TEXTURE6  Texture = 0x84C6
	TEXTURE7  Texture = 0x84C7
	TEXTURE8  Texture = 0x84C8
	TEXTURE9  Texture = 0x84C9
	TEXTURE10 Texture = 0x84CA
	TEXTURE11 Texture = 0x84CB
	TEXTURE12 Texture = 0x84CC
	TEXTURE13 Texture = 0x84CD
	TEXTURE14 Texture = 0x84CE
	TEXTURE15 Texture = 0x84CF
	TEXTURE16 Texture = 0x84D0
	TEXTURE17 Texture = 0x84D1
	TEXTURE18 Texture = 0x84D2
	TEXTURE19 Texture = 0x84D3
	TEXTURE20 Texture = 0x84D4
	TEXTURE21 Texture = 0x84D5
	TEXTURE22 Texture = 0x84D6
	TEXTURE23 Texture = 0x84D7
	TEXTURE24 Texture = 0x84D8
	TEXTURE25 Texture = 0x84D9
	TEXTURE26 Texture = 0x84DA
	TEXTURE27 Texture = 0x84DB
	TEXTURE28 Texture = 0x84DC
	TEXTURE29 Texture = 0x84DD
	TEXTURE30 Texture = 0x84DE
	TEXTURE31 Texture = 0x84DF
)

const (
	// (0) index of the lowest defined mipmap level
	TEXTURE_BASE_LEVEL   TexParam = 0x813C
	// values that should be used for border texels
	TEXTURE_BORDER_COLOR TexParam = 0x1004
	// comparison when GL_TEXTURE_COMPARE_MODE is GL_COMPARE_REF_TO_TEXTURE
	TEXTURE_COMPARE_FUNC TexParam = 0x884D
	// comparison mode for currently bound depth textures
	TEXTURE_COMPARE_MODE TexParam = 0x884C
	// value that is to be added to the level-of-detail parameter for the texture before texture sampling
	TEXTURE_LOD_BIAS     TexParam = 0x8501 
	// Texture magnification function
	TEXTURE_MAG_FILTER   TexParam = 0x2800
	// Index of the highest defined mipmap level
	TEXTURE_MAX_LEVEL    TexParam = 0x813D
	// Sets the maximum level-of-detail parameter
	TEXTURE_MAX_LOD      TexParam = 0x813B
	// The texture minifying function
	TEXTURE_MIN_FILTER   TexParam = 0x2801
	// Sets the minimum level-of-detail parameter
	TEXTURE_MIN_LOD      TexParam = 0x813A
	TEXTURE_WRAP_R       TexParam = 0x8072
	TEXTURE_WRAP_S       TexParam = 0x2802
	TEXTURE_WRAP_T       TexParam = 0x2803
)

// GenTextures returns a slice of n Texture names.
// https://www.opengl.org/sdk/docs/man3/xhtml/glGenTextures.xml
func GenTextures(n int) []Texture {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Texture, n)
	C.goglGenTextures(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return buf
}

// Object deletion

func DeleteTextures(b []Texture) {
	if b == nil {
		C.goglDeleteTextures(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func ActiveTexture(tex Texture) error {
	C.goglActiveTexture(C.GLenum(tex))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid texture %v", tex)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func BindTexture(t Target, b Texture) error {
	C.goglBindTexture(C.GLenum(t), C.GLuint(b))
	err := C.goglGetError()
	switch err {
	case gl_INVALID_ENUM:
		return fmt.Errorf("Invalid target %v", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%x is not a Texture", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

func TexImage2D(tgt Target, lvl int, internal TexFormat, width, height int, format TexFormat, data interface{}) {
	typ, _, _, ptr := sliceInfo(data)
	C.goglTexImage2D(
		C.GLenum(tgt),
		C.GLint(lvl),
		C.GLint(internal),
		C.GLsizei(width),
		C.GLsizei(height),
		0,
		C.GLenum(format),
		C.GLenum(typ),
		ptr,
	)
}

func TexParameter(tgt Target, param TexParam, val interface{}) {
	switch val := val.(type) {
	case int:
		C.goglTexParameteri(C.GLenum(tgt), C.GLenum(param), C.GLint(val))
	case Enum:
		C.goglTexParameteri(C.GLenum(tgt), C.GLenum(param), C.GLint(val))
	case float32:
		C.goglTexParameterf(C.GLenum(tgt), C.GLenum(param), C.GLfloat(val))
	case []int:
		ptr := (*C.GLint)(nil)
		if len(val) > 0 {
			ptr = (*C.GLint)(unsafe.Pointer(&val[0]))
		}
		C.goglTexParameteriv(C.GLenum(tgt), C.GLenum(param), (*C.GLint)(ptr))
	case []float32:
		ptr := (*C.GLfloat)(nil)
		if len(val) > 0 {
			ptr = (*C.GLfloat)(unsafe.Pointer(&val[0]))
		}
		C.goglTexParameterfv(C.GLenum(tgt), C.GLenum(param), (*C.GLfloat)(ptr))
	default:
		panic(fmt.Sprintf("Invalid TexParameter value type %T", val))
	}
}
