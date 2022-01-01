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
	listComplexParametrs := lib.ProcessingStrTolist2dComplex(cmd)
	lib.GlobaloAlphaModel(listComplexParametrs, 1e-10, 10, 3000, 16, 60)

	//fmt.Println(listComplexParametrs)
	//fmt.Println(len(listComplexParametrs))
	cmd2, er := exec.Command("python3", "Main/Painter.py", "77").Output()
	if er == nil {
		fmt.Println("Done")
		fmt.Println(string(cmd2))
	} else {
		fmt.Println(er)
	}
}
