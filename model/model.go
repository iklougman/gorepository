package model

import "time"

type TodoItem struct {
	Id          int       `bson: "_id,omitempty"`
	Task        string    `bson: "task,omitempty"`
	Completed   bool      `bson: "completed,omitempty"`
	CreateAt    time.Time `bson: "createdAt,omitempty"`
	CompletedAt time.Time
}

type TodoItemsList []TodoItem

type TodoRepository interface {
	FindAll() (*TodoItemsList, error)
	Insert(item *TodoItemsList) error
	// Update(Id int) error
}
