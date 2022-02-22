package service

import "github.com/goltsev/task-management/entities"

type ProjectRepository interface {
	CreateProject(*entities.Project) (int, error)
	ReadProject(int) (*entities.Project, error)
	UpdateProject(*entities.Project) error
	DeleteProject(int) error
	ListProjects() ([]*entities.Project, error)
}

type ColumnRepository interface {
	CreateColumn(*entities.Column) (int, error)
	ReadColumn(int) (*entities.Column, error)
	UpdateColumn(*entities.Column) error
	DeleteColumn(int) error
}

type TaskRepository interface {
	CreateTask(*entities.Task) (int, error)
	ReadTask(int) (*entities.Task, error)
	UpdateTask(*entities.Task) error
	DeleteTask(int) error
}

type CommentRepository interface {
	CreateComment(*entities.Comment) (int, error)
	ReadComment(int) (*entities.Comment, error)
	UpdateComment(*entities.Comment) error
	DeleteComment(int) error
}

type Database interface {
	ProjectRepository
	ColumnRepository
	TaskRepository
	CommentRepository
}
