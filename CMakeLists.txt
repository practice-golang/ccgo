cmake_minimum_required(VERSION 3.0.0)

project(ccgoFooPractice)

# set( CMAKE_C_FLAGS "-ansi" )

add_library( foo c_cpp/foo.cpp c_cpp/cwrap.cc )

set(C_INCLUDE_PATH ${PROJECT_SOURCE_DIR}/c_cpp)
set(C_LIBRARY_NAMES foo stdc++)
set(C_LIBRARY_PATHS ./)

set(CFLAGS "-I${C_INCLUDE_PATH}")
set(LDFLAGS "-L${C_LIBRARY_PATHS}")

foreach (clib_name ${C_LIBRARY_NAMES})
    set(LDFLAGS "${LDFLAGS} -l${clib_name}")
endforeach()

set(INFILE ${PROJECT_SOURCE_DIR}/foo_vanilla.go)
set(OUTFILE ${PROJECT_BINARY_DIR}/foo.go)

configure_file( ${INFILE} ${OUTFILE} )

add_custom_target(
    ${PROJECT_NAME} ALL
    COMMAND go build ${OUTFILE}
)
