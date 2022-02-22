package entities

import (
	"fmt"
	"sort"
)

const (
	MinProjName = 1
	MaxProjName = 500

	MinProjDesc = 0
	MaxProjDesc = 1000

	DefaultColName = "default"
)

type Project struct {
	Columns     []*Column `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"-"`
	ID          int
	Name        string
	Description string
}

// CreateProject returns Project struct and validates name and description
func NewProject(name, description string) (*Project, error) {
	p := &Project{
		Name:        name,
		Description: description,
	}
	col, err := NewColumn(DefaultColName)
	if err != nil {
		return nil, err
	}
	if err := p.AddColumn(col); err != nil {
		return nil, err
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

// Validate checks Project Name and Description length
func (p *Project) Validate() error {
	if len(p.Name) > MaxProjName || len(p.Name) < MinProjName {
		return fmt.Errorf("name has to be (%d-%d characters)",
			MinProjName, MaxProjName)
	}
	if len(p.Description) > MaxProjDesc || len(p.Description) < MinProjDesc {
		return fmt.Errorf("description has to be (%d-%d characters)",
			MinProjDesc, MaxProjDesc)
	}
	if len(p.Columns) < 1 {
		return fmt.Errorf("project has to contain at least one column")
	}
	return nil
}

// AddColumn adds a Column and returns error if it's Name is not unique
func (p *Project) AddColumn(col *Column) error {
	for _, v := range p.Columns {
		if v.Name == col.Name {
			return fmt.Errorf("column name has to be unique")
		}
	}
	p.Columns = append(p.Columns, col)
	col.Position = len(p.Columns) - 1
	col.ProjectID = p.ID
	return nil
}

// GetColumnsList returns list of columns sorted by position
func (p *Project) SortColumnList() []*Column {
	sort.Slice(p.Columns, func(i, j int) bool {
		return p.Columns[i].Position < p.Columns[j].Position
	})
	return p.Columns
}

// SwapPositions swaps positions of two columns
func (p *Project) SwapColumns(op int, np int) error {
	oc, err := p.GetColumnPosition(op)
	if err != nil {
		return err
	}
	nc, err := p.GetColumnPosition(np)
	if err != nil {
		return err
	}
	oc.Position, nc.Position = nc.Position, oc.Position
	return nil
}

func (p *Project) MoveColumn(op int, np int) error {
	d := 1
	if op > np {
		d = -1
	}
	for i := op; i != np; i += d {
		if err := p.SwapColumns(i, i+d); err != nil {
			return err
		}
	}
	return nil
}

// GetColumnPosition returns column with specified ID
func (p *Project) GetColumnID(id int) *Column {
	cols := p.Columns
	for _, c := range cols {
		if c.ID == id {
			return c
		}
	}
	return cols[len(cols)-1]
}

// GetColumnPosition returns column with specified position
func (p *Project) GetColumnPosition(pos int) (*Column, error) {
	if pos < 0 || pos > len(p.Columns) {
		return nil, fmt.Errorf("incorrect position")
	}
	for i := 0; i < len(p.Columns); i++ {
		if p.Columns[i].Position == pos {
			return p.Columns[i], nil
		}
	}
	return nil, fmt.Errorf(
		"column position integrity is broken. column in position [%v] not found",
		pos)
}

func (p *Project) CanDeleteColumn() bool {
	if len(p.Columns) <= 1 {
		return false
	}
	return true
}

func (p *Project) DeleteColumn(pos int) error {
	if !p.CanDeleteColumn() {
		return fmt.Errorf("last column cannot be deleted")
	}
	list := p.SortColumnList()
	if pos > len(list)-1 || pos < 0 {
		return fmt.Errorf("incorrect position")
	}
	col := list[pos]
	list = append(list[:pos], list[pos+1:]...)
	for i := 0; i < len(list); i++ {
		list[i].Position = i
	}
	if pos == 0 {
		list[pos].Tasks = col.Tasks
	} else {
		list[pos-1].Tasks = col.Tasks
	}
	p.Columns = list
	return nil
}

// ChangeTaskColumn changes task with priority tprio "status"
// from column in oldpos to column in newpos
func (p *Project) MoveTask(oldpos int, newpos int, tprio int) error {
	oldcol, err := p.GetColumnPosition(oldpos)
	if err != nil {
		return err
	}
	newcol, err := p.GetColumnPosition(newpos)
	if err != nil {
		return err
	}
	t, err := oldcol.GetTask(tprio)
	if err != nil {
		return err
	}
	if err := oldcol.DeleteTask(tprio); err != nil {
		return err
	}
	newcol.AddTask(t)
	return nil
}
