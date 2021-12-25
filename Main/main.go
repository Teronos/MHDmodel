package main

import (
	"FirstPrototip/Main/lib"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Запустился main.go")
	cmd, _ := exec.Command("python3", "Main/genMass.py").Output()
	fmt.Println(string(cmd))
	ar := lib.ProcessingStrTolist2dComplex(cmd)
	fmt.Println(ar)
	fmt.Println(len(ar))
}
