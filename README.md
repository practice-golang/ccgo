# Practice for Cgo with Cpp
I tried using `MinGW`, `CMake` on `Windows`

## Build
```sh
# Windows
cmake -G "MinGW Makefiles"
cmake --build .
# or
mingw32-make.exe

#Linux
cmake -G "Unix Makefiles"
cmake --build .
# or
make

```

## Clean
```sh
rm -rf CMakeFiles
rm -rf cmake_install.cmake
rm -rf CMakeCache.txt
rm -rf Makefile
rm -rf libfoo.a

# Windows
rm -rf gofoo.exe

# Linux
rm -rf gofoo

```

## Reference
* https://github.com/burke/howto-go-with-cpp
* https://stackoverflow.com/a/1721230
