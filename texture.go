package gl

/*
#include "gogl.h"
*/
import "C"

import (
	"fmt"
)

type Texture uint32

func GenTextures(n int) []Texture {
	buf := make([]Texture, n)
	C.goglGenBuffers(C.GLsizei(n), (*C.GLuint)(&buf[0]))
	if C.goglGetError() != gl_NO_ERROR {
		panic(fmt.Sprintf("Negative value (%d) to GenBuffers", n))
	}
	return buf
}
