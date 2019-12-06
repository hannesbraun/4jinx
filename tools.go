package main

import (
	"os"
	"os/exec"
)

func mkdir(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		mkdirCommand := exec.Command("mkdir", name)
		mkdirCommand.Run()
		return
	}
}
