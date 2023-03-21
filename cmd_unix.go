//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-01-02 19:05
**/

package utils

import (
	"os/exec"
	"strings"
)

type cm int

const Cmd cm = iota

type cmd struct {
	c *exec.Cmd
}

func (cm cm) New(command string) *cmd {
	var arr = strings.Split(command, " ")
	if len(arr) == 0 {
		panic("command is empty")
	}
	var c = exec.Command(arr[0], arr[1:]...)
	return &cmd{c: c}
}

func (c *cmd) Cmd() *exec.Cmd {
	return c.c
}
