package handler

import (
	"fmt"
	"os/exec"
)

func RunPythonScript(py string) {
	c := exec.Command("python", py)
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
