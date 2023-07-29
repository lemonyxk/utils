//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

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
	"strings"
)

type Cmd struct {
	c *exec.Cmd
}

func New(command string) *Cmd {
	var arr = strings.Split(command, " ")
	if len(arr) == 0 {
		panic("command is empty")
	}
	var c = exec.Command(arr[0], arr[1:]...)
	return &Cmd{c: c}
}

func (c *Cmd) Cmd() *exec.Cmd {
	return c.c
}
