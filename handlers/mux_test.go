package handlers_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/handlers"
	"github.com/goltsev/task-management/handlers/mock_handlers"
)

func TestGetProjectList(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := mock_handlers.NewMockService(ctrl)

	p0, _ := entities.NewProject("p0", "")
	p1, _ := entities.NewProject("p1", "")
	p2, _ := entities.NewProject("p2", "")
	list := []*entities.Project{p0, p1, p2}
	s.
		EXPECT().
		GetProjectList().
		Return(list, nil)

	mux := handlers.NewMux(s)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	mux.GetProjectList(w, r)

	got, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected, _ := json.Marshal(&list)
	expected = append(expected, '\n')
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected:\n%v\ngot:\n%v\n",
			string(expected),
			string(got))
	}
}
