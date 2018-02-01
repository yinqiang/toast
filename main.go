package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"runtime"
)

type Config struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func showMsg(msg string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("msg", "*", msg).Run()
	}
	return errors.New("Unsupport")
}

func main() {
	d, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		panic(err)
		return
	}
	conf := Config{}
	if err = json.Unmarshal(d, &conf); err != nil {
		panic(err)
		return
	}
	if len(conf.IP) == 0 {
		conf.IP = "0.0.0.0"
	}

	laddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", conf.IP, conf.Port))
	if err != nil {
		panic(err)
		return
	}
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		panic(err)
		return
	}
	// log.Printf("listen %s", laddr.String())
	showMsg("Server start")

	buf := make([]byte, 1024)
	for {
		rlen, _, err := conn.ReadFromUDP(buf)
		if err == nil {
			msg := string(buf[:rlen])
			// log.Printf("msg: %s. from: %s\n", msg, remote.String())
			if err = showMsg(msg); err != nil {
				log.Println(err)
			}
		}
	}
}
