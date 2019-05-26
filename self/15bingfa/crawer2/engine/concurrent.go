package engine

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int

}

//add scheduler func
type Scheduler interface {
	Submit(Request)
}

//scheduler 实现第一步：所有worker 公用一个输入
func (e ConcurrentEngine) Run(seed ...Request) {

	for _, r := range seed {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)

	for i:=0;i< e.WorkerCount ; i++ {
		//这边的worker 做：把scheduler 的 in  处理完， 输出 out 给 engine
		//createWorker(in, out)
	}

	for  {
		//这边的out 和上面的对应， 这时候 engine 收到的 out(也就是 requests, items)
		//要给 scheduler 处理
		result := <- out
		for _,item := range result.Items{
			fmt.Printf("got item : %v", item)
		}

		//这边掉起 scheduler
		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}

}
