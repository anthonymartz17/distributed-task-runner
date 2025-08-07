package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)

func StartHTTPServer(){

	port:= ":8080"

	mux:= http.NewServeMux()

	mux.HandleFunc("/tasks",handleTasks)

	fmt.Printf("Server running on port: %v\n",port)
	err:= http.ListenAndServe(port,mux)

	if err != nil{
		fmt.Printf("failed to start server: %v\n",err)
		return
	}

}


func handleTasks(w http.ResponseWriter,req *http.Request){
	
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


	switch task.Type{
	case "word_count":

		res,err:= handleWordCount(task)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)
		if err:= json.NewEncoder(w).Encode(res); err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
    
	case "reverse_array_int":

		res,err:= handleReverseArrayInt(task)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
    w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusOK)

		if err:= json.NewEncoder(w).Encode(res); err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

	default:
		http.Error(w,"Unknown task type",http.StatusBadRequest)
	}


}


func handleWordCount(task *domain.Task) (int,error){
  var text string

	if err:= json.Unmarshal(task.Payload,&text); err != nil{
		return 0,errors.New("invalid payload")
	}
	count:= len(strings.Fields(text))
	return count,nil
}

func handleReverseArrayInt(task *domain.Task) ([]int64,error){
 
	var result []int64

	if err:= json.Unmarshal(task.Payload,&result); err != nil{
		return nil,err
	}

	l:= 0
	r:= len(result) - 1

	for l < r{
   temp:= result[l]
	 result[l] =  result[r]
	 result[r] = temp
	 l++
	 r--
	}

	return result,nil
}