package entities_test

import (
	"testing"

	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/helpers"
)

func TestSwapTasks(t *testing.T) {
	c, _ := entities.NewColumn("c0")
	t0, _ := entities.NewTask("t0", "")
	t1, _ := entities.NewTask("t1", "")
	c.AddTask(t0)
	c.AddTask(t1)
	if err := c.SwapTasks(0, 1); err != nil {
		t.Fatal(err)
	}
	if t0.Priority != 1 || t1.Priority != 0 {
		t.Errorf("tasks were not swapped properly:\n%v",
			helpers.PrintTasks(c.Tasks))
	}
}

func TestDeleteTask(t *testing.T) {
	c, _ := entities.NewColumn("c0")
	t0, _ := entities.NewTask("t0", "")
	c.AddTask(t0)
	t1, _ := entities.NewTask("t1", "")
	c.AddTask(t1)

	if err := c.DeleteTask(0); err != nil {
		t.Fatal(err)
	}
	if c.Tasks[0].Priority != 0 || c.Tasks[0].Name != "t1" {
		t.Errorf("task was not deleted properly")
	}
}
