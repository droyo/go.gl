package gl

// TODO(droyo) func ShaderBinary(bin []byte)

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
	// A Uniform is a reference to a "global" variable
	// in a GLSL program. Uniforms can be retrieved with
	// GetUniformLocation and updated with the Uniform*
	// functions
	Uniform      int
	// A Shader is reference to a GLSL program.
	Shader       uint
	// Shaders have a type that describes where they sit in
	// the rendering pipeline.
	ShaderType   uint
	// A GLSL program is a collection of GLSL programs
	// sufficient for rasterizing vertex data. Programs must
	// have at least a vertex and fragment shader, and may
	// contain geometry and tesselation shaders.
	Program      uint
)

const (
	gl_SHADER_SOURCE_LENGTH = 0x8B88
	gl_SHADER_TYPE          = 0x8B4F
	gl_INFO_LOG_LENGTH      = 0x8B84
	gl_LINK_STATUS          = 0x8B82
	gl_VALIDATE_STATUS      = 0x8B83
	gl_COMPILE_STATUS       = 0x8B81
)

const (
	VERTEX_SHADER   ShaderType = 0x8B31
	FRAGMENT_SHADER ShaderType = 0x8B30
	GEOMETRY_SHADER ShaderType = 0x8DD9
)

var errNoCompiler = errors.New("Shader compilation not supported on this platform")

// CreateShader creates a fresh name for a GLSL shader
// of the given type.
func CreateShader(t ShaderType) Shader {
	x := C.goglCreateShader(C.GLenum(t))
	if err := getError(); err != nil {
		panic(fmt.Sprintf("Invalid shader type %d: %s", t, err))
	}
	return Shader(x)
}

// CreateProgram creates a fresh name for a collection of
// GLSL shaders. 
func CreateProgram() Program {
	x := C.goglCreateProgram()
	if x == 0 {
		panic("Could not create program object")
	}
	return Program(x)
}

// DeleteShader deletes a given shader. It is legal to
// delete a shader after linking a program.
func DeleteShader(s Shader) {
	C.goglDeleteShader(C.GLuint(s))
}

// DeleteProgram deletes a GLSL program. Programs may
// not be immediately deleted; see
// https://www.opengl.org/sdk/docs/man3/xhtml/glDeleteProgram.xml
func DeleteProgram(p Program) {
	C.goglDeleteProgram(C.GLuint(p))
}

