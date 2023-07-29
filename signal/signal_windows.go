//go:build windows
// +build windows

/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-22 17:44
**/

package signal

import (
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

type Done struct {
	fn func(func(sig os.Signal))
}

func (d *Done) Done(fn func(sig os.Signal)) {
	d.fn(fn)
}

func ListenKill() *Done {
	var signalList = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, signalList...)
	return &Done{fn: func(f func(signal os.Signal)) {
		f(<-signalChan)
		signal.Stop(signalChan)
	}}
}

func Listen(sig ...os.Signal) *Done {
	var signalList = sig
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, signalList...)
	return &Done{fn: func(f func(signal os.Signal)) {
		f(<-signalChan)
		// 停止
		signal.Stop(signalChan)
	}}
}

func KillGroup(pid int, sig syscall.Signal) error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))
	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	return kill.Run()
}
