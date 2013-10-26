#include <stdlib.h>
#include <string.h>
#include "gogl.h"

static void* goglGetProcAddress(const char* name) {
#ifdef __APPLE__
	return dlsym(RTLD_DEFAULT, name);
#elif _WIN32
	void* pf = wglGetProcAddress((LPCSTR)name);
	if(pf) {
		return pf;
	}
	if(opengl32 == NULL) {
		opengl32 = LoadLibraryA("opengl32.dll");
	}
	return GetProcAddress(opengl32, (LPCSTR)name);
#else
	return glXGetProcAddress((const GLubyte*)name);
#endif
}

void (APIENTRYP pglAccum)(GLenum op, GLfloat value);
void (APIENTRYP pglActiveTexture)(GLenum texture);
void (APIENTRYP pglAlphaFunc)(GLenum func, GLfloat ref);
void (APIENTRYP pglAlphaFuncx)(GLenum func, GLfixed ref);
GLboolean (APIENTRYP pglAreTexturesResident)(GLsizei n, const GLuint* textures, GLboolean* residences);
void (APIENTRYP pglArrayElement)(GLint i);
void (APIENTRYP pglAttachShader)(GLuint program, GLuint shader);
void (APIENTRYP pglBegin)(GLenum mode);
void (APIENTRYP pglBeginQuery)(GLenum target, GLuint id);
void (APIENTRYP pglBindAttribLocation)(GLuint program, GLuint index, const GLchar* name);
void (APIENTRYP pglBindBuffer)(GLenum target, GLuint buffer);
void (APIENTRYP pglBindFramebuffer)(GLenum target, GLuint framebuffer);
void (APIENTRYP pglBindRenderbuffer)(GLenum target, GLuint renderbuffer);
void (APIENTRYP pglBindTexture)(GLenum target, GLuint texture);
void (APIENTRYP pglBitmap)(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, const GLubyte* bitmap);
void (APIENTRYP pglBlendColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
void (APIENTRYP pglBlendEquation)(GLenum mode);
void (APIENTRYP pglBlendEquationSeparate)(GLenum modeRGB, GLenum modeAlpha);
void (APIENTRYP pglBlendFunc)(GLenum sfactor, GLenum dfactor);
void (APIENTRYP pglBlendFuncSeparate)(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha);
void (APIENTRYP pglBufferData)(GLenum target, GLsizeiptr size, const GLvoid* data, GLenum usage);
void (APIENTRYP pglBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, const GLvoid* data);
void (APIENTRYP pglCallList)(GLuint list);
void (APIENTRYP pglCallLists)(GLsizei n, GLenum type, const GLvoid* lists);
GLenum (APIENTRYP pglCheckFramebufferStatus)(GLenum target);
void (APIENTRYP pglClear)(GLbitfield mask);
void (APIENTRYP pglClearAccum)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
void (APIENTRYP pglClearColor)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
void (APIENTRYP pglClearColorx)(GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha);
void (APIENTRYP pglClearDepth)(GLdouble depth);
void (APIENTRYP pglClearDepthf)(GLfloat d);
void (APIENTRYP pglClearDepthx)(GLfixed depth);
void (APIENTRYP pglClearIndex)(GLfloat c);
void (APIENTRYP pglClearStencil)(GLint s);
void (APIENTRYP pglClientActiveTexture)(GLenum texture);
void (APIENTRYP pglClipPlane)(GLenum plane, const GLdouble* equation);
void (APIENTRYP pglClipPlanef)(GLenum p, const GLfloat* eqn);
void (APIENTRYP pglClipPlanex)(GLenum plane, const GLfixed* equation);
void (APIENTRYP pglColor3b)(GLbyte red, GLbyte green, GLbyte blue);
void (APIENTRYP pglColor3bv)(const GLbyte* v);
void (APIENTRYP pglColor3d)(GLdouble red, GLdouble green, GLdouble blue);
void (APIENTRYP pglColor3dv)(const GLdouble* v);
void (APIENTRYP pglColor3f)(GLfloat red, GLfloat green, GLfloat blue);
void (APIENTRYP pglColor3fv)(const GLfloat* v);
void (APIENTRYP pglColor3i)(GLint red, GLint green, GLint blue);
void (APIENTRYP pglColor3iv)(const GLint* v);
void (APIENTRYP pglColor3s)(GLshort red, GLshort green, GLshort blue);
void (APIENTRYP pglColor3sv)(const GLshort* v);
void (APIENTRYP pglColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
void (APIENTRYP pglColor3ubv)(const GLubyte* v);
void (APIENTRYP pglColor3ui)(GLuint red, GLuint green, GLuint blue);
void (APIENTRYP pglColor3uiv)(const GLuint* v);
void (APIENTRYP pglColor3us)(GLushort red, GLushort green, GLushort blue);
void (APIENTRYP pglColor3usv)(const GLushort* v);
void (APIENTRYP pglColor4b)(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha);
void (APIENTRYP pglColor4bv)(const GLbyte* v);
void (APIENTRYP pglColor4d)(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha);
void (APIENTRYP pglColor4dv)(const GLdouble* v);
void (APIENTRYP pglColor4f)(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha);
void (APIENTRYP pglColor4fv)(const GLfloat* v);
void (APIENTRYP pglColor4i)(GLint red, GLint green, GLint blue, GLint alpha);
void (APIENTRYP pglColor4iv)(const GLint* v);
void (APIENTRYP pglColor4s)(GLshort red, GLshort green, GLshort blue, GLshort alpha);
void (APIENTRYP pglColor4sv)(const GLshort* v);
void (APIENTRYP pglColor4ub)(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha);
void (APIENTRYP pglColor4ubv)(const GLubyte* v);
void (APIENTRYP pglColor4ui)(GLuint red, GLuint green, GLuint blue, GLuint alpha);
void (APIENTRYP pglColor4uiv)(const GLuint* v);
void (APIENTRYP pglColor4us)(GLushort red, GLushort green, GLushort blue, GLushort alpha);
void (APIENTRYP pglColor4usv)(const GLushort* v);
void (APIENTRYP pglColor4x)(GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha);
void (APIENTRYP pglColorMask)(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
void (APIENTRYP pglColorMaterial)(GLenum face, GLenum mode);
void (APIENTRYP pglColorPointer)(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglCompileShader)(GLuint shader);
void (APIENTRYP pglCompressedTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCompressedTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCompressedTexImage3D)(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCompressedTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCompressedTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCompressedTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid* data);
void (APIENTRYP pglCopyPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type);
void (APIENTRYP pglCopyTexImage1D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border);
void (APIENTRYP pglCopyTexImage2D)(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
void (APIENTRYP pglCopyTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width);
void (APIENTRYP pglCopyTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
void (APIENTRYP pglCopyTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
GLuint (APIENTRYP pglCreateProgram)();
GLuint (APIENTRYP pglCreateShader)(GLenum type);
void (APIENTRYP pglCullFace)(GLenum mode);
void (APIENTRYP pglDeleteBuffers)(GLsizei n, const GLuint* buffers);
void (APIENTRYP pglDeleteFramebuffers)(GLsizei n, const GLuint* framebuffers);
void (APIENTRYP pglDeleteLists)(GLuint list, GLsizei range);
void (APIENTRYP pglDeleteProgram)(GLuint program);
void (APIENTRYP pglDeleteQueries)(GLsizei n, const GLuint* ids);
void (APIENTRYP pglDeleteRenderbuffers)(GLsizei n, const GLuint* renderbuffers);
void (APIENTRYP pglDeleteShader)(GLuint shader);
void (APIENTRYP pglDeleteTextures)(GLsizei n, const GLuint* textures);
void (APIENTRYP pglDepthFunc)(GLenum func);
void (APIENTRYP pglDepthMask)(GLboolean flag);
void (APIENTRYP pglDepthRange)(GLdouble near, GLdouble far);
void (APIENTRYP pglDepthRangef)(GLfloat n, GLfloat f);
void (APIENTRYP pglDepthRangex)(GLfixed n, GLfixed f);
void (APIENTRYP pglDetachShader)(GLuint program, GLuint shader);
void (APIENTRYP pglDisable)(GLenum cap);
void (APIENTRYP pglDisableClientState)(GLenum array);
void (APIENTRYP pglDisableVertexAttribArray)(GLuint index);
void (APIENTRYP pglDrawArrays)(GLenum mode, GLint first, GLsizei count);
void (APIENTRYP pglDrawBuffer)(GLenum mode);
void (APIENTRYP pglDrawBuffers)(GLsizei n, const GLenum* bufs);
void (APIENTRYP pglDrawElements)(GLenum mode, GLsizei count, GLenum type, const GLvoid* indices);
void (APIENTRYP pglDrawPixels)(GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglDrawRangeElements)(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, const GLvoid* indices);
void (APIENTRYP pglEdgeFlag)(GLboolean flag);
void (APIENTRYP pglEdgeFlagPointer)(GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglEdgeFlagv)(const GLboolean* flag);
void (APIENTRYP pglEnable)(GLenum cap);
void (APIENTRYP pglEnableClientState)(GLenum array);
void (APIENTRYP pglEnableVertexAttribArray)(GLuint index);
void (APIENTRYP pglEnd)();
void (APIENTRYP pglEndList)();
void (APIENTRYP pglEndQuery)(GLenum target);
void (APIENTRYP pglEvalCoord1d)(GLdouble u);
void (APIENTRYP pglEvalCoord1dv)(const GLdouble* u);
void (APIENTRYP pglEvalCoord1f)(GLfloat u);
void (APIENTRYP pglEvalCoord1fv)(const GLfloat* u);
void (APIENTRYP pglEvalCoord2d)(GLdouble u, GLdouble v);
void (APIENTRYP pglEvalCoord2dv)(const GLdouble* u);
void (APIENTRYP pglEvalCoord2f)(GLfloat u, GLfloat v);
void (APIENTRYP pglEvalCoord2fv)(const GLfloat* u);
void (APIENTRYP pglEvalMesh1)(GLenum mode, GLint i1, GLint i2);
void (APIENTRYP pglEvalMesh2)(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2);
void (APIENTRYP pglEvalPoint1)(GLint i);
void (APIENTRYP pglEvalPoint2)(GLint i, GLint j);
void (APIENTRYP pglFeedbackBuffer)(GLsizei size, GLenum type, GLfloat* buffer);
void (APIENTRYP pglFinish)();
void (APIENTRYP pglFlush)();
void (APIENTRYP pglFogCoordPointer)(GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglFogCoordd)(GLdouble coord);
void (APIENTRYP pglFogCoorddv)(const GLdouble* coord);
void (APIENTRYP pglFogCoordf)(GLfloat coord);
void (APIENTRYP pglFogCoordfv)(const GLfloat* coord);
void (APIENTRYP pglFogf)(GLenum pname, GLfloat param);
void (APIENTRYP pglFogfv)(GLenum pname, const GLfloat* params);
void (APIENTRYP pglFogi)(GLenum pname, GLint param);
void (APIENTRYP pglFogiv)(GLenum pname, const GLint* params);
void (APIENTRYP pglFogx)(GLenum pname, GLfixed param);
void (APIENTRYP pglFogxv)(GLenum pname, const GLfixed* param);
void (APIENTRYP pglFramebufferRenderbuffer)(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
void (APIENTRYP pglFramebufferTexture2D)(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
void (APIENTRYP pglFrontFace)(GLenum mode);
void (APIENTRYP pglFrustum)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
void (APIENTRYP pglFrustumf)(GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f);
void (APIENTRYP pglFrustumx)(GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f);
void (APIENTRYP pglGenBuffers)(GLsizei n, GLuint* buffers);
void (APIENTRYP pglGenFramebuffers)(GLsizei n, GLuint* framebuffers);
GLuint (APIENTRYP pglGenLists)(GLsizei range);
void (APIENTRYP pglGenQueries)(GLsizei n, GLuint* ids);
void (APIENTRYP pglGenRenderbuffers)(GLsizei n, GLuint* renderbuffers);
void (APIENTRYP pglGenTextures)(GLsizei n, GLuint* textures);
void (APIENTRYP pglGenerateMipmap)(GLenum target);
void (APIENTRYP pglGetActiveAttrib)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
void (APIENTRYP pglGetActiveUniform)(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
void (APIENTRYP pglGetAttachedShaders)(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders);
GLint (APIENTRYP pglGetAttribLocation)(GLuint program, const GLchar* name);
void (APIENTRYP pglGetBooleanv)(GLenum pname, GLboolean* params);
void (APIENTRYP pglGetBufferParameteriv)(GLenum target, GLenum pname, GLint* params);
void (APIENTRYP pglGetBufferPointerv)(GLenum target, GLenum pname, GLvoid** params);
void (APIENTRYP pglGetBufferSubData)(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data);
void (APIENTRYP pglGetClipPlane)(GLenum plane, GLdouble* equation);
void (APIENTRYP pglGetClipPlanef)(GLenum plane, GLfloat* equation);
void (APIENTRYP pglGetClipPlanex)(GLenum plane, GLfixed* equation);
void (APIENTRYP pglGetCompressedTexImage)(GLenum target, GLint level, GLvoid* img);
void (APIENTRYP pglGetDoublev)(GLenum pname, GLdouble* params);
GLenum (APIENTRYP pglGetError)();
void (APIENTRYP pglGetFixedv)(GLenum pname, GLfixed* params);
void (APIENTRYP pglGetFloatv)(GLenum pname, GLfloat* params);
void (APIENTRYP pglGetFramebufferAttachmentParameteriv)(GLenum target, GLenum attachment, GLenum pname, GLint* params);
void (APIENTRYP pglGetIntegerv)(GLenum pname, GLint* params);
void (APIENTRYP pglGetLightfv)(GLenum light, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetLightiv)(GLenum light, GLenum pname, GLint* params);
void (APIENTRYP pglGetLightxv)(GLenum light, GLenum pname, GLfixed* params);
void (APIENTRYP pglGetMapdv)(GLenum target, GLenum query, GLdouble* v);
void (APIENTRYP pglGetMapfv)(GLenum target, GLenum query, GLfloat* v);
void (APIENTRYP pglGetMapiv)(GLenum target, GLenum query, GLint* v);
void (APIENTRYP pglGetMaterialfv)(GLenum face, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetMaterialiv)(GLenum face, GLenum pname, GLint* params);
void (APIENTRYP pglGetMaterialxv)(GLenum face, GLenum pname, GLfixed* params);
void (APIENTRYP pglGetPixelMapfv)(GLenum map, GLfloat* values);
void (APIENTRYP pglGetPixelMapuiv)(GLenum map, GLuint* values);
void (APIENTRYP pglGetPixelMapusv)(GLenum map, GLushort* values);
void (APIENTRYP pglGetPointerv)(GLenum pname, GLvoid** params);
void (APIENTRYP pglGetPolygonStipple)(GLubyte* mask);
void (APIENTRYP pglGetProgramInfoLog)(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
void (APIENTRYP pglGetProgramiv)(GLuint program, GLenum pname, GLint* params);
void (APIENTRYP pglGetQueryObjectiv)(GLuint id, GLenum pname, GLint* params);
void (APIENTRYP pglGetQueryObjectuiv)(GLuint id, GLenum pname, GLuint* params);
void (APIENTRYP pglGetQueryiv)(GLenum target, GLenum pname, GLint* params);
void (APIENTRYP pglGetRenderbufferParameteriv)(GLenum target, GLenum pname, GLint* params);
void (APIENTRYP pglGetShaderInfoLog)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog);
void (APIENTRYP pglGetShaderPrecisionFormat)(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision);
void (APIENTRYP pglGetShaderSource)(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source);
void (APIENTRYP pglGetShaderiv)(GLuint shader, GLenum pname, GLint* params);
const GLubyte* (APIENTRYP pglGetString)(GLenum name);
void (APIENTRYP pglGetTexEnvfv)(GLenum target, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetTexEnviv)(GLenum target, GLenum pname, GLint* params);
void (APIENTRYP pglGetTexEnvxv)(GLenum target, GLenum pname, GLfixed* params);
void (APIENTRYP pglGetTexGendv)(GLenum coord, GLenum pname, GLdouble* params);
void (APIENTRYP pglGetTexGenfv)(GLenum coord, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetTexGeniv)(GLenum coord, GLenum pname, GLint* params);
void (APIENTRYP pglGetTexImage)(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels);
void (APIENTRYP pglGetTexLevelParameterfv)(GLenum target, GLint level, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetTexLevelParameteriv)(GLenum target, GLint level, GLenum pname, GLint* params);
void (APIENTRYP pglGetTexParameterfv)(GLenum target, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetTexParameteriv)(GLenum target, GLenum pname, GLint* params);
void (APIENTRYP pglGetTexParameterxv)(GLenum target, GLenum pname, GLfixed* params);
GLint (APIENTRYP pglGetUniformLocation)(GLuint program, const GLchar* name);
void (APIENTRYP pglGetUniformfv)(GLuint program, GLint location, GLfloat* params);
void (APIENTRYP pglGetUniformiv)(GLuint program, GLint location, GLint* params);
void (APIENTRYP pglGetVertexAttribPointerv)(GLuint index, GLenum pname, GLvoid** pointer);
void (APIENTRYP pglGetVertexAttribdv)(GLuint index, GLenum pname, GLdouble* params);
void (APIENTRYP pglGetVertexAttribfv)(GLuint index, GLenum pname, GLfloat* params);
void (APIENTRYP pglGetVertexAttribiv)(GLuint index, GLenum pname, GLint* params);
void (APIENTRYP pglHint)(GLenum target, GLenum mode);
void (APIENTRYP pglIndexMask)(GLuint mask);
void (APIENTRYP pglIndexPointer)(GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglIndexd)(GLdouble c);
void (APIENTRYP pglIndexdv)(const GLdouble* c);
void (APIENTRYP pglIndexf)(GLfloat c);
void (APIENTRYP pglIndexfv)(const GLfloat* c);
void (APIENTRYP pglIndexi)(GLint c);
void (APIENTRYP pglIndexiv)(const GLint* c);
void (APIENTRYP pglIndexs)(GLshort c);
void (APIENTRYP pglIndexsv)(const GLshort* c);
void (APIENTRYP pglIndexub)(GLubyte c);
void (APIENTRYP pglIndexubv)(const GLubyte* c);
void (APIENTRYP pglInitNames)();
void (APIENTRYP pglInterleavedArrays)(GLenum format, GLsizei stride, const GLvoid* pointer);
GLboolean (APIENTRYP pglIsBuffer)(GLuint buffer);
GLboolean (APIENTRYP pglIsEnabled)(GLenum cap);
GLboolean (APIENTRYP pglIsFramebuffer)(GLuint framebuffer);
GLboolean (APIENTRYP pglIsList)(GLuint list);
GLboolean (APIENTRYP pglIsProgram)(GLuint program);
GLboolean (APIENTRYP pglIsQuery)(GLuint id);
GLboolean (APIENTRYP pglIsRenderbuffer)(GLuint renderbuffer);
GLboolean (APIENTRYP pglIsShader)(GLuint shader);
GLboolean (APIENTRYP pglIsTexture)(GLuint texture);
void (APIENTRYP pglLightModelf)(GLenum pname, GLfloat param);
void (APIENTRYP pglLightModelfv)(GLenum pname, const GLfloat* params);
void (APIENTRYP pglLightModeli)(GLenum pname, GLint param);
void (APIENTRYP pglLightModeliv)(GLenum pname, const GLint* params);
void (APIENTRYP pglLightModelx)(GLenum pname, GLfixed param);
void (APIENTRYP pglLightModelxv)(GLenum pname, const GLfixed* param);
void (APIENTRYP pglLightf)(GLenum light, GLenum pname, GLfloat param);
void (APIENTRYP pglLightfv)(GLenum light, GLenum pname, const GLfloat* params);
void (APIENTRYP pglLighti)(GLenum light, GLenum pname, GLint param);
void (APIENTRYP pglLightiv)(GLenum light, GLenum pname, const GLint* params);
void (APIENTRYP pglLightx)(GLenum light, GLenum pname, GLfixed param);
void (APIENTRYP pglLightxv)(GLenum light, GLenum pname, const GLfixed* params);
void (APIENTRYP pglLineStipple)(GLint factor, GLushort pattern);
void (APIENTRYP pglLineWidth)(GLfloat width);
void (APIENTRYP pglLineWidthx)(GLfixed width);
void (APIENTRYP pglLinkProgram)(GLuint program);
void (APIENTRYP pglListBase)(GLuint base);
void (APIENTRYP pglLoadIdentity)();
void (APIENTRYP pglLoadMatrixd)(const GLdouble* m);
void (APIENTRYP pglLoadMatrixf)(const GLfloat* m);
void (APIENTRYP pglLoadMatrixx)(const GLfixed* m);
void (APIENTRYP pglLoadName)(GLuint name);
void (APIENTRYP pglLoadTransposeMatrixd)(const GLdouble* m);
void (APIENTRYP pglLoadTransposeMatrixf)(const GLfloat* m);
void (APIENTRYP pglLogicOp)(GLenum opcode);
void (APIENTRYP pglMap1d)(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, const GLdouble* points);
void (APIENTRYP pglMap1f)(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, const GLfloat* points);
void (APIENTRYP pglMap2d)(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, const GLdouble* points);
void (APIENTRYP pglMap2f)(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, const GLfloat* points);
void* (APIENTRYP pglMapBuffer)(GLenum target, GLenum access);
void (APIENTRYP pglMapGrid1d)(GLint un, GLdouble u1, GLdouble u2);
void (APIENTRYP pglMapGrid1f)(GLint un, GLfloat u1, GLfloat u2);
void (APIENTRYP pglMapGrid2d)(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2);
void (APIENTRYP pglMapGrid2f)(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2);
void (APIENTRYP pglMaterialf)(GLenum face, GLenum pname, GLfloat param);
void (APIENTRYP pglMaterialfv)(GLenum face, GLenum pname, const GLfloat* params);
void (APIENTRYP pglMateriali)(GLenum face, GLenum pname, GLint param);
void (APIENTRYP pglMaterialiv)(GLenum face, GLenum pname, const GLint* params);
void (APIENTRYP pglMaterialx)(GLenum face, GLenum pname, GLfixed param);
void (APIENTRYP pglMaterialxv)(GLenum face, GLenum pname, const GLfixed* param);
void (APIENTRYP pglMatrixMode)(GLenum mode);
void (APIENTRYP pglMultMatrixd)(const GLdouble* m);
void (APIENTRYP pglMultMatrixf)(const GLfloat* m);
void (APIENTRYP pglMultMatrixx)(const GLfixed* m);
void (APIENTRYP pglMultTransposeMatrixd)(const GLdouble* m);
void (APIENTRYP pglMultTransposeMatrixf)(const GLfloat* m);
void (APIENTRYP pglMultiDrawArrays)(GLenum mode, const GLint* first, const GLsizei* count, GLsizei drawcount);
void (APIENTRYP pglMultiDrawElements)(GLenum mode, const GLsizei* count, GLenum type, const GLvoid** indices, GLsizei drawcount);
void (APIENTRYP pglMultiTexCoord1d)(GLenum target, GLdouble s);
void (APIENTRYP pglMultiTexCoord1dv)(GLenum target, const GLdouble* v);
void (APIENTRYP pglMultiTexCoord1f)(GLenum target, GLfloat s);
void (APIENTRYP pglMultiTexCoord1fv)(GLenum target, const GLfloat* v);
void (APIENTRYP pglMultiTexCoord1i)(GLenum target, GLint s);
void (APIENTRYP pglMultiTexCoord1iv)(GLenum target, const GLint* v);
void (APIENTRYP pglMultiTexCoord1s)(GLenum target, GLshort s);
void (APIENTRYP pglMultiTexCoord1sv)(GLenum target, const GLshort* v);
void (APIENTRYP pglMultiTexCoord2d)(GLenum target, GLdouble s, GLdouble t);
void (APIENTRYP pglMultiTexCoord2dv)(GLenum target, const GLdouble* v);
void (APIENTRYP pglMultiTexCoord2f)(GLenum target, GLfloat s, GLfloat t);
void (APIENTRYP pglMultiTexCoord2fv)(GLenum target, const GLfloat* v);
void (APIENTRYP pglMultiTexCoord2i)(GLenum target, GLint s, GLint t);
void (APIENTRYP pglMultiTexCoord2iv)(GLenum target, const GLint* v);
void (APIENTRYP pglMultiTexCoord2s)(GLenum target, GLshort s, GLshort t);
void (APIENTRYP pglMultiTexCoord2sv)(GLenum target, const GLshort* v);
void (APIENTRYP pglMultiTexCoord3d)(GLenum target, GLdouble s, GLdouble t, GLdouble r);
void (APIENTRYP pglMultiTexCoord3dv)(GLenum target, const GLdouble* v);
void (APIENTRYP pglMultiTexCoord3f)(GLenum target, GLfloat s, GLfloat t, GLfloat r);
void (APIENTRYP pglMultiTexCoord3fv)(GLenum target, const GLfloat* v);
void (APIENTRYP pglMultiTexCoord3i)(GLenum target, GLint s, GLint t, GLint r);
void (APIENTRYP pglMultiTexCoord3iv)(GLenum target, const GLint* v);
void (APIENTRYP pglMultiTexCoord3s)(GLenum target, GLshort s, GLshort t, GLshort r);
void (APIENTRYP pglMultiTexCoord3sv)(GLenum target, const GLshort* v);
void (APIENTRYP pglMultiTexCoord4d)(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q);
void (APIENTRYP pglMultiTexCoord4dv)(GLenum target, const GLdouble* v);
void (APIENTRYP pglMultiTexCoord4f)(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q);
void (APIENTRYP pglMultiTexCoord4fv)(GLenum target, const GLfloat* v);
void (APIENTRYP pglMultiTexCoord4i)(GLenum target, GLint s, GLint t, GLint r, GLint q);
void (APIENTRYP pglMultiTexCoord4iv)(GLenum target, const GLint* v);
void (APIENTRYP pglMultiTexCoord4s)(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q);
void (APIENTRYP pglMultiTexCoord4sv)(GLenum target, const GLshort* v);
void (APIENTRYP pglMultiTexCoord4x)(GLenum texture, GLfixed s, GLfixed t, GLfixed r, GLfixed q);
void (APIENTRYP pglNewList)(GLuint list, GLenum mode);
void (APIENTRYP pglNormal3b)(GLbyte nx, GLbyte ny, GLbyte nz);
void (APIENTRYP pglNormal3bv)(const GLbyte* v);
void (APIENTRYP pglNormal3d)(GLdouble nx, GLdouble ny, GLdouble nz);
void (APIENTRYP pglNormal3dv)(const GLdouble* v);
void (APIENTRYP pglNormal3f)(GLfloat nx, GLfloat ny, GLfloat nz);
void (APIENTRYP pglNormal3fv)(const GLfloat* v);
void (APIENTRYP pglNormal3i)(GLint nx, GLint ny, GLint nz);
void (APIENTRYP pglNormal3iv)(const GLint* v);
void (APIENTRYP pglNormal3s)(GLshort nx, GLshort ny, GLshort nz);
void (APIENTRYP pglNormal3sv)(const GLshort* v);
void (APIENTRYP pglNormal3x)(GLfixed nx, GLfixed ny, GLfixed nz);
void (APIENTRYP pglNormalPointer)(GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglOrtho)(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar);
void (APIENTRYP pglOrthof)(GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f);
void (APIENTRYP pglOrthox)(GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f);
void (APIENTRYP pglPassThrough)(GLfloat token);
void (APIENTRYP pglPixelMapfv)(GLenum map, GLsizei mapsize, const GLfloat* values);
void (APIENTRYP pglPixelMapuiv)(GLenum map, GLsizei mapsize, const GLuint* values);
void (APIENTRYP pglPixelMapusv)(GLenum map, GLsizei mapsize, const GLushort* values);
void (APIENTRYP pglPixelStoref)(GLenum pname, GLfloat param);
void (APIENTRYP pglPixelStorei)(GLenum pname, GLint param);
void (APIENTRYP pglPixelTransferf)(GLenum pname, GLfloat param);
void (APIENTRYP pglPixelTransferi)(GLenum pname, GLint param);
void (APIENTRYP pglPixelZoom)(GLfloat xfactor, GLfloat yfactor);
void (APIENTRYP pglPointParameterf)(GLenum pname, GLfloat param);
void (APIENTRYP pglPointParameterfv)(GLenum pname, const GLfloat* params);
void (APIENTRYP pglPointParameteri)(GLenum pname, GLint param);
void (APIENTRYP pglPointParameteriv)(GLenum pname, const GLint* params);
void (APIENTRYP pglPointParameterx)(GLenum pname, GLfixed param);
void (APIENTRYP pglPointParameterxv)(GLenum pname, const GLfixed* params);
void (APIENTRYP pglPointSize)(GLfloat size);
void (APIENTRYP pglPointSizex)(GLfixed size);
void (APIENTRYP pglPolygonMode)(GLenum face, GLenum mode);
void (APIENTRYP pglPolygonOffset)(GLfloat factor, GLfloat units);
void (APIENTRYP pglPolygonOffsetx)(GLfixed factor, GLfixed units);
void (APIENTRYP pglPolygonStipple)(const GLubyte* mask);
void (APIENTRYP pglPopAttrib)();
void (APIENTRYP pglPopClientAttrib)();
void (APIENTRYP pglPopMatrix)();
void (APIENTRYP pglPopName)();
void (APIENTRYP pglPrioritizeTextures)(GLsizei n, const GLuint* textures, const GLfloat* priorities);
void (APIENTRYP pglPushAttrib)(GLbitfield mask);
void (APIENTRYP pglPushClientAttrib)(GLbitfield mask);
void (APIENTRYP pglPushMatrix)();
void (APIENTRYP pglPushName)(GLuint name);
void (APIENTRYP pglRasterPos2d)(GLdouble x, GLdouble y);
void (APIENTRYP pglRasterPos2dv)(const GLdouble* v);
void (APIENTRYP pglRasterPos2f)(GLfloat x, GLfloat y);
void (APIENTRYP pglRasterPos2fv)(const GLfloat* v);
void (APIENTRYP pglRasterPos2i)(GLint x, GLint y);
void (APIENTRYP pglRasterPos2iv)(const GLint* v);
void (APIENTRYP pglRasterPos2s)(GLshort x, GLshort y);
void (APIENTRYP pglRasterPos2sv)(const GLshort* v);
void (APIENTRYP pglRasterPos3d)(GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglRasterPos3dv)(const GLdouble* v);
void (APIENTRYP pglRasterPos3f)(GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglRasterPos3fv)(const GLfloat* v);
void (APIENTRYP pglRasterPos3i)(GLint x, GLint y, GLint z);
void (APIENTRYP pglRasterPos3iv)(const GLint* v);
void (APIENTRYP pglRasterPos3s)(GLshort x, GLshort y, GLshort z);
void (APIENTRYP pglRasterPos3sv)(const GLshort* v);
void (APIENTRYP pglRasterPos4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
void (APIENTRYP pglRasterPos4dv)(const GLdouble* v);
void (APIENTRYP pglRasterPos4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
void (APIENTRYP pglRasterPos4fv)(const GLfloat* v);
void (APIENTRYP pglRasterPos4i)(GLint x, GLint y, GLint z, GLint w);
void (APIENTRYP pglRasterPos4iv)(const GLint* v);
void (APIENTRYP pglRasterPos4s)(GLshort x, GLshort y, GLshort z, GLshort w);
void (APIENTRYP pglRasterPos4sv)(const GLshort* v);
void (APIENTRYP pglReadBuffer)(GLenum mode);
void (APIENTRYP pglReadPixels)(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
void (APIENTRYP pglRectd)(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2);
void (APIENTRYP pglRectdv)(const GLdouble* v1, const GLdouble* v2);
void (APIENTRYP pglRectf)(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2);
void (APIENTRYP pglRectfv)(const GLfloat* v1, const GLfloat* v2);
void (APIENTRYP pglRecti)(GLint x1, GLint y1, GLint x2, GLint y2);
void (APIENTRYP pglRectiv)(const GLint* v1, const GLint* v2);
void (APIENTRYP pglRects)(GLshort x1, GLshort y1, GLshort x2, GLshort y2);
void (APIENTRYP pglRectsv)(const GLshort* v1, const GLshort* v2);
void (APIENTRYP pglReleaseShaderCompiler)();
GLint (APIENTRYP pglRenderMode)(GLenum mode);
void (APIENTRYP pglRenderbufferStorage)(GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
void (APIENTRYP pglRotated)(GLdouble angle, GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglRotatef)(GLfloat angle, GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglRotatex)(GLfixed angle, GLfixed x, GLfixed y, GLfixed z);
void (APIENTRYP pglSampleCoverage)(GLfloat value, GLboolean invert);
void (APIENTRYP pglSampleCoveragex)(GLclampx value, GLboolean invert);
void (APIENTRYP pglScaled)(GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglScalef)(GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglScalex)(GLfixed x, GLfixed y, GLfixed z);
void (APIENTRYP pglScissor)(GLint x, GLint y, GLsizei width, GLsizei height);
void (APIENTRYP pglSecondaryColor3b)(GLbyte red, GLbyte green, GLbyte blue);
void (APIENTRYP pglSecondaryColor3bv)(const GLbyte* v);
void (APIENTRYP pglSecondaryColor3d)(GLdouble red, GLdouble green, GLdouble blue);
void (APIENTRYP pglSecondaryColor3dv)(const GLdouble* v);
void (APIENTRYP pglSecondaryColor3f)(GLfloat red, GLfloat green, GLfloat blue);
void (APIENTRYP pglSecondaryColor3fv)(const GLfloat* v);
void (APIENTRYP pglSecondaryColor3i)(GLint red, GLint green, GLint blue);
void (APIENTRYP pglSecondaryColor3iv)(const GLint* v);
void (APIENTRYP pglSecondaryColor3s)(GLshort red, GLshort green, GLshort blue);
void (APIENTRYP pglSecondaryColor3sv)(const GLshort* v);
void (APIENTRYP pglSecondaryColor3ub)(GLubyte red, GLubyte green, GLubyte blue);
void (APIENTRYP pglSecondaryColor3ubv)(const GLubyte* v);
void (APIENTRYP pglSecondaryColor3ui)(GLuint red, GLuint green, GLuint blue);
void (APIENTRYP pglSecondaryColor3uiv)(const GLuint* v);
void (APIENTRYP pglSecondaryColor3us)(GLushort red, GLushort green, GLushort blue);
void (APIENTRYP pglSecondaryColor3usv)(const GLushort* v);
void (APIENTRYP pglSecondaryColorPointer)(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglSelectBuffer)(GLsizei size, GLuint* buffer);
void (APIENTRYP pglShadeModel)(GLenum mode);
void (APIENTRYP pglShaderBinary)(GLsizei count, const GLuint* shaders, GLenum binaryformat, const GLvoid* binary, GLsizei length);
void (APIENTRYP pglShaderSource)(GLuint shader, GLsizei count, const GLchar** string, const GLint* length);
void (APIENTRYP pglStencilFunc)(GLenum func, GLint ref, GLuint mask);
void (APIENTRYP pglStencilFuncSeparate)(GLenum face, GLenum func, GLint ref, GLuint mask);
void (APIENTRYP pglStencilMask)(GLuint mask);
void (APIENTRYP pglStencilMaskSeparate)(GLenum face, GLuint mask);
void (APIENTRYP pglStencilOp)(GLenum fail, GLenum zfail, GLenum zpass);
void (APIENTRYP pglStencilOpSeparate)(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass);
void (APIENTRYP pglTexCoord1d)(GLdouble s);
void (APIENTRYP pglTexCoord1dv)(const GLdouble* v);
void (APIENTRYP pglTexCoord1f)(GLfloat s);
void (APIENTRYP pglTexCoord1fv)(const GLfloat* v);
void (APIENTRYP pglTexCoord1i)(GLint s);
void (APIENTRYP pglTexCoord1iv)(const GLint* v);
void (APIENTRYP pglTexCoord1s)(GLshort s);
void (APIENTRYP pglTexCoord1sv)(const GLshort* v);
void (APIENTRYP pglTexCoord2d)(GLdouble s, GLdouble t);
void (APIENTRYP pglTexCoord2dv)(const GLdouble* v);
void (APIENTRYP pglTexCoord2f)(GLfloat s, GLfloat t);
void (APIENTRYP pglTexCoord2fv)(const GLfloat* v);
void (APIENTRYP pglTexCoord2i)(GLint s, GLint t);
void (APIENTRYP pglTexCoord2iv)(const GLint* v);
void (APIENTRYP pglTexCoord2s)(GLshort s, GLshort t);
void (APIENTRYP pglTexCoord2sv)(const GLshort* v);
void (APIENTRYP pglTexCoord3d)(GLdouble s, GLdouble t, GLdouble r);
void (APIENTRYP pglTexCoord3dv)(const GLdouble* v);
void (APIENTRYP pglTexCoord3f)(GLfloat s, GLfloat t, GLfloat r);
void (APIENTRYP pglTexCoord3fv)(const GLfloat* v);
void (APIENTRYP pglTexCoord3i)(GLint s, GLint t, GLint r);
void (APIENTRYP pglTexCoord3iv)(const GLint* v);
void (APIENTRYP pglTexCoord3s)(GLshort s, GLshort t, GLshort r);
void (APIENTRYP pglTexCoord3sv)(const GLshort* v);
void (APIENTRYP pglTexCoord4d)(GLdouble s, GLdouble t, GLdouble r, GLdouble q);
void (APIENTRYP pglTexCoord4dv)(const GLdouble* v);
void (APIENTRYP pglTexCoord4f)(GLfloat s, GLfloat t, GLfloat r, GLfloat q);
void (APIENTRYP pglTexCoord4fv)(const GLfloat* v);
void (APIENTRYP pglTexCoord4i)(GLint s, GLint t, GLint r, GLint q);
void (APIENTRYP pglTexCoord4iv)(const GLint* v);
void (APIENTRYP pglTexCoord4s)(GLshort s, GLshort t, GLshort r, GLshort q);
void (APIENTRYP pglTexCoord4sv)(const GLshort* v);
void (APIENTRYP pglTexCoordPointer)(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglTexEnvf)(GLenum target, GLenum pname, GLfloat param);
void (APIENTRYP pglTexEnvfv)(GLenum target, GLenum pname, const GLfloat* params);
void (APIENTRYP pglTexEnvi)(GLenum target, GLenum pname, GLint param);
void (APIENTRYP pglTexEnviv)(GLenum target, GLenum pname, const GLint* params);
void (APIENTRYP pglTexEnvx)(GLenum target, GLenum pname, GLfixed param);
void (APIENTRYP pglTexEnvxv)(GLenum target, GLenum pname, const GLfixed* params);
void (APIENTRYP pglTexGend)(GLenum coord, GLenum pname, GLdouble param);
void (APIENTRYP pglTexGendv)(GLenum coord, GLenum pname, const GLdouble* params);
void (APIENTRYP pglTexGenf)(GLenum coord, GLenum pname, GLfloat param);
void (APIENTRYP pglTexGenfv)(GLenum coord, GLenum pname, const GLfloat* params);
void (APIENTRYP pglTexGeni)(GLenum coord, GLenum pname, GLint param);
void (APIENTRYP pglTexGeniv)(GLenum coord, GLenum pname, const GLint* params);
void (APIENTRYP pglTexImage1D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTexImage2D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTexImage3D)(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTexParameterf)(GLenum target, GLenum pname, GLfloat param);
void (APIENTRYP pglTexParameterfv)(GLenum target, GLenum pname, const GLfloat* params);
void (APIENTRYP pglTexParameteri)(GLenum target, GLenum pname, GLint param);
void (APIENTRYP pglTexParameteriv)(GLenum target, GLenum pname, const GLint* params);
void (APIENTRYP pglTexParameterx)(GLenum target, GLenum pname, GLfixed param);
void (APIENTRYP pglTexParameterxv)(GLenum target, GLenum pname, const GLfixed* params);
void (APIENTRYP pglTexSubImage1D)(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTexSubImage2D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTexSubImage3D)(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid* pixels);
void (APIENTRYP pglTranslated)(GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglTranslatef)(GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglTranslatex)(GLfixed x, GLfixed y, GLfixed z);
void (APIENTRYP pglUniform1f)(GLint location, GLfloat v0);
void (APIENTRYP pglUniform1fv)(GLint location, GLsizei count, const GLfloat* value);
void (APIENTRYP pglUniform1i)(GLint location, GLint v0);
void (APIENTRYP pglUniform1iv)(GLint location, GLsizei count, const GLint* value);
void (APIENTRYP pglUniform2f)(GLint location, GLfloat v0, GLfloat v1);
void (APIENTRYP pglUniform2fv)(GLint location, GLsizei count, const GLfloat* value);
void (APIENTRYP pglUniform2i)(GLint location, GLint v0, GLint v1);
void (APIENTRYP pglUniform2iv)(GLint location, GLsizei count, const GLint* value);
void (APIENTRYP pglUniform3f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2);
void (APIENTRYP pglUniform3fv)(GLint location, GLsizei count, const GLfloat* value);
void (APIENTRYP pglUniform3i)(GLint location, GLint v0, GLint v1, GLint v2);
void (APIENTRYP pglUniform3iv)(GLint location, GLsizei count, const GLint* value);
void (APIENTRYP pglUniform4f)(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3);
void (APIENTRYP pglUniform4fv)(GLint location, GLsizei count, const GLfloat* value);
void (APIENTRYP pglUniform4i)(GLint location, GLint v0, GLint v1, GLint v2, GLint v3);
void (APIENTRYP pglUniform4iv)(GLint location, GLsizei count, const GLint* value);
void (APIENTRYP pglUniformMatrix2fv)(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
void (APIENTRYP pglUniformMatrix3fv)(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
void (APIENTRYP pglUniformMatrix4fv)(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
GLboolean (APIENTRYP pglUnmapBuffer)(GLenum target);
void (APIENTRYP pglUseProgram)(GLuint program);
void (APIENTRYP pglValidateProgram)(GLuint program);
void (APIENTRYP pglVertex2d)(GLdouble x, GLdouble y);
void (APIENTRYP pglVertex2dv)(const GLdouble* v);
void (APIENTRYP pglVertex2f)(GLfloat x, GLfloat y);
void (APIENTRYP pglVertex2fv)(const GLfloat* v);
void (APIENTRYP pglVertex2i)(GLint x, GLint y);
void (APIENTRYP pglVertex2iv)(const GLint* v);
void (APIENTRYP pglVertex2s)(GLshort x, GLshort y);
void (APIENTRYP pglVertex2sv)(const GLshort* v);
void (APIENTRYP pglVertex3d)(GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglVertex3dv)(const GLdouble* v);
void (APIENTRYP pglVertex3f)(GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglVertex3fv)(const GLfloat* v);
void (APIENTRYP pglVertex3i)(GLint x, GLint y, GLint z);
void (APIENTRYP pglVertex3iv)(const GLint* v);
void (APIENTRYP pglVertex3s)(GLshort x, GLshort y, GLshort z);
void (APIENTRYP pglVertex3sv)(const GLshort* v);
void (APIENTRYP pglVertex4d)(GLdouble x, GLdouble y, GLdouble z, GLdouble w);
void (APIENTRYP pglVertex4dv)(const GLdouble* v);
void (APIENTRYP pglVertex4f)(GLfloat x, GLfloat y, GLfloat z, GLfloat w);
void (APIENTRYP pglVertex4fv)(const GLfloat* v);
void (APIENTRYP pglVertex4i)(GLint x, GLint y, GLint z, GLint w);
void (APIENTRYP pglVertex4iv)(const GLint* v);
void (APIENTRYP pglVertex4s)(GLshort x, GLshort y, GLshort z, GLshort w);
void (APIENTRYP pglVertex4sv)(const GLshort* v);
void (APIENTRYP pglVertexAttrib1d)(GLuint index, GLdouble x);
void (APIENTRYP pglVertexAttrib1dv)(GLuint index, const GLdouble* v);
void (APIENTRYP pglVertexAttrib1f)(GLuint index, GLfloat x);
void (APIENTRYP pglVertexAttrib1fv)(GLuint index, const GLfloat* v);
void (APIENTRYP pglVertexAttrib1s)(GLuint index, GLshort x);
void (APIENTRYP pglVertexAttrib1sv)(GLuint index, const GLshort* v);
void (APIENTRYP pglVertexAttrib2d)(GLuint index, GLdouble x, GLdouble y);
void (APIENTRYP pglVertexAttrib2dv)(GLuint index, const GLdouble* v);
void (APIENTRYP pglVertexAttrib2f)(GLuint index, GLfloat x, GLfloat y);
void (APIENTRYP pglVertexAttrib2fv)(GLuint index, const GLfloat* v);
void (APIENTRYP pglVertexAttrib2s)(GLuint index, GLshort x, GLshort y);
void (APIENTRYP pglVertexAttrib2sv)(GLuint index, const GLshort* v);
void (APIENTRYP pglVertexAttrib3d)(GLuint index, GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglVertexAttrib3dv)(GLuint index, const GLdouble* v);
void (APIENTRYP pglVertexAttrib3f)(GLuint index, GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglVertexAttrib3fv)(GLuint index, const GLfloat* v);
void (APIENTRYP pglVertexAttrib3s)(GLuint index, GLshort x, GLshort y, GLshort z);
void (APIENTRYP pglVertexAttrib3sv)(GLuint index, const GLshort* v);
void (APIENTRYP pglVertexAttrib4Nbv)(GLuint index, const GLbyte* v);
void (APIENTRYP pglVertexAttrib4Niv)(GLuint index, const GLint* v);
void (APIENTRYP pglVertexAttrib4Nsv)(GLuint index, const GLshort* v);
void (APIENTRYP pglVertexAttrib4Nub)(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w);
void (APIENTRYP pglVertexAttrib4Nubv)(GLuint index, const GLubyte* v);
void (APIENTRYP pglVertexAttrib4Nuiv)(GLuint index, const GLuint* v);
void (APIENTRYP pglVertexAttrib4Nusv)(GLuint index, const GLushort* v);
void (APIENTRYP pglVertexAttrib4bv)(GLuint index, const GLbyte* v);
void (APIENTRYP pglVertexAttrib4d)(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w);
void (APIENTRYP pglVertexAttrib4dv)(GLuint index, const GLdouble* v);
void (APIENTRYP pglVertexAttrib4f)(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
void (APIENTRYP pglVertexAttrib4fv)(GLuint index, const GLfloat* v);
void (APIENTRYP pglVertexAttrib4iv)(GLuint index, const GLint* v);
void (APIENTRYP pglVertexAttrib4s)(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w);
void (APIENTRYP pglVertexAttrib4sv)(GLuint index, const GLshort* v);
void (APIENTRYP pglVertexAttrib4ubv)(GLuint index, const GLubyte* v);
void (APIENTRYP pglVertexAttrib4uiv)(GLuint index, const GLuint* v);
void (APIENTRYP pglVertexAttrib4usv)(GLuint index, const GLushort* v);
void (APIENTRYP pglVertexAttribPointer)(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglVertexPointer)(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer);
void (APIENTRYP pglViewport)(GLint x, GLint y, GLsizei width, GLsizei height);
void (APIENTRYP pglWindowPos2d)(GLdouble x, GLdouble y);
void (APIENTRYP pglWindowPos2dv)(const GLdouble* v);
void (APIENTRYP pglWindowPos2f)(GLfloat x, GLfloat y);
void (APIENTRYP pglWindowPos2fv)(const GLfloat* v);
void (APIENTRYP pglWindowPos2i)(GLint x, GLint y);
void (APIENTRYP pglWindowPos2iv)(const GLint* v);
void (APIENTRYP pglWindowPos2s)(GLshort x, GLshort y);
void (APIENTRYP pglWindowPos2sv)(const GLshort* v);
void (APIENTRYP pglWindowPos3d)(GLdouble x, GLdouble y, GLdouble z);
void (APIENTRYP pglWindowPos3dv)(const GLdouble* v);
void (APIENTRYP pglWindowPos3f)(GLfloat x, GLfloat y, GLfloat z);
void (APIENTRYP pglWindowPos3fv)(const GLfloat* v);
void (APIENTRYP pglWindowPos3i)(GLint x, GLint y, GLint z);
void (APIENTRYP pglWindowPos3iv)(const GLint* v);
void (APIENTRYP pglWindowPos3s)(GLshort x, GLshort y, GLshort z);
void (APIENTRYP pglWindowPos3sv)(const GLshort* v);

void goglAccum(GLenum op, GLfloat value) {
	(*pglAccum)(op, value);
}
void goglActiveTexture(GLenum texture) {
	(*pglActiveTexture)(texture);
}
void goglAlphaFunc(GLenum func, GLfloat ref) {
	(*pglAlphaFunc)(func, ref);
}
void goglAlphaFuncx(GLenum func, GLfixed ref) {
	(*pglAlphaFuncx)(func, ref);
}
GLboolean goglAreTexturesResident(GLsizei n, const GLuint* textures, GLboolean* residences) {
	return (*pglAreTexturesResident)(n, textures, residences);
}
void goglArrayElement(GLint i) {
	(*pglArrayElement)(i);
}
void goglAttachShader(GLuint program, GLuint shader) {
	(*pglAttachShader)(program, shader);
}
void goglBegin(GLenum mode) {
	(*pglBegin)(mode);
}
void goglBeginQuery(GLenum target, GLuint id) {
	(*pglBeginQuery)(target, id);
}
void goglBindAttribLocation(GLuint program, GLuint index, const GLchar* name) {
	(*pglBindAttribLocation)(program, index, name);
}
void goglBindBuffer(GLenum target, GLuint buffer) {
	(*pglBindBuffer)(target, buffer);
}
void goglBindFramebuffer(GLenum target, GLuint framebuffer) {
	(*pglBindFramebuffer)(target, framebuffer);
}
void goglBindRenderbuffer(GLenum target, GLuint renderbuffer) {
	(*pglBindRenderbuffer)(target, renderbuffer);
}
void goglBindTexture(GLenum target, GLuint texture) {
	(*pglBindTexture)(target, texture);
}
void goglBitmap(GLsizei width, GLsizei height, GLfloat xorig, GLfloat yorig, GLfloat xmove, GLfloat ymove, const GLubyte* bitmap) {
	(*pglBitmap)(width, height, xorig, yorig, xmove, ymove, bitmap);
}
void goglBlendColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	(*pglBlendColor)(red, green, blue, alpha);
}
void goglBlendEquation(GLenum mode) {
	(*pglBlendEquation)(mode);
}
void goglBlendEquationSeparate(GLenum modeRGB, GLenum modeAlpha) {
	(*pglBlendEquationSeparate)(modeRGB, modeAlpha);
}
void goglBlendFunc(GLenum sfactor, GLenum dfactor) {
	(*pglBlendFunc)(sfactor, dfactor);
}
void goglBlendFuncSeparate(GLenum sfactorRGB, GLenum dfactorRGB, GLenum sfactorAlpha, GLenum dfactorAlpha) {
	(*pglBlendFuncSeparate)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
}
void goglBufferData(GLenum target, GLsizeiptr size, const GLvoid* data, GLenum usage) {
	(*pglBufferData)(target, size, data, usage);
}
void goglBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, const GLvoid* data) {
	(*pglBufferSubData)(target, offset, size, data);
}
void goglCallList(GLuint list) {
	(*pglCallList)(list);
}
void goglCallLists(GLsizei n, GLenum type, const GLvoid* lists) {
	(*pglCallLists)(n, type, lists);
}
GLenum goglCheckFramebufferStatus(GLenum target) {
	return (*pglCheckFramebufferStatus)(target);
}
void goglClear(GLbitfield mask) {
	(*pglClear)(mask);
}
void goglClearAccum(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	(*pglClearAccum)(red, green, blue, alpha);
}
void goglClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	(*pglClearColor)(red, green, blue, alpha);
}
void goglClearColorx(GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha) {
	(*pglClearColorx)(red, green, blue, alpha);
}
void goglClearDepth(GLdouble depth) {
	(*pglClearDepth)(depth);
}
void goglClearDepthf(GLfloat d) {
	(*pglClearDepthf)(d);
}
void goglClearDepthx(GLfixed depth) {
	(*pglClearDepthx)(depth);
}
void goglClearIndex(GLfloat c) {
	(*pglClearIndex)(c);
}
void goglClearStencil(GLint s) {
	(*pglClearStencil)(s);
}
void goglClientActiveTexture(GLenum texture) {
	(*pglClientActiveTexture)(texture);
}
void goglClipPlane(GLenum plane, const GLdouble* equation) {
	(*pglClipPlane)(plane, equation);
}
void goglClipPlanef(GLenum p, const GLfloat* eqn) {
	(*pglClipPlanef)(p, eqn);
}
void goglClipPlanex(GLenum plane, const GLfixed* equation) {
	(*pglClipPlanex)(plane, equation);
}
void goglColor3b(GLbyte red, GLbyte green, GLbyte blue) {
	(*pglColor3b)(red, green, blue);
}
void goglColor3bv(const GLbyte* v) {
	(*pglColor3bv)(v);
}
void goglColor3d(GLdouble red, GLdouble green, GLdouble blue) {
	(*pglColor3d)(red, green, blue);
}
void goglColor3dv(const GLdouble* v) {
	(*pglColor3dv)(v);
}
void goglColor3f(GLfloat red, GLfloat green, GLfloat blue) {
	(*pglColor3f)(red, green, blue);
}
void goglColor3fv(const GLfloat* v) {
	(*pglColor3fv)(v);
}
void goglColor3i(GLint red, GLint green, GLint blue) {
	(*pglColor3i)(red, green, blue);
}
void goglColor3iv(const GLint* v) {
	(*pglColor3iv)(v);
}
void goglColor3s(GLshort red, GLshort green, GLshort blue) {
	(*pglColor3s)(red, green, blue);
}
void goglColor3sv(const GLshort* v) {
	(*pglColor3sv)(v);
}
void goglColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
	(*pglColor3ub)(red, green, blue);
}
void goglColor3ubv(const GLubyte* v) {
	(*pglColor3ubv)(v);
}
void goglColor3ui(GLuint red, GLuint green, GLuint blue) {
	(*pglColor3ui)(red, green, blue);
}
void goglColor3uiv(const GLuint* v) {
	(*pglColor3uiv)(v);
}
void goglColor3us(GLushort red, GLushort green, GLushort blue) {
	(*pglColor3us)(red, green, blue);
}
void goglColor3usv(const GLushort* v) {
	(*pglColor3usv)(v);
}
void goglColor4b(GLbyte red, GLbyte green, GLbyte blue, GLbyte alpha) {
	(*pglColor4b)(red, green, blue, alpha);
}
void goglColor4bv(const GLbyte* v) {
	(*pglColor4bv)(v);
}
void goglColor4d(GLdouble red, GLdouble green, GLdouble blue, GLdouble alpha) {
	(*pglColor4d)(red, green, blue, alpha);
}
void goglColor4dv(const GLdouble* v) {
	(*pglColor4dv)(v);
}
void goglColor4f(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha) {
	(*pglColor4f)(red, green, blue, alpha);
}
void goglColor4fv(const GLfloat* v) {
	(*pglColor4fv)(v);
}
void goglColor4i(GLint red, GLint green, GLint blue, GLint alpha) {
	(*pglColor4i)(red, green, blue, alpha);
}
void goglColor4iv(const GLint* v) {
	(*pglColor4iv)(v);
}
void goglColor4s(GLshort red, GLshort green, GLshort blue, GLshort alpha) {
	(*pglColor4s)(red, green, blue, alpha);
}
void goglColor4sv(const GLshort* v) {
	(*pglColor4sv)(v);
}
void goglColor4ub(GLubyte red, GLubyte green, GLubyte blue, GLubyte alpha) {
	(*pglColor4ub)(red, green, blue, alpha);
}
void goglColor4ubv(const GLubyte* v) {
	(*pglColor4ubv)(v);
}
void goglColor4ui(GLuint red, GLuint green, GLuint blue, GLuint alpha) {
	(*pglColor4ui)(red, green, blue, alpha);
}
void goglColor4uiv(const GLuint* v) {
	(*pglColor4uiv)(v);
}
void goglColor4us(GLushort red, GLushort green, GLushort blue, GLushort alpha) {
	(*pglColor4us)(red, green, blue, alpha);
}
void goglColor4usv(const GLushort* v) {
	(*pglColor4usv)(v);
}
void goglColor4x(GLfixed red, GLfixed green, GLfixed blue, GLfixed alpha) {
	(*pglColor4x)(red, green, blue, alpha);
}
void goglColorMask(GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha) {
	(*pglColorMask)(red, green, blue, alpha);
}
void goglColorMaterial(GLenum face, GLenum mode) {
	(*pglColorMaterial)(face, mode);
}
void goglColorPointer(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglColorPointer)(size, type, stride, pointer);
}
void goglCompileShader(GLuint shader) {
	(*pglCompileShader)(shader);
}
void goglCompressedTexImage1D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexImage1D)(target, level, internalformat, width, border, imageSize, data);
}
void goglCompressedTexImage2D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexImage2D)(target, level, internalformat, width, height, border, imageSize, data);
}
void goglCompressedTexImage3D(GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexImage3D)(target, level, internalformat, width, height, depth, border, imageSize, data);
}
void goglCompressedTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexSubImage1D)(target, level, xoffset, width, format, imageSize, data);
}
void goglCompressedTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
}
void goglCompressedTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid* data) {
	(*pglCompressedTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
}
void goglCopyPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum type) {
	(*pglCopyPixels)(x, y, width, height, type);
}
void goglCopyTexImage1D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLint border) {
	(*pglCopyTexImage1D)(target, level, internalformat, x, y, width, border);
}
void goglCopyTexImage2D(GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border) {
	(*pglCopyTexImage2D)(target, level, internalformat, x, y, width, height, border);
}
void goglCopyTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLint x, GLint y, GLsizei width) {
	(*pglCopyTexSubImage1D)(target, level, xoffset, x, y, width);
}
void goglCopyTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
	(*pglCopyTexSubImage2D)(target, level, xoffset, yoffset, x, y, width, height);
}
void goglCopyTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height) {
	(*pglCopyTexSubImage3D)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
}
GLuint goglCreateProgram() {
	return (*pglCreateProgram)();
}
GLuint goglCreateShader(GLenum type) {
	return (*pglCreateShader)(type);
}
void goglCullFace(GLenum mode) {
	(*pglCullFace)(mode);
}
void goglDeleteBuffers(GLsizei n, const GLuint* buffers) {
	(*pglDeleteBuffers)(n, buffers);
}
void goglDeleteFramebuffers(GLsizei n, const GLuint* framebuffers) {
	(*pglDeleteFramebuffers)(n, framebuffers);
}
void goglDeleteLists(GLuint list, GLsizei range) {
	(*pglDeleteLists)(list, range);
}
void goglDeleteProgram(GLuint program) {
	(*pglDeleteProgram)(program);
}
void goglDeleteQueries(GLsizei n, const GLuint* ids) {
	(*pglDeleteQueries)(n, ids);
}
void goglDeleteRenderbuffers(GLsizei n, const GLuint* renderbuffers) {
	(*pglDeleteRenderbuffers)(n, renderbuffers);
}
void goglDeleteShader(GLuint shader) {
	(*pglDeleteShader)(shader);
}
void goglDeleteTextures(GLsizei n, const GLuint* textures) {
	(*pglDeleteTextures)(n, textures);
}
void goglDepthFunc(GLenum func) {
	(*pglDepthFunc)(func);
}
void goglDepthMask(GLboolean flag) {
	(*pglDepthMask)(flag);
}
void goglDepthRange(GLdouble near, GLdouble far) {
	(*pglDepthRange)(near, far);
}
void goglDepthRangef(GLfloat n, GLfloat f) {
	(*pglDepthRangef)(n, f);
}
void goglDepthRangex(GLfixed n, GLfixed f) {
	(*pglDepthRangex)(n, f);
}
void goglDetachShader(GLuint program, GLuint shader) {
	(*pglDetachShader)(program, shader);
}
void goglDisable(GLenum cap) {
	(*pglDisable)(cap);
}
void goglDisableClientState(GLenum array) {
	(*pglDisableClientState)(array);
}
void goglDisableVertexAttribArray(GLuint index) {
	(*pglDisableVertexAttribArray)(index);
}
void goglDrawArrays(GLenum mode, GLint first, GLsizei count) {
	(*pglDrawArrays)(mode, first, count);
}
void goglDrawBuffer(GLenum mode) {
	(*pglDrawBuffer)(mode);
}
void goglDrawBuffers(GLsizei n, const GLenum* bufs) {
	(*pglDrawBuffers)(n, bufs);
}
void goglDrawElements(GLenum mode, GLsizei count, GLenum type, const GLvoid* indices) {
	(*pglDrawElements)(mode, count, type, indices);
}
void goglDrawPixels(GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglDrawPixels)(width, height, format, type, pixels);
}
void goglDrawRangeElements(GLenum mode, GLuint start, GLuint end, GLsizei count, GLenum type, const GLvoid* indices) {
	(*pglDrawRangeElements)(mode, start, end, count, type, indices);
}
void goglEdgeFlag(GLboolean flag) {
	(*pglEdgeFlag)(flag);
}
void goglEdgeFlagPointer(GLsizei stride, const GLvoid* pointer) {
	(*pglEdgeFlagPointer)(stride, pointer);
}
void goglEdgeFlagv(const GLboolean* flag) {
	(*pglEdgeFlagv)(flag);
}
void goglEnable(GLenum cap) {
	(*pglEnable)(cap);
}
void goglEnableClientState(GLenum array) {
	(*pglEnableClientState)(array);
}
void goglEnableVertexAttribArray(GLuint index) {
	(*pglEnableVertexAttribArray)(index);
}
void goglEnd() {
	(*pglEnd)();
}
void goglEndList() {
	(*pglEndList)();
}
void goglEndQuery(GLenum target) {
	(*pglEndQuery)(target);
}
void goglEvalCoord1d(GLdouble u) {
	(*pglEvalCoord1d)(u);
}
void goglEvalCoord1dv(const GLdouble* u) {
	(*pglEvalCoord1dv)(u);
}
void goglEvalCoord1f(GLfloat u) {
	(*pglEvalCoord1f)(u);
}
void goglEvalCoord1fv(const GLfloat* u) {
	(*pglEvalCoord1fv)(u);
}
void goglEvalCoord2d(GLdouble u, GLdouble v) {
	(*pglEvalCoord2d)(u, v);
}
void goglEvalCoord2dv(const GLdouble* u) {
	(*pglEvalCoord2dv)(u);
}
void goglEvalCoord2f(GLfloat u, GLfloat v) {
	(*pglEvalCoord2f)(u, v);
}
void goglEvalCoord2fv(const GLfloat* u) {
	(*pglEvalCoord2fv)(u);
}
void goglEvalMesh1(GLenum mode, GLint i1, GLint i2) {
	(*pglEvalMesh1)(mode, i1, i2);
}
void goglEvalMesh2(GLenum mode, GLint i1, GLint i2, GLint j1, GLint j2) {
	(*pglEvalMesh2)(mode, i1, i2, j1, j2);
}
void goglEvalPoint1(GLint i) {
	(*pglEvalPoint1)(i);
}
void goglEvalPoint2(GLint i, GLint j) {
	(*pglEvalPoint2)(i, j);
}
void goglFeedbackBuffer(GLsizei size, GLenum type, GLfloat* buffer) {
	(*pglFeedbackBuffer)(size, type, buffer);
}
void goglFinish() {
	(*pglFinish)();
}
void goglFlush() {
	(*pglFlush)();
}
void goglFogCoordPointer(GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglFogCoordPointer)(type, stride, pointer);
}
void goglFogCoordd(GLdouble coord) {
	(*pglFogCoordd)(coord);
}
void goglFogCoorddv(const GLdouble* coord) {
	(*pglFogCoorddv)(coord);
}
void goglFogCoordf(GLfloat coord) {
	(*pglFogCoordf)(coord);
}
void goglFogCoordfv(const GLfloat* coord) {
	(*pglFogCoordfv)(coord);
}
void goglFogf(GLenum pname, GLfloat param) {
	(*pglFogf)(pname, param);
}
void goglFogfv(GLenum pname, const GLfloat* params) {
	(*pglFogfv)(pname, params);
}
void goglFogi(GLenum pname, GLint param) {
	(*pglFogi)(pname, param);
}
void goglFogiv(GLenum pname, const GLint* params) {
	(*pglFogiv)(pname, params);
}
void goglFogx(GLenum pname, GLfixed param) {
	(*pglFogx)(pname, param);
}
void goglFogxv(GLenum pname, const GLfixed* param) {
	(*pglFogxv)(pname, param);
}
void goglFramebufferRenderbuffer(GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer) {
	(*pglFramebufferRenderbuffer)(target, attachment, renderbuffertarget, renderbuffer);
}
void goglFramebufferTexture2D(GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level) {
	(*pglFramebufferTexture2D)(target, attachment, textarget, texture, level);
}
void goglFrontFace(GLenum mode) {
	(*pglFrontFace)(mode);
}
void goglFrustum(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
	(*pglFrustum)(left, right, bottom, top, zNear, zFar);
}
void goglFrustumf(GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f) {
	(*pglFrustumf)(l, r, b, t, n, f);
}
void goglFrustumx(GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f) {
	(*pglFrustumx)(l, r, b, t, n, f);
}
void goglGenBuffers(GLsizei n, GLuint* buffers) {
	(*pglGenBuffers)(n, buffers);
}
void goglGenFramebuffers(GLsizei n, GLuint* framebuffers) {
	(*pglGenFramebuffers)(n, framebuffers);
}
GLuint goglGenLists(GLsizei range) {
	return (*pglGenLists)(range);
}
void goglGenQueries(GLsizei n, GLuint* ids) {
	(*pglGenQueries)(n, ids);
}
void goglGenRenderbuffers(GLsizei n, GLuint* renderbuffers) {
	(*pglGenRenderbuffers)(n, renderbuffers);
}
void goglGenTextures(GLsizei n, GLuint* textures) {
	(*pglGenTextures)(n, textures);
}
void goglGenerateMipmap(GLenum target) {
	(*pglGenerateMipmap)(target);
}
void goglGetActiveAttrib(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
	(*pglGetActiveAttrib)(program, index, bufSize, length, size, type, name);
}
void goglGetActiveUniform(GLuint program, GLuint index, GLsizei bufSize, GLsizei* length, GLint* size, GLenum* type, GLchar* name) {
	(*pglGetActiveUniform)(program, index, bufSize, length, size, type, name);
}
void goglGetAttachedShaders(GLuint program, GLsizei maxCount, GLsizei* count, GLuint* shaders) {
	(*pglGetAttachedShaders)(program, maxCount, count, shaders);
}
GLint goglGetAttribLocation(GLuint program, const GLchar* name) {
	return (*pglGetAttribLocation)(program, name);
}
void goglGetBooleanv(GLenum pname, GLboolean* params) {
	(*pglGetBooleanv)(pname, params);
}
void goglGetBufferParameteriv(GLenum target, GLenum pname, GLint* params) {
	(*pglGetBufferParameteriv)(target, pname, params);
}
void goglGetBufferPointerv(GLenum target, GLenum pname, GLvoid** params) {
	(*pglGetBufferPointerv)(target, pname, params);
}
void goglGetBufferSubData(GLenum target, GLintptr offset, GLsizeiptr size, GLvoid* data) {
	(*pglGetBufferSubData)(target, offset, size, data);
}
void goglGetClipPlane(GLenum plane, GLdouble* equation) {
	(*pglGetClipPlane)(plane, equation);
}
void goglGetClipPlanef(GLenum plane, GLfloat* equation) {
	(*pglGetClipPlanef)(plane, equation);
}
void goglGetClipPlanex(GLenum plane, GLfixed* equation) {
	(*pglGetClipPlanex)(plane, equation);
}
void goglGetCompressedTexImage(GLenum target, GLint level, GLvoid* img) {
	(*pglGetCompressedTexImage)(target, level, img);
}
void goglGetDoublev(GLenum pname, GLdouble* params) {
	(*pglGetDoublev)(pname, params);
}
GLenum goglGetError() {
	return (*pglGetError)();
}
void goglGetFixedv(GLenum pname, GLfixed* params) {
	(*pglGetFixedv)(pname, params);
}
void goglGetFloatv(GLenum pname, GLfloat* params) {
	(*pglGetFloatv)(pname, params);
}
void goglGetFramebufferAttachmentParameteriv(GLenum target, GLenum attachment, GLenum pname, GLint* params) {
	(*pglGetFramebufferAttachmentParameteriv)(target, attachment, pname, params);
}
void goglGetIntegerv(GLenum pname, GLint* params) {
	(*pglGetIntegerv)(pname, params);
}
void goglGetLightfv(GLenum light, GLenum pname, GLfloat* params) {
	(*pglGetLightfv)(light, pname, params);
}
void goglGetLightiv(GLenum light, GLenum pname, GLint* params) {
	(*pglGetLightiv)(light, pname, params);
}
void goglGetLightxv(GLenum light, GLenum pname, GLfixed* params) {
	(*pglGetLightxv)(light, pname, params);
}
void goglGetMapdv(GLenum target, GLenum query, GLdouble* v) {
	(*pglGetMapdv)(target, query, v);
}
void goglGetMapfv(GLenum target, GLenum query, GLfloat* v) {
	(*pglGetMapfv)(target, query, v);
}
void goglGetMapiv(GLenum target, GLenum query, GLint* v) {
	(*pglGetMapiv)(target, query, v);
}
void goglGetMaterialfv(GLenum face, GLenum pname, GLfloat* params) {
	(*pglGetMaterialfv)(face, pname, params);
}
void goglGetMaterialiv(GLenum face, GLenum pname, GLint* params) {
	(*pglGetMaterialiv)(face, pname, params);
}
void goglGetMaterialxv(GLenum face, GLenum pname, GLfixed* params) {
	(*pglGetMaterialxv)(face, pname, params);
}
void goglGetPixelMapfv(GLenum map, GLfloat* values) {
	(*pglGetPixelMapfv)(map, values);
}
void goglGetPixelMapuiv(GLenum map, GLuint* values) {
	(*pglGetPixelMapuiv)(map, values);
}
void goglGetPixelMapusv(GLenum map, GLushort* values) {
	(*pglGetPixelMapusv)(map, values);
}
void goglGetPointerv(GLenum pname, GLvoid** params) {
	(*pglGetPointerv)(pname, params);
}
void goglGetPolygonStipple(GLubyte* mask) {
	(*pglGetPolygonStipple)(mask);
}
void goglGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
	(*pglGetProgramInfoLog)(program, bufSize, length, infoLog);
}
void goglGetProgramiv(GLuint program, GLenum pname, GLint* params) {
	(*pglGetProgramiv)(program, pname, params);
}
void goglGetQueryObjectiv(GLuint id, GLenum pname, GLint* params) {
	(*pglGetQueryObjectiv)(id, pname, params);
}
void goglGetQueryObjectuiv(GLuint id, GLenum pname, GLuint* params) {
	(*pglGetQueryObjectuiv)(id, pname, params);
}
void goglGetQueryiv(GLenum target, GLenum pname, GLint* params) {
	(*pglGetQueryiv)(target, pname, params);
}
void goglGetRenderbufferParameteriv(GLenum target, GLenum pname, GLint* params) {
	(*pglGetRenderbufferParameteriv)(target, pname, params);
}
void goglGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* infoLog) {
	(*pglGetShaderInfoLog)(shader, bufSize, length, infoLog);
}
void goglGetShaderPrecisionFormat(GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision) {
	(*pglGetShaderPrecisionFormat)(shadertype, precisiontype, range, precision);
}
void goglGetShaderSource(GLuint shader, GLsizei bufSize, GLsizei* length, GLchar* source) {
	(*pglGetShaderSource)(shader, bufSize, length, source);
}
void goglGetShaderiv(GLuint shader, GLenum pname, GLint* params) {
	(*pglGetShaderiv)(shader, pname, params);
}
const GLubyte* goglGetString(GLenum name) {
	return (*pglGetString)(name);
}
void goglGetTexEnvfv(GLenum target, GLenum pname, GLfloat* params) {
	(*pglGetTexEnvfv)(target, pname, params);
}
void goglGetTexEnviv(GLenum target, GLenum pname, GLint* params) {
	(*pglGetTexEnviv)(target, pname, params);
}
void goglGetTexEnvxv(GLenum target, GLenum pname, GLfixed* params) {
	(*pglGetTexEnvxv)(target, pname, params);
}
void goglGetTexGendv(GLenum coord, GLenum pname, GLdouble* params) {
	(*pglGetTexGendv)(coord, pname, params);
}
void goglGetTexGenfv(GLenum coord, GLenum pname, GLfloat* params) {
	(*pglGetTexGenfv)(coord, pname, params);
}
void goglGetTexGeniv(GLenum coord, GLenum pname, GLint* params) {
	(*pglGetTexGeniv)(coord, pname, params);
}
void goglGetTexImage(GLenum target, GLint level, GLenum format, GLenum type, GLvoid* pixels) {
	(*pglGetTexImage)(target, level, format, type, pixels);
}
void goglGetTexLevelParameterfv(GLenum target, GLint level, GLenum pname, GLfloat* params) {
	(*pglGetTexLevelParameterfv)(target, level, pname, params);
}
void goglGetTexLevelParameteriv(GLenum target, GLint level, GLenum pname, GLint* params) {
	(*pglGetTexLevelParameteriv)(target, level, pname, params);
}
void goglGetTexParameterfv(GLenum target, GLenum pname, GLfloat* params) {
	(*pglGetTexParameterfv)(target, pname, params);
}
void goglGetTexParameteriv(GLenum target, GLenum pname, GLint* params) {
	(*pglGetTexParameteriv)(target, pname, params);
}
void goglGetTexParameterxv(GLenum target, GLenum pname, GLfixed* params) {
	(*pglGetTexParameterxv)(target, pname, params);
}
GLint goglGetUniformLocation(GLuint program, const GLchar* name) {
	return (*pglGetUniformLocation)(program, name);
}
void goglGetUniformfv(GLuint program, GLint location, GLfloat* params) {
	(*pglGetUniformfv)(program, location, params);
}
void goglGetUniformiv(GLuint program, GLint location, GLint* params) {
	(*pglGetUniformiv)(program, location, params);
}
void goglGetVertexAttribPointerv(GLuint index, GLenum pname, GLvoid** pointer) {
	(*pglGetVertexAttribPointerv)(index, pname, pointer);
}
void goglGetVertexAttribdv(GLuint index, GLenum pname, GLdouble* params) {
	(*pglGetVertexAttribdv)(index, pname, params);
}
void goglGetVertexAttribfv(GLuint index, GLenum pname, GLfloat* params) {
	(*pglGetVertexAttribfv)(index, pname, params);
}
void goglGetVertexAttribiv(GLuint index, GLenum pname, GLint* params) {
	(*pglGetVertexAttribiv)(index, pname, params);
}
void goglHint(GLenum target, GLenum mode) {
	(*pglHint)(target, mode);
}
void goglIndexMask(GLuint mask) {
	(*pglIndexMask)(mask);
}
void goglIndexPointer(GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglIndexPointer)(type, stride, pointer);
}
void goglIndexd(GLdouble c) {
	(*pglIndexd)(c);
}
void goglIndexdv(const GLdouble* c) {
	(*pglIndexdv)(c);
}
void goglIndexf(GLfloat c) {
	(*pglIndexf)(c);
}
void goglIndexfv(const GLfloat* c) {
	(*pglIndexfv)(c);
}
void goglIndexi(GLint c) {
	(*pglIndexi)(c);
}
void goglIndexiv(const GLint* c) {
	(*pglIndexiv)(c);
}
void goglIndexs(GLshort c) {
	(*pglIndexs)(c);
}
void goglIndexsv(const GLshort* c) {
	(*pglIndexsv)(c);
}
void goglIndexub(GLubyte c) {
	(*pglIndexub)(c);
}
void goglIndexubv(const GLubyte* c) {
	(*pglIndexubv)(c);
}
void goglInitNames() {
	(*pglInitNames)();
}
void goglInterleavedArrays(GLenum format, GLsizei stride, const GLvoid* pointer) {
	(*pglInterleavedArrays)(format, stride, pointer);
}
GLboolean goglIsBuffer(GLuint buffer) {
	return (*pglIsBuffer)(buffer);
}
GLboolean goglIsEnabled(GLenum cap) {
	return (*pglIsEnabled)(cap);
}
GLboolean goglIsFramebuffer(GLuint framebuffer) {
	return (*pglIsFramebuffer)(framebuffer);
}
GLboolean goglIsList(GLuint list) {
	return (*pglIsList)(list);
}
GLboolean goglIsProgram(GLuint program) {
	return (*pglIsProgram)(program);
}
GLboolean goglIsQuery(GLuint id) {
	return (*pglIsQuery)(id);
}
GLboolean goglIsRenderbuffer(GLuint renderbuffer) {
	return (*pglIsRenderbuffer)(renderbuffer);
}
GLboolean goglIsShader(GLuint shader) {
	return (*pglIsShader)(shader);
}
GLboolean goglIsTexture(GLuint texture) {
	return (*pglIsTexture)(texture);
}
void goglLightModelf(GLenum pname, GLfloat param) {
	(*pglLightModelf)(pname, param);
}
void goglLightModelfv(GLenum pname, const GLfloat* params) {
	(*pglLightModelfv)(pname, params);
}
void goglLightModeli(GLenum pname, GLint param) {
	(*pglLightModeli)(pname, param);
}
void goglLightModeliv(GLenum pname, const GLint* params) {
	(*pglLightModeliv)(pname, params);
}
void goglLightModelx(GLenum pname, GLfixed param) {
	(*pglLightModelx)(pname, param);
}
void goglLightModelxv(GLenum pname, const GLfixed* param) {
	(*pglLightModelxv)(pname, param);
}
void goglLightf(GLenum light, GLenum pname, GLfloat param) {
	(*pglLightf)(light, pname, param);
}
void goglLightfv(GLenum light, GLenum pname, const GLfloat* params) {
	(*pglLightfv)(light, pname, params);
}
void goglLighti(GLenum light, GLenum pname, GLint param) {
	(*pglLighti)(light, pname, param);
}
void goglLightiv(GLenum light, GLenum pname, const GLint* params) {
	(*pglLightiv)(light, pname, params);
}
void goglLightx(GLenum light, GLenum pname, GLfixed param) {
	(*pglLightx)(light, pname, param);
}
void goglLightxv(GLenum light, GLenum pname, const GLfixed* params) {
	(*pglLightxv)(light, pname, params);
}
void goglLineStipple(GLint factor, GLushort pattern) {
	(*pglLineStipple)(factor, pattern);
}
void goglLineWidth(GLfloat width) {
	(*pglLineWidth)(width);
}
void goglLineWidthx(GLfixed width) {
	(*pglLineWidthx)(width);
}
void goglLinkProgram(GLuint program) {
	(*pglLinkProgram)(program);
}
void goglListBase(GLuint base) {
	(*pglListBase)(base);
}
void goglLoadIdentity() {
	(*pglLoadIdentity)();
}
void goglLoadMatrixd(const GLdouble* m) {
	(*pglLoadMatrixd)(m);
}
void goglLoadMatrixf(const GLfloat* m) {
	(*pglLoadMatrixf)(m);
}
void goglLoadMatrixx(const GLfixed* m) {
	(*pglLoadMatrixx)(m);
}
void goglLoadName(GLuint name) {
	(*pglLoadName)(name);
}
void goglLoadTransposeMatrixd(const GLdouble* m) {
	(*pglLoadTransposeMatrixd)(m);
}
void goglLoadTransposeMatrixf(const GLfloat* m) {
	(*pglLoadTransposeMatrixf)(m);
}
void goglLogicOp(GLenum opcode) {
	(*pglLogicOp)(opcode);
}
void goglMap1d(GLenum target, GLdouble u1, GLdouble u2, GLint stride, GLint order, const GLdouble* points) {
	(*pglMap1d)(target, u1, u2, stride, order, points);
}
void goglMap1f(GLenum target, GLfloat u1, GLfloat u2, GLint stride, GLint order, const GLfloat* points) {
	(*pglMap1f)(target, u1, u2, stride, order, points);
}
void goglMap2d(GLenum target, GLdouble u1, GLdouble u2, GLint ustride, GLint uorder, GLdouble v1, GLdouble v2, GLint vstride, GLint vorder, const GLdouble* points) {
	(*pglMap2d)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}
