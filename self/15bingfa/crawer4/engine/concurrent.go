package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int

}

//add scheduler func
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

//scheduler 实现第一步：所有worker 公用一个输入
//			 (带来问题：卡死，循环等待问题)
//scheduler 实现第2步：并发分发 request (每个request 创建一个goroutine)
			//(带来问题：控制力度小，无法控制 goroutine)
//scheduler 实现第3步：实现 requests 队列 和 worker 队列
func (e *ConcurrentEngine) Run(seed ...Request) {

	//in := make(chan Request)
	//out := make(chan ParseResult)
	//e.Scheduler.ConfigureMasterWorkChan(in)

	//v2
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:=0;i< e.WorkerCount ; i++ {
		//这边的worker 做：把scheduler 的 in  处理完， 输出 out 给 engine
		//createWorker(in, out)
		createWorker(out, e.Scheduler)
	}

	for _, r := range seed {
		e.Scheduler.Submit(r)
	}

	for  {
		//这边的out 和上面的对应， 这时候 engine 收到的 out(也就是 requests, items)
		//要给 scheduler 处理
		result := <- out
		for _,item := range result.Items{
			log.Printf("got item : %v", item)
		}

		//这边掉起 scheduler
		for _,request := range result.Requests{
			//这边是engine 送 request 给 scheduler
			//是直接调用，而不是通过channel
			e.Scheduler.Submit(request)
		}
	}

}

//func createWorker(in chan Request, out chan ParseResult)  {
func createWorker(out chan ParseResult, s Scheduler)  {

	in := make(chan Request)

	//不断从 in 收集request, 不断 out 处理
	go func() {
		for  {
			//tell scheduler i am ready
			s.WorkerReady(in)
			request := <- in
			result , err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}