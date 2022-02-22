package entities_test

import (
	"testing"

	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/helpers"
)

func TestSortCommentList(t *testing.T) {
	text0 := "comment0"
	text1 := "comment1"
	t0, _ := entities.NewTask("t0", "")
	c0, _ := entities.NewComment(text0)
	c1, _ := entities.NewComment(text1)
	c1.Date = c1.Date.AddDate(0, 0, -1)
	t0.AddComment(c0)
	t0.AddComment(c1)
	cmts := t0.SortCommentList()
	if cmts[1].Text != text1 || cmts[0].Text != text0 {
		t.Errorf("comments were not sorted by date: %v",
			helpers.PrintComments(cmts))
	}
}