void goglMap2f(GLenum target, GLfloat u1, GLfloat u2, GLint ustride, GLint uorder, GLfloat v1, GLfloat v2, GLint vstride, GLint vorder, const GLfloat* points) {
	(*pglMap2f)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
}
void* goglMapBuffer(GLenum target, GLenum access) {
	return (*pglMapBuffer)(target, access);
}
void goglMapGrid1d(GLint un, GLdouble u1, GLdouble u2) {
	(*pglMapGrid1d)(un, u1, u2);
}
void goglMapGrid1f(GLint un, GLfloat u1, GLfloat u2) {
	(*pglMapGrid1f)(un, u1, u2);
}
void goglMapGrid2d(GLint un, GLdouble u1, GLdouble u2, GLint vn, GLdouble v1, GLdouble v2) {
	(*pglMapGrid2d)(un, u1, u2, vn, v1, v2);
}
void goglMapGrid2f(GLint un, GLfloat u1, GLfloat u2, GLint vn, GLfloat v1, GLfloat v2) {
	(*pglMapGrid2f)(un, u1, u2, vn, v1, v2);
}
void goglMaterialf(GLenum face, GLenum pname, GLfloat param) {
	(*pglMaterialf)(face, pname, param);
}
void goglMaterialfv(GLenum face, GLenum pname, const GLfloat* params) {
	(*pglMaterialfv)(face, pname, params);
}
void goglMateriali(GLenum face, GLenum pname, GLint param) {
	(*pglMateriali)(face, pname, param);
}
void goglMaterialiv(GLenum face, GLenum pname, const GLint* params) {
	(*pglMaterialiv)(face, pname, params);
}
void goglMaterialx(GLenum face, GLenum pname, GLfixed param) {
	(*pglMaterialx)(face, pname, param);
}
void goglMaterialxv(GLenum face, GLenum pname, const GLfixed* param) {
	(*pglMaterialxv)(face, pname, param);
}
void goglMatrixMode(GLenum mode) {
	(*pglMatrixMode)(mode);
}
void goglMultMatrixd(const GLdouble* m) {
	(*pglMultMatrixd)(m);
}
void goglMultMatrixf(const GLfloat* m) {
	(*pglMultMatrixf)(m);
}
void goglMultMatrixx(const GLfixed* m) {
	(*pglMultMatrixx)(m);
}
void goglMultTransposeMatrixd(const GLdouble* m) {
	(*pglMultTransposeMatrixd)(m);
}
void goglMultTransposeMatrixf(const GLfloat* m) {
	(*pglMultTransposeMatrixf)(m);
}
void goglMultiDrawArrays(GLenum mode, const GLint* first, const GLsizei* count, GLsizei drawcount) {
	(*pglMultiDrawArrays)(mode, first, count, drawcount);
}
void goglMultiDrawElements(GLenum mode, const GLsizei* count, GLenum type, const GLvoid** indices, GLsizei drawcount) {
	(*pglMultiDrawElements)(mode, count, type, indices, drawcount);
}
void goglMultiTexCoord1d(GLenum target, GLdouble s) {
	(*pglMultiTexCoord1d)(target, s);
}
void goglMultiTexCoord1dv(GLenum target, const GLdouble* v) {
	(*pglMultiTexCoord1dv)(target, v);
}
void goglMultiTexCoord1f(GLenum target, GLfloat s) {
	(*pglMultiTexCoord1f)(target, s);
}
void goglMultiTexCoord1fv(GLenum target, const GLfloat* v) {
	(*pglMultiTexCoord1fv)(target, v);
}
void goglMultiTexCoord1i(GLenum target, GLint s) {
	(*pglMultiTexCoord1i)(target, s);
}
void goglMultiTexCoord1iv(GLenum target, const GLint* v) {
	(*pglMultiTexCoord1iv)(target, v);
}
void goglMultiTexCoord1s(GLenum target, GLshort s) {
	(*pglMultiTexCoord1s)(target, s);
}
void goglMultiTexCoord1sv(GLenum target, const GLshort* v) {
	(*pglMultiTexCoord1sv)(target, v);
}
void goglMultiTexCoord2d(GLenum target, GLdouble s, GLdouble t) {
	(*pglMultiTexCoord2d)(target, s, t);
}
void goglMultiTexCoord2dv(GLenum target, const GLdouble* v) {
	(*pglMultiTexCoord2dv)(target, v);
}
void goglMultiTexCoord2f(GLenum target, GLfloat s, GLfloat t) {
	(*pglMultiTexCoord2f)(target, s, t);
}
void goglMultiTexCoord2fv(GLenum target, const GLfloat* v) {
	(*pglMultiTexCoord2fv)(target, v);
}
void goglMultiTexCoord2i(GLenum target, GLint s, GLint t) {
	(*pglMultiTexCoord2i)(target, s, t);
}
void goglMultiTexCoord2iv(GLenum target, const GLint* v) {
	(*pglMultiTexCoord2iv)(target, v);
}
void goglMultiTexCoord2s(GLenum target, GLshort s, GLshort t) {
	(*pglMultiTexCoord2s)(target, s, t);
}
void goglMultiTexCoord2sv(GLenum target, const GLshort* v) {
	(*pglMultiTexCoord2sv)(target, v);
}
void goglMultiTexCoord3d(GLenum target, GLdouble s, GLdouble t, GLdouble r) {
	(*pglMultiTexCoord3d)(target, s, t, r);
}
void goglMultiTexCoord3dv(GLenum target, const GLdouble* v) {
	(*pglMultiTexCoord3dv)(target, v);
}
void goglMultiTexCoord3f(GLenum target, GLfloat s, GLfloat t, GLfloat r) {
	(*pglMultiTexCoord3f)(target, s, t, r);
}
void goglMultiTexCoord3fv(GLenum target, const GLfloat* v) {
	(*pglMultiTexCoord3fv)(target, v);
}
void goglMultiTexCoord3i(GLenum target, GLint s, GLint t, GLint r) {
	(*pglMultiTexCoord3i)(target, s, t, r);
}
void goglMultiTexCoord3iv(GLenum target, const GLint* v) {
	(*pglMultiTexCoord3iv)(target, v);
}
void goglMultiTexCoord3s(GLenum target, GLshort s, GLshort t, GLshort r) {
	(*pglMultiTexCoord3s)(target, s, t, r);
}
void goglMultiTexCoord3sv(GLenum target, const GLshort* v) {
	(*pglMultiTexCoord3sv)(target, v);
}
void goglMultiTexCoord4d(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
	(*pglMultiTexCoord4d)(target, s, t, r, q);
}
void goglMultiTexCoord4dv(GLenum target, const GLdouble* v) {
	(*pglMultiTexCoord4dv)(target, v);
}
void goglMultiTexCoord4f(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
	(*pglMultiTexCoord4f)(target, s, t, r, q);
}
void goglMultiTexCoord4fv(GLenum target, const GLfloat* v) {
	(*pglMultiTexCoord4fv)(target, v);
}
void goglMultiTexCoord4i(GLenum target, GLint s, GLint t, GLint r, GLint q) {
	(*pglMultiTexCoord4i)(target, s, t, r, q);
}
void goglMultiTexCoord4iv(GLenum target, const GLint* v) {
	(*pglMultiTexCoord4iv)(target, v);
}
void goglMultiTexCoord4s(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
	(*pglMultiTexCoord4s)(target, s, t, r, q);
}
void goglMultiTexCoord4sv(GLenum target, const GLshort* v) {
	(*pglMultiTexCoord4sv)(target, v);
}
void goglMultiTexCoord4x(GLenum texture, GLfixed s, GLfixed t, GLfixed r, GLfixed q) {
	(*pglMultiTexCoord4x)(texture, s, t, r, q);
}
void goglNewList(GLuint list, GLenum mode) {
	(*pglNewList)(list, mode);
}
void goglNormal3b(GLbyte nx, GLbyte ny, GLbyte nz) {
	(*pglNormal3b)(nx, ny, nz);
}
void goglNormal3bv(const GLbyte* v) {
	(*pglNormal3bv)(v);
}
void goglNormal3d(GLdouble nx, GLdouble ny, GLdouble nz) {
	(*pglNormal3d)(nx, ny, nz);
}
void goglNormal3dv(const GLdouble* v) {
	(*pglNormal3dv)(v);
}
void goglNormal3f(GLfloat nx, GLfloat ny, GLfloat nz) {
	(*pglNormal3f)(nx, ny, nz);
}
void goglNormal3fv(const GLfloat* v) {
	(*pglNormal3fv)(v);
}
void goglNormal3i(GLint nx, GLint ny, GLint nz) {
	(*pglNormal3i)(nx, ny, nz);
}
void goglNormal3iv(const GLint* v) {
	(*pglNormal3iv)(v);
}
void goglNormal3s(GLshort nx, GLshort ny, GLshort nz) {
	(*pglNormal3s)(nx, ny, nz);
}
void goglNormal3sv(const GLshort* v) {
	(*pglNormal3sv)(v);
}
void goglNormal3x(GLfixed nx, GLfixed ny, GLfixed nz) {
	(*pglNormal3x)(nx, ny, nz);
}
void goglNormalPointer(GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglNormalPointer)(type, stride, pointer);
}
void goglOrtho(GLdouble left, GLdouble right, GLdouble bottom, GLdouble top, GLdouble zNear, GLdouble zFar) {
	(*pglOrtho)(left, right, bottom, top, zNear, zFar);
}
void goglOrthof(GLfloat l, GLfloat r, GLfloat b, GLfloat t, GLfloat n, GLfloat f) {
	(*pglOrthof)(l, r, b, t, n, f);
}
void goglOrthox(GLfixed l, GLfixed r, GLfixed b, GLfixed t, GLfixed n, GLfixed f) {
	(*pglOrthox)(l, r, b, t, n, f);
}
void goglPassThrough(GLfloat token) {
	(*pglPassThrough)(token);
}
void goglPixelMapfv(GLenum map, GLsizei mapsize, const GLfloat* values) {
	(*pglPixelMapfv)(map, mapsize, values);
}
void goglPixelMapuiv(GLenum map, GLsizei mapsize, const GLuint* values) {
	(*pglPixelMapuiv)(map, mapsize, values);
}
void goglPixelMapusv(GLenum map, GLsizei mapsize, const GLushort* values) {
	(*pglPixelMapusv)(map, mapsize, values);
}
void goglPixelStoref(GLenum pname, GLfloat param) {
	(*pglPixelStoref)(pname, param);
}
void goglPixelStorei(GLenum pname, GLint param) {
	(*pglPixelStorei)(pname, param);
}
void goglPixelTransferf(GLenum pname, GLfloat param) {
	(*pglPixelTransferf)(pname, param);
}
void goglPixelTransferi(GLenum pname, GLint param) {
	(*pglPixelTransferi)(pname, param);
}
void goglPixelZoom(GLfloat xfactor, GLfloat yfactor) {
	(*pglPixelZoom)(xfactor, yfactor);
}
void goglPointParameterf(GLenum pname, GLfloat param) {
	(*pglPointParameterf)(pname, param);
}
void goglPointParameterfv(GLenum pname, const GLfloat* params) {
	(*pglPointParameterfv)(pname, params);
}
void goglPointParameteri(GLenum pname, GLint param) {
	(*pglPointParameteri)(pname, param);
}
void goglPointParameteriv(GLenum pname, const GLint* params) {
	(*pglPointParameteriv)(pname, params);
}
void goglPointParameterx(GLenum pname, GLfixed param) {
	(*pglPointParameterx)(pname, param);
}
void goglPointParameterxv(GLenum pname, const GLfixed* params) {
	(*pglPointParameterxv)(pname, params);
}
void goglPointSize(GLfloat size) {
	(*pglPointSize)(size);
}
void goglPointSizex(GLfixed size) {
	(*pglPointSizex)(size);
}
void goglPolygonMode(GLenum face, GLenum mode) {
	(*pglPolygonMode)(face, mode);
}
void goglPolygonOffset(GLfloat factor, GLfloat units) {
	(*pglPolygonOffset)(factor, units);
}
void goglPolygonOffsetx(GLfixed factor, GLfixed units) {
	(*pglPolygonOffsetx)(factor, units);
}
void goglPolygonStipple(const GLubyte* mask) {
	(*pglPolygonStipple)(mask);
}
void goglPopAttrib() {
	(*pglPopAttrib)();
}
void goglPopClientAttrib() {
	(*pglPopClientAttrib)();
}
void goglPopMatrix() {
	(*pglPopMatrix)();
}
void goglPopName() {
	(*pglPopName)();
}
void goglPrioritizeTextures(GLsizei n, const GLuint* textures, const GLfloat* priorities) {
	(*pglPrioritizeTextures)(n, textures, priorities);
}
void goglPushAttrib(GLbitfield mask) {
	(*pglPushAttrib)(mask);
}
void goglPushClientAttrib(GLbitfield mask) {
	(*pglPushClientAttrib)(mask);
}
void goglPushMatrix() {
	(*pglPushMatrix)();
}
void goglPushName(GLuint name) {
	(*pglPushName)(name);
}
void goglRasterPos2d(GLdouble x, GLdouble y) {
	(*pglRasterPos2d)(x, y);
}
void goglRasterPos2dv(const GLdouble* v) {
	(*pglRasterPos2dv)(v);
}
void goglRasterPos2f(GLfloat x, GLfloat y) {
	(*pglRasterPos2f)(x, y);
}
void goglRasterPos2fv(const GLfloat* v) {
	(*pglRasterPos2fv)(v);
}
void goglRasterPos2i(GLint x, GLint y) {
	(*pglRasterPos2i)(x, y);
}
void goglRasterPos2iv(const GLint* v) {
	(*pglRasterPos2iv)(v);
}
void goglRasterPos2s(GLshort x, GLshort y) {
	(*pglRasterPos2s)(x, y);
}
void goglRasterPos2sv(const GLshort* v) {
	(*pglRasterPos2sv)(v);
}
void goglRasterPos3d(GLdouble x, GLdouble y, GLdouble z) {
	(*pglRasterPos3d)(x, y, z);
}
void goglRasterPos3dv(const GLdouble* v) {
	(*pglRasterPos3dv)(v);
}
void goglRasterPos3f(GLfloat x, GLfloat y, GLfloat z) {
	(*pglRasterPos3f)(x, y, z);
}
void goglRasterPos3fv(const GLfloat* v) {
	(*pglRasterPos3fv)(v);
}
void goglRasterPos3i(GLint x, GLint y, GLint z) {
	(*pglRasterPos3i)(x, y, z);
}
void goglRasterPos3iv(const GLint* v) {
	(*pglRasterPos3iv)(v);
}
void goglRasterPos3s(GLshort x, GLshort y, GLshort z) {
	(*pglRasterPos3s)(x, y, z);
}
void goglRasterPos3sv(const GLshort* v) {
	(*pglRasterPos3sv)(v);
}
void goglRasterPos4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
	(*pglRasterPos4d)(x, y, z, w);
}
void goglRasterPos4dv(const GLdouble* v) {
	(*pglRasterPos4dv)(v);
}
void goglRasterPos4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
	(*pglRasterPos4f)(x, y, z, w);
}
void goglRasterPos4fv(const GLfloat* v) {
	(*pglRasterPos4fv)(v);
}
void goglRasterPos4i(GLint x, GLint y, GLint z, GLint w) {
	(*pglRasterPos4i)(x, y, z, w);
}
void goglRasterPos4iv(const GLint* v) {
	(*pglRasterPos4iv)(v);
}
void goglRasterPos4s(GLshort x, GLshort y, GLshort z, GLshort w) {
	(*pglRasterPos4s)(x, y, z, w);
}
void goglRasterPos4sv(const GLshort* v) {
	(*pglRasterPos4sv)(v);
}
void goglReadBuffer(GLenum mode) {
	(*pglReadBuffer)(mode);
}
void goglReadPixels(GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels) {
	(*pglReadPixels)(x, y, width, height, format, type, pixels);
}
void goglRectd(GLdouble x1, GLdouble y1, GLdouble x2, GLdouble y2) {
	(*pglRectd)(x1, y1, x2, y2);
}
void goglRectdv(const GLdouble* v1, const GLdouble* v2) {
	(*pglRectdv)(v1, v2);
}
void goglRectf(GLfloat x1, GLfloat y1, GLfloat x2, GLfloat y2) {
	(*pglRectf)(x1, y1, x2, y2);
}
void goglRectfv(const GLfloat* v1, const GLfloat* v2) {
	(*pglRectfv)(v1, v2);
}
void goglRecti(GLint x1, GLint y1, GLint x2, GLint y2) {
	(*pglRecti)(x1, y1, x2, y2);
}
void goglRectiv(const GLint* v1, const GLint* v2) {
	(*pglRectiv)(v1, v2);
}
void goglRects(GLshort x1, GLshort y1, GLshort x2, GLshort y2) {
	(*pglRects)(x1, y1, x2, y2);
}
void goglRectsv(const GLshort* v1, const GLshort* v2) {
	(*pglRectsv)(v1, v2);
}
void goglReleaseShaderCompiler() {
	(*pglReleaseShaderCompiler)();
}
GLint goglRenderMode(GLenum mode) {
	return (*pglRenderMode)(mode);
}
void goglRenderbufferStorage(GLenum target, GLenum internalformat, GLsizei width, GLsizei height) {
	(*pglRenderbufferStorage)(target, internalformat, width, height);
}
void goglRotated(GLdouble angle, GLdouble x, GLdouble y, GLdouble z) {
	(*pglRotated)(angle, x, y, z);
}
void goglRotatef(GLfloat angle, GLfloat x, GLfloat y, GLfloat z) {
	(*pglRotatef)(angle, x, y, z);
}
void goglRotatex(GLfixed angle, GLfixed x, GLfixed y, GLfixed z) {
	(*pglRotatex)(angle, x, y, z);
}
void goglSampleCoverage(GLfloat value, GLboolean invert) {
	(*pglSampleCoverage)(value, invert);
}
void goglSampleCoveragex(GLclampx value, GLboolean invert) {
	(*pglSampleCoveragex)(value, invert);
}
void goglScaled(GLdouble x, GLdouble y, GLdouble z) {
	(*pglScaled)(x, y, z);
}
void goglScalef(GLfloat x, GLfloat y, GLfloat z) {
	(*pglScalef)(x, y, z);
}
void goglScalex(GLfixed x, GLfixed y, GLfixed z) {
	(*pglScalex)(x, y, z);
}
void goglScissor(GLint x, GLint y, GLsizei width, GLsizei height) {
	(*pglScissor)(x, y, width, height);
}
void goglSecondaryColor3b(GLbyte red, GLbyte green, GLbyte blue) {
	(*pglSecondaryColor3b)(red, green, blue);
}
void goglSecondaryColor3bv(const GLbyte* v) {
	(*pglSecondaryColor3bv)(v);
}
void goglSecondaryColor3d(GLdouble red, GLdouble green, GLdouble blue) {
	(*pglSecondaryColor3d)(red, green, blue);
}
void goglSecondaryColor3dv(const GLdouble* v) {
	(*pglSecondaryColor3dv)(v);
}
void goglSecondaryColor3f(GLfloat red, GLfloat green, GLfloat blue) {
	(*pglSecondaryColor3f)(red, green, blue);
}
void goglSecondaryColor3fv(const GLfloat* v) {
	(*pglSecondaryColor3fv)(v);
}
void goglSecondaryColor3i(GLint red, GLint green, GLint blue) {
	(*pglSecondaryColor3i)(red, green, blue);
}
void goglSecondaryColor3iv(const GLint* v) {
	(*pglSecondaryColor3iv)(v);
}
void goglSecondaryColor3s(GLshort red, GLshort green, GLshort blue) {
	(*pglSecondaryColor3s)(red, green, blue);
}
void goglSecondaryColor3sv(const GLshort* v) {
	(*pglSecondaryColor3sv)(v);
}
void goglSecondaryColor3ub(GLubyte red, GLubyte green, GLubyte blue) {
	(*pglSecondaryColor3ub)(red, green, blue);
}
void goglSecondaryColor3ubv(const GLubyte* v) {
	(*pglSecondaryColor3ubv)(v);
}
void goglSecondaryColor3ui(GLuint red, GLuint green, GLuint blue) {
	(*pglSecondaryColor3ui)(red, green, blue);
}
void goglSecondaryColor3uiv(const GLuint* v) {
	(*pglSecondaryColor3uiv)(v);
}
void goglSecondaryColor3us(GLushort red, GLushort green, GLushort blue) {
	(*pglSecondaryColor3us)(red, green, blue);
}
void goglSecondaryColor3usv(const GLushort* v) {
	(*pglSecondaryColor3usv)(v);
}
void goglSecondaryColorPointer(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglSecondaryColorPointer)(size, type, stride, pointer);
}
void goglSelectBuffer(GLsizei size, GLuint* buffer) {
	(*pglSelectBuffer)(size, buffer);
}
void goglShadeModel(GLenum mode) {
	(*pglShadeModel)(mode);
}
void goglShaderBinary(GLsizei count, const GLuint* shaders, GLenum binaryformat, const GLvoid* binary, GLsizei length) {
	(*pglShaderBinary)(count, shaders, binaryformat, binary, length);
}
void goglShaderSource(GLuint shader, GLsizei count, const GLchar** string, const GLint* length) {
	(*pglShaderSource)(shader, count, string, length);
}
void goglStencilFunc(GLenum func, GLint ref, GLuint mask) {
	(*pglStencilFunc)(func, ref, mask);
}
void goglStencilFuncSeparate(GLenum face, GLenum func, GLint ref, GLuint mask) {
	(*pglStencilFuncSeparate)(face, func, ref, mask);
}
void goglStencilMask(GLuint mask) {
	(*pglStencilMask)(mask);
}
void goglStencilMaskSeparate(GLenum face, GLuint mask) {
	(*pglStencilMaskSeparate)(face, mask);
}
void goglStencilOp(GLenum fail, GLenum zfail, GLenum zpass) {
	(*pglStencilOp)(fail, zfail, zpass);
}
void goglStencilOpSeparate(GLenum face, GLenum sfail, GLenum dpfail, GLenum dppass) {
	(*pglStencilOpSeparate)(face, sfail, dpfail, dppass);
}
void goglTexCoord1d(GLdouble s) {
	(*pglTexCoord1d)(s);
}
void goglTexCoord1dv(const GLdouble* v) {
	(*pglTexCoord1dv)(v);
}
void goglTexCoord1f(GLfloat s) {
	(*pglTexCoord1f)(s);
}
void goglTexCoord1fv(const GLfloat* v) {
	(*pglTexCoord1fv)(v);
}
void goglTexCoord1i(GLint s) {
	(*pglTexCoord1i)(s);
}
void goglTexCoord1iv(const GLint* v) {
	(*pglTexCoord1iv)(v);
}
void goglTexCoord1s(GLshort s) {
	(*pglTexCoord1s)(s);
}
void goglTexCoord1sv(const GLshort* v) {
	(*pglTexCoord1sv)(v);
}
void goglTexCoord2d(GLdouble s, GLdouble t) {
	(*pglTexCoord2d)(s, t);
}
void goglTexCoord2dv(const GLdouble* v) {
	(*pglTexCoord2dv)(v);
}
void goglTexCoord2f(GLfloat s, GLfloat t) {
	(*pglTexCoord2f)(s, t);
}
void goglTexCoord2fv(const GLfloat* v) {
	(*pglTexCoord2fv)(v);
}
void goglTexCoord2i(GLint s, GLint t) {
	(*pglTexCoord2i)(s, t);
}
void goglTexCoord2iv(const GLint* v) {
	(*pglTexCoord2iv)(v);
}
void goglTexCoord2s(GLshort s, GLshort t) {
	(*pglTexCoord2s)(s, t);
}
void goglTexCoord2sv(const GLshort* v) {
	(*pglTexCoord2sv)(v);
}
void goglTexCoord3d(GLdouble s, GLdouble t, GLdouble r) {
	(*pglTexCoord3d)(s, t, r);
}
void goglTexCoord3dv(const GLdouble* v) {
	(*pglTexCoord3dv)(v);
}
void goglTexCoord3f(GLfloat s, GLfloat t, GLfloat r) {
	(*pglTexCoord3f)(s, t, r);
}
void goglTexCoord3fv(const GLfloat* v) {
	(*pglTexCoord3fv)(v);
}
void goglTexCoord3i(GLint s, GLint t, GLint r) {
	(*pglTexCoord3i)(s, t, r);
}
void goglTexCoord3iv(const GLint* v) {
	(*pglTexCoord3iv)(v);
}
void goglTexCoord3s(GLshort s, GLshort t, GLshort r) {
	(*pglTexCoord3s)(s, t, r);
}
void goglTexCoord3sv(const GLshort* v) {
	(*pglTexCoord3sv)(v);
}
void goglTexCoord4d(GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
	(*pglTexCoord4d)(s, t, r, q);
}
void goglTexCoord4dv(const GLdouble* v) {
	(*pglTexCoord4dv)(v);
}
void goglTexCoord4f(GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
	(*pglTexCoord4f)(s, t, r, q);
}
void goglTexCoord4fv(const GLfloat* v) {
	(*pglTexCoord4fv)(v);
}
void goglTexCoord4i(GLint s, GLint t, GLint r, GLint q) {
	(*pglTexCoord4i)(s, t, r, q);
}
void goglTexCoord4iv(const GLint* v) {
	(*pglTexCoord4iv)(v);
}
void goglTexCoord4s(GLshort s, GLshort t, GLshort r, GLshort q) {
	(*pglTexCoord4s)(s, t, r, q);
}
void goglTexCoord4sv(const GLshort* v) {
	(*pglTexCoord4sv)(v);
}
void goglTexCoordPointer(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglTexCoordPointer)(size, type, stride, pointer);
}
void goglTexEnvf(GLenum target, GLenum pname, GLfloat param) {
	(*pglTexEnvf)(target, pname, param);
}
void goglTexEnvfv(GLenum target, GLenum pname, const GLfloat* params) {
	(*pglTexEnvfv)(target, pname, params);
}
void goglTexEnvi(GLenum target, GLenum pname, GLint param) {
	(*pglTexEnvi)(target, pname, param);
}
void goglTexEnviv(GLenum target, GLenum pname, const GLint* params) {
	(*pglTexEnviv)(target, pname, params);
}
void goglTexEnvx(GLenum target, GLenum pname, GLfixed param) {
	(*pglTexEnvx)(target, pname, param);
}
void goglTexEnvxv(GLenum target, GLenum pname, const GLfixed* params) {
	(*pglTexEnvxv)(target, pname, params);
}
void goglTexGend(GLenum coord, GLenum pname, GLdouble param) {
	(*pglTexGend)(coord, pname, param);
}
void goglTexGendv(GLenum coord, GLenum pname, const GLdouble* params) {
	(*pglTexGendv)(coord, pname, params);
}
void goglTexGenf(GLenum coord, GLenum pname, GLfloat param) {
	(*pglTexGenf)(coord, pname, param);
}
void goglTexGenfv(GLenum coord, GLenum pname, const GLfloat* params) {
	(*pglTexGenfv)(coord, pname, params);
}
void goglTexGeni(GLenum coord, GLenum pname, GLint param) {
	(*pglTexGeni)(coord, pname, param);
}
void goglTexGeniv(GLenum coord, GLenum pname, const GLint* params) {
	(*pglTexGeniv)(coord, pname, params);
}
void goglTexImage1D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLint border, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexImage1D)(target, level, internalformat, width, border, format, type, pixels);
}
void goglTexImage2D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexImage2D)(target, level, internalformat, width, height, border, format, type, pixels);
}
void goglTexImage3D(GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexImage3D)(target, level, internalformat, width, height, depth, border, format, type, pixels);
}
void goglTexParameterf(GLenum target, GLenum pname, GLfloat param) {
	(*pglTexParameterf)(target, pname, param);
}
void goglTexParameterfv(GLenum target, GLenum pname, const GLfloat* params) {
	(*pglTexParameterfv)(target, pname, params);
}
void goglTexParameteri(GLenum target, GLenum pname, GLint param) {
	(*pglTexParameteri)(target, pname, param);
}
void goglTexParameteriv(GLenum target, GLenum pname, const GLint* params) {
	(*pglTexParameteriv)(target, pname, params);
}
void goglTexParameterx(GLenum target, GLenum pname, GLfixed param) {
	(*pglTexParameterx)(target, pname, param);
}
void goglTexParameterxv(GLenum target, GLenum pname, const GLfixed* params) {
	(*pglTexParameterxv)(target, pname, params);
}
void goglTexSubImage1D(GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexSubImage1D)(target, level, xoffset, width, format, type, pixels);
}
void goglTexSubImage2D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexSubImage2D)(target, level, xoffset, yoffset, width, height, format, type, pixels);
}
void goglTexSubImage3D(GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid* pixels) {
	(*pglTexSubImage3D)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
}
void goglTranslated(GLdouble x, GLdouble y, GLdouble z) {
	(*pglTranslated)(x, y, z);
}
void goglTranslatef(GLfloat x, GLfloat y, GLfloat z) {
	(*pglTranslatef)(x, y, z);
}
void goglTranslatex(GLfixed x, GLfixed y, GLfixed z) {
	(*pglTranslatex)(x, y, z);
}
void goglUniform1f(GLint location, GLfloat v0) {
	(*pglUniform1f)(location, v0);
}
void goglUniform1fv(GLint location, GLsizei count, const GLfloat* value) {
	(*pglUniform1fv)(location, count, value);
}
void goglUniform1i(GLint location, GLint v0) {
	(*pglUniform1i)(location, v0);
}
void goglUniform1iv(GLint location, GLsizei count, const GLint* value) {
	(*pglUniform1iv)(location, count, value);
}
void goglUniform2f(GLint location, GLfloat v0, GLfloat v1) {
	(*pglUniform2f)(location, v0, v1);
}
void goglUniform2fv(GLint location, GLsizei count, const GLfloat* value) {
	(*pglUniform2fv)(location, count, value);
}
void goglUniform2i(GLint location, GLint v0, GLint v1) {
	(*pglUniform2i)(location, v0, v1);
}
void goglUniform2iv(GLint location, GLsizei count, const GLint* value) {
	(*pglUniform2iv)(location, count, value);
}
void goglUniform3f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2) {
	(*pglUniform3f)(location, v0, v1, v2);
}
void goglUniform3fv(GLint location, GLsizei count, const GLfloat* value) {
	(*pglUniform3fv)(location, count, value);
}
void goglUniform3i(GLint location, GLint v0, GLint v1, GLint v2) {
	(*pglUniform3i)(location, v0, v1, v2);
}
void goglUniform3iv(GLint location, GLsizei count, const GLint* value) {
	(*pglUniform3iv)(location, count, value);
}
void goglUniform4f(GLint location, GLfloat v0, GLfloat v1, GLfloat v2, GLfloat v3) {
	(*pglUniform4f)(location, v0, v1, v2, v3);
}
void goglUniform4fv(GLint location, GLsizei count, const GLfloat* value) {
	(*pglUniform4fv)(location, count, value);
}
void goglUniform4i(GLint location, GLint v0, GLint v1, GLint v2, GLint v3) {
	(*pglUniform4i)(location, v0, v1, v2, v3);
}
void goglUniform4iv(GLint location, GLsizei count, const GLint* value) {
	(*pglUniform4iv)(location, count, value);
}
void goglUniformMatrix2fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value) {
	(*pglUniformMatrix2fv)(location, count, transpose, value);
}
void goglUniformMatrix3fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value) {
	(*pglUniformMatrix3fv)(location, count, transpose, value);
}
void goglUniformMatrix4fv(GLint location, GLsizei count, GLboolean transpose, const GLfloat* value) {
	(*pglUniformMatrix4fv)(location, count, transpose, value);
}
GLboolean goglUnmapBuffer(GLenum target) {
	return (*pglUnmapBuffer)(target);
}
void goglUseProgram(GLuint program) {
	(*pglUseProgram)(program);
}
void goglValidateProgram(GLuint program) {
	(*pglValidateProgram)(program);
}
void goglVertex2d(GLdouble x, GLdouble y) {
	(*pglVertex2d)(x, y);
}
void goglVertex2dv(const GLdouble* v) {
	(*pglVertex2dv)(v);
}
void goglVertex2f(GLfloat x, GLfloat y) {
	(*pglVertex2f)(x, y);
}
void goglVertex2fv(const GLfloat* v) {
	(*pglVertex2fv)(v);
}
void goglVertex2i(GLint x, GLint y) {
	(*pglVertex2i)(x, y);
}
void goglVertex2iv(const GLint* v) {
	(*pglVertex2iv)(v);
}
void goglVertex2s(GLshort x, GLshort y) {
	(*pglVertex2s)(x, y);
}
void goglVertex2sv(const GLshort* v) {
	(*pglVertex2sv)(v);
}
void goglVertex3d(GLdouble x, GLdouble y, GLdouble z) {
	(*pglVertex3d)(x, y, z);
}
void goglVertex3dv(const GLdouble* v) {
	(*pglVertex3dv)(v);
}
void goglVertex3f(GLfloat x, GLfloat y, GLfloat z) {
	(*pglVertex3f)(x, y, z);
}
void goglVertex3fv(const GLfloat* v) {
	(*pglVertex3fv)(v);
}
void goglVertex3i(GLint x, GLint y, GLint z) {
	(*pglVertex3i)(x, y, z);
}
void goglVertex3iv(const GLint* v) {
	(*pglVertex3iv)(v);
}
void goglVertex3s(GLshort x, GLshort y, GLshort z) {
	(*pglVertex3s)(x, y, z);
}
void goglVertex3sv(const GLshort* v) {
	(*pglVertex3sv)(v);
}
void goglVertex4d(GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
	(*pglVertex4d)(x, y, z, w);
}
void goglVertex4dv(const GLdouble* v) {
	(*pglVertex4dv)(v);
}
void goglVertex4f(GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
	(*pglVertex4f)(x, y, z, w);
}
void goglVertex4fv(const GLfloat* v) {
	(*pglVertex4fv)(v);
}
void goglVertex4i(GLint x, GLint y, GLint z, GLint w) {
	(*pglVertex4i)(x, y, z, w);
}
void goglVertex4iv(const GLint* v) {
	(*pglVertex4iv)(v);
}
void goglVertex4s(GLshort x, GLshort y, GLshort z, GLshort w) {
	(*pglVertex4s)(x, y, z, w);
}
void goglVertex4sv(const GLshort* v) {
	(*pglVertex4sv)(v);
}
void goglVertexAttrib1d(GLuint index, GLdouble x) {
	(*pglVertexAttrib1d)(index, x);
}
void goglVertexAttrib1dv(GLuint index, const GLdouble* v) {
	(*pglVertexAttrib1dv)(index, v);
}
void goglVertexAttrib1f(GLuint index, GLfloat x) {
	(*pglVertexAttrib1f)(index, x);
}
void goglVertexAttrib1fv(GLuint index, const GLfloat* v) {
	(*pglVertexAttrib1fv)(index, v);
}
void goglVertexAttrib1s(GLuint index, GLshort x) {
	(*pglVertexAttrib1s)(index, x);
}
void goglVertexAttrib1sv(GLuint index, const GLshort* v) {
	(*pglVertexAttrib1sv)(index, v);
}
void goglVertexAttrib2d(GLuint index, GLdouble x, GLdouble y) {
	(*pglVertexAttrib2d)(index, x, y);
}
void goglVertexAttrib2dv(GLuint index, const GLdouble* v) {
	(*pglVertexAttrib2dv)(index, v);
}
void goglVertexAttrib2f(GLuint index, GLfloat x, GLfloat y) {
	(*pglVertexAttrib2f)(index, x, y);
}
void goglVertexAttrib2fv(GLuint index, const GLfloat* v) {
	(*pglVertexAttrib2fv)(index, v);
}
void goglVertexAttrib2s(GLuint index, GLshort x, GLshort y) {
	(*pglVertexAttrib2s)(index, x, y);
}
void goglVertexAttrib2sv(GLuint index, const GLshort* v) {
	(*pglVertexAttrib2sv)(index, v);
}
void goglVertexAttrib3d(GLuint index, GLdouble x, GLdouble y, GLdouble z) {
	(*pglVertexAttrib3d)(index, x, y, z);
}
void goglVertexAttrib3dv(GLuint index, const GLdouble* v) {
	(*pglVertexAttrib3dv)(index, v);
}
void goglVertexAttrib3f(GLuint index, GLfloat x, GLfloat y, GLfloat z) {
	(*pglVertexAttrib3f)(index, x, y, z);
}
void goglVertexAttrib3fv(GLuint index, const GLfloat* v) {
	(*pglVertexAttrib3fv)(index, v);
}
void goglVertexAttrib3s(GLuint index, GLshort x, GLshort y, GLshort z) {
	(*pglVertexAttrib3s)(index, x, y, z);
}
void goglVertexAttrib3sv(GLuint index, const GLshort* v) {
	(*pglVertexAttrib3sv)(index, v);
}
void goglVertexAttrib4Nbv(GLuint index, const GLbyte* v) {
	(*pglVertexAttrib4Nbv)(index, v);
}
void goglVertexAttrib4Niv(GLuint index, const GLint* v) {
	(*pglVertexAttrib4Niv)(index, v);
}
void goglVertexAttrib4Nsv(GLuint index, const GLshort* v) {
	(*pglVertexAttrib4Nsv)(index, v);
}
void goglVertexAttrib4Nub(GLuint index, GLubyte x, GLubyte y, GLubyte z, GLubyte w) {
	(*pglVertexAttrib4Nub)(index, x, y, z, w);
}
void goglVertexAttrib4Nubv(GLuint index, const GLubyte* v) {
	(*pglVertexAttrib4Nubv)(index, v);
}
void goglVertexAttrib4Nuiv(GLuint index, const GLuint* v) {
	(*pglVertexAttrib4Nuiv)(index, v);
}
void goglVertexAttrib4Nusv(GLuint index, const GLushort* v) {
	(*pglVertexAttrib4Nusv)(index, v);
}
void goglVertexAttrib4bv(GLuint index, const GLbyte* v) {
	(*pglVertexAttrib4bv)(index, v);
}
void goglVertexAttrib4d(GLuint index, GLdouble x, GLdouble y, GLdouble z, GLdouble w) {
	(*pglVertexAttrib4d)(index, x, y, z, w);
}
void goglVertexAttrib4dv(GLuint index, const GLdouble* v) {
	(*pglVertexAttrib4dv)(index, v);
}
void goglVertexAttrib4f(GLuint index, GLfloat x, GLfloat y, GLfloat z, GLfloat w) {
	(*pglVertexAttrib4f)(index, x, y, z, w);
}
void goglVertexAttrib4fv(GLuint index, const GLfloat* v) {
	(*pglVertexAttrib4fv)(index, v);
}
void goglVertexAttrib4iv(GLuint index, const GLint* v) {
	(*pglVertexAttrib4iv)(index, v);
}
void goglVertexAttrib4s(GLuint index, GLshort x, GLshort y, GLshort z, GLshort w) {
	(*pglVertexAttrib4s)(index, x, y, z, w);
}
void goglVertexAttrib4sv(GLuint index, const GLshort* v) {
	(*pglVertexAttrib4sv)(index, v);
}
void goglVertexAttrib4ubv(GLuint index, const GLubyte* v) {
	(*pglVertexAttrib4ubv)(index, v);
}
void goglVertexAttrib4uiv(GLuint index, const GLuint* v) {
	(*pglVertexAttrib4uiv)(index, v);
}
void goglVertexAttrib4usv(GLuint index, const GLushort* v) {
	(*pglVertexAttrib4usv)(index, v);
}
void goglVertexAttribPointer(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const GLvoid* pointer) {
	(*pglVertexAttribPointer)(index, size, type, normalized, stride, pointer);
}
void goglVertexPointer(GLint size, GLenum type, GLsizei stride, const GLvoid* pointer) {
	(*pglVertexPointer)(size, type, stride, pointer);
}
void goglViewport(GLint x, GLint y, GLsizei width, GLsizei height) {
	(*pglViewport)(x, y, width, height);
}
void goglWindowPos2d(GLdouble x, GLdouble y) {
	(*pglWindowPos2d)(x, y);
}
void goglWindowPos2dv(const GLdouble* v) {
	(*pglWindowPos2dv)(v);
}
void goglWindowPos2f(GLfloat x, GLfloat y) {
	(*pglWindowPos2f)(x, y);
}
void goglWindowPos2fv(const GLfloat* v) {
	(*pglWindowPos2fv)(v);
}
void goglWindowPos2i(GLint x, GLint y) {
	(*pglWindowPos2i)(x, y);
}
void goglWindowPos2iv(const GLint* v) {
	(*pglWindowPos2iv)(v);
}
void goglWindowPos2s(GLshort x, GLshort y) {
	(*pglWindowPos2s)(x, y);
}
void goglWindowPos2sv(const GLshort* v) {
	(*pglWindowPos2sv)(v);
}
void goglWindowPos3d(GLdouble x, GLdouble y, GLdouble z) {
	(*pglWindowPos3d)(x, y, z);
}
void goglWindowPos3dv(const GLdouble* v) {
	(*pglWindowPos3dv)(v);
}
void goglWindowPos3f(GLfloat x, GLfloat y, GLfloat z) {
	(*pglWindowPos3f)(x, y, z);
}
void goglWindowPos3fv(const GLfloat* v) {
	(*pglWindowPos3fv)(v);
}
void goglWindowPos3i(GLint x, GLint y, GLint z) {
	(*pglWindowPos3i)(x, y, z);
}
void goglWindowPos3iv(const GLint* v) {
	(*pglWindowPos3iv)(v);
}
void goglWindowPos3s(GLshort x, GLshort y, GLshort z) {
	(*pglWindowPos3s)(x, y, z);
}
void goglWindowPos3sv(const GLshort* v) {
	(*pglWindowPos3sv)(v);
}

