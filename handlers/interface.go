package handlers

import "github.com/goltsev/task-management/entities"

type Service interface {
	CreateProject(name, desc string) (int, error)
	GetProject(id int) (*entities.Project, error)
	GetProjectList() ([]*entities.Project, error)
	UpdateProject(p *entities.Project) error
	UpdateProjectInfo(id int, name string, desc string) error
	DeleteProject(id int) error

	CreateColumn(pid int, name string) (int, error)
	GetColumn(cid int) (*entities.Column, error)
	GetColumnList(pid int) ([]*entities.Column, error)
	UpdateColumn(c *entities.Column) error
	UpdateColumnInfo(id int, name string) error
	MoveColumn(cid int, pos int) error
	DeleteColumn(cid int) error

	GetProjectTaskList(pid int) ([]*entities.Task, error)
	GetTaskList(cid int) ([]*entities.Task, error)
	CreateTask(cid int, name string, desc string) (int, error)
	GetTask(tid int) (*entities.Task, error)
	UpdateTask(t *entities.Task) error
	UpdateTaskInfo(tid int, name string, desc string) error
	ChangeTaskPriority(tid int, prio int) error
	MoveTaskTo(tid int, colpos int) error
	DeleteTask(tid int) error

	GetCommentList(tid int) ([]*entities.Comment, error)
	CreateComment(tid int, text string) (int, error)
	GetComment(cid int) (*entities.Comment, error)
	UpdateComment(c *entities.Comment) error
	UpdateCommentInfo(cid int, text string) error
	DeleteComment(cid int) error
}
