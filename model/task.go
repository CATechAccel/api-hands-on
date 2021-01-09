package model

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	_           time.Time `json:"created_at"`
	_           time.Time `json:"updated_at"`
}

type Tasks []Task
