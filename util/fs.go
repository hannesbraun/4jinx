package util

import (
	"os"
	"os/exec"
)

// File system related utilities

func Mkdir(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		mkdirCommand := exec.Command("mkdir", name)
		mkdirCommand.Run()
		return
	}
}
