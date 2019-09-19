package server

import (
	"bufio"
	"os/exec"
)

type commander struct {
}

func newCommander() commander {
	return commander{}
}

func (c commander) start(cmd *exec.Cmd) *bufio.Reader {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	return bufio.NewReader(stdout)
}
