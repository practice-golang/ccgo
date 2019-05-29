# Practice for Cgo with Cpp
Worked using `MinGW`, `CMake` on `Windows`

## Build
```sh
cmake -G "MinGW Makefiles"

mingw32-make.exe
# or
cmake --build .
```

## Clean
```sh
rm -rf CMakeFiles
rm -rf cmake_install.cmake
rm -rf CMakeCache.txt
rm -rf Makefile
rm -rf libfoo.a
rm -rf gofoo.exe

```

## Reference
* https://github.com/burke/howto-go-with-cpp
* https://stackoverflow.com/a/1721230
