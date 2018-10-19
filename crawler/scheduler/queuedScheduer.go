package scheduler

import "Go_Spider/crawler/engine"

type QueuedSchduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueuedSchduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedSchduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedSchduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueuedSchduler) Run() {

	q.workChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
