package main

import (
	"runtime"

	"github.com/anthonymartz17/distributed-task-runner/internal/server"
	"github.com/anthonymartz17/distributed-task-runner/internal/store"
	workerpool "github.com/anthonymartz17/distributed-task-runner/internal/workerPool"
)

func main(){
	
	store:= store.NewStore()
	wp:= workerpool.NewWorkerPool(runtime.NumCPU(),store)

	wp.Start()
  server.StartHTTPServer(wp)

}