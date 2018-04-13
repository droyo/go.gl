load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["draw_triangle.go"],
    importpath = "aqwari.net/exp/gl/demo/draw_triangle",
    visibility = ["//visibility:private"],
    deps = [
        "//aqwari.net/exp/display:go_default_library",
        "//aqwari.net/exp/gl:go_default_library",
    ],
)

go_binary(
    name = "draw_triangle",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
