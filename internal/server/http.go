package server

import (
	"fmt"
	"net/http"

	workerpool "github.com/anthonymartz17/distributed-task-runner/internal/workerPool"
)

func StartHTTPServer(wp *workerpool.WorkerPool){

	port:= ":8080"

	mux:= http.NewServeMux()

	mux.HandleFunc("/tasks",wp.HandleTasks)
	mux.HandleFunc("/tasks/{id}",wp.HandleGetTask)

	fmt.Printf("Server running on port: %v\n",port)
	err:= http.ListenAndServe(port,mux)

	if err != nil{
		fmt.Printf("failed to start server: %v\n",err)
		return
	}

}





