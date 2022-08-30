package main

import (
	"github.com/go-demo-1/define"
	"github.com/go-demo-1/server"
	"os"
	"os/exec"
	"os/signal"
<<<<<<< HEAD
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chBrowserDie := make(chan struct{})
	chBackendDie := make(chan struct{})

	wg.Add(1)
	go server.Run()

	// 启动 Chrome Broswer
	go func() {
		chromePath := "C:\\Program Files (x86)\\Microsoft\\EdgeCore\\104.0.1293.70\\msedge.exe"
		cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+define.Port+"/static/index.html")
		cmd.Start()

		go func() {
			<-chBackendDie
			cmd.Process.Kill()
		}()

		go func() {
			cmd.Wait()
			chBrowserDie <- struct{}{}
		}()

	}()

	// 捕捉系统中断信号
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	// 浏览器退出 或者 系统中断
	select {
	case <-chSignal:
		chBackendDie <- struct{}{}
	case <-chBrowserDie:
		wg.Done()
	}
	wg.Wait()
}
=======
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

>>>>>>> c875f6b522271ef460e14ac2f54da591f25d97e1
