package domain

import "time"


type Result struct{
	TaskId string `json:"task_id"`
	Result interface{} `json:"result"`
	Error string `json:"error,omitempty"`
	CompletedAt time.Time `json:"completed_at"`
}