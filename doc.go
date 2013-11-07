/*
Package gl provides a set of modern OpenGL bindings. The bindings
follow the OpenGL specification closely, with some adjustments to make
it easier to use with Go's type system and capabilites. Only modern
OpenGL functions that use the programmable rendering pipeline are
provided.

Additional changes are added for safety and convenience. OpenGL enumerations
are classified into types based on their usage, so using a value where it
should not be, such as

	gl.CreateShader(gl.ARRAY_BUFFER)

will result in a compile-time error. Initialization functions that are not
often in performance-critical code paths perform error checking. The shader
compilation routine will fetch the shader information log for you if an error
occurs.

There are few, but some, deviations from the OpenGL naming schemes. Functions
such as Uniform1f, Uniform2f, etc are collapsed into a single function of
multiple arguments. Type enumerations follow after those of the reflect.Kind
rather than the OpenGL identifiers. uintptr is used for offset arguments rather
than unsafe.Pointer. Functions that upload data, such as BufferData, take fewer
arguments and use reflection to discover the type and size of the data.

//BUG(droyo) This library is still in an experimental state. The API is subject to
// change.
*/
package gl
