/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-12-26 19:29
**/

package system

import (
	"fmt"
	"os/exec"
	"runtime"
)

var ch = make(chan int)

func OpenBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func Exit(code int) {
	ch <- code
}

func Block() int {
	return <-ch
}
