package common

import (
	"os/exec"
	"strings"
)

func ExecScript(script string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", script)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(output)
	return str, nil
}

func ExecShell(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := strings.TrimSuffix(strings.TrimSpace(string(output)), "\n")
	return str, nil
}
