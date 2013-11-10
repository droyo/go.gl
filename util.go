package gl

/*
#include "gogl.h"
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
	"reflect"
)

const (
	gl_NO_ERROR                      = 0
	gl_INVALID_VALUE                 = 0x0501
	gl_INVALID_OPERATION             = 0x0502
	gl_INVALID_FRAMEBUFFER_OPERATION = 0x0506
	gl_INVALID_ENUM                  = 0x0500
	gl_OUT_OF_MEMORY                 = 0x0505
)

func glBool(b bool) C.GLboolean {
	if b {
		return 1
	}
	return 0
}

func sliceInfo(v interface{}) (format Type, size uintptr, count int, ptr unsafe.Pointer) {
        val := reflect.ValueOf(v)
        typ := reflect.TypeOf(v)

		switch typ.Kind() {
		case reflect.Array, reflect.Slice:
			switch typ.Elem().Kind() {
			case reflect.Float32:
				format = Float32
			case reflect.Int:
				//NOTE(droyo) consider panicking for unqualified Int types
				format = Int32
			case reflect.Int8:
				format = Int8
			case reflect.Int16:
				format = Int16
			case reflect.Int32:
				format = Int32
			case reflect.Uint:
				format = Uint32
			case reflect.Uint8:
				format = Uint8
			case reflect.Uint16:
				format = Uint16
			case reflect.Uint32:
				format = Uint32
			default:
				panic("Invalid element type " + typ.Elem().String())
			}
			size = typ.Elem().Size()
			count = val.Len()
			if val.Len() > 0 {
				ptr = unsafe.Pointer(val.Index(0).UnsafeAddr())
			} else {
				ptr = nil
			}
		default:
			panic("value must be a slice or array")
		}
		return
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
