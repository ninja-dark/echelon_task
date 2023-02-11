package handler

import (
	"context"
	"fmt"
	"github.com/ninja-dark/echelon_task/internal/command"
)

type Hanlder struct {
	c *command.Executor
}

func NewHandler(c *command.Executor) *Hanlder {
	r := &Hanlder{
		c: c,
	}
	return r
}

func (rt *Hanlder) ExecuteCommand(ctx context.Context, cmd string, os string, stdin string) (command.Result, error) {
	answer, err := rt.c.ExecutingCommand(ctx, cmd, os, stdin)
	if err != nil {
		return command.Result{}, fmt.Errorf("can't execute command: %w", err)
	}

	return answer, nil
}
