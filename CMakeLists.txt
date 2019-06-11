cmake_minimum_required(VERSION 3.0.0)

project(gofoo)

# set( CMAKE_C_FLAGS "-ansi" )

add_library( foo c_cpp/foo.cpp c_cpp/cwrap.cc )

add_custom_target(
    ${PROJECT_NAME} ALL
    COMMAND go build .
)
