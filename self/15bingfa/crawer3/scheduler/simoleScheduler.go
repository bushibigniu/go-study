package scheduler

import "go-study/self/15bingfa/crawer3/engine"

//实现 scheduler interface

type SimpleScheduler struct {
	//send request down to worker chan
	workerChan chan engine.Request


}
//收集engine 送 过来的 request
func (s *SimpleScheduler) Submit(r engine.Request)  {

	//v1 s.workerChan <- r

	//v2 解决卡死，循环等待问题
	//这行关键，每个 request 创建一个 goroutine
	//存在问题：控制力度小，无法控制 goroutine
	go func(){ s.workerChan <- r }()

}

//ConfigureMasterWorkChan 会改变 scheduler 内容 所以用 指针
func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request)  {
	s.workerChan = c
}