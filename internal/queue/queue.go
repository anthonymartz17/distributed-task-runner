package queue

import (
	"encoding/json"
	"sync"
)

// Queue is a generic thread-safe FIFO queue backed by a slice
type Queue [T comparable] struct{
 data []T
 mu sync.Mutex
}

// QueueInterface defines the methods for the generic thread-safe FIFO queue
type QueueInterface[T comparable] interface{
	Enqueue(T) 
	Dequeue() (T, bool)
  Size() int
  IsEmpty() bool
}


func NewQueue[T comparable]()*Queue[T]{
  return &Queue[T]{
		data: []T{},
	}
}


//Enqueue adds an item to the queue
func (q *Queue[T])Enqueue(task T){
  q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data,task)
}
//Dequeue removes and returns first item in line from queue
func (q *Queue[T])Dequeue()(T, bool){
  var zeroValOf T

	q.mu.Lock()
	defer q.mu.Unlock()
	
	if len(q.data) == 0{
		return zeroValOf,false
	}

	removedItem:= q.data[0]
	q.data = q.data[1:]

	return removedItem,true
}


func (q *Queue[T])PrintQueue() (string,error){

	q.mu.Lock()
	defer q.mu.Unlock()

  items,err:= json.MarshalIndent(q.data,""," ")
	if err != nil{
		return "",err
	}
   
	return string(items),nil
}