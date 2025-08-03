package domain

import "time"


type Task struct{
	Id string `json:"id"`
	Type string `json:"type"`
	Payload interface{} `json:"payload"`
	CreatedAt time.Time `json:"create_at"`
}