char *goglInit(char *buf, size_t n) {
	const char *err;
	if((pglAccum = goglGetProcAddress(err="glAccum")) == NULL) goto fail;
	if((pglActiveTexture = goglGetProcAddress(err="glActiveTexture")) == NULL) goto fail;
	if((pglAlphaFunc = goglGetProcAddress(err="glAlphaFunc")) == NULL) goto fail;
//	if((pglAlphaFuncx = goglGetProcAddress(err="glAlphaFuncx")) == NULL) goto fail;
	if((pglAreTexturesResident = goglGetProcAddress(err="glAreTexturesResident")) == NULL) goto fail;
	if((pglArrayElement = goglGetProcAddress(err="glArrayElement")) == NULL) goto fail;
	if((pglAttachShader = goglGetProcAddress(err="glAttachShader")) == NULL) goto fail;
	if((pglBegin = goglGetProcAddress(err="glBegin")) == NULL) goto fail;
	if((pglBeginQuery = goglGetProcAddress(err="glBeginQuery")) == NULL) goto fail;
	if((pglBindAttribLocation = goglGetProcAddress(err="glBindAttribLocation")) == NULL) goto fail;
	if((pglBindBuffer = goglGetProcAddress(err="glBindBuffer")) == NULL) goto fail;
	if((pglBindFramebuffer = goglGetProcAddress(err="glBindFramebuffer")) == NULL) goto fail;
	if((pglBindRenderbuffer = goglGetProcAddress(err="glBindRenderbuffer")) == NULL) goto fail;
	if((pglBindTexture = goglGetProcAddress(err="glBindTexture")) == NULL) goto fail;
	if((pglBitmap = goglGetProcAddress(err="glBitmap")) == NULL) goto fail;
	if((pglBlendColor = goglGetProcAddress(err="glBlendColor")) == NULL) goto fail;
	if((pglBlendEquation = goglGetProcAddress(err="glBlendEquation")) == NULL) goto fail;
	if((pglBlendEquationSeparate = goglGetProcAddress(err="glBlendEquationSeparate")) == NULL) goto fail;
	if((pglBlendFunc = goglGetProcAddress(err="glBlendFunc")) == NULL) goto fail;
	if((pglBlendFuncSeparate = goglGetProcAddress(err="glBlendFuncSeparate")) == NULL) goto fail;
	if((pglBufferData = goglGetProcAddress(err="glBufferData")) == NULL) goto fail;
	if((pglBufferSubData = goglGetProcAddress(err="glBufferSubData")) == NULL) goto fail;
	if((pglCallList = goglGetProcAddress(err="glCallList")) == NULL) goto fail;
	if((pglCallLists = goglGetProcAddress(err="glCallLists")) == NULL) goto fail;
	if((pglCheckFramebufferStatus = goglGetProcAddress(err="glCheckFramebufferStatus")) == NULL) goto fail;
	if((pglClear = goglGetProcAddress(err="glClear")) == NULL) goto fail;
	if((pglClearAccum = goglGetProcAddress(err="glClearAccum")) == NULL) goto fail;
	if((pglClearColor = goglGetProcAddress(err="glClearColor")) == NULL) goto fail;
//	if((pglClearColorx = goglGetProcAddress(err="glClearColorx")) == NULL) goto fail;
	if((pglClearDepth = goglGetProcAddress(err="glClearDepth")) == NULL) goto fail;
	if((pglClearDepthf = goglGetProcAddress(err="glClearDepthf")) == NULL) goto fail;
//	if((pglClearDepthx = goglGetProcAddress(err="glClearDepthx")) == NULL) goto fail;
	if((pglClearIndex = goglGetProcAddress(err="glClearIndex")) == NULL) goto fail;
	if((pglClearStencil = goglGetProcAddress(err="glClearStencil")) == NULL) goto fail;
	if((pglClientActiveTexture = goglGetProcAddress(err="glClientActiveTexture")) == NULL) goto fail;
	if((pglClipPlane = goglGetProcAddress(err="glClipPlane")) == NULL) goto fail;
//	if((pglClipPlanef = goglGetProcAddress(err="glClipPlanef")) == NULL) goto fail;
//	if((pglClipPlanex = goglGetProcAddress(err="glClipPlanex")) == NULL) goto fail;
	if((pglColor3b = goglGetProcAddress(err="glColor3b")) == NULL) goto fail;
	if((pglColor3bv = goglGetProcAddress(err="glColor3bv")) == NULL) goto fail;
	if((pglColor3d = goglGetProcAddress(err="glColor3d")) == NULL) goto fail;
	if((pglColor3dv = goglGetProcAddress(err="glColor3dv")) == NULL) goto fail;
	if((pglColor3f = goglGetProcAddress(err="glColor3f")) == NULL) goto fail;
	if((pglColor3fv = goglGetProcAddress(err="glColor3fv")) == NULL) goto fail;
	if((pglColor3i = goglGetProcAddress(err="glColor3i")) == NULL) goto fail;
	if((pglColor3iv = goglGetProcAddress(err="glColor3iv")) == NULL) goto fail;
	if((pglColor3s = goglGetProcAddress(err="glColor3s")) == NULL) goto fail;
	if((pglColor3sv = goglGetProcAddress(err="glColor3sv")) == NULL) goto fail;
	if((pglColor3ub = goglGetProcAddress(err="glColor3ub")) == NULL) goto fail;
	if((pglColor3ubv = goglGetProcAddress(err="glColor3ubv")) == NULL) goto fail;
	if((pglColor3ui = goglGetProcAddress(err="glColor3ui")) == NULL) goto fail;
	if((pglColor3uiv = goglGetProcAddress(err="glColor3uiv")) == NULL) goto fail;
	if((pglColor3us = goglGetProcAddress(err="glColor3us")) == NULL) goto fail;
	if((pglColor3usv = goglGetProcAddress(err="glColor3usv")) == NULL) goto fail;
	if((pglColor4b = goglGetProcAddress(err="glColor4b")) == NULL) goto fail;
	if((pglColor4bv = goglGetProcAddress(err="glColor4bv")) == NULL) goto fail;
	if((pglColor4d = goglGetProcAddress(err="glColor4d")) == NULL) goto fail;
	if((pglColor4dv = goglGetProcAddress(err="glColor4dv")) == NULL) goto fail;
	if((pglColor4f = goglGetProcAddress(err="glColor4f")) == NULL) goto fail;
	if((pglColor4fv = goglGetProcAddress(err="glColor4fv")) == NULL) goto fail;
	if((pglColor4i = goglGetProcAddress(err="glColor4i")) == NULL) goto fail;
	if((pglColor4iv = goglGetProcAddress(err="glColor4iv")) == NULL) goto fail;
	if((pglColor4s = goglGetProcAddress(err="glColor4s")) == NULL) goto fail;
	if((pglColor4sv = goglGetProcAddress(err="glColor4sv")) == NULL) goto fail;
	if((pglColor4ub = goglGetProcAddress(err="glColor4ub")) == NULL) goto fail;
	if((pglColor4ubv = goglGetProcAddress(err="glColor4ubv")) == NULL) goto fail;
	if((pglColor4ui = goglGetProcAddress(err="glColor4ui")) == NULL) goto fail;
	if((pglColor4uiv = goglGetProcAddress(err="glColor4uiv")) == NULL) goto fail;
	if((pglColor4us = goglGetProcAddress(err="glColor4us")) == NULL) goto fail;
	if((pglColor4usv = goglGetProcAddress(err="glColor4usv")) == NULL) goto fail;
//	if((pglColor4x = goglGetProcAddress(err="glColor4x")) == NULL) goto fail;
	if((pglColorMask = goglGetProcAddress(err="glColorMask")) == NULL) goto fail;
	if((pglColorMaterial = goglGetProcAddress(err="glColorMaterial")) == NULL) goto fail;
	if((pglColorPointer = goglGetProcAddress(err="glColorPointer")) == NULL) goto fail;
	if((pglCompileShader = goglGetProcAddress(err="glCompileShader")) == NULL) goto fail;
	if((pglCompressedTexImage1D = goglGetProcAddress(err="glCompressedTexImage1D")) == NULL) goto fail;
	if((pglCompressedTexImage2D = goglGetProcAddress(err="glCompressedTexImage2D")) == NULL) goto fail;
	if((pglCompressedTexImage3D = goglGetProcAddress(err="glCompressedTexImage3D")) == NULL) goto fail;
	if((pglCompressedTexSubImage1D = goglGetProcAddress(err="glCompressedTexSubImage1D")) == NULL) goto fail;
	if((pglCompressedTexSubImage2D = goglGetProcAddress(err="glCompressedTexSubImage2D")) == NULL) goto fail;
	if((pglCompressedTexSubImage3D = goglGetProcAddress(err="glCompressedTexSubImage3D")) == NULL) goto fail;
	if((pglCopyPixels = goglGetProcAddress(err="glCopyPixels")) == NULL) goto fail;
	if((pglCopyTexImage1D = goglGetProcAddress(err="glCopyTexImage1D")) == NULL) goto fail;
	if((pglCopyTexImage2D = goglGetProcAddress(err="glCopyTexImage2D")) == NULL) goto fail;
	if((pglCopyTexSubImage1D = goglGetProcAddress(err="glCopyTexSubImage1D")) == NULL) goto fail;
	if((pglCopyTexSubImage2D = goglGetProcAddress(err="glCopyTexSubImage2D")) == NULL) goto fail;
	if((pglCopyTexSubImage3D = goglGetProcAddress(err="glCopyTexSubImage3D")) == NULL) goto fail;
	if((pglCreateProgram = goglGetProcAddress(err="glCreateProgram")) == NULL) goto fail;
	if((pglCreateShader = goglGetProcAddress(err="glCreateShader")) == NULL) goto fail;
	if((pglCullFace = goglGetProcAddress(err="glCullFace")) == NULL) goto fail;
	if((pglDeleteBuffers = goglGetProcAddress(err="glDeleteBuffers")) == NULL) goto fail;
	if((pglDeleteFramebuffers = goglGetProcAddress(err="glDeleteFramebuffers")) == NULL) goto fail;
	if((pglDeleteLists = goglGetProcAddress(err="glDeleteLists")) == NULL) goto fail;
	if((pglDeleteProgram = goglGetProcAddress(err="glDeleteProgram")) == NULL) goto fail;
	if((pglDeleteQueries = goglGetProcAddress(err="glDeleteQueries")) == NULL) goto fail;
	if((pglDeleteRenderbuffers = goglGetProcAddress(err="glDeleteRenderbuffers")) == NULL) goto fail;
	if((pglDeleteShader = goglGetProcAddress(err="glDeleteShader")) == NULL) goto fail;
	if((pglDeleteTextures = goglGetProcAddress(err="glDeleteTextures")) == NULL) goto fail;
	if((pglDepthFunc = goglGetProcAddress(err="glDepthFunc")) == NULL) goto fail;
	if((pglDepthMask = goglGetProcAddress(err="glDepthMask")) == NULL) goto fail;
	if((pglDepthRange = goglGetProcAddress(err="glDepthRange")) == NULL) goto fail;
	if((pglDepthRangef = goglGetProcAddress(err="glDepthRangef")) == NULL) goto fail;
//	if((pglDepthRangex = goglGetProcAddress(err="glDepthRangex")) == NULL) goto fail;
	if((pglDetachShader = goglGetProcAddress(err="glDetachShader")) == NULL) goto fail;
	if((pglDisable = goglGetProcAddress(err="glDisable")) == NULL) goto fail;
	if((pglDisableClientState = goglGetProcAddress(err="glDisableClientState")) == NULL) goto fail;
	if((pglDisableVertexAttribArray = goglGetProcAddress(err="glDisableVertexAttribArray")) == NULL) goto fail;
	if((pglDrawArrays = goglGetProcAddress(err="glDrawArrays")) == NULL) goto fail;
	if((pglDrawBuffer = goglGetProcAddress(err="glDrawBuffer")) == NULL) goto fail;
	if((pglDrawBuffers = goglGetProcAddress(err="glDrawBuffers")) == NULL) goto fail;
	if((pglDrawElements = goglGetProcAddress(err="glDrawElements")) == NULL) goto fail;
	if((pglDrawPixels = goglGetProcAddress(err="glDrawPixels")) == NULL) goto fail;
	if((pglDrawRangeElements = goglGetProcAddress(err="glDrawRangeElements")) == NULL) goto fail;
	if((pglEdgeFlag = goglGetProcAddress(err="glEdgeFlag")) == NULL) goto fail;
	if((pglEdgeFlagPointer = goglGetProcAddress(err="glEdgeFlagPointer")) == NULL) goto fail;
	if((pglEdgeFlagv = goglGetProcAddress(err="glEdgeFlagv")) == NULL) goto fail;
	if((pglEnable = goglGetProcAddress(err="glEnable")) == NULL) goto fail;
	if((pglEnableClientState = goglGetProcAddress(err="glEnableClientState")) == NULL) goto fail;
	if((pglEnableVertexAttribArray = goglGetProcAddress(err="glEnableVertexAttribArray")) == NULL) goto fail;
	if((pglEnd = goglGetProcAddress(err="glEnd")) == NULL) goto fail;
	if((pglEndList = goglGetProcAddress(err="glEndList")) == NULL) goto fail;
	if((pglEndQuery = goglGetProcAddress(err="glEndQuery")) == NULL) goto fail;
	if((pglEvalCoord1d = goglGetProcAddress(err="glEvalCoord1d")) == NULL) goto fail;
	if((pglEvalCoord1dv = goglGetProcAddress(err="glEvalCoord1dv")) == NULL) goto fail;
	if((pglEvalCoord1f = goglGetProcAddress(err="glEvalCoord1f")) == NULL) goto fail;
	if((pglEvalCoord1fv = goglGetProcAddress(err="glEvalCoord1fv")) == NULL) goto fail;
	if((pglEvalCoord2d = goglGetProcAddress(err="glEvalCoord2d")) == NULL) goto fail;
	if((pglEvalCoord2dv = goglGetProcAddress(err="glEvalCoord2dv")) == NULL) goto fail;
	if((pglEvalCoord2f = goglGetProcAddress(err="glEvalCoord2f")) == NULL) goto fail;
	if((pglEvalCoord2fv = goglGetProcAddress(err="glEvalCoord2fv")) == NULL) goto fail;
	if((pglEvalMesh1 = goglGetProcAddress(err="glEvalMesh1")) == NULL) goto fail;
	if((pglEvalMesh2 = goglGetProcAddress(err="glEvalMesh2")) == NULL) goto fail;
	if((pglEvalPoint1 = goglGetProcAddress(err="glEvalPoint1")) == NULL) goto fail;
	if((pglEvalPoint2 = goglGetProcAddress(err="glEvalPoint2")) == NULL) goto fail;
	if((pglFeedbackBuffer = goglGetProcAddress(err="glFeedbackBuffer")) == NULL) goto fail;
	if((pglFinish = goglGetProcAddress(err="glFinish")) == NULL) goto fail;
	if((pglFlush = goglGetProcAddress(err="glFlush")) == NULL) goto fail;
	if((pglFogCoordPointer = goglGetProcAddress(err="glFogCoordPointer")) == NULL) goto fail;
	if((pglFogCoordd = goglGetProcAddress(err="glFogCoordd")) == NULL) goto fail;
	if((pglFogCoorddv = goglGetProcAddress(err="glFogCoorddv")) == NULL) goto fail;
	if((pglFogCoordf = goglGetProcAddress(err="glFogCoordf")) == NULL) goto fail;
	if((pglFogCoordfv = goglGetProcAddress(err="glFogCoordfv")) == NULL) goto fail;
	if((pglFogf = goglGetProcAddress(err="glFogf")) == NULL) goto fail;
	if((pglFogfv = goglGetProcAddress(err="glFogfv")) == NULL) goto fail;
	if((pglFogi = goglGetProcAddress(err="glFogi")) == NULL) goto fail;
	if((pglFogiv = goglGetProcAddress(err="glFogiv")) == NULL) goto fail;
//	if((pglFogx = goglGetProcAddress(err="glFogx")) == NULL) goto fail;
//	if((pglFogxv = goglGetProcAddress(err="glFogxv")) == NULL) goto fail;
	if((pglFramebufferRenderbuffer = goglGetProcAddress(err="glFramebufferRenderbuffer")) == NULL) goto fail;
	if((pglFramebufferTexture2D = goglGetProcAddress(err="glFramebufferTexture2D")) == NULL) goto fail;
	if((pglFrontFace = goglGetProcAddress(err="glFrontFace")) == NULL) goto fail;
	if((pglFrustum = goglGetProcAddress(err="glFrustum")) == NULL) goto fail;
//	if((pglFrustumf = goglGetProcAddress(err="glFrustumf")) == NULL) goto fail;
//	if((pglFrustumx = goglGetProcAddress(err="glFrustumx")) == NULL) goto fail;
	if((pglGenBuffers = goglGetProcAddress(err="glGenBuffers")) == NULL) goto fail;
	if((pglGenFramebuffers = goglGetProcAddress(err="glGenFramebuffers")) == NULL) goto fail;
	if((pglGenLists = goglGetProcAddress(err="glGenLists")) == NULL) goto fail;
	if((pglGenQueries = goglGetProcAddress(err="glGenQueries")) == NULL) goto fail;
	if((pglGenRenderbuffers = goglGetProcAddress(err="glGenRenderbuffers")) == NULL) goto fail;
	if((pglGenTextures = goglGetProcAddress(err="glGenTextures")) == NULL) goto fail;
	if((pglGenerateMipmap = goglGetProcAddress(err="glGenerateMipmap")) == NULL) goto fail;
	if((pglGetActiveAttrib = goglGetProcAddress(err="glGetActiveAttrib")) == NULL) goto fail;
	if((pglGetActiveUniform = goglGetProcAddress(err="glGetActiveUniform")) == NULL) goto fail;
	if((pglGetAttachedShaders = goglGetProcAddress(err="glGetAttachedShaders")) == NULL) goto fail;
	if((pglGetAttribLocation = goglGetProcAddress(err="glGetAttribLocation")) == NULL) goto fail;
	if((pglGetBooleanv = goglGetProcAddress(err="glGetBooleanv")) == NULL) goto fail;
	if((pglGetBufferParameteriv = goglGetProcAddress(err="glGetBufferParameteriv")) == NULL) goto fail;
	if((pglGetBufferPointerv = goglGetProcAddress(err="glGetBufferPointerv")) == NULL) goto fail;
	if((pglGetBufferSubData = goglGetProcAddress(err="glGetBufferSubData")) == NULL) goto fail;
	if((pglGetClipPlane = goglGetProcAddress(err="glGetClipPlane")) == NULL) goto fail;
//	if((pglGetClipPlanef = goglGetProcAddress(err="glGetClipPlanef")) == NULL) goto fail;
//	if((pglGetClipPlanex = goglGetProcAddress(err="glGetClipPlanex")) == NULL) goto fail;
	if((pglGetCompressedTexImage = goglGetProcAddress(err="glGetCompressedTexImage")) == NULL) goto fail;
	if((pglGetDoublev = goglGetProcAddress(err="glGetDoublev")) == NULL) goto fail;
	if((pglGetError = goglGetProcAddress(err="glGetError")) == NULL) goto fail;
//	if((pglGetFixedv = goglGetProcAddress(err="glGetFixedv")) == NULL) goto fail;
	if((pglGetFloatv = goglGetProcAddress(err="glGetFloatv")) == NULL) goto fail;
	if((pglGetFramebufferAttachmentParameteriv = goglGetProcAddress(err="glGetFramebufferAttachmentParameteriv")) == NULL) goto fail;
	if((pglGetIntegerv = goglGetProcAddress(err="glGetIntegerv")) == NULL) goto fail;
	if((pglGetLightfv = goglGetProcAddress(err="glGetLightfv")) == NULL) goto fail;
	if((pglGetLightiv = goglGetProcAddress(err="glGetLightiv")) == NULL) goto fail;
//	if((pglGetLightxv = goglGetProcAddress(err="glGetLightxv")) == NULL) goto fail;
	if((pglGetMapdv = goglGetProcAddress(err="glGetMapdv")) == NULL) goto fail;
	if((pglGetMapfv = goglGetProcAddress(err="glGetMapfv")) == NULL) goto fail;
	if((pglGetMapiv = goglGetProcAddress(err="glGetMapiv")) == NULL) goto fail;
	if((pglGetMaterialfv = goglGetProcAddress(err="glGetMaterialfv")) == NULL) goto fail;
	if((pglGetMaterialiv = goglGetProcAddress(err="glGetMaterialiv")) == NULL) goto fail;
//	if((pglGetMaterialxv = goglGetProcAddress(err="glGetMaterialxv")) == NULL) goto fail;
	if((pglGetPixelMapfv = goglGetProcAddress(err="glGetPixelMapfv")) == NULL) goto fail;
	if((pglGetPixelMapuiv = goglGetProcAddress(err="glGetPixelMapuiv")) == NULL) goto fail;
	if((pglGetPixelMapusv = goglGetProcAddress(err="glGetPixelMapusv")) == NULL) goto fail;
	if((pglGetPointerv = goglGetProcAddress(err="glGetPointerv")) == NULL) goto fail;
	if((pglGetPolygonStipple = goglGetProcAddress(err="glGetPolygonStipple")) == NULL) goto fail;
	if((pglGetProgramInfoLog = goglGetProcAddress(err="glGetProgramInfoLog")) == NULL) goto fail;
	if((pglGetProgramiv = goglGetProcAddress(err="glGetProgramiv")) == NULL) goto fail;
	if((pglGetQueryObjectiv = goglGetProcAddress(err="glGetQueryObjectiv")) == NULL) goto fail;
	if((pglGetQueryObjectuiv = goglGetProcAddress(err="glGetQueryObjectuiv")) == NULL) goto fail;
	if((pglGetQueryiv = goglGetProcAddress(err="glGetQueryiv")) == NULL) goto fail;
	if((pglGetRenderbufferParameteriv = goglGetProcAddress(err="glGetRenderbufferParameteriv")) == NULL) goto fail;
	if((pglGetShaderInfoLog = goglGetProcAddress(err="glGetShaderInfoLog")) == NULL) goto fail;
	if((pglGetShaderPrecisionFormat = goglGetProcAddress(err="glGetShaderPrecisionFormat")) == NULL) goto fail;
	if((pglGetShaderSource = goglGetProcAddress(err="glGetShaderSource")) == NULL) goto fail;
	if((pglGetShaderiv = goglGetProcAddress(err="glGetShaderiv")) == NULL) goto fail;
	if((pglGetString = goglGetProcAddress(err="glGetString")) == NULL) goto fail;
	if((pglGetTexEnvfv = goglGetProcAddress(err="glGetTexEnvfv")) == NULL) goto fail;
	if((pglGetTexEnviv = goglGetProcAddress(err="glGetTexEnviv")) == NULL) goto fail;
//	if((pglGetTexEnvxv = goglGetProcAddress(err="glGetTexEnvxv")) == NULL) goto fail;
	if((pglGetTexGendv = goglGetProcAddress(err="glGetTexGendv")) == NULL) goto fail;
	if((pglGetTexGenfv = goglGetProcAddress(err="glGetTexGenfv")) == NULL) goto fail;
	if((pglGetTexGeniv = goglGetProcAddress(err="glGetTexGeniv")) == NULL) goto fail;
	if((pglGetTexImage = goglGetProcAddress(err="glGetTexImage")) == NULL) goto fail;
	if((pglGetTexLevelParameterfv = goglGetProcAddress(err="glGetTexLevelParameterfv")) == NULL) goto fail;
	if((pglGetTexLevelParameteriv = goglGetProcAddress(err="glGetTexLevelParameteriv")) == NULL) goto fail;
	if((pglGetTexParameterfv = goglGetProcAddress(err="glGetTexParameterfv")) == NULL) goto fail;
	if((pglGetTexParameteriv = goglGetProcAddress(err="glGetTexParameteriv")) == NULL) goto fail;
//	if((pglGetTexParameterxv = goglGetProcAddress(err="glGetTexParameterxv")) == NULL) goto fail;
	if((pglGetUniformLocation = goglGetProcAddress(err="glGetUniformLocation")) == NULL) goto fail;
	if((pglGetUniformfv = goglGetProcAddress(err="glGetUniformfv")) == NULL) goto fail;
	if((pglGetUniformiv = goglGetProcAddress(err="glGetUniformiv")) == NULL) goto fail;
	if((pglGetVertexAttribPointerv = goglGetProcAddress(err="glGetVertexAttribPointerv")) == NULL) goto fail;
	if((pglGetVertexAttribdv = goglGetProcAddress(err="glGetVertexAttribdv")) == NULL) goto fail;
	if((pglGetVertexAttribfv = goglGetProcAddress(err="glGetVertexAttribfv")) == NULL) goto fail;
	if((pglGetVertexAttribiv = goglGetProcAddress(err="glGetVertexAttribiv")) == NULL) goto fail;
	if((pglHint = goglGetProcAddress(err="glHint")) == NULL) goto fail;
	if((pglIndexMask = goglGetProcAddress(err="glIndexMask")) == NULL) goto fail;
	if((pglIndexPointer = goglGetProcAddress(err="glIndexPointer")) == NULL) goto fail;
	if((pglIndexd = goglGetProcAddress(err="glIndexd")) == NULL) goto fail;
	if((pglIndexdv = goglGetProcAddress(err="glIndexdv")) == NULL) goto fail;
	if((pglIndexf = goglGetProcAddress(err="glIndexf")) == NULL) goto fail;
	if((pglIndexfv = goglGetProcAddress(err="glIndexfv")) == NULL) goto fail;
	if((pglIndexi = goglGetProcAddress(err="glIndexi")) == NULL) goto fail;
	if((pglIndexiv = goglGetProcAddress(err="glIndexiv")) == NULL) goto fail;
	if((pglIndexs = goglGetProcAddress(err="glIndexs")) == NULL) goto fail;
	if((pglIndexsv = goglGetProcAddress(err="glIndexsv")) == NULL) goto fail;
	if((pglIndexub = goglGetProcAddress(err="glIndexub")) == NULL) goto fail;
	if((pglIndexubv = goglGetProcAddress(err="glIndexubv")) == NULL) goto fail;
	if((pglInitNames = goglGetProcAddress(err="glInitNames")) == NULL) goto fail;
	if((pglInterleavedArrays = goglGetProcAddress(err="glInterleavedArrays")) == NULL) goto fail;
	if((pglIsBuffer = goglGetProcAddress(err="glIsBuffer")) == NULL) goto fail;
	if((pglIsEnabled = goglGetProcAddress(err="glIsEnabled")) == NULL) goto fail;
	if((pglIsFramebuffer = goglGetProcAddress(err="glIsFramebuffer")) == NULL) goto fail;
	if((pglIsList = goglGetProcAddress(err="glIsList")) == NULL) goto fail;
	if((pglIsProgram = goglGetProcAddress(err="glIsProgram")) == NULL) goto fail;
	if((pglIsQuery = goglGetProcAddress(err="glIsQuery")) == NULL) goto fail;
	if((pglIsRenderbuffer = goglGetProcAddress(err="glIsRenderbuffer")) == NULL) goto fail;
	if((pglIsShader = goglGetProcAddress(err="glIsShader")) == NULL) goto fail;
	if((pglIsTexture = goglGetProcAddress(err="glIsTexture")) == NULL) goto fail;
	if((pglLightModelf = goglGetProcAddress(err="glLightModelf")) == NULL) goto fail;
	if((pglLightModelfv = goglGetProcAddress(err="glLightModelfv")) == NULL) goto fail;
	if((pglLightModeli = goglGetProcAddress(err="glLightModeli")) == NULL) goto fail;
	if((pglLightModeliv = goglGetProcAddress(err="glLightModeliv")) == NULL) goto fail;
//	if((pglLightModelx = goglGetProcAddress(err="glLightModelx")) == NULL) goto fail;
//	if((pglLightModelxv = goglGetProcAddress(err="glLightModelxv")) == NULL) goto fail;
	if((pglLightf = goglGetProcAddress(err="glLightf")) == NULL) goto fail;
	if((pglLightfv = goglGetProcAddress(err="glLightfv")) == NULL) goto fail;
	if((pglLighti = goglGetProcAddress(err="glLighti")) == NULL) goto fail;
	if((pglLightiv = goglGetProcAddress(err="glLightiv")) == NULL) goto fail;
//	if((pglLightx = goglGetProcAddress(err="glLightx")) == NULL) goto fail;
//	if((pglLightxv = goglGetProcAddress(err="glLightxv")) == NULL) goto fail;
	if((pglLineStipple = goglGetProcAddress(err="glLineStipple")) == NULL) goto fail;
	if((pglLineWidth = goglGetProcAddress(err="glLineWidth")) == NULL) goto fail;
//	if((pglLineWidthx = goglGetProcAddress(err="glLineWidthx")) == NULL) goto fail;
	if((pglLinkProgram = goglGetProcAddress(err="glLinkProgram")) == NULL) goto fail;
	if((pglListBase = goglGetProcAddress(err="glListBase")) == NULL) goto fail;
	if((pglLoadIdentity = goglGetProcAddress(err="glLoadIdentity")) == NULL) goto fail;
	if((pglLoadMatrixd = goglGetProcAddress(err="glLoadMatrixd")) == NULL) goto fail;
	if((pglLoadMatrixf = goglGetProcAddress(err="glLoadMatrixf")) == NULL) goto fail;
//	if((pglLoadMatrixx = goglGetProcAddress(err="glLoadMatrixx")) == NULL) goto fail;
	if((pglLoadName = goglGetProcAddress(err="glLoadName")) == NULL) goto fail;
	if((pglLoadTransposeMatrixd = goglGetProcAddress(err="glLoadTransposeMatrixd")) == NULL) goto fail;
	if((pglLoadTransposeMatrixf = goglGetProcAddress(err="glLoadTransposeMatrixf")) == NULL) goto fail;
	if((pglLogicOp = goglGetProcAddress(err="glLogicOp")) == NULL) goto fail;
	if((pglMap1d = goglGetProcAddress(err="glMap1d")) == NULL) goto fail;
	if((pglMap1f = goglGetProcAddress(err="glMap1f")) == NULL) goto fail;
	if((pglMap2d = goglGetProcAddress(err="glMap2d")) == NULL) goto fail;
	if((pglMap2f = goglGetProcAddress(err="glMap2f")) == NULL) goto fail;
	if((pglMapBuffer = goglGetProcAddress(err="glMapBuffer")) == NULL) goto fail;
	if((pglMapGrid1d = goglGetProcAddress(err="glMapGrid1d")) == NULL) goto fail;
	if((pglMapGrid1f = goglGetProcAddress(err="glMapGrid1f")) == NULL) goto fail;
	if((pglMapGrid2d = goglGetProcAddress(err="glMapGrid2d")) == NULL) goto fail;
	if((pglMapGrid2f = goglGetProcAddress(err="glMapGrid2f")) == NULL) goto fail;
	if((pglMaterialf = goglGetProcAddress(err="glMaterialf")) == NULL) goto fail;
	if((pglMaterialfv = goglGetProcAddress(err="glMaterialfv")) == NULL) goto fail;
	if((pglMateriali = goglGetProcAddress(err="glMateriali")) == NULL) goto fail;
	if((pglMaterialiv = goglGetProcAddress(err="glMaterialiv")) == NULL) goto fail;
//	if((pglMaterialx = goglGetProcAddress(err="glMaterialx")) == NULL) goto fail;
//	if((pglMaterialxv = goglGetProcAddress(err="glMaterialxv")) == NULL) goto fail;
	if((pglMatrixMode = goglGetProcAddress(err="glMatrixMode")) == NULL) goto fail;
	if((pglMultMatrixd = goglGetProcAddress(err="glMultMatrixd")) == NULL) goto fail;
	if((pglMultMatrixf = goglGetProcAddress(err="glMultMatrixf")) == NULL) goto fail;
//	if((pglMultMatrixx = goglGetProcAddress(err="glMultMatrixx")) == NULL) goto fail;
	if((pglMultTransposeMatrixd = goglGetProcAddress(err="glMultTransposeMatrixd")) == NULL) goto fail;
	if((pglMultTransposeMatrixf = goglGetProcAddress(err="glMultTransposeMatrixf")) == NULL) goto fail;
	if((pglMultiDrawArrays = goglGetProcAddress(err="glMultiDrawArrays")) == NULL) goto fail;
	if((pglMultiDrawElements = goglGetProcAddress(err="glMultiDrawElements")) == NULL) goto fail;
	if((pglMultiTexCoord1d = goglGetProcAddress(err="glMultiTexCoord1d")) == NULL) goto fail;
	if((pglMultiTexCoord1dv = goglGetProcAddress(err="glMultiTexCoord1dv")) == NULL) goto fail;
	if((pglMultiTexCoord1f = goglGetProcAddress(err="glMultiTexCoord1f")) == NULL) goto fail;
	if((pglMultiTexCoord1fv = goglGetProcAddress(err="glMultiTexCoord1fv")) == NULL) goto fail;
	if((pglMultiTexCoord1i = goglGetProcAddress(err="glMultiTexCoord1i")) == NULL) goto fail;
	if((pglMultiTexCoord1iv = goglGetProcAddress(err="glMultiTexCoord1iv")) == NULL) goto fail;
	if((pglMultiTexCoord1s = goglGetProcAddress(err="glMultiTexCoord1s")) == NULL) goto fail;
	if((pglMultiTexCoord1sv = goglGetProcAddress(err="glMultiTexCoord1sv")) == NULL) goto fail;
	if((pglMultiTexCoord2d = goglGetProcAddress(err="glMultiTexCoord2d")) == NULL) goto fail;
	if((pglMultiTexCoord2dv = goglGetProcAddress(err="glMultiTexCoord2dv")) == NULL) goto fail;
	if((pglMultiTexCoord2f = goglGetProcAddress(err="glMultiTexCoord2f")) == NULL) goto fail;
	if((pglMultiTexCoord2fv = goglGetProcAddress(err="glMultiTexCoord2fv")) == NULL) goto fail;
	if((pglMultiTexCoord2i = goglGetProcAddress(err="glMultiTexCoord2i")) == NULL) goto fail;
	if((pglMultiTexCoord2iv = goglGetProcAddress(err="glMultiTexCoord2iv")) == NULL) goto fail;
	if((pglMultiTexCoord2s = goglGetProcAddress(err="glMultiTexCoord2s")) == NULL) goto fail;
	if((pglMultiTexCoord2sv = goglGetProcAddress(err="glMultiTexCoord2sv")) == NULL) goto fail;
	if((pglMultiTexCoord3d = goglGetProcAddress(err="glMultiTexCoord3d")) == NULL) goto fail;
	if((pglMultiTexCoord3dv = goglGetProcAddress(err="glMultiTexCoord3dv")) == NULL) goto fail;
	if((pglMultiTexCoord3f = goglGetProcAddress(err="glMultiTexCoord3f")) == NULL) goto fail;
	if((pglMultiTexCoord3fv = goglGetProcAddress(err="glMultiTexCoord3fv")) == NULL) goto fail;
	if((pglMultiTexCoord3i = goglGetProcAddress(err="glMultiTexCoord3i")) == NULL) goto fail;
	if((pglMultiTexCoord3iv = goglGetProcAddress(err="glMultiTexCoord3iv")) == NULL) goto fail;
	if((pglMultiTexCoord3s = goglGetProcAddress(err="glMultiTexCoord3s")) == NULL) goto fail;
	if((pglMultiTexCoord3sv = goglGetProcAddress(err="glMultiTexCoord3sv")) == NULL) goto fail;
	if((pglMultiTexCoord4d = goglGetProcAddress(err="glMultiTexCoord4d")) == NULL) goto fail;
	if((pglMultiTexCoord4dv = goglGetProcAddress(err="glMultiTexCoord4dv")) == NULL) goto fail;
	if((pglMultiTexCoord4f = goglGetProcAddress(err="glMultiTexCoord4f")) == NULL) goto fail;
	if((pglMultiTexCoord4fv = goglGetProcAddress(err="glMultiTexCoord4fv")) == NULL) goto fail;
	if((pglMultiTexCoord4i = goglGetProcAddress(err="glMultiTexCoord4i")) == NULL) goto fail;
	if((pglMultiTexCoord4iv = goglGetProcAddress(err="glMultiTexCoord4iv")) == NULL) goto fail;
	if((pglMultiTexCoord4s = goglGetProcAddress(err="glMultiTexCoord4s")) == NULL) goto fail;
	if((pglMultiTexCoord4sv = goglGetProcAddress(err="glMultiTexCoord4sv")) == NULL) goto fail;
//	if((pglMultiTexCoord4x = goglGetProcAddress(err="glMultiTexCoord4x")) == NULL) goto fail;
	if((pglNewList = goglGetProcAddress(err="glNewList")) == NULL) goto fail;
	if((pglNormal3b = goglGetProcAddress(err="glNormal3b")) == NULL) goto fail;
	if((pglNormal3bv = goglGetProcAddress(err="glNormal3bv")) == NULL) goto fail;
	if((pglNormal3d = goglGetProcAddress(err="glNormal3d")) == NULL) goto fail;
	if((pglNormal3dv = goglGetProcAddress(err="glNormal3dv")) == NULL) goto fail;
	if((pglNormal3f = goglGetProcAddress(err="glNormal3f")) == NULL) goto fail;
	if((pglNormal3fv = goglGetProcAddress(err="glNormal3fv")) == NULL) goto fail;
	if((pglNormal3i = goglGetProcAddress(err="glNormal3i")) == NULL) goto fail;
	if((pglNormal3iv = goglGetProcAddress(err="glNormal3iv")) == NULL) goto fail;
	if((pglNormal3s = goglGetProcAddress(err="glNormal3s")) == NULL) goto fail;
	if((pglNormal3sv = goglGetProcAddress(err="glNormal3sv")) == NULL) goto fail;
//	if((pglNormal3x = goglGetProcAddress(err="glNormal3x")) == NULL) goto fail;
	if((pglNormalPointer = goglGetProcAddress(err="glNormalPointer")) == NULL) goto fail;
	if((pglOrtho = goglGetProcAddress(err="glOrtho")) == NULL) goto fail;
//	if((pglOrthof = goglGetProcAddress(err="glOrthof")) == NULL) goto fail;
//	if((pglOrthox = goglGetProcAddress(err="glOrthox")) == NULL) goto fail;
	if((pglPassThrough = goglGetProcAddress(err="glPassThrough")) == NULL) goto fail;
	if((pglPixelMapfv = goglGetProcAddress(err="glPixelMapfv")) == NULL) goto fail;
	if((pglPixelMapuiv = goglGetProcAddress(err="glPixelMapuiv")) == NULL) goto fail;
	if((pglPixelMapusv = goglGetProcAddress(err="glPixelMapusv")) == NULL) goto fail;
	if((pglPixelStoref = goglGetProcAddress(err="glPixelStoref")) == NULL) goto fail;
	if((pglPixelStorei = goglGetProcAddress(err="glPixelStorei")) == NULL) goto fail;
	if((pglPixelTransferf = goglGetProcAddress(err="glPixelTransferf")) == NULL) goto fail;
	if((pglPixelTransferi = goglGetProcAddress(err="glPixelTransferi")) == NULL) goto fail;
	if((pglPixelZoom = goglGetProcAddress(err="glPixelZoom")) == NULL) goto fail;
	if((pglPointParameterf = goglGetProcAddress(err="glPointParameterf")) == NULL) goto fail;
	if((pglPointParameterfv = goglGetProcAddress(err="glPointParameterfv")) == NULL) goto fail;
	if((pglPointParameteri = goglGetProcAddress(err="glPointParameteri")) == NULL) goto fail;
	if((pglPointParameteriv = goglGetProcAddress(err="glPointParameteriv")) == NULL) goto fail;
//	if((pglPointParameterx = goglGetProcAddress(err="glPointParameterx")) == NULL) goto fail;
//	if((pglPointParameterxv = goglGetProcAddress(err="glPointParameterxv")) == NULL) goto fail;
	if((pglPointSize = goglGetProcAddress(err="glPointSize")) == NULL) goto fail;
//	if((pglPointSizex = goglGetProcAddress(err="glPointSizex")) == NULL) goto fail;
	if((pglPolygonMode = goglGetProcAddress(err="glPolygonMode")) == NULL) goto fail;
	if((pglPolygonOffset = goglGetProcAddress(err="glPolygonOffset")) == NULL) goto fail;
//	if((pglPolygonOffsetx = goglGetProcAddress(err="glPolygonOffsetx")) == NULL) goto fail;
	if((pglPolygonStipple = goglGetProcAddress(err="glPolygonStipple")) == NULL) goto fail;
	if((pglPopAttrib = goglGetProcAddress(err="glPopAttrib")) == NULL) goto fail;
	if((pglPopClientAttrib = goglGetProcAddress(err="glPopClientAttrib")) == NULL) goto fail;
//	if((pglPopMatrix = goglGetProcAddress(err="glPopMatrix")) == NULL) goto fail;
	if((pglPopName = goglGetProcAddress(err="glPopName")) == NULL) goto fail;
	if((pglPrioritizeTextures = goglGetProcAddress(err="glPrioritizeTextures")) == NULL) goto fail;
	if((pglPushAttrib = goglGetProcAddress(err="glPushAttrib")) == NULL) goto fail;
	if((pglPushClientAttrib = goglGetProcAddress(err="glPushClientAttrib")) == NULL) goto fail;
//	if((pglPushMatrix = goglGetProcAddress(err="glPushMatrix")) == NULL) goto fail;
	if((pglPushName = goglGetProcAddress(err="glPushName")) == NULL) goto fail;
	if((pglRasterPos2d = goglGetProcAddress(err="glRasterPos2d")) == NULL) goto fail;
	if((pglRasterPos2dv = goglGetProcAddress(err="glRasterPos2dv")) == NULL) goto fail;
	if((pglRasterPos2f = goglGetProcAddress(err="glRasterPos2f")) == NULL) goto fail;
	if((pglRasterPos2fv = goglGetProcAddress(err="glRasterPos2fv")) == NULL) goto fail;
	if((pglRasterPos2i = goglGetProcAddress(err="glRasterPos2i")) == NULL) goto fail;
	if((pglRasterPos2iv = goglGetProcAddress(err="glRasterPos2iv")) == NULL) goto fail;
	if((pglRasterPos2s = goglGetProcAddress(err="glRasterPos2s")) == NULL) goto fail;
	if((pglRasterPos2sv = goglGetProcAddress(err="glRasterPos2sv")) == NULL) goto fail;
	if((pglRasterPos3d = goglGetProcAddress(err="glRasterPos3d")) == NULL) goto fail;
	if((pglRasterPos3dv = goglGetProcAddress(err="glRasterPos3dv")) == NULL) goto fail;
	if((pglRasterPos3f = goglGetProcAddress(err="glRasterPos3f")) == NULL) goto fail;
	if((pglRasterPos3fv = goglGetProcAddress(err="glRasterPos3fv")) == NULL) goto fail;
	if((pglRasterPos3i = goglGetProcAddress(err="glRasterPos3i")) == NULL) goto fail;
	if((pglRasterPos3iv = goglGetProcAddress(err="glRasterPos3iv")) == NULL) goto fail;
	if((pglRasterPos3s = goglGetProcAddress(err="glRasterPos3s")) == NULL) goto fail;
	if((pglRasterPos3sv = goglGetProcAddress(err="glRasterPos3sv")) == NULL) goto fail;
	if((pglRasterPos4d = goglGetProcAddress(err="glRasterPos4d")) == NULL) goto fail;
	if((pglRasterPos4dv = goglGetProcAddress(err="glRasterPos4dv")) == NULL) goto fail;
	if((pglRasterPos4f = goglGetProcAddress(err="glRasterPos4f")) == NULL) goto fail;
	if((pglRasterPos4fv = goglGetProcAddress(err="glRasterPos4fv")) == NULL) goto fail;
	if((pglRasterPos4i = goglGetProcAddress(err="glRasterPos4i")) == NULL) goto fail;
	if((pglRasterPos4iv = goglGetProcAddress(err="glRasterPos4iv")) == NULL) goto fail;
	if((pglRasterPos4s = goglGetProcAddress(err="glRasterPos4s")) == NULL) goto fail;
	if((pglRasterPos4sv = goglGetProcAddress(err="glRasterPos4sv")) == NULL) goto fail;
	if((pglReadBuffer = goglGetProcAddress(err="glReadBuffer")) == NULL) goto fail;
	if((pglReadPixels = goglGetProcAddress(err="glReadPixels")) == NULL) goto fail;
	if((pglRectd = goglGetProcAddress(err="glRectd")) == NULL) goto fail;
	if((pglRectdv = goglGetProcAddress(err="glRectdv")) == NULL) goto fail;
	if((pglRectf = goglGetProcAddress(err="glRectf")) == NULL) goto fail;
	if((pglRectfv = goglGetProcAddress(err="glRectfv")) == NULL) goto fail;
	if((pglRecti = goglGetProcAddress(err="glRecti")) == NULL) goto fail;
	if((pglRectiv = goglGetProcAddress(err="glRectiv")) == NULL) goto fail;
	if((pglRects = goglGetProcAddress(err="glRects")) == NULL) goto fail;
	if((pglRectsv = goglGetProcAddress(err="glRectsv")) == NULL) goto fail;
	if((pglReleaseShaderCompiler = goglGetProcAddress(err="glReleaseShaderCompiler")) == NULL) goto fail;
	if((pglRenderMode = goglGetProcAddress(err="glRenderMode")) == NULL) goto fail;
	if((pglRenderbufferStorage = goglGetProcAddress(err="glRenderbufferStorage")) == NULL) goto fail;
	if((pglRotated = goglGetProcAddress(err="glRotated")) == NULL) goto fail;
	if((pglRotatef = goglGetProcAddress(err="glRotatef")) == NULL) goto fail;
//	if((pglRotatex = goglGetProcAddress(err="glRotatex")) == NULL) goto fail;
	if((pglSampleCoverage = goglGetProcAddress(err="glSampleCoverage")) == NULL) goto fail;
//	if((pglSampleCoveragex = goglGetProcAddress(err="glSampleCoveragex")) == NULL) goto fail;
	if((pglScaled = goglGetProcAddress(err="glScaled")) == NULL) goto fail;
	if((pglScalef = goglGetProcAddress(err="glScalef")) == NULL) goto fail;
//	if((pglScalex = goglGetProcAddress(err="glScalex")) == NULL) goto fail;
	if((pglScissor = goglGetProcAddress(err="glScissor")) == NULL) goto fail;
	if((pglSecondaryColor3b = goglGetProcAddress(err="glSecondaryColor3b")) == NULL) goto fail;
	if((pglSecondaryColor3bv = goglGetProcAddress(err="glSecondaryColor3bv")) == NULL) goto fail;
	if((pglSecondaryColor3d = goglGetProcAddress(err="glSecondaryColor3d")) == NULL) goto fail;
	if((pglSecondaryColor3dv = goglGetProcAddress(err="glSecondaryColor3dv")) == NULL) goto fail;
	if((pglSecondaryColor3f = goglGetProcAddress(err="glSecondaryColor3f")) == NULL) goto fail;
	if((pglSecondaryColor3fv = goglGetProcAddress(err="glSecondaryColor3fv")) == NULL) goto fail;
	if((pglSecondaryColor3i = goglGetProcAddress(err="glSecondaryColor3i")) == NULL) goto fail;
	if((pglSecondaryColor3iv = goglGetProcAddress(err="glSecondaryColor3iv")) == NULL) goto fail;
	if((pglSecondaryColor3s = goglGetProcAddress(err="glSecondaryColor3s")) == NULL) goto fail;
	if((pglSecondaryColor3sv = goglGetProcAddress(err="glSecondaryColor3sv")) == NULL) goto fail;
	if((pglSecondaryColor3ub = goglGetProcAddress(err="glSecondaryColor3ub")) == NULL) goto fail;
	if((pglSecondaryColor3ubv = goglGetProcAddress(err="glSecondaryColor3ubv")) == NULL) goto fail;
	if((pglSecondaryColor3ui = goglGetProcAddress(err="glSecondaryColor3ui")) == NULL) goto fail;
	if((pglSecondaryColor3uiv = goglGetProcAddress(err="glSecondaryColor3uiv")) == NULL) goto fail;
	if((pglSecondaryColor3us = goglGetProcAddress(err="glSecondaryColor3us")) == NULL) goto fail;
	if((pglSecondaryColor3usv = goglGetProcAddress(err="glSecondaryColor3usv")) == NULL) goto fail;
	if((pglSecondaryColorPointer = goglGetProcAddress(err="glSecondaryColorPointer")) == NULL) goto fail;
	if((pglSelectBuffer = goglGetProcAddress(err="glSelectBuffer")) == NULL) goto fail;
	if((pglShadeModel = goglGetProcAddress(err="glShadeModel")) == NULL) goto fail;
	if((pglShaderBinary = goglGetProcAddress(err="glShaderBinary")) == NULL) goto fail;
	if((pglShaderSource = goglGetProcAddress(err="glShaderSource")) == NULL) goto fail;
	if((pglStencilFunc = goglGetProcAddress(err="glStencilFunc")) == NULL) goto fail;
	if((pglStencilFuncSeparate = goglGetProcAddress(err="glStencilFuncSeparate")) == NULL) goto fail;
	if((pglStencilMask = goglGetProcAddress(err="glStencilMask")) == NULL) goto fail;
	if((pglStencilMaskSeparate = goglGetProcAddress(err="glStencilMaskSeparate")) == NULL) goto fail;
	if((pglStencilOp = goglGetProcAddress(err="glStencilOp")) == NULL) goto fail;
	if((pglStencilOpSeparate = goglGetProcAddress(err="glStencilOpSeparate")) == NULL) goto fail;
	if((pglTexCoord1d = goglGetProcAddress(err="glTexCoord1d")) == NULL) goto fail;
	if((pglTexCoord1dv = goglGetProcAddress(err="glTexCoord1dv")) == NULL) goto fail;
	if((pglTexCoord1f = goglGetProcAddress(err="glTexCoord1f")) == NULL) goto fail;
	if((pglTexCoord1fv = goglGetProcAddress(err="glTexCoord1fv")) == NULL) goto fail;
	if((pglTexCoord1i = goglGetProcAddress(err="glTexCoord1i")) == NULL) goto fail;
	if((pglTexCoord1iv = goglGetProcAddress(err="glTexCoord1iv")) == NULL) goto fail;
	if((pglTexCoord1s = goglGetProcAddress(err="glTexCoord1s")) == NULL) goto fail;
	if((pglTexCoord1sv = goglGetProcAddress(err="glTexCoord1sv")) == NULL) goto fail;
	if((pglTexCoord2d = goglGetProcAddress(err="glTexCoord2d")) == NULL) goto fail;
	if((pglTexCoord2dv = goglGetProcAddress(err="glTexCoord2dv")) == NULL) goto fail;
	if((pglTexCoord2f = goglGetProcAddress(err="glTexCoord2f")) == NULL) goto fail;
	if((pglTexCoord2fv = goglGetProcAddress(err="glTexCoord2fv")) == NULL) goto fail;
	if((pglTexCoord2i = goglGetProcAddress(err="glTexCoord2i")) == NULL) goto fail;
	if((pglTexCoord2iv = goglGetProcAddress(err="glTexCoord2iv")) == NULL) goto fail;
	if((pglTexCoord2s = goglGetProcAddress(err="glTexCoord2s")) == NULL) goto fail;
	if((pglTexCoord2sv = goglGetProcAddress(err="glTexCoord2sv")) == NULL) goto fail;
	if((pglTexCoord3d = goglGetProcAddress(err="glTexCoord3d")) == NULL) goto fail;
	if((pglTexCoord3dv = goglGetProcAddress(err="glTexCoord3dv")) == NULL) goto fail;
	if((pglTexCoord3f = goglGetProcAddress(err="glTexCoord3f")) == NULL) goto fail;
	if((pglTexCoord3fv = goglGetProcAddress(err="glTexCoord3fv")) == NULL) goto fail;
	if((pglTexCoord3i = goglGetProcAddress(err="glTexCoord3i")) == NULL) goto fail;
	if((pglTexCoord3iv = goglGetProcAddress(err="glTexCoord3iv")) == NULL) goto fail;
	if((pglTexCoord3s = goglGetProcAddress(err="glTexCoord3s")) == NULL) goto fail;
	if((pglTexCoord3sv = goglGetProcAddress(err="glTexCoord3sv")) == NULL) goto fail;
	if((pglTexCoord4d = goglGetProcAddress(err="glTexCoord4d")) == NULL) goto fail;
	if((pglTexCoord4dv = goglGetProcAddress(err="glTexCoord4dv")) == NULL) goto fail;
	if((pglTexCoord4f = goglGetProcAddress(err="glTexCoord4f")) == NULL) goto fail;
	if((pglTexCoord4fv = goglGetProcAddress(err="glTexCoord4fv")) == NULL) goto fail;
	if((pglTexCoord4i = goglGetProcAddress(err="glTexCoord4i")) == NULL) goto fail;
	if((pglTexCoord4iv = goglGetProcAddress(err="glTexCoord4iv")) == NULL) goto fail;
	if((pglTexCoord4s = goglGetProcAddress(err="glTexCoord4s")) == NULL) goto fail;
	if((pglTexCoord4sv = goglGetProcAddress(err="glTexCoord4sv")) == NULL) goto fail;
	if((pglTexCoordPointer = goglGetProcAddress(err="glTexCoordPointer")) == NULL) goto fail;
	if((pglTexEnvf = goglGetProcAddress(err="glTexEnvf")) == NULL) goto fail;
	if((pglTexEnvfv = goglGetProcAddress(err="glTexEnvfv")) == NULL) goto fail;
	if((pglTexEnvi = goglGetProcAddress(err="glTexEnvi")) == NULL) goto fail;
	if((pglTexEnviv = goglGetProcAddress(err="glTexEnviv")) == NULL) goto fail;
//	if((pglTexEnvx = goglGetProcAddress(err="glTexEnvx")) == NULL) goto fail;
//	if((pglTexEnvxv = goglGetProcAddress(err="glTexEnvxv")) == NULL) goto fail;
	if((pglTexGend = goglGetProcAddress(err="glTexGend")) == NULL) goto fail;
	if((pglTexGendv = goglGetProcAddress(err="glTexGendv")) == NULL) goto fail;
	if((pglTexGenf = goglGetProcAddress(err="glTexGenf")) == NULL) goto fail;
	if((pglTexGenfv = goglGetProcAddress(err="glTexGenfv")) == NULL) goto fail;
	if((pglTexGeni = goglGetProcAddress(err="glTexGeni")) == NULL) goto fail;
	if((pglTexGeniv = goglGetProcAddress(err="glTexGeniv")) == NULL) goto fail;
	if((pglTexImage1D = goglGetProcAddress(err="glTexImage1D")) == NULL) goto fail;
	if((pglTexImage2D = goglGetProcAddress(err="glTexImage2D")) == NULL) goto fail;
	if((pglTexImage3D = goglGetProcAddress(err="glTexImage3D")) == NULL) goto fail;
	if((pglTexParameterf = goglGetProcAddress(err="glTexParameterf")) == NULL) goto fail;
	if((pglTexParameterfv = goglGetProcAddress(err="glTexParameterfv")) == NULL) goto fail;
	if((pglTexParameteri = goglGetProcAddress(err="glTexParameteri")) == NULL) goto fail;
	if((pglTexParameteriv = goglGetProcAddress(err="glTexParameteriv")) == NULL) goto fail;
//	if((pglTexParameterx = goglGetProcAddress(err="glTexParameterx")) == NULL) goto fail;
//	if((pglTexParameterxv = goglGetProcAddress(err="glTexParameterxv")) == NULL) goto fail;
	if((pglTexSubImage1D = goglGetProcAddress(err="glTexSubImage1D")) == NULL) goto fail;
	if((pglTexSubImage2D = goglGetProcAddress(err="glTexSubImage2D")) == NULL) goto fail;
	if((pglTexSubImage3D = goglGetProcAddress(err="glTexSubImage3D")) == NULL) goto fail;
	if((pglTranslated = goglGetProcAddress(err="glTranslated")) == NULL) goto fail;
	if((pglTranslatef = goglGetProcAddress(err="glTranslatef")) == NULL) goto fail;
//	if((pglTranslatex = goglGetProcAddress(err="glTranslatex")) == NULL) goto fail;
	if((pglUniform1f = goglGetProcAddress(err="glUniform1f")) == NULL) goto fail;
	if((pglUniform1fv = goglGetProcAddress(err="glUniform1fv")) == NULL) goto fail;
	if((pglUniform1i = goglGetProcAddress(err="glUniform1i")) == NULL) goto fail;
	if((pglUniform1iv = goglGetProcAddress(err="glUniform1iv")) == NULL) goto fail;
	if((pglUniform2f = goglGetProcAddress(err="glUniform2f")) == NULL) goto fail;
	if((pglUniform2fv = goglGetProcAddress(err="glUniform2fv")) == NULL) goto fail;
	if((pglUniform2i = goglGetProcAddress(err="glUniform2i")) == NULL) goto fail;
	if((pglUniform2iv = goglGetProcAddress(err="glUniform2iv")) == NULL) goto fail;
	if((pglUniform3f = goglGetProcAddress(err="glUniform3f")) == NULL) goto fail;
	if((pglUniform3fv = goglGetProcAddress(err="glUniform3fv")) == NULL) goto fail;
	if((pglUniform3i = goglGetProcAddress(err="glUniform3i")) == NULL) goto fail;
	if((pglUniform3iv = goglGetProcAddress(err="glUniform3iv")) == NULL) goto fail;
	if((pglUniform4f = goglGetProcAddress(err="glUniform4f")) == NULL) goto fail;
	if((pglUniform4fv = goglGetProcAddress(err="glUniform4fv")) == NULL) goto fail;
	if((pglUniform4i = goglGetProcAddress(err="glUniform4i")) == NULL) goto fail;
	if((pglUniform4iv = goglGetProcAddress(err="glUniform4iv")) == NULL) goto fail;
	if((pglUniformMatrix2fv = goglGetProcAddress(err="glUniformMatrix2fv")) == NULL) goto fail;
	if((pglUniformMatrix3fv = goglGetProcAddress(err="glUniformMatrix3fv")) == NULL) goto fail;
	if((pglUniformMatrix4fv = goglGetProcAddress(err="glUniformMatrix4fv")) == NULL) goto fail;
	if((pglUnmapBuffer = goglGetProcAddress(err="glUnmapBuffer")) == NULL) goto fail;
	if((pglUseProgram = goglGetProcAddress(err="glUseProgram")) == NULL) goto fail;
	if((pglValidateProgram = goglGetProcAddress(err="glValidateProgram")) == NULL) goto fail;
	if((pglVertex2d = goglGetProcAddress(err="glVertex2d")) == NULL) goto fail;
	if((pglVertex2dv = goglGetProcAddress(err="glVertex2dv")) == NULL) goto fail;
	if((pglVertex2f = goglGetProcAddress(err="glVertex2f")) == NULL) goto fail;
	if((pglVertex2fv = goglGetProcAddress(err="glVertex2fv")) == NULL) goto fail;
	if((pglVertex2i = goglGetProcAddress(err="glVertex2i")) == NULL) goto fail;
	if((pglVertex2iv = goglGetProcAddress(err="glVertex2iv")) == NULL) goto fail;
	if((pglVertex2s = goglGetProcAddress(err="glVertex2s")) == NULL) goto fail;
	if((pglVertex2sv = goglGetProcAddress(err="glVertex2sv")) == NULL) goto fail;
	if((pglVertex3d = goglGetProcAddress(err="glVertex3d")) == NULL) goto fail;
	if((pglVertex3dv = goglGetProcAddress(err="glVertex3dv")) == NULL) goto fail;
	if((pglVertex3f = goglGetProcAddress(err="glVertex3f")) == NULL) goto fail;
	if((pglVertex3fv = goglGetProcAddress(err="glVertex3fv")) == NULL) goto fail;
	if((pglVertex3i = goglGetProcAddress(err="glVertex3i")) == NULL) goto fail;
	if((pglVertex3iv = goglGetProcAddress(err="glVertex3iv")) == NULL) goto fail;
	if((pglVertex3s = goglGetProcAddress(err="glVertex3s")) == NULL) goto fail;
	if((pglVertex3sv = goglGetProcAddress(err="glVertex3sv")) == NULL) goto fail;
	if((pglVertex4d = goglGetProcAddress(err="glVertex4d")) == NULL) goto fail;
	if((pglVertex4dv = goglGetProcAddress(err="glVertex4dv")) == NULL) goto fail;
	if((pglVertex4f = goglGetProcAddress(err="glVertex4f")) == NULL) goto fail;
	if((pglVertex4fv = goglGetProcAddress(err="glVertex4fv")) == NULL) goto fail;
	if((pglVertex4i = goglGetProcAddress(err="glVertex4i")) == NULL) goto fail;
	if((pglVertex4iv = goglGetProcAddress(err="glVertex4iv")) == NULL) goto fail;
	if((pglVertex4s = goglGetProcAddress(err="glVertex4s")) == NULL) goto fail;
	if((pglVertex4sv = goglGetProcAddress(err="glVertex4sv")) == NULL) goto fail;
	if((pglVertexAttrib1d = goglGetProcAddress(err="glVertexAttrib1d")) == NULL) goto fail;
	if((pglVertexAttrib1dv = goglGetProcAddress(err="glVertexAttrib1dv")) == NULL) goto fail;
	if((pglVertexAttrib1f = goglGetProcAddress(err="glVertexAttrib1f")) == NULL) goto fail;
	if((pglVertexAttrib1fv = goglGetProcAddress(err="glVertexAttrib1fv")) == NULL) goto fail;
	if((pglVertexAttrib1s = goglGetProcAddress(err="glVertexAttrib1s")) == NULL) goto fail;
	if((pglVertexAttrib1sv = goglGetProcAddress(err="glVertexAttrib1sv")) == NULL) goto fail;
	if((pglVertexAttrib2d = goglGetProcAddress(err="glVertexAttrib2d")) == NULL) goto fail;
	if((pglVertexAttrib2dv = goglGetProcAddress(err="glVertexAttrib2dv")) == NULL) goto fail;
	if((pglVertexAttrib2f = goglGetProcAddress(err="glVertexAttrib2f")) == NULL) goto fail;
	if((pglVertexAttrib2fv = goglGetProcAddress(err="glVertexAttrib2fv")) == NULL) goto fail;
	if((pglVertexAttrib2s = goglGetProcAddress(err="glVertexAttrib2s")) == NULL) goto fail;
	if((pglVertexAttrib2sv = goglGetProcAddress(err="glVertexAttrib2sv")) == NULL) goto fail;
	if((pglVertexAttrib3d = goglGetProcAddress(err="glVertexAttrib3d")) == NULL) goto fail;
	if((pglVertexAttrib3dv = goglGetProcAddress(err="glVertexAttrib3dv")) == NULL) goto fail;
	if((pglVertexAttrib3f = goglGetProcAddress(err="glVertexAttrib3f")) == NULL) goto fail;
	if((pglVertexAttrib3fv = goglGetProcAddress(err="glVertexAttrib3fv")) == NULL) goto fail;
	if((pglVertexAttrib3s = goglGetProcAddress(err="glVertexAttrib3s")) == NULL) goto fail;
	if((pglVertexAttrib3sv = goglGetProcAddress(err="glVertexAttrib3sv")) == NULL) goto fail;
	if((pglVertexAttrib4Nbv = goglGetProcAddress(err="glVertexAttrib4Nbv")) == NULL) goto fail;
	if((pglVertexAttrib4Niv = goglGetProcAddress(err="glVertexAttrib4Niv")) == NULL) goto fail;
	if((pglVertexAttrib4Nsv = goglGetProcAddress(err="glVertexAttrib4Nsv")) == NULL) goto fail;
	if((pglVertexAttrib4Nub = goglGetProcAddress(err="glVertexAttrib4Nub")) == NULL) goto fail;
	if((pglVertexAttrib4Nubv = goglGetProcAddress(err="glVertexAttrib4Nubv")) == NULL) goto fail;
	if((pglVertexAttrib4Nuiv = goglGetProcAddress(err="glVertexAttrib4Nuiv")) == NULL) goto fail;
	if((pglVertexAttrib4Nusv = goglGetProcAddress(err="glVertexAttrib4Nusv")) == NULL) goto fail;
	if((pglVertexAttrib4bv = goglGetProcAddress(err="glVertexAttrib4bv")) == NULL) goto fail;
	if((pglVertexAttrib4d = goglGetProcAddress(err="glVertexAttrib4d")) == NULL) goto fail;
	if((pglVertexAttrib4dv = goglGetProcAddress(err="glVertexAttrib4dv")) == NULL) goto fail;
	if((pglVertexAttrib4f = goglGetProcAddress(err="glVertexAttrib4f")) == NULL) goto fail;
	if((pglVertexAttrib4fv = goglGetProcAddress(err="glVertexAttrib4fv")) == NULL) goto fail;
	if((pglVertexAttrib4iv = goglGetProcAddress(err="glVertexAttrib4iv")) == NULL) goto fail;
	if((pglVertexAttrib4s = goglGetProcAddress(err="glVertexAttrib4s")) == NULL) goto fail;
	if((pglVertexAttrib4sv = goglGetProcAddress(err="glVertexAttrib4sv")) == NULL) goto fail;
	if((pglVertexAttrib4ubv = goglGetProcAddress(err="glVertexAttrib4ubv")) == NULL) goto fail;
	if((pglVertexAttrib4uiv = goglGetProcAddress(err="glVertexAttrib4uiv")) == NULL) goto fail;
	if((pglVertexAttrib4usv = goglGetProcAddress(err="glVertexAttrib4usv")) == NULL) goto fail;
	if((pglVertexAttribPointer = goglGetProcAddress(err="glVertexAttribPointer")) == NULL) goto fail;
	if((pglVertexPointer = goglGetProcAddress(err="glVertexPointer")) == NULL) goto fail;
	if((pglViewport = goglGetProcAddress(err="glViewport")) == NULL) goto fail;
	if((pglWindowPos2d = goglGetProcAddress(err="glWindowPos2d")) == NULL) goto fail;
	if((pglWindowPos2dv = goglGetProcAddress(err="glWindowPos2dv")) == NULL) goto fail;
	if((pglWindowPos2f = goglGetProcAddress(err="glWindowPos2f")) == NULL) goto fail;
	if((pglWindowPos2fv = goglGetProcAddress(err="glWindowPos2fv")) == NULL) goto fail;
	if((pglWindowPos2i = goglGetProcAddress(err="glWindowPos2i")) == NULL) goto fail;
	if((pglWindowPos2iv = goglGetProcAddress(err="glWindowPos2iv")) == NULL) goto fail;
	if((pglWindowPos2s = goglGetProcAddress(err="glWindowPos2s")) == NULL) goto fail;
	if((pglWindowPos2sv = goglGetProcAddress(err="glWindowPos2sv")) == NULL) goto fail;
	if((pglWindowPos3d = goglGetProcAddress(err="glWindowPos3d")) == NULL) goto fail;
	if((pglWindowPos3dv = goglGetProcAddress(err="glWindowPos3dv")) == NULL) goto fail;
	if((pglWindowPos3f = goglGetProcAddress(err="glWindowPos3f")) == NULL) goto fail;
	if((pglWindowPos3fv = goglGetProcAddress(err="glWindowPos3fv")) == NULL) goto fail;
	if((pglWindowPos3i = goglGetProcAddress(err="glWindowPos3i")) == NULL) goto fail;
	if((pglWindowPos3iv = goglGetProcAddress(err="glWindowPos3iv")) == NULL) goto fail;
	if((pglWindowPos3s = goglGetProcAddress(err="glWindowPos3s")) == NULL) goto fail;
	if((pglWindowPos3sv = goglGetProcAddress(err="glWindowPos3sv")) == NULL) goto fail;
	return NULL;
fail:
	return strdup(err);
}
