package workerpool

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/anthonymartz17/distributed-task-runner/internal/domain"
)

func HandleWordCount(task *domain.Task) *domain.Result{
  var text string

	if err:= json.Unmarshal(task.Payload,&text); err != nil{

		return&domain.Result{
			TaskId: task.Id,
			Error: err.Error(),
			CompletedAt: time.Now(),
		}
	}

	count:= len(strings.Fields(text))

	return&domain.Result{
		TaskId: task.Id,
		Result: count,
		CompletedAt: time.Now(),
	}

}

func HandleReverseArrayInt(task *domain.Task) *domain.Result{
 
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
   temp:= result[l]
	 result[l] =  result[r]
	 result[r] = temp
	 l++
	 r--
	}

	return&domain.Result{
		TaskId: task.Id,
		Result: result,
		CompletedAt: time.Now(),
	}
}

