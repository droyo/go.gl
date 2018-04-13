load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "aqwari.net/exp/gl/demo/conway",
    visibility = ["//visibility:private"],
    deps = [
        "//aqwari.net/exp/display:go_default_library",
        "//aqwari.net/exp/gl:go_default_library",
    ],
)

go_binary(
    name = "conway",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
