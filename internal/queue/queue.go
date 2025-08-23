package queue

import (
	"encoding/json"
	"sync"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)

// Queue is a generic thread-safe FIFO queue backed by a slice
type Queue  struct{
 data [] *domain.Task
 mu sync.Mutex
}

// TaskQueueer defines the methods for the generic thread-safe FIFO queue

type TaskQueueer interface{
	Enqueue(*domain.Task) 
	Dequeue() (*domain.Task, bool)
  Size() int
  IsEmpty() bool
}


func NewQueue()*Queue{
  return &Queue{
		data: []*domain.Task{},
	}
}


//Enqueue adds an item to the queue
func (q *Queue)Enqueue(task *domain.Task){
  q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data,task)
}
//Dequeue removes and returns first item in line from queue
func (q *Queue)Dequeue()(*domain.Task, bool){

	q.mu.Lock()
	defer q.mu.Unlock()
	
	if len(q.data) == 0{
		return nil,false
	}

	removedItem:= q.data[0]
	q.data = q.data[1:]

	return removedItem,true
}

func(q *Queue)Size() int{
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.data)
}

func(q *Queue)IsEmpty() bool{
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.data) == 0
}


func (q *Queue)PrintQueue() (string,error){

	q.mu.Lock()
	defer q.mu.Unlock()

  items,err:= json.MarshalIndent(q.data,""," ")
	if err != nil{
		return "",err
	}
   
	return string(items),nil
}