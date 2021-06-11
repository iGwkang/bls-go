### Usage:
```
mkdir build && cd build
```

- windows build
```
cmake -G "MinGW Makefiles" ..
cmake --build . -- -j8
cd ..
go build .
```

- linux/mac build
```
cmake ..
cmake --build . -- -j8
cd ..
go build .
```