package gl

/*
#include <stdlib.h>
#include "gogl.h"
*/
import "C"

import (
	"fmt"
	"errors"
	"unsafe"
)

var (
	errNoCompiler = errors.New("Shader compilation not supported on this platform")
)

type (
	Program uint32
	Shader     uint32
	ShaderType uint32
	Uniform    uint32
	Attrib uint32
)

const (
	VertexShader ShaderType = gl_VERTEX_SHADER
	FragmentShader ShaderType = gl_FRAGMENT_SHADER
)

func CreateShader(t ShaderType) Shader {
	x := C.goglCreateShader(C.GLenum(t))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Invalid shader type %d", t))
	}
	return Shader(x)
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

func (s Shader) SetSource(src ...string) error {
	if C.goglIsShader(C.GLuint(s)) != gl_TRUE {
		return fmt.Errorf("%x is not a shader")
	}
	if len(src) == 0 {
		panic("At least one argument must be provided to SetSource")
	}
	data := make([]*C.GLchar, len(src))
	length := make([]C.GLint, len(src))
	
	for i := range src {
		data[i] = (*C.GLchar)(C.CString(src[i]))
		length[i] = C.GLint(len(src[i]))
	}
	C.goglShaderSource(
		C.GLuint(s),
		C.GLsizei(len(src)),
		&data[0],
		&length[0])
		
	for i := range data {
		C.free(unsafe.Pointer(data[i]))
	}
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

func (s Shader) Type() ShaderType {
	var typ C.GLint
	C.goglGetShaderiv(C.GLuint(s), gl_SHADER_TYPE, &typ)
	if C.goglGetError() != gl_NO_ERROR {
		return 0
	}
	return ShaderType(typ)
}

func (s Shader) Source() ([]byte, error) {
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

func (s Shader) Compile() error {
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
	if C.GLenum(status) != gl_TRUE {
		C.goglGetShaderiv(C.GLuint(s), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetShaderInfoLog(C.GLuint(s), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

func (s Shader) Delete() {
	C.goglDeleteShader(C.GLuint(s))
}

// TODO(droyo) func (s Shader) SetBinary(src []byte)

func (p Program) String() string {
	if p == 0 || C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Sprintf("<invalid program %x>", p)
	}
	return fmt.Sprintf("<program %x>", p)
}

func CreateProgram() Program {
	x := C.goglCreateProgram()
	if x == 0 {
		panic("Could not create program object")
	}
	return Program(x)
}

func (p Program) Attach(s Shader) error {
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

func (p Program) Detach(s Shader) error {
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

func (p Program) Link() error {
	var n,status C.GLint
	C.goglLinkProgram(C.GLuint(p))
	if err := getError(); err != nil {
		return err
	}
	C.goglGetProgramiv(C.GLuint(p), gl_LINK_STATUS, &status)
	if C.GLenum(status) != gl_TRUE {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		bufp := unsafe.Pointer(&buf[0])
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&n), (*C.GLchar)(bufp))
		return errors.New(C.GoStringN((*C.char)(bufp), C.int(n)))
	}
	return nil
}

func (p Program) Use() error {
	if C.goglIsProgram(C.GLuint(p)) != gl_TRUE {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglUseProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("Could not make %v part of current rendering state", p)
	}
	return nil
}

func (p Program) Delete() {
	C.goglDeleteProgram(C.GLuint(p))
}

func (p Program) Validate() error {
	var n,status C.GLint
	C.goglValidateProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglGetProgramiv(C.GLuint(p), gl_LINK_STATUS, &status)
	if C.GLenum(status) != gl_TRUE {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		bufp := unsafe.Pointer(&buf[0])
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&n), (*C.GLchar)(bufp))
		return errors.New(C.GoStringN((*C.char)(bufp), C.int(n)))
	}
	return nil
}

func (p Program) Uniform(name string) (Uniform, error) {
	q := C.CString(name)
	defer C.free(unsafe.Pointer(q))
	
	u := C.goglGetUniformLocation(C.GLuint(p), (*C.GLchar)(q))
	if u < 0 {
		return 0, getError()
	}
	return Uniform(u), nil
}

func (p Program) Attrib(name string) (Attrib, error) {
	q := C.CString(name)
	defer C.free(unsafe.Pointer(q))
	
	a := C.goglGetAttribLocation(C.GLuint(p), (*C.GLchar)(q))
	if a < 0 {
		return 0, getError()
	}
	return Attrib(a), nil
}

func (a Attrib) Enable() {
	C.goglEnableVertexAttribArray(C.GLuint(a))
}

func (a Attrib) Disable() {
	C.goglDisableVertexAttribArray(C.GLuint(a))
}

func (a Attrib) Pointer(size int, typ Type, normalize bool, stride, offset uintptr) error {
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

func (u Uniform) SetInt(v ...int) {
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

func (u Uniform) SetFloat(v ...float32) {
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
