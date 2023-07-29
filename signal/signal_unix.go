//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

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
	"os/signal"
	"syscall"
)

type Done struct {
	fn func(func(sig os.Signal))
}

func (d *Done) Done(fn func(sig os.Signal)) {
	d.fn(fn)
}

// ListenAll listen all signal
func ListenAll() *Done {
	var signalList = []os.Signal{
		syscall.SIGABRT, syscall.SIGALRM, syscall.SIGBUS, syscall.SIGCHLD, syscall.SIGCONT,
		syscall.SIGFPE, syscall.SIGHUP, syscall.SIGILL, syscall.SIGINT, syscall.SIGIO,
		syscall.SIGIOT, syscall.SIGKILL, syscall.SIGPIPE, syscall.SIGPROF, syscall.SIGQUIT,
		syscall.SIGSEGV, syscall.SIGSTOP, syscall.SIGSYS, syscall.SIGTERM, syscall.SIGTRAP,
		syscall.SIGTSTP, syscall.SIGTTIN, syscall.SIGTTOU, syscall.SIGURG, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGVTALRM, syscall.SIGWINCH, syscall.SIGXCPU, syscall.SIGXFSZ,
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, signalList...)
	return &Done{fn: func(f func(signal os.Signal)) {
		f(<-signalChan)
		signal.Stop(signalChan)
	}}
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
		signal.Stop(signalChan)
	}}
}

func Signal(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}

func KillGroup(pid int, sig syscall.Signal) error {
	return syscall.Kill(-pid, sig)
}
