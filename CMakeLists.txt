cmake_minimum_required(VERSION 3.14)
project(bls-go)

add_subdirectory(./bls bls)

#add_custom_command(
#  OUTPUT "bls-go"
#  COMMAND go build ${CMAKE_CURRENT_SOURCE_DIR}
# )