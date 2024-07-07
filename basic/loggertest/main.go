package main

import "gotest1111/basic/mylogger"

func main() {
	log := mylogger.NewLog("Fatal")
	var s = "11111"
	var d = 12121
	log.Debug("这是一条debug信息%s %d", s, d)
	// mylogger.GetInfo(1)
}
