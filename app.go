package main

import (
)
import "github.com/vvotm/apiHahajok/boot"

func main() {
	app := boot.GetApp().AppInst
	app.Logger.Errorf("系统错误: %v", app.Start(":1223"))
}
