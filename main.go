package main

import (
	"errors"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func showMsg(msg string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("powershell", "msg *", msg).Run()
	}
	return errors.New("Unsupport")
}

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8703")
	if err != nil {
		panic(err)
		return
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		panic(err)
		return
	}
	log.Printf("listen %s", laddr.String())

	buf := make([]byte, 1024)
	for {
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err == nil {
			msg := string(buf[:rlen])
			log.Printf("msg: %s. from: %s\n", msg, remote.String())
			showMsg(msg)
		}
	}
}
