package model

import "time"

type TodoItem struct {
	Id          int    `json: "id"`
	Task        string `json: "task"`
	Completed   bool   `json: "completed"`
	CreateAt    time.Time
	CompletedAt time.Time
}

type TodoItemsList []TodoItem

type TodoRepository interface {
	FindAll() (*TodoItemsList, error)
	// Save(item *TodoItem) error
	// Update(Id int) error
}
