package gp

import "fmt"

//GoroutinePool ...
type GoroutinePool struct {
	Queue  chan func() error
	Number int
	Total  int

	result         chan error
	finishCallback func()
}

//Init 初始化
func (g *GoroutinePool) Init(number int, total int) {
	g.Queue = make(chan func() error, total)
	g.Number = number
	g.Total = total
	g.result = make(chan error, total)
}

//Start 开门接客
func (g *GoroutinePool) Start() {
	// 开启Number个goroutine
	for i := 0; i < g.Number; i++ {
		go func() {
			for {
				task, ok := <-g.Queue
				if !ok {
					break
				}

				err := task()
				g.result <- err
			}
		}()
	}

	// 获得每个work的执行结果
	for j := 0; j < g.Total; j++ {
		res, ok := <-g.result
		if !ok {
			break
		}

		if res != nil {
			fmt.Println(res)
		}
	}

	// 所有任务都执行完成，回调函数
	if g.finishCallback != nil {
		g.finishCallback()
	}
}

//Stop 关门送客
func (g *GoroutinePool) Stop() {
	close(g.Queue)
	close(g.result)
}

//AddTask 添加任务
func (g *GoroutinePool) AddTask(task func() error) {
	g.Queue <- task
}

//SetFinishCallback 设置结束回调
func (g *GoroutinePool) SetFinishCallback(callback func()) {
	g.finishCallback = callback
}
