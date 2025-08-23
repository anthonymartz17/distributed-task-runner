package workerpool

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)

func HandleWordCount(task *domain.Task,ctx context.Context) (*domain.Result,error){
  var text string
	
	if err:= json.Unmarshal(task.Payload,&text); err != nil{
		return nil,err
	}
  words:=  strings.Fields(text)
  var count int64  

	for i:= 0; i < len(words); i++{
    
		select{
		case <- ctx.Done():
			return nil, errors.New("process canceled")

		default:
			count++
		}
	}

	return&domain.Result{
		TaskId: task.Id,
		Result: count,
		CompletedAt: time.Now(),
	},nil

}

func HandleReverseArrayInt(task *domain.Task, ctx context.Context) (*domain.Result,error){
 
	var result []int64

	if err:= json.Unmarshal(task.Payload,&result); err != nil{
		return  nil,err
	}

	l:= 0
	r:= len(result) - 1

	for l < r{

   select{
	 case <- ctx.Done():
		return nil, errors.New("process canceled")
	default:
		temp:= result[l]
		result[l] =  result[r]
		result[r] = temp
		l++
		r--
	}
	}
	

	return &domain.Result{
		TaskId: task.Id,
		Result: result,
		CompletedAt: time.Now(),
	},nil
}


