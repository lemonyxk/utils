//go:build windows
// +build windows

/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2020-01-02 19:05
**/

package cmd

import (
	"os/exec"
)

type Cmd struct {
	c *exec.Cmd
}

func New(command string) *Cmd {
	var c = exec.Command("cmd", "/c", command)
	return &Cmd{c: c}
}

func (c *Cmd) Cmd() *exec.Cmd {
	return c.c
}
