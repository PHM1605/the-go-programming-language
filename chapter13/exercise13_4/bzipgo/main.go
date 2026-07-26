package bzipgo

import (
	"io"
	"os/exec"
)

// writer: has 2 component
// - command-executor (with "bzip2") from Go library
// -
type writer struct {
	cmd  *exec.Cmd
	pipe io.WriteCloser // pipe from "Stdin" => command "cmd" (bzip2)
}

func (w *writer) Write(data []byte) (int, error) {
	return w.pipe.Write(data)
}

func (w *writer) Close() error {
	// tell io.WriteCloser "in" that we have no more input
	if err := w.pipe.Close(); err != nil {
		return err
	}
	// wait for command-executor to finish
	return w.cmd.Wait()
}

func NewWriter(out io.Writer) io.WriteCloser {
	// NEW: create command-executor, from command of Go library "bzip2"
	// Flow:
	cmd := exec.Command("bzip2")
	// compressed output goes here
	cmd.Stdout = out

	// create a pipe from "Stdin" and this command "cmd" (bzip2)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	return &writer{
		cmd:  cmd,
		pipe: pipe,
	}
}
