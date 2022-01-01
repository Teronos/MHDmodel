package main

import (
	"FirstPrototip/Main/lib"
	"fmt"

	"os/exec"
)

func main() {
	fmt.Println("lauching main.go")
	cmd, _ := exec.Command("python3", "Main/genMass.py").Output()
	fmt.Println(string(cmd))
	listComplexParametrs := lib.ProcessingStrTolist2dComplex(cmd)
	data := lib.GlobaloAlphaModel(listComplexParametrs, 1e-10, 10, 3000, 16, 60)
	//fmt.Println(string(data))
	cmd2, er := exec.Command("python3", "Main/Painter.py", string(data)).Output()
	if er == nil {
		//fmt.Println("done")
		fmt.Println(string(cmd2))
	} else {
		fmt.Println(er)

	}
}
