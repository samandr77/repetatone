package main

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Task   string `json:"task"`       // Наш сервер будет ожидать json c полем text
	IsDone bool   `json:"is_done"`    // В GO используем CamelCase, в Json - snake
	ID     uint   `gorm:"primarykey"` //индефикатор задачи
}
