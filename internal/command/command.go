package command

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type Executor struct {
}

func NewExecutor() *Executor {
	return &Executor{}
}

type Result struct {
	Stdout string
	Stderr string
}

func (c *Executor) ExecutingCommand(ctx context.Context, cmd string, os string, stdin string) (Result, error) {
	if err := c.checkOS(os); err != nil {
		return Result{}, err
	}
	stdout := &strings.Builder{}
	stderr := &strings.Builder{}
	name, args := c.parseArgs(cmd)
	command := exec.CommandContext(ctx, name, args...)
	command.Stdin = strings.NewReader(stdin)
	command.Stdout = stdout
	command.Stderr = stderr
	err := command.Run()
	if err != nil {
		return Result{}, fmt.Errorf("execute command: %w", err)
	}

	return Result{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}, nil
}

func (c *Executor) checkOS(os string) error {
	// проверить что текущ.os == os
	if runtime.GOOS != os {
		return fmt.Errorf("detect OS: %s", runtime.GOOS)
	}
	return nil
}

func (c *Executor) parseArgs(cmd string) (string, []string) {
	s := strings.Fields(cmd)
	cmd = s[0]
	s = s[1:]
	return cmd, s
}
