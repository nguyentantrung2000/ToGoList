package model

import (
	"time"
	"togolist/db"
)

type Task struct {
	ID          int        `json:"id" gorm:"column:id;primary_key"`
	Title       string     `json:"title" gorm:"column:title"`
	Description string     `json:"description" gorm:"column:description"`
	Image       string     `json:"image" gorm:"column:image"`
	Status      string     `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (Task) TableName() string {
	return "todo_items"
}

type TaskCreation struct {
	ID          int    `json:"id" gorm:"column:id;primary_key"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Image       string `json:"image" gorm:"column:image"`
	Status      string `json:"status" gorm:"column:status"`
}

func (TaskCreation) TableName() string {
	return Task{}.TableName()
}

func GetAllTasks() ([]Task, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()

	var tasks []Task
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func CreateTask(task *TaskCreation) error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	result := db.Create(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetTask(task *Task) error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	result := db.Where("id = ?").First(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
