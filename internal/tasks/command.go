package tasks

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type CommandTask struct {
	cmd string
}

func NewCommandTask(command string) *CommandTask {
	return &CommandTask{
		cmd: command,
	}
}

func (ct *CommandTask) Run() error {
	// println(ct.cmd)
	args := strings.Split(ct.cmd, " ")
	cmd := exec.Command(args[0], args[1:]...)
	// if runtime.GOOS == "windows" {
	// 	cmd = exec.Command("tasklist")
	// }

	// var stdoutBuf, stderrBuf bytes.Buffer
	// cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	// cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// outStr, errStr := ), string(stderrBuf.Bytes())
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", stdoutBuf.String(), stderrBuf.String())
	return nil
}
