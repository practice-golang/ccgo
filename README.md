# Practice for Cgo with Cpp
I tried using `MinGW`, `CMake` on `Windows`

## Build
```sh
# Windows
cmake -G "MinGW Makefiles"
cmake --build .

#Linux
cmake -G "Unix Makefiles"
cmake --build .

```

## Clean
```sh
rm -rf CMakeFiles cmake_install.cmake CMakeCache.txt Makefile libfoo.a foo.go

# Windows
rm -f foo.exe

# Linux
rm -f foo

```

## Reference
* https://github.com/burke/howto-go-with-cpp
* https://github.com/donbright/go-hello-static-world
* https://stackoverflow.com/a/1721230