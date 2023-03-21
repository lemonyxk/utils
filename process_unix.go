//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-01-02 19:33
**/

package utils

import (
	"os"
	"strings"
	"sync"
	"syscall"
)

type pro int

const Process pro = iota

type process struct {
	Cmd *cmd
}

var pMux sync.Mutex

var worker []*process

var workerNumber int

var managerPid int

func (p pro) Fork(fn func(), number int) {

	switch os.Getenv("FORK_CHILD") {
	case "":
		managerPid = os.Getpid()
		_ = os.Setenv("FORK_CHILD", "TRUE")
		workerNumber = number
		p.run()
	default:
		go fn()
		Signal.ListenKill().Done(func(sig os.Signal) {
			os.Exit(0)
		})
	}
}

func (p pro) run() {
	pMux.Lock()
	defer pMux.Unlock()
	for i := 0; i < workerNumber; i++ {
		var c = Cmd.New(strings.Join(os.Args, " "))
		c.c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		err := c.c.Start()
		if err != nil {
			panic(err)
		}
		go func() { _, _ = c.c.Process.Wait() }()
		worker = append(worker, &process{Cmd: c})
	}
}

func (p pro) Kill(pid int) {
	pMux.Lock()
	defer pMux.Unlock()
	if managerPid == 0 {
		return
	}
	_ = Signal.KillGroup(pid, syscall.SIGTERM)
	for i := 0; i < len(worker); i++ {
		if worker[i].Cmd.c.Process.Pid == pid {
			worker = append(worker[0:i], worker[i+1:]...)
		}
	}
}

func (p pro) Reload() {
	if managerPid == 0 {
		return
	}
	for _, ps := range worker {
		p.Kill(ps.Cmd.c.Process.Pid)
	}
	p.run()
}

func (p pro) Manager() int {
	return managerPid
}

func (p pro) Worker() []*process {
	return worker
}
