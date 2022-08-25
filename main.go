package main

import (
	"github.com/go-demo-1/define"
	"github.com/go-demo-1/server"
	"os"
	"os/exec"
	"os/signal"
)


func main() {

	go server.Run()

	cmd := startChrome()

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	<-chSignal
	cmd.Process.Kill()
}

func startChrome() *exec.Cmd {
	chromePath := "C:\\Program Files (x86)\\Microsoft\\EdgeCore\\104.0.1293.63\\msedge.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+ define.Port + "/static/index.html")
	cmd.Start()
	return cmd
}

