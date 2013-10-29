package gl

/*
#include "gogl.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"errors"
	"unsafe"
)

type (
	Program      uint32
	Shader       uint32
	ShaderType   uint32
	Uniform      int32
	Attrib       uint32
	VertexArray  uint32
	Buffer       uint32
	Framebuffer  uint32
	Renderbuffer uint32
	Texture      uint32
)

type (
	Type       uint
	Target     uint
	UsageHint  uint
	DrawMode   uint
	Attachment uint
	TexFormat  uint
	TexParam   uint
)

const (
	DEPTH_ATTACHMENT   Attachment = gl_DEPTH_ATTACHMENT
	STENCIL_ATTACHMENT Attachment = gl_STENCIL_ATTACHMENT
	COLOR_ATTACHMENT0  Attachment = gl_COLOR_ATTACHMENT0
)

const (
	POINTS        DrawMode = gl_POINTS
	LINE_STRIP    DrawMode = gl_LINE_STRIP
	LINE_LOOP     DrawMode = gl_LINE_LOOP
	LINES         DrawMode = gl_LINES
	TRIAGLE_STRIP DrawMode = gl_TRIANGLE_STRIP
	TRIANGLE_FAN  DrawMode = gl_TRIANGLE_FAN
	TRIANGLES     DrawMode = gl_TRIANGLES
)

const (
	DEPTH_COMPONENT       TexFormat = gl_DEPTH_COMPONENT
	DEPTH_COMPONENT16     TexFormat = gl_DEPTH_COMPONENT16
	DEPTH_COMPONENT24     TexFormat = gl_DEPTH_COMPONENT24
	DEPTH_COMPONENT32     TexFormat = gl_DEPTH_COMPONENT32
	RGB                   TexFormat = gl_RGB
	RGBA                  TexFormat = gl_RGBA
)

const (
	Float  Type = gl_FLOAT
	Int    Type = gl_INT
	Int16  Type = gl_SHORT
	Int8   Type = gl_BYTE
	Uint   Type = gl_UNSIGNED_INT
	Uint16 Type = gl_UNSIGNED_SHORT
	Uint8  Type = gl_UNSIGNED_BYTE
)

const (
	ARRAY_BUFFER                 Target = gl_ARRAY_BUFFER
	ELEMENT_ARRAY_BUFFER         Target = gl_ELEMENT_ARRAY_BUFFER
	FRAMEBUFFER                  Target = gl_FRAMEBUFFER
	RENDERBUFFER                 Target = gl_RENDERBUFFER
	TEXTURE_1D                   Target = gl_TEXTURE_1D
	TEXTURE_2D                   Target = gl_TEXTURE_2D
	TEXTURE_3D                   Target = gl_TEXTURE_3D
	TEXTURE_BUFFER               Target = 0x8C2A
	TEXTURE_CUBE_MAP             Target = gl_TEXTURE_CUBE_MAP
	TEXTURE_RECTANGLE            Target = 0x84F5
	UNIFORM_BUFFER               Target = 0x8A11
)

const (
	STREAM_DRAW  UsageHint = 0x88E0 // Modify once, use seldom.
	STATIC_DRAW  UsageHint = 0x88E4 // Modify once, use often.
	DYNAMIC_DRAW UsageHint = 0x88E8 // Modify often, use often.
)

const (
	VERTEX_SHADER   ShaderType = gl_VERTEX_SHADER
	FRAGMENT_SHADER ShaderType = gl_FRAGMENT_SHADER
)

const (
	TEXTURE_BASE_LEVEL   TexParam = gl_TEXTURE_BASE_LEVEL
	TEXTURE_BORDER_COLOR TexParam = gl_TEXTURE_BORDER_COLOR
	TEXTURE_COMPARE_FUNC TexParam = gl_TEXTURE_COMPARE_FUNC
	TEXTURE_COMPARE_MODE TexParam = gl_TEXTURE_COMPARE_MODE
	TEXTURE_LOD_BIAS     TexParam = gl_TEXTURE_LOD_BIAS
	TEXTURE_MAG_FILTER   TexParam = gl_TEXTURE_MAG_FILTER
	TEXTURE_MAX_LEVEL    TexParam = gl_TEXTURE_MAX_LEVEL
	TEXTURE_MAX_LOD      TexParam = gl_TEXTURE_MAX_LOD
	TEXTURE_MIN_FILTER   TexParam = gl_TEXTURE_MIN_FILTER
	TEXTURE_MIN_LOD      TexParam = gl_TEXTURE_MIN_LOD
	TEXTURE_WRAP_R       TexParam = gl_TEXTURE_WRAP_R
	TEXTURE_WRAP_S       TexParam = gl_TEXTURE_WRAP_S
	TEXTURE_WRAP_T       TexParam = gl_TEXTURE_WRAP_T
)

var (
	errNoCompiler = errors.New("Shader compilation not supported on this platform")
)

// Object Creation

func GenTextures(n int) []Texture {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Texture, n)
	C.goglGenTextures(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return buf
}

func GenBuffers(n int) []Buffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Buffer, n)
	C.goglGenBuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Buffer(buf)
}

func GenFramebuffers(n int) []Framebuffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Framebuffer, n)
	C.goglGenFramebuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Framebuffer(buf)
}

func GenRenderbuffers(n int) []Renderbuffer {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]Renderbuffer, n)
	C.goglGenRenderbuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []Renderbuffer(buf)
}

func GenVertexArrays(n int) []VertexArray {
	if n < 0 {
		panic("Cannot pass negative value to Gen* functions")
	}
	buf := make([]VertexArray, n)
	C.goglGenVertexArrays(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	return []VertexArray(buf)
}

func CreateShader(t ShaderType) Shader {
	x := C.goglCreateShader(C.GLenum(t))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Invalid shader type %d", t))
	}
	return Shader(x)
}

func CreateProgram() Program {
	x := C.goglCreateProgram()
	if x == 0 {
		panic("Could not create program object")
	}
	return Program(x)
}

// Object deletion

// DeleteBuffers deletes several buffers at once.
func DeleteBuffers(b []Buffer) {
	if b != nil {
		C.goglDeleteBuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func DeleteFramebuffers(b []Framebuffer) {
	if b != nil {
		C.goglDeleteFramebuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func DeleteVertexArrays(b []VertexArray) {
	if b != nil {
		C.goglDeleteVertexArrays(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func DeleteRenderbuffers(b []Renderbuffer) {
	if b != nil {
		C.goglDeleteRenderbuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func DeleteTextures(b []Texture) {
	if b == nil {
		C.goglDeleteTextures(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
	}
}

func DeleteShader(s Shader) {
	C.goglDeleteShader(C.GLuint(s))
}

func DeleteProgram(p Program) {
	C.goglDeleteProgram(C.GLuint(p))
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

// Framebuffers

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

// Textures

func TexImage(tgt Target, lvl int, internal TexFormat, format TexFormat, data interface{}, shape ...int) error {
	typ, _, count, ptr := sliceInfo(data)
	{
		internal := C.GLint(internal)
		tgt := C.GLenum(tgt)
		lvl := C.GLint(lvl)
		format := C.GLenum(format)
		typ := C.GLenum(typ)
		
		switch len(shape) {
		case 0:
			shape = []int{count}
			fallthrough
		case 1:
			width := C.GLsizei(shape[0])
			C.goglTexImage1D(tgt, lvl, internal, width, 0, format, typ, ptr)
		case 2:
			width, height := C.GLsizei(shape[0]), C.GLsizei(shape[1])
			C.goglTexImage2D(tgt, lvl, internal, width, height, 0, format, typ, ptr)
		case 3:
			width, height, depth := C.GLsizei(shape[0]), C.GLsizei(shape[1]), C.GLsizei(shape[2])
			C.goglTexImage3D(tgt, lvl, internal, width, height, depth, 0, format, typ, ptr)
		default:
			panic(fmt.Sprintf("Wrong number of dimensions given to TexImage: %d", len(shape)))
		}
	}
	return getError()
}

func TexParameter(tgt Target, param TexParam, val interface{}) {
	switch val := val.(type) {
	case int:
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

// GLSL Programs

// TODO(droyo) func ShaderBinary(bin []byte)
func ShaderSource(s Shader, src []byte) error {
	if C.goglIsShader(C.GLuint(s)) != gl_TRUE {
		return fmt.Errorf("%x is not a shader")
	}
	
	srcp := unsafe.Pointer(&src[0])
	length := []C.GLint{C.GLint(len(src))}
	data := []*C.GLchar{(*C.GLchar)(srcp)}
	
	C.goglShaderSource(C.GLuint(s), 1, &data[0], &length[0])
	code := C.goglGetError()
	switch code {
	case gl_INVALID_OPERATION:
		return errNoCompiler
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&code)
	}
}

func GetShaderSource(s Shader) ([]byte, error) {
	var n C.GLint
	var sz C.GLsizei
	
	if C.goglIsShader(C.GLuint(s)) != gl_TRUE {
		return nil, fmt.Errorf("%x is not a Shader", s)
	}
	C.goglGetShaderiv(C.GLuint(s), gl_SHADER_SOURCE_LENGTH, &n)
	code := C.goglGetError()
	switch code {
	case gl_NO_ERROR:
		break
	case gl_INVALID_OPERATION:
		return nil, errNoCompiler
	default:
		return nil, (*Error)(&code)
	}
	src := make([]byte, int(n))
	srcp := (*C.GLchar)(unsafe.Pointer(&src[0]))
	C.goglGetShaderSource(C.GLuint(s), C.GLsizei(n), &sz, srcp)
	if err := getError(); err != nil {
		return nil, err
	}
	src = src[:int(sz)]
	return src, nil
}

func (s Shader) String() string {
	switch s.Type() {
	case VERTEX_SHADER:
		return fmt.Sprintf("<vert shader %x>", uint32(s))
	case FRAGMENT_SHADER:
		return fmt.Sprintf("<frag shader %x>", uint32(s))
	default:
		return fmt.Sprintf("<invalid shader %x>", uint32(s))
	}
}

func (s Shader) Type() ShaderType {
	var typ C.GLint
	C.goglGetShaderiv(C.GLuint(s), gl_SHADER_TYPE, &typ)
	if C.goglGetError() != gl_NO_ERROR {
		return 0
	}
	return ShaderType(typ)
}

func (p Program) String() string {
	if p == 0 || C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Sprintf("<invalid program %x>", uint32(p))
	}
	return fmt.Sprintf("<program %x>", uint32(p))
}

func AttachShader(p Program, s Shader) error {
	if C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Errorf("%x is not a Program", p)
	}
	if C.goglIsShader(C.GLuint(s)) != gl_TRUE {
		return fmt.Errorf("%x is not a Shader", s)
	}
	C.goglAttachShader(C.GLuint(p), C.GLuint(s))
	code := C.goglGetError()
	switch code {
	case gl_NO_ERROR:
		return nil
	case gl_INVALID_OPERATION:
		return fmt.Errorf("%v or Shader of same type already attached to %v",s,p)
	default:
		return (*Error)(&code)
	}
}

func DetachShader(p Program, s Shader) error {
	if C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Errorf("%x is not a Program", p)
	}
	if C.goglIsShader(C.GLuint(s)) != gl_TRUE {
		return fmt.Errorf("%x is not a Shader", s)
	}
	C.goglDetachShader(C.GLuint(p), C.GLuint(s))
	code := C.goglGetError()
	switch code {
	case gl_NO_ERROR:
		return nil
	case gl_INVALID_OPERATION:
		return fmt.Errorf("%v not attached to %v", s, p)
	default:
		return (*Error)(&code)
	}
}

func CompileShader(s Shader) error {
	var n,status C.GLint
	var sz C.GLsizei
	
	C.goglCompileShader(C.GLuint(s))
	code := C.goglGetError()
	switch code {
	case gl_INVALID_OPERATION:
		return errNoCompiler
	case gl_NO_ERROR:
		break
	case gl_INVALID_VALUE:
		return fmt.Errorf("%x is not a Shader", s)
	}
	C.goglGetShaderiv(C.GLuint(s), gl_COMPILE_STATUS, &status)
	if status != gl_TRUE {
		C.goglGetShaderiv(C.GLuint(s), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetShaderInfoLog(C.GLuint(s), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

func LinkProgram(p Program) error {
	var n,status C.GLint
	var sz C.GLsizei
	
	C.goglLinkProgram(C.GLuint(p))
	if err := getError(); err != nil {
		return err
	}
	C.goglGetProgramiv(C.GLuint(p), gl_LINK_STATUS, &status)
	if status != gl_TRUE {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

func ValidateProgram(p Program) error {
	var n,status C.GLint
	var sz C.GLsizei
	
	C.goglValidateProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglGetProgramiv(C.GLuint(p), gl_VALIDATE_STATUS, &status)
	if status != gl_TRUE {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

func UseProgram(p Program) error {
	if C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglUseProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("Could not make %v part of current rendering state", p)
	}
	return nil
}

func GetUniformLocation(p Program, name string) (Uniform, error) {
	q := C.CString(name)
	defer C.free(unsafe.Pointer(q))
	
	u := C.goglGetUniformLocation(C.GLuint(p), (*C.GLchar)(q))
	if u < 0 {
		return 0, getError()
	}
	return Uniform(u), nil
}

func GetAttribLocation(p Program, name string) (Attrib, error) {
	q := C.CString(name)
	defer C.free(unsafe.Pointer(q))
	
	a := C.goglGetAttribLocation(C.GLuint(p), (*C.GLchar)(q))
	if a < 0 {
		return 0, getError()
	}
	return Attrib(a), nil
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

func Uniformi(u Uniform, v ...int) {
	switch len(v) {
	case 4:
		C.goglUniform4i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]), C.GLint(v[2]), C.GLint(v[3]))
	case 3:
		C.goglUniform3i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]), C.GLint(v[2]))
	case 2:
		C.goglUniform2i(C.GLint(u), C.GLint(v[0]), C.GLint(v[1]))
	case 1:
		C.goglUniform1i(C.GLint(u), C.GLint(v[0]))
	default:
		panic("Uniforms must have 1-4 values")
	}
}

func Uniformf(u Uniform, v ...float32) {
	switch len(v) {
	case 4:
		C.goglUniform4f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]), C.GLfloat(v[2]), C.GLfloat(v[3]))
	case 3:
		C.goglUniform3f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]), C.GLfloat(v[2]))
	case 2:
		C.goglUniform2f(C.GLint(u), C.GLfloat(v[0]), C.GLfloat(v[1]))
	case 1:
		C.goglUniform1f(C.GLint(u), C.GLfloat(v[0]))
	default:
		panic("Uniforms must have 1-4 values")
	}
}
