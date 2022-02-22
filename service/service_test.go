package service_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/helpers"

	"github.com/goltsev/task-management/service"
	"github.com/goltsev/task-management/service/mock_service"
)

func TestMoveColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock_service.NewMockDatabase(ctrl)

	p, _ := entities.NewProject("project", "")
	c0, _ := p.GetColumnPosition(0)
	c1, _ := entities.NewColumn("col0")
	c2, _ := entities.NewColumn("col1")
	p.AddColumn(c1)
	p.AddColumn(c2)

	id := 2
	newpos := 0

	db.
		EXPECT().
		ReadProject(0).
		Return(p, nil).AnyTimes()
	db.
		EXPECT().
		UpdateProject(p).
		Return(nil).AnyTimes()

	db.
		EXPECT().
		ReadColumn(id).
		Return(c2, nil).AnyTimes()
	db.
		EXPECT().
		UpdateColumn(c2).
		Return(nil).AnyTimes()

	s := service.NewService(db)
	if err := s.MoveColumn(id, newpos); err != nil {
		t.Error(err)
	}
	if c0.Position != 1 || c1.Position != 2 || c2.Position != 0 {
		t.Errorf("column was not moved properly: %v",
			helpers.PrintProject(p))
	}
}

func TestGetProjectTaskList(t *testing.T) {
	t0, _ := entities.NewTask("t0", "")
	t1, _ := entities.NewTask("t1", "")
	c0, _ := entities.NewColumn("c0")
	c1, _ := entities.NewColumn("c1")
	p, _ := entities.NewProject("p", "")
	p.AddColumn(c0)
	p.AddColumn(c1)
	c0.AddTask(t0)
	c1.AddTask(t1)

	ctrl := gomock.NewController(t)
	db := mock_service.NewMockDatabase(ctrl)
	db.
		EXPECT().
		ReadProject(0).
		Return(p, nil).AnyTimes()

	s := service.NewService(db)
	list, err := s.GetProjectTaskList(0)
	if err != nil {
		t.Fatal(err)
	}
	expected := []*entities.Task{t0, t1}
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("\nexpected:\n%v\ngot:\n%v\n",
			helpers.PrintTasks(expected),
			helpers.PrintTasks(list))
	}
}
