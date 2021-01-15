# Practice for Cgo with CMake
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

## Run
```sh
# run on linux
./foo
# or on windows
foo.exe

# result
foo.cpp - data : Hello World!!
foo.cpp - mytest[1] : World!!
cwrap.cc from foo.cpp return : myReturn ^-^
foo.go from cwrap.cc return : myReturn ^-^

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
