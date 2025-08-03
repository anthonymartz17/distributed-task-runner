package worker

import (
	"context"
	"sync"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
	"github.com/anthonymartz17/distributed-task-runner/internal/store"
)
type WorkerPool struct{
	wg *sync.WaitGroup
	numWorkers int
	taskCh chan *domain.Task
	store store.Storer
	ctx context.Context
	cancel  context.CancelFunc
}

func NewWorkerPool(worker_count int) *WorkerPool{
	ctx,cancel:= context.WithCancel(context.Background())
	return&WorkerPool{
	 wg: &sync.WaitGroup{},
   numWorkers: worker_count,
	 taskCh: make(chan *domain.Task),
	 store: store.NewStore(),
	 ctx: ctx,
	 cancel: cancel,
	}
}

func(wp *WorkerPool)Start(){
	wp.wg.Add(wp.numWorkers)
	
	for i:= 0; i < wp.numWorkers; i++{
		
		go func(){
			defer wp.wg.Done()
			for{
				select {
				case <- wp.ctx.Done():
					return
				case task:= <- wp.taskCh:
					  //process task and save to store 

					
				}
			}

		}()

	}



}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
}
