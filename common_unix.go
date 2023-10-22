package system_shutdown

import (
	"bytes"
	"errors"
	"os/exec"
)

func runCommand(command string, args ...string) (output string, err error) {
	cmd := exec.Command(command, args...)

	var stdoutBytes, stderrBytes bytes.Buffer
	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes

	err = cmd.Run()
	if err != nil {
		return "", err
	}
	if stdoutBytes.Len() > 0 {
		return "", errors.New(stderrBytes.String())
	}

	output = stdoutBytes.String()
	return output, err
}
