package utils

import (
	"log"
	"os"
	"os/exec"
)

func GetLogger() *log.Logger {
	l := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	return l
}

func runCommand(command string, args ...string) ([]byte, error) {

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}

	return output, nil
}
