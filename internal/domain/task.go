package domain

import "time"



type Task struct{
	Id string `json:"id"`
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
	CreatedAt time.Time `json:"create_at"`
}

type Result struct{
	TaskId string `json:"task_id"`
	Result interface{} `json:"result"`
	Error string `json:"error,omitempty"`
	CompletedAt time.Time `json:"completed_at"`
}