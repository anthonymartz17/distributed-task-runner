package workerpool

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

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

func NewWorkerPool(worker_count int,store store.Storer) *WorkerPool{
	ctx,cancel:= context.WithCancel(context.Background())

	return&WorkerPool{
	 wg: &sync.WaitGroup{},
   numWorkers: worker_count,
	 taskCh: make(chan *domain.Task, worker_count * 4),
	 store: store,
	 ctx: ctx,
	 cancel: cancel,
	}
}

func(wp *WorkerPool) Start(){
	wp.wg.Add(wp.numWorkers)
	
	for i:= 0; i < wp.numWorkers; i++{
		go func(){
			defer wp.wg.Done()

			for{
				select{
				case <- wp.ctx.Done():
					return
				case task:= <- wp.taskCh:
					res:= wp.processTask(task)
					wp.store.Set(res.TaskId,res)
				}
			}

		}()
	}


}

func (wp *WorkerPool) Stop() {
	wp.cancel()
	wp.wg.Wait()
}

func (wp *WorkerPool) HandleGetTask(w http.ResponseWriter,req *http.Request){
	if req.Method != http.MethodGet{
		http.Error(w,"Only GET is allowed", http.StatusMethodNotAllowed)
	}
  
	URLParts:= strings.Split(req.RequestURI, "/")
  id:= URLParts[len(URLParts) - 1]

	res,ok:= wp.store.Get(id)

	if !ok {
		http.Error(w, "Task not ready", http.StatusAccepted)
		return
	}
	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
 
}

func(wp *WorkerPool) HandleTasks(w http.ResponseWriter,req *http.Request){
	
	if req.Method != http.MethodPost{
		http.Error(w,"Only POST is allowed",http.StatusMethodNotAllowed)
		return
	}

	body,err:= io.ReadAll(req.Body)
   
	if err != nil{
		http.Error(w,"Failed to read body",http.StatusBadRequest)
		return
	}

	var task *domain.Task

	if err:= json.Unmarshal(body,&task); err != nil{
    http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	select {
	case wp.taskCh <- task:
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Task accepted"))
	default:
		http.Error(w, "Worker pool is full", http.StatusServiceUnavailable)
	}

	
}

func (wp *WorkerPool) processTask(task *domain.Task) *domain.Result {
   
	switch task.Type{
	case "word_count":
		return HandleWordCount(task)
	case "reverse_array_int":
		return HandleReverseArrayInt(task)

	default:
	 return &domain.Result{
		TaskId: task.Id,
		Error: "unknown task type",
		CompletedAt: time.Now(),

	 }
	}
	

}