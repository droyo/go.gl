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
	"image/color"
	"reflect"
)

// func Accum(op glt.Enum, value float32) {
// 	C.goglAccum((C.GLenum)(op), (C.GLfloat)(value))
// }
// func ActiveTexture(texture glt.Enum) {
// 	C.goglActiveTexture((C.GLenum)(texture))
// }
// func AlphaFunc(glfunc glt.Enum, ref float32) {
// 	C.goglAlphaFunc((C.GLenum)(glfunc), (C.GLfloat)(ref))
// }
// func AlphaFuncx(glfunc glt.Enum, ref int32) {
// 	C.goglAlphaFuncx((C.GLenum)(glfunc), (C.GLfixed)(ref))
// }
// func AreTexturesResident(n uint32, textures *uint32, residences *byte) bool {
//	return glt.GLBoolean(C.goglAreTexturesResident((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLboolean)(residences)))
// }
// func ArrayElement(i int32) {
// 	C.goglArrayElement((C.GLint)(i))
// }
// func BindAttribLocation(program uint32, index uint32, name *int8) {
// 	C.goglBindAttribLocation((C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(name))
// }
// func BindTexture(target glt.Enum, texture uint32) {
// 	C.goglBindTexture((C.GLenum)(target), (C.GLuint)(texture))
// }
// func BlendColor(red float32, green float32, blue float32, alpha float32) {
// 	C.goglBlendColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
// }
// func BlendEquation(mode glt.Enum) {
// 	C.goglBlendEquation((C.GLenum)(mode))
// }
// func BlendEquationSeparate(modeRGB glt.Enum, modeAlpha glt.Enum) {
// 	C.goglBlendEquationSeparate((C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
// }
// func BlendFunc(sfactor glt.Enum, dfactor glt.Enum) {
// 	C.goglBlendFunc((C.GLenum)(sfactor), (C.GLenum)(dfactor))
// }
// func BlendFuncSeparate(sfactorRGB glt.Enum, dfactorRGB glt.Enum, sfactorAlpha glt.Enum, dfactorAlpha glt.Enum) {
// 	C.goglBlendFuncSeparate((C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
// }
// func BufferSubData(target glt.Enum, offset int, size int, data glt.Pointer) {
// 	C.goglBufferSubData((C.GLenum)(target), (int)(offset), (int)(size), glt.Pointer(data))
// }
// func CallList(list uint32) {
// 	C.goglCallList((C.GLuint)(list))
// }
// func CallLists(n uint32, gltype glt.Enum, lists glt.Pointer) {
// 	C.goglCallLists((C.GLsizei)(n), (C.GLenum)(gltype), glt.Pointer(lists))
// }
// func CheckFramebufferStatus(target glt.Enum) glt.Enum {
// 	return (C.GLenum)(C.goglCheckFramebufferStatus((C.GLenum)(target)))
// }
// func Clear(mask glt.Bitfieled) {
// 	C.goglClear((C.GLbitfield)(mask))
// }
// func ClearAccum(red float32, green float32, blue float32, alpha float32) {
// 	C.goglClearAccum((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
// }
// func ClearColor(red float32, green float32, blue float32, alpha float32) {
// 	C.goglClearColor((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
// }
// func ClearColorx(red int32, green int32, blue int32, alpha int32) {
// 	C.goglClearColorx((C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
// }
// func ClearDepth(depth float64) {
// 	C.goglClearDepth((C.GLdouble)(depth))
// }
// func ClearDepthf(d float32) {
// 	C.goglClearDepthf((C.GLfloat)(d))
// }
// func ClearDepthx(depth int32) {
// 	C.goglClearDepthx((C.GLfixed)(depth))
// }
// func ClearIndex(c float32) {
// 	C.goglClearIndex((C.GLfloat)(c))
// }
// func ClearStencil(s int32) {
// 	C.goglClearStencil((C.GLint)(s))
// }
// func ClientActiveTexture(texture glt.Enum) {
// 	C.goglClientActiveTexture((C.GLenum)(texture))
// }
// func ClipPlane(plane glt.Enum, equation *float64) {
// 	C.goglClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
// }
// func ClipPlanef(p glt.Enum, eqn *float32) {
// 	C.goglClipPlanef((C.GLenum)(p), (*C.GLfloat)(eqn))
// }
// func ClipPlanex(plane glt.Enum, equation *int32) {
// 	C.goglClipPlanex((C.GLenum)(plane), (*C.GLfixed)(equation))
// }
// func Color3b(red int8, green int8, blue int8) {
// 	C.goglColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
// }
// func Color3bv(v *int8) {
// 	C.goglColor3bv((*C.GLbyte)(v))
// }
// func Color3d(red float64, green float64, blue float64) {
// 	C.goglColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
// }
// func Color3dv(v *float64) {
// 	C.goglColor3dv((*C.GLdouble)(v))
// }
// func Color3f(red float32, green float32, blue float32) {
// 	C.goglColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
// }
// func Color3fv(v *float32) {
// 	C.goglColor3fv((*C.GLfloat)(v))
// }
// func Color3i(red int32, green int32, blue int32) {
// 	C.goglColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
// }
// func Color3iv(v *int32) {
// 	C.goglColor3iv((*C.GLint)(v))
// }
// func Color3s(red int16, green int16, blue int16) {
// 	C.goglColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
// }
// func Color3sv(v *int16) {
// 	C.goglColor3sv((*C.GLshort)(v))
// }
// func Color3ub(red uint8, green uint8, blue uint8) {
// 	C.goglColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
// }
// func Color3ubv(v *uint8) {
// 	C.goglColor3ubv((*C.GLubyte)(v))
// }
// func Color3ui(red uint32, green uint32, blue uint32) {
// 	C.goglColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
// }
// func Color3uiv(v *uint32) {
// 	C.goglColor3uiv((*C.GLuint)(v))
// }
// func Color3us(red uint16, green uint16, blue uint16) {
// 	C.goglColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
// }
// func Color3usv(v *uint16) {
// 	C.goglColor3usv((*C.GLushort)(v))
// }
// func Color4b(red int8, green int8, blue int8, alpha int8) {
// 	C.goglColor4b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue), (C.GLbyte)(alpha))
// }
// func Color4bv(v *int8) {
// 	C.goglColor4bv((*C.GLbyte)(v))
// }
// func Color4d(red float64, green float64, blue float64, alpha float64) {
// 	C.goglColor4d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue), (C.GLdouble)(alpha))
// }
// func Color4dv(v *float64) {
// 	C.goglColor4dv((*C.GLdouble)(v))
// }
// func Color4f(red float32, green float32, blue float32, alpha float32) {
// 	C.goglColor4f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
// }
// func Color4fv(v *float32) {
// 	C.goglColor4fv((*C.GLfloat)(v))
// }
// func Color4i(red int32, green int32, blue int32, alpha int32) {
// 	C.goglColor4i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue), (C.GLint)(alpha))
// }
// func Color4iv(v *int32) {
// 	C.goglColor4iv((*C.GLint)(v))
// }
// func Color4s(red int16, green int16, blue int16, alpha int16) {
// 	C.goglColor4s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue), (C.GLshort)(alpha))
// }
// func Color4sv(v *int16) {
// 	C.goglColor4sv((*C.GLshort)(v))
// }
// func Color4ub(red uint8, green uint8, blue uint8, alpha uint8) {
// 	C.goglColor4ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue), (C.GLubyte)(alpha))
// }
// func Color4ubv(v *uint8) {
// 	C.goglColor4ubv((*C.GLubyte)(v))
// }
// func Color4ui(red uint32, green uint32, blue uint32, alpha uint32) {
// 	C.goglColor4ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue), (C.GLuint)(alpha))
// }
// func Color4uiv(v *uint32) {
// 	C.goglColor4uiv((*C.GLuint)(v))
// }
// func Color4us(red uint16, green uint16, blue uint16, alpha uint16) {
// 	C.goglColor4us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue), (C.GLushort)(alpha))
// }
// func Color4usv(v *uint16) {
// 	C.goglColor4usv((*C.GLushort)(v))
// }
// func Color4x(red int32, green int32, blue int32, alpha int32) {
// 	C.goglColor4x((C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
// }
// func ColorMask(red bool, green bool, blue bool, alpha bool) {
// 	C.goglColorMask(glt.GLBool(red), glt.GLBool(green), glt.GLBool(blue), glt.GLBool(alpha))
// }
// func ColorMaterial(face glt.Enum, mode glt.Enum) {
// 	C.goglColorMaterial((C.GLenum)(face), (C.GLenum)(mode))
// }
// func ColorPointer(size int32, gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglColorPointer((C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func CompileShader(shader uint32) {
// 	C.goglCompileShader((C.GLuint)(shader))
// }
// func CompressedTexImage1D(target glt.Enum, level int32, internalformat glt.Enum, width uint32, border int32, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CompressedTexImage2D(target glt.Enum, level int32, internalformat glt.Enum, width uint32, height uint32, border int32, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CompressedTexImage3D(target glt.Enum, level int32, internalformat glt.Enum, width uint32, height uint32, depth uint32, border int32, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CompressedTexSubImage1D(target glt.Enum, level int32, xoffset int32, width uint32, format glt.Enum, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CompressedTexSubImage2D(target glt.Enum, level int32, xoffset int32, yoffset int32, width uint32, height uint32, format glt.Enum, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CompressedTexSubImage3D(target glt.Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width uint32, height uint32, depth uint32, format glt.Enum, imageSize uint32, data glt.Pointer) {
// 	C.goglCompressedTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), glt.Pointer(data))
// }
// func CopyPixels(x int32, y int32, width uint32, height uint32, gltype glt.Enum) {
// 	C.goglCopyPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(gltype))
// }
// func CopyTexImage1D(target glt.Enum, level int32, internalformat glt.Enum, x int32, y int32, width uint32, border int32) {
// 	C.goglCopyTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
// }
// func CopyTexImage2D(target glt.Enum, level int32, internalformat glt.Enum, x int32, y int32, width uint32, height uint32, border int32) {
// 	C.goglCopyTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
// }
// func CopyTexSubImage1D(target glt.Enum, level int32, xoffset int32, x int32, y int32, width uint32) {
// 	C.goglCopyTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
// }
// func CopyTexSubImage2D(target glt.Enum, level int32, xoffset int32, yoffset int32, x int32, y int32, width uint32, height uint32) {
// 	C.goglCopyTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
// }
// func CopyTexSubImage3D(target glt.Enum, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width uint32, height uint32) {
// 	C.goglCopyTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
// }
// func CullFace(mode glt.Enum) {
// 	C.goglCullFace((C.GLenum)(mode))
// }
// func DeleteLists(list uint32, glrange uint32) {
// 	C.goglDeleteLists((C.GLuint)(list), (C.GLsizei)(glrange))
// }
// func DeleteProgram(program uint32) {
// 	C.goglDeleteProgram((C.GLuint)(program))
// }
// func DeleteQueries(n uint32, ids *uint32) {
// 	C.goglDeleteQueries((C.GLsizei)(n), (*C.GLuint)(ids))
// }
// func DeleteShader(shader uint32) {
// 	C.goglDeleteShader((C.GLuint)(shader))
// }
// func DeleteTextures(n uint32, textures *uint32) {
// 	C.goglDeleteTextures((C.GLsizei)(n), (*C.GLuint)(textures))
// }
// func DepthFunc(glfunc glt.Enum) {
// 	C.goglDepthFunc((C.GLenum)(glfunc))
// }
// func DepthMask(flag bool) {
// 	C.goglDepthMask(glt.GLBool(flag))
// }
// func DepthRange(near float64, far float64) {
// 	C.goglDepthRange((C.GLdouble)(near), (C.GLdouble)(far))
// }
// func DepthRangef(n float32, f float32) {
// 	C.goglDepthRangef((C.GLfloat)(n), (C.GLfloat)(f))
// }
// func DepthRangex(n int32, f int32) {
// 	C.goglDepthRangex((C.GLfixed)(n), (C.GLfixed)(f))
// }
// func DetachShader(program uint32, shader uint32) {
// 	C.goglDetachShader((C.GLuint)(program), (C.GLuint)(shader))
// }
// func Disable(cap glt.Enum) {
// 	C.goglDisable((C.GLenum)(cap))
// }
// func DisableClientState(array glt.Enum) {
// 	C.goglDisableClientState((C.GLenum)(array))
// }
// func DisableVertexAttribArray(index uint32) {
// 	C.goglDisableVertexAttribArray((C.GLuint)(index))
// }
// func DrawArrays(mode glt.Enum, first int32, count uint32) {
// 	C.goglDrawArrays((C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
// }
// func DrawBuffer(mode glt.Enum) {
// 	C.goglDrawBuffer((C.GLenum)(mode))
// }
// func DrawBuffers(n uint32, bufs *glt.Enum) {
// 	C.goglDrawBuffers((C.GLsizei)(n), (*C.GLenum)(bufs))
// }
// func DrawElements(mode glt.Enum, count uint32, gltype glt.Enum, indices glt.Pointer) {
// 	C.goglDrawElements((C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(gltype), glt.Pointer(indices))
// }
// func DrawPixels(width uint32, height uint32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglDrawPixels((C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func DrawRangeElements(mode glt.Enum, start uint32, end uint32, count uint32, gltype glt.Enum, indices glt.Pointer) {
// 	C.goglDrawRangeElements((C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(gltype), glt.Pointer(indices))
// }
// func EdgeFlag(flag bool) {
// 	C.goglEdgeFlag(glt.GLBool(flag))
// }
// func EdgeFlagPointer(stride uint32, pointer glt.Pointer) {
// 	C.goglEdgeFlagPointer((C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func EdgeFlagv(flag *byte) {
// 	C.goglEdgeFlagv((*C.GLboolean)(flag))
// }
// func Enable(cap glt.Enum) {
// 	C.goglEnable((C.GLenum)(cap))
// }
// func EnableClientState(array glt.Enum) {
// 	C.goglEnableClientState((C.GLenum)(array))
// }
// func EnableVertexAttribArray(index uint32) {
// 	C.goglEnableVertexAttribArray((C.GLuint)(index))
// }
// func End() {
// 	C.goglEnd()
// }
// func EndList() {
// 	C.goglEndList()
// }
// func EndQuery(target glt.Enum) {
// 	C.goglEndQuery((C.GLenum)(target))
// }
// func EvalCoord1d(u float64) {
// 	C.goglEvalCoord1d((C.GLdouble)(u))
// }
// func EvalCoord1dv(u *float64) {
// 	C.goglEvalCoord1dv((*C.GLdouble)(u))
// }
// func EvalCoord1f(u float32) {
// 	C.goglEvalCoord1f((C.GLfloat)(u))
// }
// func EvalCoord1fv(u *float32) {
// 	C.goglEvalCoord1fv((*C.GLfloat)(u))
// }
// func EvalCoord2d(u float64, v float64) {
// 	C.goglEvalCoord2d((C.GLdouble)(u), (C.GLdouble)(v))
// }
// func EvalCoord2dv(u *float64) {
// 	C.goglEvalCoord2dv((*C.GLdouble)(u))
// }
// func EvalCoord2f(u float32, v float32) {
// 	C.goglEvalCoord2f((C.GLfloat)(u), (C.GLfloat)(v))
// }
// func EvalCoord2fv(u *float32) {
// 	C.goglEvalCoord2fv((*C.GLfloat)(u))
// }
// func EvalMesh1(mode glt.Enum, i1 int32, i2 int32) {
// 	C.goglEvalMesh1((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2))
// }
// func EvalMesh2(mode glt.Enum, i1 int32, i2 int32, j1 int32, j2 int32) {
// 	C.goglEvalMesh2((C.GLenum)(mode), (C.GLint)(i1), (C.GLint)(i2), (C.GLint)(j1), (C.GLint)(j2))
// }
// func EvalPoint1(i int32) {
// 	C.goglEvalPoint1((C.GLint)(i))
// }
// func EvalPoint2(i int32, j int32) {
// 	C.goglEvalPoint2((C.GLint)(i), (C.GLint)(j))
// }
// func FeedbackBuffer(size uint32, gltype glt.Enum, buffer *float32) {
// 	C.goglFeedbackBuffer((C.GLsizei)(size), (C.GLenum)(gltype), (*C.GLfloat)(buffer))
// }
// func Finish() {
// 	C.goglFinish()
// }
// func Flush() {
// 	C.goglFlush()
// }
// func FogCoordPointer(gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglFogCoordPointer((C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func FogCoordd(coord float64) {
// 	C.goglFogCoordd((C.GLdouble)(coord))
// }
// func FogCoorddv(coord *float64) {
// 	C.goglFogCoorddv((*C.GLdouble)(coord))
// }
// func FogCoordf(coord float32) {
// 	C.goglFogCoordf((C.GLfloat)(coord))
// }
// func FogCoordfv(coord *float32) {
// 	C.goglFogCoordfv((*C.GLfloat)(coord))
// }
// func Fogf(pname glt.Enum, param float32) {
// 	C.goglFogf((C.GLenum)(pname), (C.GLfloat)(param))
// }
// func Fogfv(pname glt.Enum, params *float32) {
// 	C.goglFogfv((C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func Fogi(pname glt.Enum, param int32) {
// 	C.goglFogi((C.GLenum)(pname), (C.GLint)(param))
// }
// func Fogiv(pname glt.Enum, params *int32) {
// 	C.goglFogiv((C.GLenum)(pname), (*C.GLint)(params))
// }
// func Fogx(pname glt.Enum, param int32) {
// 	C.goglFogx((C.GLenum)(pname), (C.GLfixed)(param))
// }
// func Fogxv(pname glt.Enum, param *int32) {
// 	C.goglFogxv((C.GLenum)(pname), (*C.GLfixed)(param))
// }
// func FramebufferRenderbuffer(target glt.Enum, attachment glt.Enum, renderbuffertarget glt.Enum, renderbuffer uint32) {
// 	C.goglFramebufferRenderbuffer((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
// }
// func FramebufferTexture2D(target glt.Enum, attachment glt.Enum, textarget glt.Enum, texture uint32, level int32) {
// 	C.goglFramebufferTexture2D((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
// }
// func FrontFace(mode glt.Enum) {
// 	C.goglFrontFace((C.GLenum)(mode))
// }
// func Frustum(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
// 	C.goglFrustum((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
// }
// func Frustumf(l float32, r float32, b float32, t float32, n float32, f float32) {
// 	C.goglFrustumf((C.GLfloat)(l), (C.GLfloat)(r), (C.GLfloat)(b), (C.GLfloat)(t), (C.GLfloat)(n), (C.GLfloat)(f))
// }
// func Frustumx(l int32, r int32, b int32, t int32, n int32, f int32) {
// 	C.goglFrustumx((C.GLfixed)(l), (C.GLfixed)(r), (C.GLfixed)(b), (C.GLfixed)(t), (C.GLfixed)(n), (C.GLfixed)(f))
// }
// func GenLists(glrange uint32) uint32 {
// 	return (C.GLuint)(C.goglGenLists((C.GLsizei)(glrange)))
// }
// func GenQueries(n uint32, ids *uint32) {
// 	C.goglGenQueries((C.GLsizei)(n), (*C.GLuint)(ids))
// }
// func GenerateMipmap(target glt.Enum) {
// 	C.goglGenerateMipmap((C.GLenum)(target))
// }
// func GetActiveAttrib(program uint32, index uint32, bufSize uint32, length *uint32, size *int32, gltype *glt.Enum, name *int8) {
// 	C.goglGetActiveAttrib((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(gltype), (*C.GLchar)(name))
// }
// func GetActiveUniform(program uint32, index uint32, bufSize uint32, length *uint32, size *int32, gltype *glt.Enum, name *int8) {
// 	C.goglGetActiveUniform((C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLint)(size), (*C.GLenum)(gltype), (*C.GLchar)(name))
// }
// func GetAttachedShaders(program uint32, maxCount uint32, count *uint32, shaders *uint32) {
// 	C.goglGetAttachedShaders((C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(count), (*C.GLuint)(shaders))
// }
// func GetAttribLocation(program uint32, name *int8) int32 {
// 	return (C.GLint)(C.goglGetAttribLocation((C.GLuint)(program), (*C.GLchar)(name)))
// }
// func GetBooleanv(pname glt.Enum, params *byte) {
// 	C.goglGetBooleanv((C.GLenum)(pname), (*C.GLboolean)(params))
// }
// func GetBufferParameteriv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetBufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetBufferPointerv(target glt.Enum, pname glt.Enum, params *glt.Pointer) {
// 	C.goglGetBufferPointerv((C.GLenum)(target), (C.GLenum)(pname), glt.Pointer(params))
// }
// func GetBufferSubData(target glt.Enum, offset int, size int, data glt.Pointer) {
// 	C.goglGetBufferSubData((C.GLenum)(target), (int)(offset), (int)(size), glt.Pointer(data))
// }
// func GetClipPlane(plane glt.Enum, equation *float64) {
// 	C.goglGetClipPlane((C.GLenum)(plane), (*C.GLdouble)(equation))
// }
// func GetClipPlanef(plane glt.Enum, equation *float32) {
// 	C.goglGetClipPlanef((C.GLenum)(plane), (*C.GLfloat)(equation))
// }
// func GetClipPlanex(plane glt.Enum, equation *int32) {
// 	C.goglGetClipPlanex((C.GLenum)(plane), (*C.GLfixed)(equation))
// }
// func GetCompressedTexImage(target glt.Enum, level int32, img glt.Pointer) {
// 	C.goglGetCompressedTexImage((C.GLenum)(target), (C.GLint)(level), glt.Pointer(img))
// }
// func GetDoublev(pname glt.Enum, params *float64) {
// 	C.goglGetDoublev((C.GLenum)(pname), (*C.GLdouble)(params))
// }
// func GetError() glt.Enum {
// 	return (C.GLenum)(C.goglGetError())
// }
// func GetFixedv(pname glt.Enum, params *int32) {
// 	C.goglGetFixedv((C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func GetFloatv(pname glt.Enum, params *float32) {
// 	C.goglGetFloatv((C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetFramebufferAttachmentParameteriv(target glt.Enum, attachment glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetFramebufferAttachmentParameteriv((C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetIntegerv(pname glt.Enum, params *int32) {
// 	C.goglGetIntegerv((C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetLightfv(light glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglGetLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetLightiv(light glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetLightxv(light glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetLightxv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func GetMapdv(target glt.Enum, query glt.Enum, v *float64) {
// 	C.goglGetMapdv((C.GLenum)(target), (C.GLenum)(query), (*C.GLdouble)(v))
// }
// func GetMapfv(target glt.Enum, query glt.Enum, v *float32) {
// 	C.goglGetMapfv((C.GLenum)(target), (C.GLenum)(query), (*C.GLfloat)(v))
// }
// func GetMapiv(target glt.Enum, query glt.Enum, v *int32) {
// 	C.goglGetMapiv((C.GLenum)(target), (C.GLenum)(query), (*C.GLint)(v))
// }
// func GetMaterialfv(face glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglGetMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetMaterialiv(face glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetMaterialxv(face glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetMaterialxv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func GetPixelMapfv(glmap glt.Enum, values *float32) {
// 	C.goglGetPixelMapfv((C.GLenum)(glmap), (*C.GLfloat)(values))
// }
// func GetPixelMapuiv(glmap glt.Enum, values *uint32) {
// 	C.goglGetPixelMapuiv((C.GLenum)(glmap), (*C.GLuint)(values))
// }
// func GetPixelMapusv(glmap glt.Enum, values *uint16) {
// 	C.goglGetPixelMapusv((C.GLenum)(glmap), (*C.GLushort)(values))
// }
// func GetPointerv(pname glt.Enum, params *glt.Pointer) {
// 	C.goglGetPointerv((C.GLenum)(pname), glt.Pointer(params))
// }
// func GetPolygonStipple(mask *uint8) {
// 	C.goglGetPolygonStipple((*C.GLubyte)(mask))
// }
// func GetProgramInfoLog(program uint32, bufSize uint32, length *uint32, infoLog *int8) {
// 	C.goglGetProgramInfoLog((C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
// }
// func GetProgramiv(program uint32, pname glt.Enum, params *int32) {
// 	C.goglGetProgramiv((C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetQueryObjectiv(id uint32, pname glt.Enum, params *int32) {
// 	C.goglGetQueryObjectiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetQueryObjectuiv(id uint32, pname glt.Enum, params *uint32) {
// 	C.goglGetQueryObjectuiv((C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(params))
// }
// func GetQueryiv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetQueryiv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetRenderbufferParameteriv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetRenderbufferParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetShaderInfoLog(shader uint32, bufSize uint32, length *uint32, infoLog *int8) {
// 	C.goglGetShaderInfoLog((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(infoLog))
// }
// func GetShaderPrecisionFormat(shadertype glt.Enum, precisiontype glt.Enum, glrange *int32, precision *int32) {
// 	C.goglGetShaderPrecisionFormat((C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(glrange), (*C.GLint)(precision))
// }
// func GetShaderSource(shader uint32, bufSize uint32, length *uint32, source *int8) {
// 	C.goglGetShaderSource((C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(length), (*C.GLchar)(source))
// }
// func GetShaderiv(shader uint32, pname glt.Enum, params *int32) {
// 	C.goglGetShaderiv((C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetString(name glt.Enum) *uint8 {
// 	return (*C.GLubyte)(C.goglGetString((C.GLenum)(name)))
// }
// func GetTexEnvfv(target glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglGetTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetTexEnviv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetTexEnvxv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetTexEnvxv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func GetTexGendv(coord glt.Enum, pname glt.Enum, params *float64) {
// 	C.goglGetTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
// }
// func GetTexGenfv(coord glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglGetTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetTexGeniv(coord glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetTexImage(target glt.Enum, level int32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglGetTexImage((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func GetTexLevelParameterfv(target glt.Enum, level int32, pname glt.Enum, params *float32) {
// 	C.goglGetTexLevelParameterfv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetTexLevelParameteriv(target glt.Enum, level int32, pname glt.Enum, params *int32) {
// 	C.goglGetTexLevelParameteriv((C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetTexParameterfv(target glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglGetTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetTexParameteriv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func GetTexParameterxv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglGetTexParameterxv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func GetUniformLocation(program uint32, name *int8) int32 {
// 	return (C.GLint)(C.goglGetUniformLocation((C.GLuint)(program), (*C.GLchar)(name)))
// }
// func GetUniformfv(program uint32, location int32, params *float32) {
// 	C.goglGetUniformfv((C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(params))
// }
// func GetUniformiv(program uint32, location int32, params *int32) {
// 	C.goglGetUniformiv((C.GLuint)(program), (C.GLint)(location), (*C.GLint)(params))
// }
// func GetVertexAttribPointerv(index uint32, pname glt.Enum, pointer *glt.Pointer) {
// 	C.goglGetVertexAttribPointerv((C.GLuint)(index), (C.GLenum)(pname), glt.Pointer(pointer))
// }
// func GetVertexAttribdv(index uint32, pname glt.Enum, params *float64) {
// 	C.goglGetVertexAttribdv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(params))
// }
// func GetVertexAttribfv(index uint32, pname glt.Enum, params *float32) {
// 	C.goglGetVertexAttribfv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func GetVertexAttribiv(index uint32, pname glt.Enum, params *int32) {
// 	C.goglGetVertexAttribiv((C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func Hint(target glt.Enum, mode glt.Enum) {
// 	C.goglHint((C.GLenum)(target), (C.GLenum)(mode))
// }
// func IndexMask(mask uint32) {
// 	C.goglIndexMask((C.GLuint)(mask))
// }
// func IndexPointer(gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglIndexPointer((C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func Indexd(c float64) {
// 	C.goglIndexd((C.GLdouble)(c))
// }
// func Indexdv(c *float64) {
// 	C.goglIndexdv((*C.GLdouble)(c))
// }
// func Indexf(c float32) {
// 	C.goglIndexf((C.GLfloat)(c))
// }
// func Indexfv(c *float32) {
// 	C.goglIndexfv((*C.GLfloat)(c))
// }
// func Indexi(c int32) {
// 	C.goglIndexi((C.GLint)(c))
// }
// func Indexiv(c *int32) {
// 	C.goglIndexiv((*C.GLint)(c))
// }
// func Indexs(c int16) {
// 	C.goglIndexs((C.GLshort)(c))
// }
// func Indexsv(c *int16) {
// 	C.goglIndexsv((*C.GLshort)(c))
// }
// func Indexub(c uint8) {
// 	C.goglIndexub((C.GLubyte)(c))
// }
// func Indexubv(c *uint8) {
// 	C.goglIndexubv((*C.GLubyte)(c))
// }
// func InitNames() {
// 	C.goglInitNames()
// }
// func InterleavedArrays(format glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglInterleavedArrays((C.GLenum)(format), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func IsBuffer(buffer uint32) bool {
// 	return glt.GLBoolean(C.goglIsBuffer((C.GLuint)(buffer)))
// }
// func IsEnabled(cap glt.Enum) bool {
// 	return glt.GLBoolean(C.goglIsEnabled((C.GLenum)(cap)))
// }
// func IsFramebuffer(framebuffer uint32) bool {
// 	return glt.GLBoolean(C.goglIsFramebuffer((C.GLuint)(framebuffer)))
// }
// func IsList(list uint32) bool {
// 	return glt.GLBoolean(C.goglIsList((C.GLuint)(list)))
// }
// func IsProgram(program uint32) bool {
// 	return glt.GLBoolean(C.goglIsProgram((C.GLuint)(program)))
// }
// func IsQuery(id uint32) bool {
// 	return glt.GLBoolean(C.goglIsQuery((C.GLuint)(id)))
// }
// func IsRenderbuffer(renderbuffer uint32) bool {
// 	return glt.GLBoolean(C.goglIsRenderbuffer((C.GLuint)(renderbuffer)))
// }
// func IsShader(shader uint32) bool {
// 	return glt.GLBoolean(C.goglIsShader((C.GLuint)(shader)))
// }
// func IsTexture(texture uint32) bool {
// 	return glt.GLBoolean(C.goglIsTexture((C.GLuint)(texture)))
// }
// func LightModelf(pname glt.Enum, param float32) {
// 	C.goglLightModelf((C.GLenum)(pname), (C.GLfloat)(param))
// }
// func LightModelfv(pname glt.Enum, params *float32) {
// 	C.goglLightModelfv((C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func LightModeli(pname glt.Enum, param int32) {
// 	C.goglLightModeli((C.GLenum)(pname), (C.GLint)(param))
// }
// func LightModeliv(pname glt.Enum, params *int32) {
// 	C.goglLightModeliv((C.GLenum)(pname), (*C.GLint)(params))
// }
// func LightModelx(pname glt.Enum, param int32) {
// 	C.goglLightModelx((C.GLenum)(pname), (C.GLfixed)(param))
// }
// func LightModelxv(pname glt.Enum, param *int32) {
// 	C.goglLightModelxv((C.GLenum)(pname), (*C.GLfixed)(param))
// }
// func Lightf(light glt.Enum, pname glt.Enum, param float32) {
// 	C.goglLightf((C.GLenum)(light), (C.GLenum)(pname), (C.GLfloat)(param))
// }
// func Lightfv(light glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglLightfv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func Lighti(light glt.Enum, pname glt.Enum, param int32) {
// 	C.goglLighti((C.GLenum)(light), (C.GLenum)(pname), (C.GLint)(param))
// }
// func Lightiv(light glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglLightiv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func Lightx(light glt.Enum, pname glt.Enum, param int32) {
// 	C.goglLightx((C.GLenum)(light), (C.GLenum)(pname), (C.GLfixed)(param))
// }
// func Lightxv(light glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglLightxv((C.GLenum)(light), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func LineStipple(factor int32, pattern uint16) {
// 	C.goglLineStipple((C.GLint)(factor), (C.GLushort)(pattern))
// }
// func LineWidth(width float32) {
// 	C.goglLineWidth((C.GLfloat)(width))
// }
// func LineWidthx(width int32) {
// 	C.goglLineWidthx((C.GLfixed)(width))
// }
// func LinkProgram(program uint32) {
// 	C.goglLinkProgram((C.GLuint)(program))
// }
// func ListBase(base uint32) {
// 	C.goglListBase((C.GLuint)(base))
// }
// func LoadIdentity() {
// 	C.goglLoadIdentity()
// }
// func LoadMatrixd(m *float64) {
// 	C.goglLoadMatrixd((*C.GLdouble)(m))
// }
// func LoadMatrixf(m *float32) {
// 	C.goglLoadMatrixf((*C.GLfloat)(m))
// }
// func LoadMatrixx(m *int32) {
// 	C.goglLoadMatrixx((*C.GLfixed)(m))
// }
// func LoadName(name uint32) {
// 	C.goglLoadName((C.GLuint)(name))
// }
// func LoadTransposeMatrixd(m *float64) {
// 	C.goglLoadTransposeMatrixd((*C.GLdouble)(m))
// }
// func LoadTransposeMatrixf(m *float32) {
// 	C.goglLoadTransposeMatrixf((*C.GLfloat)(m))
// }
// func LogicOp(opcode glt.Enum) {
// 	C.goglLogicOp((C.GLenum)(opcode))
// }
// func Map1d(target glt.Enum, u1 float64, u2 float64, stride int32, order int32, points *float64) {
// 	C.goglMap1d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLdouble)(points))
// }
// func Map1f(target glt.Enum, u1 float32, u2 float32, stride int32, order int32, points *float32) {
// 	C.goglMap1f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(stride), (C.GLint)(order), (*C.GLfloat)(points))
// }
// func Map2d(target glt.Enum, u1 float64, u2 float64, ustride int32, uorder int32, v1 float64, v2 float64, vstride int32, vorder int32, points *float64) {
// 	C.goglMap2d((C.GLenum)(target), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLdouble)(points))
// }
// func Map2f(target glt.Enum, u1 float32, u2 float32, ustride int32, uorder int32, v1 float32, v2 float32, vstride int32, vorder int32, points *float32) {
// 	C.goglMap2f((C.GLenum)(target), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (*C.GLfloat)(points))
// }
// func MapBuffer(target glt.Enum, access glt.Enum) glt.Pointer {
// 	return glt.Pointer(C.goglMapBuffer((C.GLenum)(target), (C.GLenum)(access)))
// }
// func MapGrid1d(un int32, u1 float64, u2 float64) {
// 	C.goglMapGrid1d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2))
// }
// func MapGrid1f(un int32, u1 float32, u2 float32) {
// 	C.goglMapGrid1f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2))
// }
// func MapGrid2d(un int32, u1 float64, u2 float64, vn int32, v1 float64, v2 float64) {
// 	C.goglMapGrid2d((C.GLint)(un), (C.GLdouble)(u1), (C.GLdouble)(u2), (C.GLint)(vn), (C.GLdouble)(v1), (C.GLdouble)(v2))
// }
// func MapGrid2f(un int32, u1 float32, u2 float32, vn int32, v1 float32, v2 float32) {
// 	C.goglMapGrid2f((C.GLint)(un), (C.GLfloat)(u1), (C.GLfloat)(u2), (C.GLint)(vn), (C.GLfloat)(v1), (C.GLfloat)(v2))
// }
// func Materialf(face glt.Enum, pname glt.Enum, param float32) {
// 	C.goglMaterialf((C.GLenum)(face), (C.GLenum)(pname), (C.GLfloat)(param))
// }
// func Materialfv(face glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglMaterialfv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func Materiali(face glt.Enum, pname glt.Enum, param int32) {
// 	C.goglMateriali((C.GLenum)(face), (C.GLenum)(pname), (C.GLint)(param))
// }
// func Materialiv(face glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglMaterialiv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func Materialx(face glt.Enum, pname glt.Enum, param int32) {
// 	C.goglMaterialx((C.GLenum)(face), (C.GLenum)(pname), (C.GLfixed)(param))
// }
// func Materialxv(face glt.Enum, pname glt.Enum, param *int32) {
// 	C.goglMaterialxv((C.GLenum)(face), (C.GLenum)(pname), (*C.GLfixed)(param))
// }
// func MatrixMode(mode glt.Enum) {
// 	C.goglMatrixMode((C.GLenum)(mode))
// }
// func MultMatrixd(m *float64) {
// 	C.goglMultMatrixd((*C.GLdouble)(m))
// }
// func MultMatrixf(m *float32) {
// 	C.goglMultMatrixf((*C.GLfloat)(m))
// }
// func MultMatrixx(m *int32) {
// 	C.goglMultMatrixx((*C.GLfixed)(m))
// }
// func MultTransposeMatrixd(m *float64) {
// 	C.goglMultTransposeMatrixd((*C.GLdouble)(m))
// }
// func MultTransposeMatrixf(m *float32) {
// 	C.goglMultTransposeMatrixf((*C.GLfloat)(m))
// }
// func MultiDrawArrays(mode glt.Enum, first *int32, count *uint32, drawcount uint32) {
// 	C.goglMultiDrawArrays((C.GLenum)(mode), (*C.GLint)(first), (*C.GLsizei)(count), (C.GLsizei)(drawcount))
// }
// func MultiDrawElements(mode glt.Enum, count *uint32, gltype glt.Enum, indices *glt.Pointer, drawcount uint32) {
// 	C.goglMultiDrawElements((C.GLenum)(mode), (*C.GLsizei)(count), (C.GLenum)(gltype), glt.Pointer(indices), (C.GLsizei)(drawcount))
// }
// func MultiTexCoord1d(target glt.Enum, s float64) {
// 	C.goglMultiTexCoord1d((C.GLenum)(target), (C.GLdouble)(s))
// }
// func MultiTexCoord1dv(target glt.Enum, v *float64) {
// 	C.goglMultiTexCoord1dv((C.GLenum)(target), (*C.GLdouble)(v))
// }
// func MultiTexCoord1f(target glt.Enum, s float32) {
// 	C.goglMultiTexCoord1f((C.GLenum)(target), (C.GLfloat)(s))
// }
// func MultiTexCoord1fv(target glt.Enum, v *float32) {
// 	C.goglMultiTexCoord1fv((C.GLenum)(target), (*C.GLfloat)(v))
// }
// func MultiTexCoord1i(target glt.Enum, s int32) {
// 	C.goglMultiTexCoord1i((C.GLenum)(target), (C.GLint)(s))
// }
// func MultiTexCoord1iv(target glt.Enum, v *int32) {
// 	C.goglMultiTexCoord1iv((C.GLenum)(target), (*C.GLint)(v))
// }
// func MultiTexCoord1s(target glt.Enum, s int16) {
// 	C.goglMultiTexCoord1s((C.GLenum)(target), (C.GLshort)(s))
// }
// func MultiTexCoord1sv(target glt.Enum, v *int16) {
// 	C.goglMultiTexCoord1sv((C.GLenum)(target), (*C.GLshort)(v))
// }
// func MultiTexCoord2d(target glt.Enum, s float64, t float64) {
// 	C.goglMultiTexCoord2d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t))
// }
// func MultiTexCoord2dv(target glt.Enum, v *float64) {
// 	C.goglMultiTexCoord2dv((C.GLenum)(target), (*C.GLdouble)(v))
// }
// func MultiTexCoord2f(target glt.Enum, s float32, t float32) {
// 	C.goglMultiTexCoord2f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t))
// }
// func MultiTexCoord2fv(target glt.Enum, v *float32) {
// 	C.goglMultiTexCoord2fv((C.GLenum)(target), (*C.GLfloat)(v))
// }
// func MultiTexCoord2i(target glt.Enum, s int32, t int32) {
// 	C.goglMultiTexCoord2i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t))
// }
// func MultiTexCoord2iv(target glt.Enum, v *int32) {
// 	C.goglMultiTexCoord2iv((C.GLenum)(target), (*C.GLint)(v))
// }
// func MultiTexCoord2s(target glt.Enum, s int16, t int16) {
// 	C.goglMultiTexCoord2s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t))
// }
// func MultiTexCoord2sv(target glt.Enum, v *int16) {
// 	C.goglMultiTexCoord2sv((C.GLenum)(target), (*C.GLshort)(v))
// }
// func MultiTexCoord3d(target glt.Enum, s float64, t float64, r float64) {
// 	C.goglMultiTexCoord3d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
// }
// func MultiTexCoord3dv(target glt.Enum, v *float64) {
// 	C.goglMultiTexCoord3dv((C.GLenum)(target), (*C.GLdouble)(v))
// }
// func MultiTexCoord3f(target glt.Enum, s float32, t float32, r float32) {
// 	C.goglMultiTexCoord3f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
// }
// func MultiTexCoord3fv(target glt.Enum, v *float32) {
// 	C.goglMultiTexCoord3fv((C.GLenum)(target), (*C.GLfloat)(v))
// }
// func MultiTexCoord3i(target glt.Enum, s int32, t int32, r int32) {
// 	C.goglMultiTexCoord3i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
// }
// func MultiTexCoord3iv(target glt.Enum, v *int32) {
// 	C.goglMultiTexCoord3iv((C.GLenum)(target), (*C.GLint)(v))
// }
// func MultiTexCoord3s(target glt.Enum, s int16, t int16, r int16) {
// 	C.goglMultiTexCoord3s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
// }
// func MultiTexCoord3sv(target glt.Enum, v *int16) {
// 	C.goglMultiTexCoord3sv((C.GLenum)(target), (*C.GLshort)(v))
// }
// func MultiTexCoord4d(target glt.Enum, s float64, t float64, r float64, q float64) {
// 	C.goglMultiTexCoord4d((C.GLenum)(target), (C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
// }
// func MultiTexCoord4dv(target glt.Enum, v *float64) {
// 	C.goglMultiTexCoord4dv((C.GLenum)(target), (*C.GLdouble)(v))
// }
// func MultiTexCoord4f(target glt.Enum, s float32, t float32, r float32, q float32) {
// 	C.goglMultiTexCoord4f((C.GLenum)(target), (C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
// }
// func MultiTexCoord4fv(target glt.Enum, v *float32) {
// 	C.goglMultiTexCoord4fv((C.GLenum)(target), (*C.GLfloat)(v))
// }
// func MultiTexCoord4i(target glt.Enum, s int32, t int32, r int32, q int32) {
// 	C.goglMultiTexCoord4i((C.GLenum)(target), (C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
// }
// func MultiTexCoord4iv(target glt.Enum, v *int32) {
// 	C.goglMultiTexCoord4iv((C.GLenum)(target), (*C.GLint)(v))
// }
// func MultiTexCoord4s(target glt.Enum, s int16, t int16, r int16, q int16) {
// 	C.goglMultiTexCoord4s((C.GLenum)(target), (C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
// }
// func MultiTexCoord4sv(target glt.Enum, v *int16) {
// 	C.goglMultiTexCoord4sv((C.GLenum)(target), (*C.GLshort)(v))
// }
// func MultiTexCoord4x(texture glt.Enum, s int32, t int32, r int32, q int32) {
// 	C.goglMultiTexCoord4x((C.GLenum)(texture), (C.GLfixed)(s), (C.GLfixed)(t), (C.GLfixed)(r), (C.GLfixed)(q))
// }
// func NewList(list uint32, mode glt.Enum) {
// 	C.goglNewList((C.GLuint)(list), (C.GLenum)(mode))
// }
// func Normal3b(nx int8, ny int8, nz int8) {
// 	C.goglNormal3b((C.GLbyte)(nx), (C.GLbyte)(ny), (C.GLbyte)(nz))
// }
// func Normal3bv(v *int8) {
// 	C.goglNormal3bv((*C.GLbyte)(v))
// }
// func Normal3d(nx float64, ny float64, nz float64) {
// 	C.goglNormal3d((C.GLdouble)(nx), (C.GLdouble)(ny), (C.GLdouble)(nz))
// }
// func Normal3dv(v *float64) {
// 	C.goglNormal3dv((*C.GLdouble)(v))
// }
// func Normal3f(nx float32, ny float32, nz float32) {
// 	C.goglNormal3f((C.GLfloat)(nx), (C.GLfloat)(ny), (C.GLfloat)(nz))
// }
// func Normal3fv(v *float32) {
// 	C.goglNormal3fv((*C.GLfloat)(v))
// }
// func Normal3i(nx int32, ny int32, nz int32) {
// 	C.goglNormal3i((C.GLint)(nx), (C.GLint)(ny), (C.GLint)(nz))
// }
// func Normal3iv(v *int32) {
// 	C.goglNormal3iv((*C.GLint)(v))
// }
// func Normal3s(nx int16, ny int16, nz int16) {
// 	C.goglNormal3s((C.GLshort)(nx), (C.GLshort)(ny), (C.GLshort)(nz))
// }
// func Normal3sv(v *int16) {
// 	C.goglNormal3sv((*C.GLshort)(v))
// }
// func Normal3x(nx int32, ny int32, nz int32) {
// 	C.goglNormal3x((C.GLfixed)(nx), (C.GLfixed)(ny), (C.GLfixed)(nz))
// }
// func NormalPointer(gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglNormalPointer((C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func Ortho(left float64, right float64, bottom float64, top float64, zNear float64, zFar float64) {
// 	C.goglOrtho((C.GLdouble)(left), (C.GLdouble)(right), (C.GLdouble)(bottom), (C.GLdouble)(top), (C.GLdouble)(zNear), (C.GLdouble)(zFar))
// }
// func Orthof(l float32, r float32, b float32, t float32, n float32, f float32) {
// 	C.goglOrthof((C.GLfloat)(l), (C.GLfloat)(r), (C.GLfloat)(b), (C.GLfloat)(t), (C.GLfloat)(n), (C.GLfloat)(f))
// }
// func Orthox(l int32, r int32, b int32, t int32, n int32, f int32) {
// 	C.goglOrthox((C.GLfixed)(l), (C.GLfixed)(r), (C.GLfixed)(b), (C.GLfixed)(t), (C.GLfixed)(n), (C.GLfixed)(f))
// }
// func PassThrough(token float32) {
// 	C.goglPassThrough((C.GLfloat)(token))
// }
// func PixelMapfv(glmap glt.Enum, mapsize uint32, values *float32) {
// 	C.goglPixelMapfv((C.GLenum)(glmap), (C.GLsizei)(mapsize), (*C.GLfloat)(values))
// }
// func PixelMapuiv(glmap glt.Enum, mapsize uint32, values *uint32) {
// 	C.goglPixelMapuiv((C.GLenum)(glmap), (C.GLsizei)(mapsize), (*C.GLuint)(values))
// }
// func PixelMapusv(glmap glt.Enum, mapsize uint32, values *uint16) {
// 	C.goglPixelMapusv((C.GLenum)(glmap), (C.GLsizei)(mapsize), (*C.GLushort)(values))
// }
// func PixelStoref(pname glt.Enum, param float32) {
// 	C.goglPixelStoref((C.GLenum)(pname), (C.GLfloat)(param))
// }
// func PixelStorei(pname glt.Enum, param int32) {
// 	C.goglPixelStorei((C.GLenum)(pname), (C.GLint)(param))
// }
// func PixelTransferf(pname glt.Enum, param float32) {
// 	C.goglPixelTransferf((C.GLenum)(pname), (C.GLfloat)(param))
// }
// func PixelTransferi(pname glt.Enum, param int32) {
// 	C.goglPixelTransferi((C.GLenum)(pname), (C.GLint)(param))
// }
// func PixelZoom(xfactor float32, yfactor float32) {
// 	C.goglPixelZoom((C.GLfloat)(xfactor), (C.GLfloat)(yfactor))
// }
// func PointParameterf(pname glt.Enum, param float32) {
// 	C.goglPointParameterf((C.GLenum)(pname), (C.GLfloat)(param))
// }
// func PointParameterfv(pname glt.Enum, params *float32) {
// 	C.goglPointParameterfv((C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func PointParameteri(pname glt.Enum, param int32) {
// 	C.goglPointParameteri((C.GLenum)(pname), (C.GLint)(param))
// }
// func PointParameteriv(pname glt.Enum, params *int32) {
// 	C.goglPointParameteriv((C.GLenum)(pname), (*C.GLint)(params))
// }
// func PointParameterx(pname glt.Enum, param int32) {
// 	C.goglPointParameterx((C.GLenum)(pname), (C.GLfixed)(param))
// }
// func PointParameterxv(pname glt.Enum, params *int32) {
// 	C.goglPointParameterxv((C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func PointSize(size float32) {
// 	C.goglPointSize((C.GLfloat)(size))
// }
// func PointSizex(size int32) {
// 	C.goglPointSizex((C.GLfixed)(size))
// }
// func PolygonMode(face glt.Enum, mode glt.Enum) {
// 	C.goglPolygonMode((C.GLenum)(face), (C.GLenum)(mode))
// }
// func PolygonOffset(factor float32, units float32) {
// 	C.goglPolygonOffset((C.GLfloat)(factor), (C.GLfloat)(units))
// }
// func PolygonOffsetx(factor int32, units int32) {
// 	C.goglPolygonOffsetx((C.GLfixed)(factor), (C.GLfixed)(units))
// }
// func PolygonStipple(mask *uint8) {
// 	C.goglPolygonStipple((*C.GLubyte)(mask))
// }
// func PopAttrib() {
// 	C.goglPopAttrib()
// }
// func PopClientAttrib() {
// 	C.goglPopClientAttrib()
// }
// func PopMatrix() {
// 	C.goglPopMatrix()
// }
// func PopName() {
// 	C.goglPopName()
// }
// func PrioritizeTextures(n uint32, textures *uint32, priorities *float32) {
// 	C.goglPrioritizeTextures((C.GLsizei)(n), (*C.GLuint)(textures), (*C.GLfloat)(priorities))
// }
// func PushAttrib(mask glt.Bitfieled) {
// 	C.goglPushAttrib((C.GLbitfield)(mask))
// }
// func PushClientAttrib(mask glt.Bitfieled) {
// 	C.goglPushClientAttrib((C.GLbitfield)(mask))
// }
// func PushMatrix() {
// 	C.goglPushMatrix()
// }
// func PushName(name uint32) {
// 	C.goglPushName((C.GLuint)(name))
// }
// func RasterPos2d(x float64, y float64) {
// 	C.goglRasterPos2d((C.GLdouble)(x), (C.GLdouble)(y))
// }
// func RasterPos2dv(v *float64) {
// 	C.goglRasterPos2dv((*C.GLdouble)(v))
// }
// func RasterPos2f(x float32, y float32) {
// 	C.goglRasterPos2f((C.GLfloat)(x), (C.GLfloat)(y))
// }
// func RasterPos2fv(v *float32) {
// 	C.goglRasterPos2fv((*C.GLfloat)(v))
// }
// func RasterPos2i(x int32, y int32) {
// 	C.goglRasterPos2i((C.GLint)(x), (C.GLint)(y))
// }
// func RasterPos2iv(v *int32) {
// 	C.goglRasterPos2iv((*C.GLint)(v))
// }
// func RasterPos2s(x int16, y int16) {
// 	C.goglRasterPos2s((C.GLshort)(x), (C.GLshort)(y))
// }
// func RasterPos2sv(v *int16) {
// 	C.goglRasterPos2sv((*C.GLshort)(v))
// }
// func RasterPos3d(x float64, y float64, z float64) {
// 	C.goglRasterPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func RasterPos3dv(v *float64) {
// 	C.goglRasterPos3dv((*C.GLdouble)(v))
// }
// func RasterPos3f(x float32, y float32, z float32) {
// 	C.goglRasterPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func RasterPos3fv(v *float32) {
// 	C.goglRasterPos3fv((*C.GLfloat)(v))
// }
// func RasterPos3i(x int32, y int32, z int32) {
// 	C.goglRasterPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
// }
// func RasterPos3iv(v *int32) {
// 	C.goglRasterPos3iv((*C.GLint)(v))
// }
// func RasterPos3s(x int16, y int16, z int16) {
// 	C.goglRasterPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
// }
// func RasterPos3sv(v *int16) {
// 	C.goglRasterPos3sv((*C.GLshort)(v))
// }
// func RasterPos4d(x float64, y float64, z float64, w float64) {
// 	C.goglRasterPos4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
// }
// func RasterPos4dv(v *float64) {
// 	C.goglRasterPos4dv((*C.GLdouble)(v))
// }
// func RasterPos4f(x float32, y float32, z float32, w float32) {
// 	C.goglRasterPos4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
// }
// func RasterPos4fv(v *float32) {
// 	C.goglRasterPos4fv((*C.GLfloat)(v))
// }
// func RasterPos4i(x int32, y int32, z int32, w int32) {
// 	C.goglRasterPos4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
// }
// func RasterPos4iv(v *int32) {
// 	C.goglRasterPos4iv((*C.GLint)(v))
// }
// func RasterPos4s(x int16, y int16, z int16, w int16) {
// 	C.goglRasterPos4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
// }
// func RasterPos4sv(v *int16) {
// 	C.goglRasterPos4sv((*C.GLshort)(v))
// }
// func ReadBuffer(mode glt.Enum) {
// 	C.goglReadBuffer((C.GLenum)(mode))
// }
// func ReadPixels(x int32, y int32, width uint32, height uint32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglReadPixels((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func Rectd(x1 float64, y1 float64, x2 float64, y2 float64) {
// 	C.goglRectd((C.GLdouble)(x1), (C.GLdouble)(y1), (C.GLdouble)(x2), (C.GLdouble)(y2))
// }
// func Rectdv(v1 *float64, v2 *float64) {
// 	C.goglRectdv((*C.GLdouble)(v1), (*C.GLdouble)(v2))
// }
// func Rectf(x1 float32, y1 float32, x2 float32, y2 float32) {
// 	C.goglRectf((C.GLfloat)(x1), (C.GLfloat)(y1), (C.GLfloat)(x2), (C.GLfloat)(y2))
// }
// func Rectfv(v1 *float32, v2 *float32) {
// 	C.goglRectfv((*C.GLfloat)(v1), (*C.GLfloat)(v2))
// }
// func Recti(x1 int32, y1 int32, x2 int32, y2 int32) {
// 	C.goglRecti((C.GLint)(x1), (C.GLint)(y1), (C.GLint)(x2), (C.GLint)(y2))
// }
// func Rectiv(v1 *int32, v2 *int32) {
// 	C.goglRectiv((*C.GLint)(v1), (*C.GLint)(v2))
// }
// func Rects(x1 int16, y1 int16, x2 int16, y2 int16) {
// 	C.goglRects((C.GLshort)(x1), (C.GLshort)(y1), (C.GLshort)(x2), (C.GLshort)(y2))
// }
// func Rectsv(v1 *int16, v2 *int16) {
// 	C.goglRectsv((*C.GLshort)(v1), (*C.GLshort)(v2))
// }
// func ReleaseShaderCompiler() {
// 	C.goglReleaseShaderCompiler()
// }
// func RenderMode(mode glt.Enum) int32 {
// 	return (C.GLint)(C.goglRenderMode((C.GLenum)(mode)))
// }
// func RenderbufferStorage(target glt.Enum, internalformat glt.Enum, width uint32, height uint32) {
// 	C.goglRenderbufferStorage((C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
// }
// func Rotated(angle float64, x float64, y float64, z float64) {
// 	C.goglRotated((C.GLdouble)(angle), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func Rotatef(angle float32, x float32, y float32, z float32) {
// 	C.goglRotatef((C.GLfloat)(angle), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func Rotatex(angle int32, x int32, y int32, z int32) {
// 	C.goglRotatex((C.GLfixed)(angle), (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
// }
// func SampleCoverage(value float32, invert bool) {
// 	C.goglSampleCoverage((C.GLfloat)(value), glt.GLBool(invert))
// }
// func SampleCoveragex(value , invert bool) {
// 	C.goglSampleCoveragex((C.GLclampx)(value), glt.GLBool(invert))
// }
// func Scaled(x float64, y float64, z float64) {
// 	C.goglScaled((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func Scalef(x float32, y float32, z float32) {
// 	C.goglScalef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func Scalex(x int32, y int32, z int32) {
// 	C.goglScalex((C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
// }
// func Scissor(x int32, y int32, width uint32, height uint32) {
// 	C.goglScissor((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
// }
// func SecondaryColor3b(red int8, green int8, blue int8) {
// 	C.goglSecondaryColor3b((C.GLbyte)(red), (C.GLbyte)(green), (C.GLbyte)(blue))
// }
// func SecondaryColor3bv(v *int8) {
// 	C.goglSecondaryColor3bv((*C.GLbyte)(v))
// }
// func SecondaryColor3d(red float64, green float64, blue float64) {
// 	C.goglSecondaryColor3d((C.GLdouble)(red), (C.GLdouble)(green), (C.GLdouble)(blue))
// }
// func SecondaryColor3dv(v *float64) {
// 	C.goglSecondaryColor3dv((*C.GLdouble)(v))
// }
// func SecondaryColor3f(red float32, green float32, blue float32) {
// 	C.goglSecondaryColor3f((C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue))
// }
// func SecondaryColor3fv(v *float32) {
// 	C.goglSecondaryColor3fv((*C.GLfloat)(v))
// }
// func SecondaryColor3i(red int32, green int32, blue int32) {
// 	C.goglSecondaryColor3i((C.GLint)(red), (C.GLint)(green), (C.GLint)(blue))
// }
// func SecondaryColor3iv(v *int32) {
// 	C.goglSecondaryColor3iv((*C.GLint)(v))
// }
// func SecondaryColor3s(red int16, green int16, blue int16) {
// 	C.goglSecondaryColor3s((C.GLshort)(red), (C.GLshort)(green), (C.GLshort)(blue))
// }
// func SecondaryColor3sv(v *int16) {
// 	C.goglSecondaryColor3sv((*C.GLshort)(v))
// }
// func SecondaryColor3ub(red uint8, green uint8, blue uint8) {
// 	C.goglSecondaryColor3ub((C.GLubyte)(red), (C.GLubyte)(green), (C.GLubyte)(blue))
// }
// func SecondaryColor3ubv(v *uint8) {
// 	C.goglSecondaryColor3ubv((*C.GLubyte)(v))
// }
// func SecondaryColor3ui(red uint32, green uint32, blue uint32) {
// 	C.goglSecondaryColor3ui((C.GLuint)(red), (C.GLuint)(green), (C.GLuint)(blue))
// }
// func SecondaryColor3uiv(v *uint32) {
// 	C.goglSecondaryColor3uiv((*C.GLuint)(v))
// }
// func SecondaryColor3us(red uint16, green uint16, blue uint16) {
// 	C.goglSecondaryColor3us((C.GLushort)(red), (C.GLushort)(green), (C.GLushort)(blue))
// }
// func SecondaryColor3usv(v *uint16) {
// 	C.goglSecondaryColor3usv((*C.GLushort)(v))
// }
// func SecondaryColorPointer(size int32, gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglSecondaryColorPointer((C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func SelectBuffer(size uint32, buffer *uint32) {
// 	C.goglSelectBuffer((C.GLsizei)(size), (*C.GLuint)(buffer))
// }
// func ShadeModel(mode glt.Enum) {
// 	C.goglShadeModel((C.GLenum)(mode))
// }
// func ShaderBinary(count uint32, shaders *uint32, binaryformat glt.Enum, binary glt.Pointer, length uint32) {
// 	C.goglShaderBinary((C.GLsizei)(count), (*C.GLuint)(shaders), (C.GLenum)(binaryformat), glt.Pointer(binary), (C.GLsizei)(length))
// }
// func ShaderSource(shader uint32, count uint32, glstring **int8, length *int32) {
// 	C.goglShaderSource((C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(glstring), (*C.GLint)(length))
// }
// func StencilFunc(glfunc glt.Enum, ref int32, mask uint32) {
// 	C.goglStencilFunc((C.GLenum)(glfunc), (C.GLint)(ref), (C.GLuint)(mask))
// }
// func StencilFuncSeparate(face glt.Enum, glfunc glt.Enum, ref int32, mask uint32) {
// 	C.goglStencilFuncSeparate((C.GLenum)(face), (C.GLenum)(glfunc), (C.GLint)(ref), (C.GLuint)(mask))
// }
// func StencilMask(mask uint32) {
// 	C.goglStencilMask((C.GLuint)(mask))
// }
// func StencilMaskSeparate(face glt.Enum, mask uint32) {
// 	C.goglStencilMaskSeparate((C.GLenum)(face), (C.GLuint)(mask))
// }
// func StencilOp(fail glt.Enum, zfail glt.Enum, zpass glt.Enum) {
// 	C.goglStencilOp((C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
// }
// func StencilOpSeparate(face glt.Enum, sfail glt.Enum, dpfail glt.Enum, dppass glt.Enum) {
// 	C.goglStencilOpSeparate((C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
// }
// func TexCoord1d(s float64) {
// 	C.goglTexCoord1d((C.GLdouble)(s))
// }
// func TexCoord1dv(v *float64) {
// 	C.goglTexCoord1dv((*C.GLdouble)(v))
// }
// func TexCoord1f(s float32) {
// 	C.goglTexCoord1f((C.GLfloat)(s))
// }
// func TexCoord1fv(v *float32) {
// 	C.goglTexCoord1fv((*C.GLfloat)(v))
// }
// func TexCoord1i(s int32) {
// 	C.goglTexCoord1i((C.GLint)(s))
// }
// func TexCoord1iv(v *int32) {
// 	C.goglTexCoord1iv((*C.GLint)(v))
// }
// func TexCoord1s(s int16) {
// 	C.goglTexCoord1s((C.GLshort)(s))
// }
// func TexCoord1sv(v *int16) {
// 	C.goglTexCoord1sv((*C.GLshort)(v))
// }
// func TexCoord2d(s float64, t float64) {
// 	C.goglTexCoord2d((C.GLdouble)(s), (C.GLdouble)(t))
// }
// func TexCoord2dv(v *float64) {
// 	C.goglTexCoord2dv((*C.GLdouble)(v))
// }
// func TexCoord2f(s float32, t float32) {
// 	C.goglTexCoord2f((C.GLfloat)(s), (C.GLfloat)(t))
// }
// func TexCoord2fv(v *float32) {
// 	C.goglTexCoord2fv((*C.GLfloat)(v))
// }
// func TexCoord2i(s int32, t int32) {
// 	C.goglTexCoord2i((C.GLint)(s), (C.GLint)(t))
// }
// func TexCoord2iv(v *int32) {
// 	C.goglTexCoord2iv((*C.GLint)(v))
// }
// func TexCoord2s(s int16, t int16) {
// 	C.goglTexCoord2s((C.GLshort)(s), (C.GLshort)(t))
// }
// func TexCoord2sv(v *int16) {
// 	C.goglTexCoord2sv((*C.GLshort)(v))
// }
// func TexCoord3d(s float64, t float64, r float64) {
// 	C.goglTexCoord3d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r))
// }
// func TexCoord3dv(v *float64) {
// 	C.goglTexCoord3dv((*C.GLdouble)(v))
// }
// func TexCoord3f(s float32, t float32, r float32) {
// 	C.goglTexCoord3f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r))
// }
// func TexCoord3fv(v *float32) {
// 	C.goglTexCoord3fv((*C.GLfloat)(v))
// }
// func TexCoord3i(s int32, t int32, r int32) {
// 	C.goglTexCoord3i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r))
// }
// func TexCoord3iv(v *int32) {
// 	C.goglTexCoord3iv((*C.GLint)(v))
// }
// func TexCoord3s(s int16, t int16, r int16) {
// 	C.goglTexCoord3s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r))
// }
// func TexCoord3sv(v *int16) {
// 	C.goglTexCoord3sv((*C.GLshort)(v))
// }
// func TexCoord4d(s float64, t float64, r float64, q float64) {
// 	C.goglTexCoord4d((C.GLdouble)(s), (C.GLdouble)(t), (C.GLdouble)(r), (C.GLdouble)(q))
// }
// func TexCoord4dv(v *float64) {
// 	C.goglTexCoord4dv((*C.GLdouble)(v))
// }
// func TexCoord4f(s float32, t float32, r float32, q float32) {
// 	C.goglTexCoord4f((C.GLfloat)(s), (C.GLfloat)(t), (C.GLfloat)(r), (C.GLfloat)(q))
// }
// func TexCoord4fv(v *float32) {
// 	C.goglTexCoord4fv((*C.GLfloat)(v))
// }
// func TexCoord4i(s int32, t int32, r int32, q int32) {
// 	C.goglTexCoord4i((C.GLint)(s), (C.GLint)(t), (C.GLint)(r), (C.GLint)(q))
// }
// func TexCoord4iv(v *int32) {
// 	C.goglTexCoord4iv((*C.GLint)(v))
// }
// func TexCoord4s(s int16, t int16, r int16, q int16) {
// 	C.goglTexCoord4s((C.GLshort)(s), (C.GLshort)(t), (C.GLshort)(r), (C.GLshort)(q))
// }
// func TexCoord4sv(v *int16) {
// 	C.goglTexCoord4sv((*C.GLshort)(v))
// }
// func TexCoordPointer(size int32, gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglTexCoordPointer((C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func TexEnvf(target glt.Enum, pname glt.Enum, param float32) {
// 	C.goglTexEnvf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
// }
// func TexEnvfv(target glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglTexEnvfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func TexEnvi(target glt.Enum, pname glt.Enum, param int32) {
// 	C.goglTexEnvi((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
// }
// func TexEnviv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglTexEnviv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func TexEnvx(target glt.Enum, pname glt.Enum, param int32) {
// 	C.goglTexEnvx((C.GLenum)(target), (C.GLenum)(pname), (C.GLfixed)(param))
// }
// func TexEnvxv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglTexEnvxv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func TexGend(coord glt.Enum, pname glt.Enum, param float64) {
// 	C.goglTexGend((C.GLenum)(coord), (C.GLenum)(pname), (C.GLdouble)(param))
// }
// func TexGendv(coord glt.Enum, pname glt.Enum, params *float64) {
// 	C.goglTexGendv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLdouble)(params))
// }
// func TexGenf(coord glt.Enum, pname glt.Enum, param float32) {
// 	C.goglTexGenf((C.GLenum)(coord), (C.GLenum)(pname), (C.GLfloat)(param))
// }
// func TexGenfv(coord glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglTexGenfv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func TexGeni(coord glt.Enum, pname glt.Enum, param int32) {
// 	C.goglTexGeni((C.GLenum)(coord), (C.GLenum)(pname), (C.GLint)(param))
// }
// func TexGeniv(coord glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglTexGeniv((C.GLenum)(coord), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func TexImage1D(target glt.Enum, level int32, internalformat int32, width uint32, border int32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func TexImage2D(target glt.Enum, level int32, internalformat int32, width uint32, height uint32, border int32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func TexImage3D(target glt.Enum, level int32, internalformat int32, width uint32, height uint32, depth uint32, border int32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func TexParameterf(target glt.Enum, pname glt.Enum, param float32) {
// 	C.goglTexParameterf((C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
// }
// func TexParameterfv(target glt.Enum, pname glt.Enum, params *float32) {
// 	C.goglTexParameterfv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(params))
// }
// func TexParameteri(target glt.Enum, pname glt.Enum, param int32) {
// 	C.goglTexParameteri((C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
// }
// func TexParameteriv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglTexParameteriv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(params))
// }
// func TexParameterx(target glt.Enum, pname glt.Enum, param int32) {
// 	C.goglTexParameterx((C.GLenum)(target), (C.GLenum)(pname), (C.GLfixed)(param))
// }
// func TexParameterxv(target glt.Enum, pname glt.Enum, params *int32) {
// 	C.goglTexParameterxv((C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(params))
// }
// func TexSubImage1D(target glt.Enum, level int32, xoffset int32, width uint32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexSubImage1D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func TexSubImage2D(target glt.Enum, level int32, xoffset int32, yoffset int32, width uint32, height uint32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexSubImage2D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func TexSubImage3D(target glt.Enum, level int32, xoffset int32, yoffset int32, zoffset int32, width uint32, height uint32, depth uint32, format glt.Enum, gltype glt.Enum, pixels glt.Pointer) {
// 	C.goglTexSubImage3D((C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(gltype), glt.Pointer(pixels))
// }
// func Translated(x float64, y float64, z float64) {
// 	C.goglTranslated((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func Translatef(x float32, y float32, z float32) {
// 	C.goglTranslatef((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func Translatex(x int32, y int32, z int32) {
// 	C.goglTranslatex((C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
// }
// func Uniform1f(location int32, v0 float32) {
// 	C.goglUniform1f((C.GLint)(location), (C.GLfloat)(v0))
// }
// func Uniform1fv(location int32, count uint32, value *float32) {
// 	C.goglUniform1fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
// }
// func Uniform1i(location int32, v0 int32) {
// 	C.goglUniform1i((C.GLint)(location), (C.GLint)(v0))
// }
// func Uniform1iv(location int32, count uint32, value *int32) {
// 	C.goglUniform1iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
// }
// func Uniform2f(location int32, v0 float32, v1 float32) {
// 	C.goglUniform2f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
// }
// func Uniform2fv(location int32, count uint32, value *float32) {
// 	C.goglUniform2fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
// }
// func Uniform2i(location int32, v0 int32, v1 int32) {
// 	C.goglUniform2i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
// }
// func Uniform2iv(location int32, count uint32, value *int32) {
// 	C.goglUniform2iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
// }
// func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
// 	C.goglUniform3f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
// }
// func Uniform3fv(location int32, count uint32, value *float32) {
// 	C.goglUniform3fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
// }
// func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
// 	C.goglUniform3i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
// }
// func Uniform3iv(location int32, count uint32, value *int32) {
// 	C.goglUniform3iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
// }
// func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
// 	C.goglUniform4f((C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
// }
// func Uniform4fv(location int32, count uint32, value *float32) {
// 	C.goglUniform4fv((C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(value))
// }
// func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
// 	C.goglUniform4i((C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
// }
// func Uniform4iv(location int32, count uint32, value *int32) {
// 	C.goglUniform4iv((C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(value))
// }
// func UniformMatrix2fv(location int32, count uint32, transpose bool, value *float32) {
// 	C.goglUniformMatrix2fv((C.GLint)(location), (C.GLsizei)(count), glt.GLBool(transpose), (*C.GLfloat)(value))
// }
// func UniformMatrix3fv(location int32, count uint32, transpose bool, value *float32) {
// 	C.goglUniformMatrix3fv((C.GLint)(location), (C.GLsizei)(count), glt.GLBool(transpose), (*C.GLfloat)(value))
// }
// func UniformMatrix4fv(location int32, count uint32, transpose bool, value *float32) {
// 	C.goglUniformMatrix4fv((C.GLint)(location), (C.GLsizei)(count), glt.GLBool(transpose), (*C.GLfloat)(value))
// }
// func UnmapBuffer(target glt.Enum) bool {
// 	return glt.GLBoolean(C.goglUnmapBuffer((C.GLenum)(target)))
// }
// func UseProgram(program uint32) {
// 	C.goglUseProgram((C.GLuint)(program))
// }
// func ValidateProgram(program uint32) {
// 	C.goglValidateProgram((C.GLuint)(program))
// }
// func Vertex2d(x float64, y float64) {
// 	C.goglVertex2d((C.GLdouble)(x), (C.GLdouble)(y))
// }
// func Vertex2dv(v *float64) {
// 	C.goglVertex2dv((*C.GLdouble)(v))
// }
// func Vertex2f(x float32, y float32) {
// 	C.goglVertex2f((C.GLfloat)(x), (C.GLfloat)(y))
// }
// func Vertex2fv(v *float32) {
// 	C.goglVertex2fv((*C.GLfloat)(v))
// }
// func Vertex2i(x int32, y int32) {
// 	C.goglVertex2i((C.GLint)(x), (C.GLint)(y))
// }
// func Vertex2iv(v *int32) {
// 	C.goglVertex2iv((*C.GLint)(v))
// }
// func Vertex2s(x int16, y int16) {
// 	C.goglVertex2s((C.GLshort)(x), (C.GLshort)(y))
// }
// func Vertex2sv(v *int16) {
// 	C.goglVertex2sv((*C.GLshort)(v))
// }
// func Vertex3d(x float64, y float64, z float64) {
// 	C.goglVertex3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func Vertex3dv(v *float64) {
// 	C.goglVertex3dv((*C.GLdouble)(v))
// }
// func Vertex3f(x float32, y float32, z float32) {
// 	C.goglVertex3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func Vertex3fv(v *float32) {
// 	C.goglVertex3fv((*C.GLfloat)(v))
// }
// func Vertex3i(x int32, y int32, z int32) {
// 	C.goglVertex3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
// }
// func Vertex3iv(v *int32) {
// 	C.goglVertex3iv((*C.GLint)(v))
// }
// func Vertex3s(x int16, y int16, z int16) {
// 	C.goglVertex3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
// }
// func Vertex3sv(v *int16) {
// 	C.goglVertex3sv((*C.GLshort)(v))
// }
// func Vertex4d(x float64, y float64, z float64, w float64) {
// 	C.goglVertex4d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
// }
// func Vertex4dv(v *float64) {
// 	C.goglVertex4dv((*C.GLdouble)(v))
// }
// func Vertex4f(x float32, y float32, z float32, w float32) {
// 	C.goglVertex4f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
// }
// func Vertex4fv(v *float32) {
// 	C.goglVertex4fv((*C.GLfloat)(v))
// }
// func Vertex4i(x int32, y int32, z int32, w int32) {
// 	C.goglVertex4i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
// }
// func Vertex4iv(v *int32) {
// 	C.goglVertex4iv((*C.GLint)(v))
// }
// func Vertex4s(x int16, y int16, z int16, w int16) {
// 	C.goglVertex4s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
// }
// func Vertex4sv(v *int16) {
// 	C.goglVertex4sv((*C.GLshort)(v))
// }
// func VertexAttrib1d(index uint32, x float64) {
// 	C.goglVertexAttrib1d((C.GLuint)(index), (C.GLdouble)(x))
// }
// func VertexAttrib1dv(index uint32, v *float64) {
// 	C.goglVertexAttrib1dv((C.GLuint)(index), (*C.GLdouble)(v))
// }
// func VertexAttrib1f(index uint32, x float32) {
// 	C.goglVertexAttrib1f((C.GLuint)(index), (C.GLfloat)(x))
// }
// func VertexAttrib1fv(index uint32, v *float32) {
// 	C.goglVertexAttrib1fv((C.GLuint)(index), (*C.GLfloat)(v))
// }
// func VertexAttrib1s(index uint32, x int16) {
// 	C.goglVertexAttrib1s((C.GLuint)(index), (C.GLshort)(x))
// }
// func VertexAttrib1sv(index uint32, v *int16) {
// 	C.goglVertexAttrib1sv((C.GLuint)(index), (*C.GLshort)(v))
// }
// func VertexAttrib2d(index uint32, x float64, y float64) {
// 	C.goglVertexAttrib2d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
// }
// func VertexAttrib2dv(index uint32, v *float64) {
// 	C.goglVertexAttrib2dv((C.GLuint)(index), (*C.GLdouble)(v))
// }
// func VertexAttrib2f(index uint32, x float32, y float32) {
// 	C.goglVertexAttrib2f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
// }
// func VertexAttrib2fv(index uint32, v *float32) {
// 	C.goglVertexAttrib2fv((C.GLuint)(index), (*C.GLfloat)(v))
// }
// func VertexAttrib2s(index uint32, x int16, y int16) {
// 	C.goglVertexAttrib2s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
// }
// func VertexAttrib2sv(index uint32, v *int16) {
// 	C.goglVertexAttrib2sv((C.GLuint)(index), (*C.GLshort)(v))
// }
// func VertexAttrib3d(index uint32, x float64, y float64, z float64) {
// 	C.goglVertexAttrib3d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func VertexAttrib3dv(index uint32, v *float64) {
// 	C.goglVertexAttrib3dv((C.GLuint)(index), (*C.GLdouble)(v))
// }
// func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
// 	C.goglVertexAttrib3f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func VertexAttrib3fv(index uint32, v *float32) {
// 	C.goglVertexAttrib3fv((C.GLuint)(index), (*C.GLfloat)(v))
// }
// func VertexAttrib3s(index uint32, x int16, y int16, z int16) {
// 	C.goglVertexAttrib3s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
// }
// func VertexAttrib3sv(index uint32, v *int16) {
// 	C.goglVertexAttrib3sv((C.GLuint)(index), (*C.GLshort)(v))
// }
// func VertexAttrib4Nbv(index uint32, v *int8) {
// 	C.goglVertexAttrib4Nbv((C.GLuint)(index), (*C.GLbyte)(v))
// }
// func VertexAttrib4Niv(index uint32, v *int32) {
// 	C.goglVertexAttrib4Niv((C.GLuint)(index), (*C.GLint)(v))
// }
// func VertexAttrib4Nsv(index uint32, v *int16) {
// 	C.goglVertexAttrib4Nsv((C.GLuint)(index), (*C.GLshort)(v))
// }
// func VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8) {
// 	C.goglVertexAttrib4Nub((C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
// }
// func VertexAttrib4Nubv(index uint32, v *uint8) {
// 	C.goglVertexAttrib4Nubv((C.GLuint)(index), (*C.GLubyte)(v))
// }
// func VertexAttrib4Nuiv(index uint32, v *uint32) {
// 	C.goglVertexAttrib4Nuiv((C.GLuint)(index), (*C.GLuint)(v))
// }
// func VertexAttrib4Nusv(index uint32, v *uint16) {
// 	C.goglVertexAttrib4Nusv((C.GLuint)(index), (*C.GLushort)(v))
// }
// func VertexAttrib4bv(index uint32, v *int8) {
// 	C.goglVertexAttrib4bv((C.GLuint)(index), (*C.GLbyte)(v))
// }
// func VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64) {
// 	C.goglVertexAttrib4d((C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
// }
// func VertexAttrib4dv(index uint32, v *float64) {
// 	C.goglVertexAttrib4dv((C.GLuint)(index), (*C.GLdouble)(v))
// }
// func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
// 	C.goglVertexAttrib4f((C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
// }
// func VertexAttrib4fv(index uint32, v *float32) {
// 	C.goglVertexAttrib4fv((C.GLuint)(index), (*C.GLfloat)(v))
// }
// func VertexAttrib4iv(index uint32, v *int32) {
// 	C.goglVertexAttrib4iv((C.GLuint)(index), (*C.GLint)(v))
// }
// func VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16) {
// 	C.goglVertexAttrib4s((C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
// }
// func VertexAttrib4sv(index uint32, v *int16) {
// 	C.goglVertexAttrib4sv((C.GLuint)(index), (*C.GLshort)(v))
// }
// func VertexAttrib4ubv(index uint32, v *uint8) {
// 	C.goglVertexAttrib4ubv((C.GLuint)(index), (*C.GLubyte)(v))
// }
// func VertexAttrib4uiv(index uint32, v *uint32) {
// 	C.goglVertexAttrib4uiv((C.GLuint)(index), (*C.GLuint)(v))
// }
// func VertexAttrib4usv(index uint32, v *uint16) {
// 	C.goglVertexAttrib4usv((C.GLuint)(index), (*C.GLushort)(v))
// }
// func VertexAttribPointer(index uint32, size int32, gltype glt.Enum, normalized bool, stride uint32, pointer glt.Pointer) {
// 	C.goglVertexAttribPointer((C.GLuint)(index), (C.GLint)(size), (C.GLenum)(gltype), glt.GLBool(normalized), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func VertexPointer(size int32, gltype glt.Enum, stride uint32, pointer glt.Pointer) {
// 	C.goglVertexPointer((C.GLint)(size), (C.GLenum)(gltype), (C.GLsizei)(stride), glt.Pointer(pointer))
// }
// func Viewport(x int32, y int32, width uint32, height uint32) {
// 	C.goglViewport((C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
// }
// func WindowPos2d(x float64, y float64) {
// 	C.goglWindowPos2d((C.GLdouble)(x), (C.GLdouble)(y))
// }
// func WindowPos2dv(v *float64) {
// 	C.goglWindowPos2dv((*C.GLdouble)(v))
// }
// func WindowPos2f(x float32, y float32) {
// 	C.goglWindowPos2f((C.GLfloat)(x), (C.GLfloat)(y))
// }
// func WindowPos2fv(v *float32) {
// 	C.goglWindowPos2fv((*C.GLfloat)(v))
// }
// func WindowPos2i(x int32, y int32) {
// 	C.goglWindowPos2i((C.GLint)(x), (C.GLint)(y))
// }
// func WindowPos2iv(v *int32) {
// 	C.goglWindowPos2iv((*C.GLint)(v))
// }
// func WindowPos2s(x int16, y int16) {
// 	C.goglWindowPos2s((C.GLshort)(x), (C.GLshort)(y))
// }
// func WindowPos2sv(v *int16) {
// 	C.goglWindowPos2sv((*C.GLshort)(v))
// }
// func WindowPos3d(x float64, y float64, z float64) {
// 	C.goglWindowPos3d((C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
// }
// func WindowPos3dv(v *float64) {
// 	C.goglWindowPos3dv((*C.GLdouble)(v))
// }
// func WindowPos3f(x float32, y float32, z float32) {
// 	C.goglWindowPos3f((C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
// }
// func WindowPos3fv(v *float32) {
// 	C.goglWindowPos3fv((*C.GLfloat)(v))
// }
// func WindowPos3i(x int32, y int32, z int32) {
// 	C.goglWindowPos3i((C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
// }
// func WindowPos3iv(v *int32) {
// 	C.goglWindowPos3iv((*C.GLint)(v))
// }
// func WindowPos3s(x int16, y int16, z int16) {
// 	C.goglWindowPos3s((C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
// }
// func WindowPos3sv(v *int16) {
// 	C.goglWindowPos3sv((*C.GLshort)(v))
// }
//
// 

