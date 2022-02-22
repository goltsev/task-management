package entities_test

import (
	"testing"

	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/helpers"
)

var (
	nameLenOK  = 10
	nameLenBAD = 501
	descLenOK  = 200
	descLenBAD = 1001

	gudname string
	badname string

	guddesc string
	baddesc string
)

func init() {
	var err error
	gudname, err = helpers.GenerateString(nameLenOK)
	if err != nil {
		panic(err)
	}
	badname, _ = helpers.GenerateString(nameLenBAD)
	if err != nil {
		panic(err)
	}

	guddesc, _ = helpers.GenerateString(descLenOK)
	if err != nil {
		panic(err)
	}
	baddesc, _ = helpers.GenerateString(descLenBAD)
	if err != nil {
		panic(err)
	}
}

func TestProjectValidate(t *testing.T) {
	p := &entities.Project{
		Name:        gudname,
		Description: guddesc,
		Columns:     []*entities.Column{{}},
	}
	if err := p.Validate(); err != nil {
		t.Errorf("no error expected, got: %v", err)
	}
}

func TestProjectValidateErrors(t *testing.T) {
	var p *entities.Project
	p = &entities.Project{
		Name:        gudname,
		Description: baddesc,
		Columns:     []*entities.Column{{}},
	}
	if err := p.Validate(); err == nil {
		t.Errorf("expected error, got nothing")
	}
	p.Name = badname
	p.Description = guddesc
	if err := p.Validate(); err == nil {
		t.Errorf("expected error, got nothing")
	}
	p.Name = gudname
	p.Columns = nil
	if err := p.Validate(); err == nil {
		t.Errorf("expected error, got nothing")
	}
}

func TestSwapColumns(t *testing.T) {
	pname := "name"
	p, _ := entities.NewProject(pname, "")
	cname := "col1"
	c, _ := entities.NewColumn(cname)
	p.AddColumn(c)
	if err := p.SwapColumns(0, 1); err != nil {
		t.Fatal(err)
	}
	if c.Position != 0 || p.Columns[0].Position != 1 {
		t.Errorf("columns were not swapped properly. got:\n%v",
			helpers.PrintColumnList(p.Columns))
	}
}

func TestMoveColumn(t *testing.T) {
	p, _ := entities.NewProject("proj", "")
	c0, _ := p.GetColumnPosition(0)
	c1, _ := entities.NewColumn("col1")
	c2, _ := entities.NewColumn("col2")
	c3, _ := entities.NewColumn("col3")
	c4, _ := entities.NewColumn("col4")
	p.AddColumn(c1)
	p.AddColumn(c2)
	p.AddColumn(c3)
	p.AddColumn(c4)
	op := 4
	np := 0
	if err := p.MoveColumn(op, np); err != nil {
		t.Fatal(err)
	}
	if c0.Position != 1 || c4.Position != 0 {
		t.Errorf("columns are not in specified positions:\n%v",
			helpers.PrintProject(p))
	}
}

func TestProjectDeleteColumn(t *testing.T) {
	p, err := entities.NewProject("project", "")
	if err != nil {
		t.Fatal(err)
	}

	c0, _ := p.GetColumnPosition(0)
	c1, _ := entities.NewColumn("col1")
	c2, _ := entities.NewColumn("col2")
	t0, _ := entities.NewTask("task0", "")
	t1, _ := entities.NewTask("task1", "")
	c1.AddTask(t0)
	c1.AddTask(t1)

	p.AddColumn(c1)
	p.AddColumn(c2)
	if err := p.DeleteColumn(c1.Position); err != nil {
		t.Fatal(err)
	}
	if c0.Tasks == nil {
		t.Errorf("column 0 has no tasks from deleted column 1")
	}
}

func TestMoveTask(t *testing.T) {
	p, _ := entities.NewProject("proj", "")
	c0, _ := p.GetColumnPosition(0)
	c1, _ := entities.NewColumn("stage0")
	t0, _ := entities.NewTask("task0", "")
	t1, _ := entities.NewTask("task1", "")

	p.AddColumn(c1)
	c1.AddTask(t0)
	c1.AddTask(t1)

	if err := p.MoveTask(1, 0, t0.Priority); err != nil {
		t.Fatal(err)
	}
	if c0.Tasks[0] != t0 || c1.Tasks[0] != t1 {
		t.Errorf("task was not moved properly:\n%v", helpers.PrintProject(p))
	}
	if err := p.MoveTask(1, 0, t0.Priority); err != nil {
		t.Fatal(err)
	}
	if err := p.MoveTask(1, 0, t0.Priority); err == nil {
		t.Errorf("error expected, got nothing")
	}
}
