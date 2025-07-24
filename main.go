package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
	"github.com/anthonymartz17/distributed-task-runner/internal/queue"
	"github.com/google/uuid"
)

func main(){
q:= queue.NewQueue[*domain.Task]()

payload:= "test of first method"

task1:= &domain.Task{
	Id: uuid.NewString(),
	Type: "word_count",
	Payload: payload,
	CreatedAt: time.Now(),
}

q.Enqueue(task1)

task,ok:= q.Dequeue()

if !ok {
	fmt.Println("q is empty")
	return
}

removedTask,err:= json.MarshalIndent(task,""," ")

if err != nil{
	fmt.Println(err)
	return
}

fmt.Println(string(removedTask))


// fmt.Println(q.PrintQueue())
}