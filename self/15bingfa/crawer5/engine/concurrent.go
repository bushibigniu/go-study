package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int

}

/**
	这边虽然定义了interface，
	但go 的interface 是使用者定义的
	所以，只需要实现 Scheduler，下面的所有方法即可
	不需要 Scheduler，ReadyScheduler 两个 都实现

 */

//add scheduler func
type Scheduler interface {
	Submit(Request)
	//ConfigureMasterWorkChan(chan Request)
	Run()
	//我有一个worker ,给我那个chan 呢？
	WorkerChan() chan Request
	ReadyScheduler
}

//将 WorkerReady 独立出来，不然传整个Scheduler 太重了
type ReadyScheduler interface {
	//每个 worker 公用一个chan,还是每个chan 自己一个chan,这个只有scheduler 知道
	WorkerReady(chan Request)
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

		createWorker(e.Scheduler.WorkerChan() ,out, e.Scheduler)//向scheduler 要一个chan
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

//v1 func createWorker(in chan Request, out chan ParseResult)  {
//v2 func createWorker(out chan ParseResult, s Scheduler)  {
func createWorker(in chan Request,out chan ParseResult, ready ReadyScheduler)  {

	//in := make(chan Request)

	//不断从 in 收集request, 不断 out 处理
	go func() {
		for  {
			//tell scheduler i am ready
			ready.WorkerReady(in)
			request := <- in
			result , err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}