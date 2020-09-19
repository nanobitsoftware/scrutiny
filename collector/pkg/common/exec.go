package common

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"path"
)

func ExecCmd(cmdName string, cmdArgs []string, workingDir string, environ []string) (string, error) {

	cmd := exec.Command(cmdName, cmdArgs...)
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if environ != nil {
		cmd.Env = environ
	}
	if workingDir != "" && path.IsAbs(workingDir) {
		cmd.Dir = workingDir
	} else if workingDir != "" {
		return "", errors.New("Working Directory must be an absolute path")
	}

	err := cmd.Run()
	return stdBuffer.String(), err

}
