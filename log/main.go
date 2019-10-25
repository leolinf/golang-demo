package main

import "golang-demo/log/logger"

func main() {

	// 设置是从控制台输出还是日志文件
	logger.SetConsole(true)
	// 设置日志存放路径
	logger.SetRollingDaily("/Users/fanlin/Golang/src/golang-demo/log", "test.log", "test.error.log")
	// 设置日志级别
	loglevel, _ := logger.LoggerLevelIndex("DEBUG")
	logger.SetLevel(loglevel)

	logger.Debugf("Debug")
	logger.Infof("Info")
	logger.Warnf("Warn")
	logger.Errorf("Error")
	logger.Fatalf("Fatal")
	logger.Infof("Infof")
}
