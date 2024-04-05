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

type TaskUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Image       *string `json:"image" gorm:"column:image"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TaskUpdate) TableName() string {
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

	result := db.First(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateTask(id int, task *TaskUpdate) error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	var existingTask Task
	if err := db.First(&existingTask, id).Error; err != nil {
		return err
	}

	result := db.Model(&existingTask).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteTask deletes a task by setting its status to "Deleted"
func DeleteTask(id int) (Task, error) {
	db, err := db.NewDB()
	if err != nil {
		return Task{}, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return Task{}, err
	}
	defer sqlDB.Close()
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	result := db.Where("id = ?", task.ID).Updates(Task{Status: "Deleted"})
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

// DeleteTask deletes a task by deleting it from the database

func DeleteTaskDatabase(id int) error {
	db, err := db.NewDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()
	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}
	result := db.Where("id = ?", task.ID).Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