func ClearColor(col color.Color) {
	r, g, b, a := col.RGBA()
	C.goglClearColor(
		C.GLfloat(r)/255.0,
		C.GLfloat(g)/255.0,
		C.GLfloat(b)/255.0,
		C.GLfloat(a)/255.0,
	)
}

type ClearFlags uint32
const (
	ColorBufferBit ClearFlags = gl_COLOR_BUFFER_BIT
	DepthBufferBit ClearFlags = gl_DEPTH_BUFFER_BIT
	StencilBufferBit ClearFlags = gl_STENCIL_BUFFER_BIT
)

func Clear(bit ClearFlags) {
	C.goglClear(C.GLbitfield(bit))
}

func DrawArrays(draw Mode, first, count int) {
	C.goglDrawArrays(C.GLenum(draw), C.GLint(first), C.GLsizei(count))
}

// BufferData uploads vertex data from the system into server memory.
// The argument data must be a non-nil slice of numeric values. The
// data must be a non-nil slice or array of numeric elements. Passing
// a nil or forbidden type will result in a run-time panic.
func BufferData(t Target, data interface{}, hint UsageHint) error {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		if val.IsNil() {
			panic("Nil data argument")
		}
		switch typ.Elem().Kind() {
		case reflect.Int,reflect.Int32,reflect.Int64,reflect.Uint,reflect.Uint32,reflect.Uint64,reflect.Float32,reflect.Float64:
			C.goglBufferData(
				C.GLenum(t),
				C.GLsizeiptr(typ.Elem().Size()) * C.GLsizeiptr(val.Len()),
				unsafe.Pointer(val.Index(0).UnsafeAddr()),
				C.GLenum(hint))
		default:
			panic("Non-numeric data type " + typ.Elem().String())
		}
	default:
		panic("Non-slice or array data type " + typ.String())
	}
	return getError()
}

func Init(version string) error {
	if err := C.goglInit(); err != nil {
		defer C.free(unsafe.Pointer(err))
		return errors.New("Could not initialize function " + C.GoString(err))
	}
	return nil
}
