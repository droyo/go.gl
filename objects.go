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
	Buffer       uint32
	Framebuffer  uint32
	Renderbuffer uint32
	Texture      uint32
)

type (
	Type      uint
	Target    uint
	UsageHint uint
	Mode      uint
)

const (
	Points        Mode = gl_POINTS
	LineStrip     Mode = gl_LINE_STRIP
	LineLoop      Mode = gl_LINE_LOOP
	Lines         Mode = gl_LINES
	TriangleStrip Mode = gl_TRIANGLE_STRIP
	TriangleFan   Mode = gl_TRIANGLE_FAN
	Triangles     Mode = gl_TRIANGLES
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
	ArrayBuffer        Target = gl_ARRAY_BUFFER
	ElementArrayBuffer Target = gl_ELEMENT_ARRAY_BUFFER
)

const (
	StreamDraw  UsageHint = 0x88E0 // Modify once, use seldom.
	StaticDraw  UsageHint = 0x88E4 // Modify once, use often.
	DynamicDraw UsageHint = 0x88E8 // Modify often, use often.
)

const (
	VertexShader   ShaderType = gl_VERTEX_SHADER
	FragmentShader ShaderType = gl_FRAGMENT_SHADER
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
	if b == nil {
		panic("nil Buffer slice")
	}
	C.goglDeleteBuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

func DeleteFramebuffers(b []Framebuffer) {
	if b == nil {
		panic("nil Framebuffer slice")
	}
	C.goglDeleteFramebuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

func DeleteRenderbuffers(b []Renderbuffer) {
	if b == nil {
		panic("nil Renderbuffer slice")
	}
	C.goglDeleteRenderbuffers(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
}

func DeleteTextures(b []Texture) {
	if b == nil {
		panic("nil Texture slice")
	}
	C.goglDeleteTextures(C.GLsizei(len(b)), (*C.GLuint)(&b[0]))
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
		return fmt.Errorf("Invalid target %d", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%d is not a Buffer", b)
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
		return fmt.Errorf("Invalid target %d", t)
	case gl_INVALID_VALUE:
		return fmt.Errorf("%d is not a Framebuffer", b)
	case gl_NO_ERROR:
		return nil
	default:
		return (*Error)(&err)
	}
}

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
	case VertexShader:
		return fmt.Sprintf("<vert shader %x>", s)
	case FragmentShader:
		return fmt.Sprintf("<frag shader %x>", s)
	default:
		return fmt.Sprintf("<invalid shader %x>", s)
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
		return fmt.Sprintf("<invalid program %x>", p)
	}
	return fmt.Sprintf("<program %x>", p)
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
