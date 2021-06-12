cmake_minimum_required(VERSION 3.14)
project(bls-go)

# C/C++项目
add_subdirectory(./bls bls)

add_custom_target(${PROJECT_NAME}
# 编译C/C++
COMMAND cmake --build . -- -j8
# 编译go
COMMAND go build ${CMAKE_CURRENT_SOURCE_DIR}
)
