package entities

import (
	"fmt"
	"sort"
)

const (
	MinColName = 1
	MaxColName = 255
)

type Column struct {
	Tasks     []*Task `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	ID        int
	ProjectID int
	Position  int
	Name      string
}

// NewColumn returns Column with specified Name
// and returns error if Name is not validated
func NewColumn(name string) (*Column, error) {
	c := &Column{
		Name: name,
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// Validate checks Column Name length
func (c *Column) Validate() error {
	if len(c.Name) > MaxColName || len(c.Name) < MinColName {
		return fmt.Errorf("name has to be (%d-%d characters)",
			MinColName, MaxColName)
	}
	return nil
}

func (c *Column) AddTask(t *Task) {
	c.Tasks = append(c.Tasks, t)
	t.Priority = len(c.Tasks) - 1
	t.ColumnID = c.ID
}

func (c *Column) GetTask(tprio int) (*Task, error) {
	if tprio < 0 || tprio > len(c.Tasks)-1 {
		return nil, fmt.Errorf("wrong task priority")
	}
	return c.Tasks[tprio], nil
}

// SortColumnList sorts taks by priority and returns them
func (c *Column) SortTaskList() []*Task {
	sort.Slice(c.Tasks, func(i, j int) bool {
		return c.Tasks[i].Priority < c.Tasks[j].Priority
	})
	return c.Tasks
}

func (c *Column) SwapTasks(op int, np int) error {
	oc, err := c.GetTask(op)
	if err != nil {
		return err
	}
	nc, err := c.GetTask(np)
	if err != nil {
		return err
	}
	oc.Priority, nc.Priority = nc.Priority, oc.Priority
	return nil
}

func (c *Column) MoveTask(op int, np int) error {
	d := 1
	if op > np {
		d = -1
	}
	for i := op; i != np; i += d {
		if err := c.SwapTasks(i, i+d); err != nil {
			return err
		}
	}
	return nil
}

func (c *Column) DeleteTask(prio int) error {
	if prio > len(c.Tasks)-1 || prio < 0 {
		return fmt.Errorf("incorrect position")
	}
	tasks := c.SortTaskList()
	tasks = append(tasks[:prio], tasks[prio+1:]...)
	for i := 0; i < len(tasks); i++ {
		tasks[i].Priority = i
	}
	c.Tasks = tasks
	return nil
}
