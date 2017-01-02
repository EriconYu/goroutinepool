package main

import (
	"fmt"
	"gp"
	"time"

	"github.com/cihub/seelog"
)

func test() error {
	fmt.Println("time is ", time.Now().String())
	return nil
}

func main() {
	seelog.Info("Hello goroutinepool!!!")
	pool := new(gp.GoroutinePool)
	pool.Init(3, 15)
	for i := 0; i < 15; i++ {
		pool.AddTask(func() error {
			return test()
		})
		time.Sleep(10 * time.Microsecond)
	}

	isFinish := false

	pool.SetFinishCallback(func() {
		func(isFinish *bool) {
			*isFinish = true
			seelog.Info("Job is finished")
		}(&isFinish)
	})
	pool.Start()
	for {
		if isFinish == true {
			break
		} else {
			time.Sleep(time.Second)
		}
	}
	pool.Stop()

	//
	seelog.Info("Bye goroutinepool!!!")
	seelog.Flush()
}
