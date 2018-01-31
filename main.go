package main

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

func showMsg(msg string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("powershell", "msg *", msg).Run()
	}
	return errors.New("Unsupport")
}

func main() {
	showMsg("hi")

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	<-chSig
}
