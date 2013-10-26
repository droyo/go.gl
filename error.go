package gl

/*
#include "gogl.h"
*/
import "C"
import (
	"fmt"
)

type Error C.GLenum
func (e *Error) Error() string {
	if e != nil {
		switch *e {
		case gl_NO_ERROR:
		case gl_OUT_OF_MEMORY:
			return "Out of memory"
		case gl_INVALID_OPERATION:
			return "Invalid operation"
		case gl_INVALID_ENUM:
			return "Invalid enum"
		case gl_INVALID_VALUE:
			return "Invalid value"
		case gl_INVALID_FRAMEBUFFER_OPERATION:
			return "Invalid FrameBuffer operation"
		default:
			return fmt.Sprintf("Unknown error code %d", *e)
		}
	}
	return "No error"
}

func getError() error {
	if e := C.goglGetError(); e != gl_NO_ERROR {
		return (*Error)(&e)
	}
	return nil
}
