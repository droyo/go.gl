load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "buffers.go",
        "doc.go",
        "error.go",
        "gl.go",
        "gogl.c",
        "gogl.h",
        "objects.go",
        "shaders.go",
        "util.go",
    ],
    cgo = True,
    clinkopts = select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "-framework OpenGL",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "-lGL",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "-lopengl32",
        ],
        "//conditions:default": [],
    }),
    importpath = "aqwari.net/exp/gl",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["gl_test.go"],
    embed = [":go_default_library"],
    deps = ["//aqwari.net/exp/display:go_default_library"],
)
