package common

import (
	"os/exec"
)

func ExecScript(script string) (string, error) {
	cmd := exec.Command("CMD", "/c", "type", script)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(output)
	return str, nil
}
