package containers

import (
	"bytes"
	"os/exec"

	"codeberg.org/iklabib/markisa/model"
)

type Minijail struct {
	Path       string
	ConfigPath string
}

func NewMinijail() Minijail {
	path, err := exec.LookPath("minijail0")
	if err != nil {
		panic(err)
	}

	return Minijail{
		Path:       path,
		ConfigPath: "./configs/minijail.cfg",
	}
}

func (mn Minijail) argsBuilder(dir string, commands []string) []string {
	// keep in mind that minijail need absolute path
	// is there a way for it to look in path without bash invocation?
	args := []string{"--config", mn.ConfigPath, "-C", dir, "--"}
	return append(args, commands...)
}

func (mn Minijail) ExecConfined(dir string, commands []string) model.SandboxExecResult {
	args := mn.argsBuilder(dir, commands)

	var stdoutBuff bytes.Buffer
	var stderrBuff bytes.Buffer

	cmd := exec.Command(mn.Path, args...)
	cmd.Stdout = &stdoutBuff
	cmd.Stderr = &stderrBuff
	cmd.Dir = dir
	err := cmd.Run()

	return model.SandboxExecResult{
		Error:  err,
		Stdout: stdoutBuff.String(),
		Stderr: stderrBuff.String(),
	}
}