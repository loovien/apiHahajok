package main

import (
	"github.com/vvotm/apiHahajok/boot"
	"os/signal"
	"os"
)

func main() {
	app := boot.GetApp().AppInst
	go func() {
		app.Logger.Errorf("服务停止鸟: %v", app.Start(":1223"))
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	signalNum := <- sigChan
	app.Logger.Infof("接受系统信号: %v", signalNum)
}