// ShaderSource uploads the GLSL source code to a shader
// allocated with CreateShader. It does not attempt to
// check the source for errors.
func ShaderSource(s Shader, src []byte) error {
	if C.goglIsShader(C.GLuint(s)) != 1 {
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

// GetShaderSource retrieves the source code for an
// existing GLSL shader program.
func GetShaderSource(s Shader) ([]byte, error) {
	var n C.GLint
	var sz C.GLsizei
	
	if C.goglIsShader(C.GLuint(s)) != 1 {
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

func (s Shader) Type() ShaderType {
	var typ C.GLint
	C.goglGetShaderiv(C.GLuint(s), gl_SHADER_TYPE, &typ)
	if C.goglGetError() != gl_NO_ERROR {
		return 0
	}
	return ShaderType(typ)
}

// AttachShader binds a shader to a program. The shader
// should be initialized with ShaderSource or ShaderBinary
// before attaching.
func AttachShader(p Program, s Shader) error {
	if C.goglIsProgram(C.GLuint(p)) != 1 {
		return fmt.Errorf("%x is not a Program", p)
	}
	if C.goglIsShader(C.GLuint(s)) != 1 {
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

// DetachShader detaches a shader from a Program; once a
// Program has been linked with LinkProgram, it is independent
// of its Shader objects and will continue to function if a
// shader is detached.
func DetachShader(p Program, s Shader) error {
	if C.goglIsProgram(C.GLuint(p)) != 1 {
		return fmt.Errorf("%x is not a Program", p)
	}
	if C.goglIsShader(C.GLuint(s)) != 1 {
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

// CompileShader compiles a shader program to a binary format.
// If the shader source is invalid, the content of the shader's
// information log is returned as an error.
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
	if status != 1 {
		C.goglGetShaderiv(C.GLuint(s), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetShaderInfoLog(C.GLuint(s), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

// LinkProgram creates an executable to be run on the vertex
// processor. The Program must have compiled shader objects
// attached to it. Programs must be linked before calling UseProgram
// and any subsequent drawing calls.
func LinkProgram(p Program) error {
	var n,status C.GLint
	var sz C.GLsizei
	
	C.goglLinkProgram(C.GLuint(p))
	if err := getError(); err != nil {
		return err
	}
	C.goglGetProgramiv(C.GLuint(p), gl_LINK_STATUS, &status)
	if status != 1 {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

// ValidateProgram checks to see whether the executables contained in
// the Program can execute given the current OpenGL state. If an error
// occurs, the contents of the program information log are returned as
// an error.
func ValidateProgram(p Program) error {
	var n,status C.GLint
	var sz C.GLsizei
	
	C.goglValidateProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglGetProgramiv(C.GLuint(p), gl_VALIDATE_STATUS, &status)
	if status != 1 {
		C.goglGetProgramiv(C.GLuint(p), gl_INFO_LOG_LENGTH, &n)
		buf := make([]C.GLchar, int(n))
		C.goglGetProgramInfoLog(C.GLuint(p), C.GLsizei(n), (*C.GLsizei)(&sz), &buf[0])
		return errors.New(C.GoStringN((*C.char)(&buf[0]), C.int(sz)))
	}
	return nil
}

// UseProgram installs the Program as part of the current rendering state.
// only one program may be used at a time.
func UseProgram(p Program) error {
	if C.goglIsProgram(C.GLuint(p)) != 1 {
		return fmt.Errorf("%x is not a Program", p)
	}
	C.goglUseProgram(C.GLuint(p))
	if C.goglGetError() != gl_NO_ERROR {
		return fmt.Errorf("Could not make %v part of current rendering state", p)
	}
	return nil
}

// GetUniformLocation retrieves a reference to a GLSL uniform by
// matching its name in the GLSL source code. The return value is
// suitable as an argument to Uniformf, UniformMatrixf et. al.
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

func Uniform1iv(u Uniform, v []int) {
	C.goglUniform1iv(C.GLint(u), C.GLsizei(len(v)), (*C.GLint)(unsafe.Pointer(&v[0])))
}

func Uniform2iv(u Uniform, v []int) {
	C.goglUniform2iv(C.GLint(u), C.GLsizei(len(v))/2, (*C.GLint)(unsafe.Pointer(&v[0])))
}

func Uniform3iv(u Uniform, v []int) {
	C.goglUniform3iv(C.GLint(u), C.GLsizei(len(v))/3, (*C.GLint)(unsafe.Pointer(&v[0])))
}

func Uniform4iv(u Uniform, v []int) {
	C.goglUniform4iv(C.GLint(u), C.GLsizei(len(v))/4, (*C.GLint)(unsafe.Pointer(&v[0])))
}

func Uniform1fv(u Uniform, v []float32) {
	C.goglUniform1fv(C.GLint(u), C.GLsizei(len(v)), (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

func Uniform2fv(u Uniform, v []float32) {
	C.goglUniform2fv(C.GLint(u), C.GLsizei(len(v))/2, (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

func Uniform3fv(u Uniform, v []float32) {
	C.goglUniform3fv(C.GLint(u), C.GLsizei(len(v))/3, (*C.GLfloat)(unsafe.Pointer(&v[0])))
}

func Uniform4fv(u Uniform, v []float32) {
	C.goglUniform4fv(C.GLint(u), C.GLsizei(len(v))/4, (*C.GLfloat)(unsafe.Pointer(&v[0])))
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

func UniformMatrix2fv(u Uniform, transpose bool, mat []float32) {
	if len(mat) % (2*2) != 0 {
		panic("UniformMatrix2fv expects slice size to be multiple of 4")
	}
	C.goglUniformMatrix2fv(
		C.GLint(u),
		C.GLsizei(len(mat)/(2*2)),
		glBool(transpose),
		(*C.GLfloat)(&mat[0]))
}

func UniformMatrix3fv(u Uniform, transpose bool, mat []float32) {
	if len(mat) % (3*3) != 0 {
		panic("UniformMatrix2fv expects slice size to be multiple of 9")
	}
	C.goglUniformMatrix3fv(
		C.GLint(u),
		C.GLsizei(len(mat)/(3*3)),
		glBool(transpose),
		(*C.GLfloat)(&mat[0]))
}

func UniformMatrix4fv(u Uniform, transpose bool, mat []float32) {
	if len(mat) % (4*4) != 0 {
		panic("UniformMatrix2fv expects slice size to be multiple of 16")
	}
	C.goglUniformMatrix4fv(
		C.GLint(u),
		C.GLsizei(len(mat)/(4*4)),
		glBool(transpose),
		(*C.GLfloat)(&mat[0]))
}
