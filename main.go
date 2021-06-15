package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("install bls-go")
	command := exec.Command("cmake", cmakeArgs...)
	err := command.Run()
	if err != nil {
		log.Fatalln(err)
	}

	command = exec.Command("cmake", "--build", ".")
	err = command.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
