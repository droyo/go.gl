WORK IN PROGRESS - INCOMPLETE

This is a set of OpenGL bindings meant to be easy to use from Go.
They were made by generating bindings using (Christoph Schunk's
GoGL2 utility)[0] and modifying the output to be easier to use with
Go. For example, setup functions like BindBuffer and CompileShader,
that are not expected to be in a tight loop, return errors, and
functions like glBufferData expect a slice of numbers and use
reflection to figure out the other arguments. Additional type safety is
achieved by explicitly defining types for OpenGL objects (Buffer, Program, etc).

[0]: https://github.com/chsc/gogl2
