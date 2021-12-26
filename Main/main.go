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
	cmd2, er := exec.Command("python3", "Main/Painter.py", "77").Output()
	if er == nil {
		fmt.Println("Done")
		fmt.Println(string(cmd2))
	} else {
		fmt.Println(er)
	}
}
