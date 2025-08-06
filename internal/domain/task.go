package domain

import (
	"encoding/json"


	// "github.com/google/uuid"
)


type Task struct{
	Id string `json:"id"`
	Type string `json:"type"`
	Payload json.RawMessage`json:"payload"`
	CreatedAt int64 `json:"created_at"`
}

// func newTask() *Task{
// 	return &Task{
// 		Id: uuid.NewString(),
// 	}
// }
