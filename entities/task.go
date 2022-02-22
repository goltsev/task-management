package entities

import (
	"fmt"
	"sort"
)

const (
	MinTaskName = 1
	MaxTaskName = 500

	MinTaskDesc = 0
	MaxTaskDesc = 5000
)

type Task struct {
	Comments    []*Comment `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	ID          int
	ColumnID    int
	Priority    int
	Name        string
	Description string
}

func NewTask(name, desc string) (*Task, error) {
	t := &Task{
		Name:        name,
		Description: desc,
	}
	if err := t.Validate(); err != nil {
		return nil, err
	}
	return t, nil
}

// Validate checks Taskect Name and Description length
func (t *Task) Validate() error {
	if len(t.Name) > MaxTaskName || len(t.Name) < MinTaskName {
		return fmt.Errorf("name has to be (%d-%d characters)",
			MinTaskName, MaxTaskName)
	}
	if len(t.Description) > MaxTaskDesc || len(t.Description) < MinTaskDesc {
		return fmt.Errorf("description has to be (%d-%d characters)",
			MinTaskDesc, MaxTaskDesc)
	}
	return nil
}

func (t *Task) AddComment(c *Comment) {
	t.Comments = append(t.Comments, c)
}

func (t *Task) SortCommentList() []*Comment {
	sort.Slice(t.Comments, func(i, j int) bool {
		return t.Comments[i].Date.After(t.Comments[j].Date)
	})
	return t.Comments
}
