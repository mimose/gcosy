package lib

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"syscall"
)

type Program struct {
	SimpleName string
	FileName   string
	Path       string
}

func Search(fileName string, simpleName string) (*Program, error) {
	path, err := exec.LookPath(fileName)
	if err != nil {
		return nil, err
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &Program{
		SimpleName: simpleName,
		FileName:   fileName,
		Path:       path,
	}, nil
}

func (p *Program) Command(args ...string) (stdout, stderr string, exitCode int, err error) {
	// build command with program's filePath and args
	command := exec.Command(p.Path, args...)
	// set the out&error about after command Run
	var stdo, stde bytes.Buffer
	command.Stdout = &stdo
	command.Stderr = &stde
	// go run command
	err = command.Run()
	// get the out&error
	stdout = string(stdo.Bytes())
	stderr = string(stde.Bytes())

	if err != nil {
		// err is instance of exec.ExitError, try to get exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = 1
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// not error, exist code will be 0
		ws := command.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}
