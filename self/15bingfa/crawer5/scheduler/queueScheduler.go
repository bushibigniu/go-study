package scheduler

import "go-study/self/15bingfa/crawer5/engine"

type QueueScheduler struct {
	//request 类型的chan
	RequestChan chan engine.Request
	//workerchan  is chan of work(chan engine.Reques)
	//chan 类型的chan
	workerChan chan chan engine.Request
}

//scheduler 实现第一步：所有worker 公用一个输入
//			 (带来问题：卡死，循环等待问题)
//scheduler 实现第2步：并发分发 request (每个request 创建一个goroutine)
//(带来问题：控制力度小，无法控制 goroutine)
//scheduler 实现第3步：实现 requests 队列 和 worker 队列

func (s *QueueScheduler) Submit(r engine.Request)  {
	s.RequestChan <- r
}

//func (s *QueueScheduler) ConfigureMasterWorkChan(c chan engine.Request)  {
//	panic("todo :do not write")
//}

func (s *QueueScheduler) WorkerReady(w chan engine.Request)  {
	//s.workerChan <- w
}

func (s *QueueScheduler) WorkerChan() chan engine.Request{
	return make(chan engine.Request)
}

func (s *QueueScheduler) Run()  {
	//make 改变了它的内容，所以必须使用指针接收者
	s.workerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		//chan of workQ
		var workerQ []chan engine.Request

		for  {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <- s.RequestChan:
				//send to request a ?
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				//sebd ? next_request to w
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}