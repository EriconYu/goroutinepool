package main

import "github.com/cihub/seelog"

func main() {
	seelog.Info("Hello goroutinepool!!!")

	seelog.Info("Bye goroutinepool!!!")
	seelog.Flush()
}
