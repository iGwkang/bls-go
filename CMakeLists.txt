cmake_minimum_required(VERSION 3.14)
project(bls-go)

# C/C++œÓƒø
add_subdirectory(./bls bls)

add_custom_target(${PROJECT_NAME}
# ±‡“ÎC/C++
COMMAND cmake --build . -- -j8
# ±‡“Îgo
COMMAND go build ${CMAKE_CURRENT_SOURCE_DIR}
)
