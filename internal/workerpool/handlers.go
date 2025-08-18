package workerpool

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)

func HandleWordCount(task *domain.Task,ctx context.Context) *domain.Result{
  var text string
	
	if err:= json.Unmarshal(task.Payload,&text); err != nil{

		return&domain.Result{
			TaskId: task.Id,
			Error: err.Error(),
			CompletedAt: time.Now(),
		}
	}
  words:=  strings.Fields(text)
  var count int64  

	for i:= 0; i < len(words); i++{
    
		select{
		case <- ctx.Done():
			return&domain.Result{
				TaskId: task.Id,
				Error: ctx.Err().Error(),
				CompletedAt: time.Now(),
			}

		default:
			count++
		}
	}

	return&domain.Result{
		TaskId: task.Id,
		Result: count,
		CompletedAt: time.Now(),
	}

}

func HandleReverseArrayInt(task *domain.Task, ctx context.Context) *domain.Result{
 
	var result []int64

	if err:= json.Unmarshal(task.Payload,&result); err != nil{
		return&domain.Result{
			TaskId: task.Id,
			Error: err.Error(),
			CompletedAt: time.Now(),
		}
	}

	l:= 0
	r:= len(result) - 1

	for l < r{

   select{
	 case <- ctx.Done():
		return&domain.Result{
			TaskId: task.Id,
			Error: ctx.Err().Error(),
			CompletedAt: time.Now(),
		}
	default:
		temp:= result[l]
		result[l] =  result[r]
		result[r] = temp
		l++
		r--
	}
	}

	return&domain.Result{
		TaskId: task.Id,
		Result: result,
		CompletedAt: time.Now(),
	}
}

