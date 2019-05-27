package scheduler

import "go-study/self/15bingfa/crawer4/engine"

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan chan chan engine.Request
}

func (s *QueueScheduler) Submit(r engine.Request)  {
	s.RequestChan <- r
}

func (s *QueueScheduler) ConfigureMasterWorkChan(c chan engine.Request)  {
	panic("todo :do not write")
}

func (s *QueueScheduler) WorkerReady(w chan engine.Request)  {
	s.WorkerChan <- w
}

func (s *QueueScheduler) Run()  {
	//make 改变了它的内容，所以必须使用指针接收者
	s.WorkerChan = make(chan chan engine.Request)
	s.RequestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
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
			case w := <- s.WorkerChan:
				//sebd ? next_request to w
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}