package handlers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	_ "github.com/goltsev/task-management/docs"
)

type Mux struct {
	mux     http.Handler
	service Service
}

func NewMux(s Service) *Mux {
	mux := httprouter.New()
	r := &Mux{
		mux:     mux,
		service: s,
	}
	mux.HandlerFunc("GET", "/api/v1/projects", r.GetProjectList)
	mux.HandlerFunc("POST", "/api/v1/projects", r.CreateProject)

	mux.HandlerFunc("GET", "/api/v1/projects/:projectid", r.GetProject)
	mux.HandlerFunc("POST", "/api/v1/projects/:projectid", r.UpdateProject)
	mux.HandlerFunc("DELETE", "/api/v1/projects/:projectid", r.DeleteProject)
	mux.HandlerFunc("GET", "/api/v1/projects/:projectid/columns", r.GetColumnList)
	mux.HandlerFunc("POST", "/api/v1/projects/:projectid/columns", r.CreateColumn)
	mux.HandlerFunc("GET", "/api/v1/projects/:projectid/tasks", r.GetProjectTaskList)

	mux.HandlerFunc("GET", "/api/v1/columns/:columnid", r.GetColumn)
	mux.HandlerFunc("POST", "/api/v1/columns/:columnid", r.UpdateColumn)
	mux.HandlerFunc("DELETE", "/api/v1/columns/:columnid", r.DeleteColumn)
	mux.HandlerFunc("GET", "/api/v1/columns/:columnid/tasks", r.GetTaskList)
	mux.HandlerFunc("POST", "/api/v1/columns/:columnid/tasks", r.CreateTask)

	mux.HandlerFunc("GET", "/api/v1/tasks/:taskid", r.GetTask)
	mux.HandlerFunc("POST", "/api/v1/tasks/:taskid", r.UpdateTask)
	mux.HandlerFunc("DELETE", "/api/v1/tasks/:taskid", r.DeleteTask)
	mux.HandlerFunc("GET", "/api/v1/tasks/:taskid/comments", r.GetCommentList)
	mux.HandlerFunc("POST", "/api/v1/tasks/:taskid/comments", r.CreateComment)

	mux.HandlerFunc("GET", "/api/v1/comments/:commentid", r.GetComment)
	mux.HandlerFunc("POST", "/api/v1/comments/:commentid", r.UpdateComment)
	mux.HandlerFunc("DELETE", "/api/v1/comments/:commentid", r.DeleteComment)

	return r
}

// GetProjectList godoc
// @Summary Get Project list
// @Tags projects
// @Produce json
// @Success 200
// @Router /projects [get]
func (m *Mux) GetProjectList(w http.ResponseWriter, r *http.Request) {
	list, err := m.service.GetProjectList()
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, list)
}

// CreateProject godoc
// @Summary Create Project
// @Tags projects
// @Accept json
// @Produce json
// @Param info body Project true "Project parameters"
// @Success 200
// @Router /projects [post]
func (m *Mux) CreateProject(w http.ResponseWriter, r *http.Request) {
	proj := &Project{}
	if err := decodeJSON(w, r, proj); err != nil {
		handleError(w, r, err)
		return
	}
	id, err := m.service.CreateProject(proj.Name, proj.Description)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, map[string]int{"ID": id})
}

// GetProject godoc
// @Summary Get Project with ID
// @Tags projects
// @Produce json
// @Param projectid path int true "Project ID"
// @Success 200
// @Router /projects/{projectid} [get]
func (m *Mux) GetProject(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	projectid := params.ByName("projectid")
	id, err := strconv.Atoi(projectid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	project, err := m.service.GetProject(id)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, project)
}

// UpdateProject godoc
// @Summary Update Project
// @Tags projects
// @Produce json
// @Param projectid path int true "Project ID"
// @Param info body Project true "Project information"
// @Success 200
// @Router /projects/{projectid} [post]
func (m *Mux) UpdateProject(w http.ResponseWriter, r *http.Request) {
	pid, err := getHTTPRouterParamID(r, "projectid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	proj := &Project{}
	if err := decodeJSON(w, r, proj); err != nil {
		handleError(w, r, err)
		return
	}
	err = m.service.UpdateProjectInfo(pid, proj.Name, proj.Description)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

// DeleteProject godoc
// @Summary Delete Project
// @Tags projects
// @Produce json
// @Param projectid path int true "Project ID"
// @Success 200
// @Router /projects/{projectid} [delete]
func (m *Mux) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	projectid := params.ByName("projectid")
	id, err := strconv.Atoi(projectid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	err = m.service.DeleteProject(id)
	if err != nil {
		handleError(w, r, err)
		return
	}
}

// GetColumnList godoc
// @Summary Get list of columns in project with ID
// @Tags columns
// @Produce json
// @Param projectid path int true "Project ID"
// @Success 200
// @Router /projects/{projectid}/columns [get]
func (m *Mux) GetColumnList(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	projectid := params.ByName("projectid")
	id, err := strconv.Atoi(projectid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	list, err := m.service.GetColumnList(id)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, list)
}

// CreateColumn godoc
// @Summary Create column in project
// @Tags columns
// @Produce json
// @Param projectid path int true "Project ID"
// @Param column body Column true "Column info"
// @Success 200
// @Router /projects/{projectid}/columns [post]
func (m *Mux) CreateColumn(w http.ResponseWriter, r *http.Request) {
	pid, err := getHTTPRouterParamID(r, "projectid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	col := &Column{}
	if err := decodeJSON(w, r, col); err != nil {
		handleError(w, r, err)
		return
	}
	cid, err := m.service.CreateColumn(pid, col.Name)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, map[string]int{"ID": cid})
}

// GetColumn godoc
// @Summary Get column with ID
// @Tags columns
// @Produce json
// @Param columnid path int true "Column ID"
// @Success 200
// @Router /columns/{columnid} [get]
func (m *Mux) GetColumn(w http.ResponseWriter, r *http.Request) {
	id, err := getHTTPRouterParamID(r, "columnid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	c, err := m.service.GetColumn(id)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, c)
}

// UpdateColumn godoc
// @Summary Update column with ID
// @Tags columns
// @Produce json
// @Param columnid path int true "Column ID"
// @Param column body Column true "Column info"
// @Success 200
// @Router /columns/{columnid} [post]
func (m *Mux) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	id, err := getHTTPRouterParamID(r, "columnid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	col := &Column{}
	if err := decodeJSON(w, r, col); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.UpdateColumnInfo(id, col.Name); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.MoveColumn(id, col.Position); err != nil {
		handleError(w, r, err)
		return
	}
}

// DeleteColumn godoc
// @Summary Delete column with ID
// @Tags columns
// @Produce json
// @Param columnid path int true "Column ID"
// @Success 200
// @Router /columns/{columnid} [delete]
func (m *Mux) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "columnid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.DeleteColumn(cid); err != nil {
		handleError(w, r, err)
		return
	}
}

// GetTaskList godoc
// @Summary Get list of tasks from column with ID
// @Tags tasks
// @Produce json
// @Param columnid path int true "Column ID"
// @Success 200
// @Router /columns/{columnid}/tasks [get]
func (m *Mux) GetTaskList(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "columnid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	list, err := m.service.GetTaskList(cid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, list)
}

// GetProjectTaskList godoc
// @Summary Get list of tasks for project with ID
// @Tags tasks
// @Produce json
// @Param projectid path int true "Project ID"
// @Success 200
// @Router /projects/{projectid}/tasks [get]
func (m *Mux) GetProjectTaskList(w http.ResponseWriter, r *http.Request) {
	pid, err := getHTTPRouterParamID(r, "projectid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	list, err := m.service.GetProjectTaskList(pid)
	if err != nil {
		handleError(w, r, err)
	}
	encodeJSON(w, r, list)
}

// CreateTask godoc
// @Summary Create task in column with ID
// @Tags tasks
// @Produce json
// @Param columnid path int true "Column ID"
// @Param task body Task true "Task info"
// @Success 200
// @Router /columns/{columnid}/tasks [post]
func (m *Mux) CreateTask(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "columnid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	task := &Task{}
	if err := decodeJSON(w, r, task); err != nil {
		handleError(w, r, err)
		return
	}
	tid, err := m.service.CreateTask(cid, task.Name, task.Description)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, map[string]int{"ID": tid})
}

// GetTask godoc
// @Summary Get task with ID
// @Tags tasks
// @Produce json
// @Param taskid path int true "Task ID"
// @Success 200
// @Router /tasks/{taskid} [get]
func (m *Mux) GetTask(w http.ResponseWriter, r *http.Request) {
	tid, err := getHTTPRouterParamID(r, "taskid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	t, err := m.service.GetTask(tid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, t)
}

// UpdateTask godoc
// @Summary Update task with ID
// @Tags tasks
// @Produce json
// @Param taskid path int true "Task ID"
// @Param task body Task true "Task info"
// @Success 200
// @Router /tasks/{taskid} [post]
func (m *Mux) UpdateTask(w http.ResponseWriter, r *http.Request) {
	tid, err := getHTTPRouterParamID(r, "taskid")
	if err != nil {
		handleError(w, r, err)
	}
	task := &Task{}
	if err := decodeJSON(w, r, task); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.UpdateTaskInfo(tid, task.Name, task.Description); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.MoveTaskTo(tid, task.Column); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.ChangeTaskPriority(tid, task.Priority); err != nil {
		handleError(w, r, err)
	}
}

// DeleteTask godoc
// @Summary Delete task with ID
// @Tags tasks
// @Produce json
// @Param taskid path int true "Task ID"
// @Success 200
// @Router /tasks/{taskid} [delete]
func (m *Mux) DeleteTask(w http.ResponseWriter, r *http.Request) {
	tid, err := getHTTPRouterParamID(r, "taskid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.DeleteTask(tid); err != nil {
		handleError(w, r, err)
		return
	}
}

// GetCommentList godoc
// @Summary Get list of comments from task with ID
// @Tags comments
// @Produce json
// @Param taskid path int true "Task ID"
// @Success 200
// @Router /tasks/{taskid}/comments [get]
func (m *Mux) GetCommentList(w http.ResponseWriter, r *http.Request) {
	tid, err := getHTTPRouterParamID(r, "taskid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	list, err := m.service.GetCommentList(tid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, list)
}

// CreateComment godoc
// @Summary Create comment for task with ID
// @Tags comments
// @Produce json
// @Param taskid path int true "Task ID"
// @Param comment body Comment true "Comment text"
// @Success 200
// @Router /tasks/{taskid}/comments [post]
func (m *Mux) CreateComment(w http.ResponseWriter, r *http.Request) {
	tid, err := getHTTPRouterParamID(r, "taskid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	c := &Comment{}
	if err := decodeJSON(w, r, c); err != nil {
		handleError(w, r, err)
		return
	}
	cid, err := m.service.CreateComment(tid, c.Text)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, map[string]int{"ID": cid})
}

// GetComment godoc
// @Summary Get comment with ID
// @Tags comments
// @Produce json
// @Param commentid path int true "Comment ID"
// @Success 200
// @Router /comments/{commentid} [get]
func (m *Mux) GetComment(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "commentid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	c, err := m.service.GetComment(cid)
	if err != nil {
		handleError(w, r, err)
		return
	}
	encodeJSON(w, r, c)
}

// UpdateComment godoc
// @Summary Update comment with ID
// @Tags comments
// @Produce json
// @Param commentid path int true "Comment ID"
// @Param comment body Comment true "Comment text"
// @Success 200
// @Router /comments/{commentid} [post]
func (m *Mux) UpdateComment(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "commentid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	c := &Comment{}
	if err := decodeJSON(w, r, c); err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.UpdateCommentInfo(cid, c.Text); err != nil {
		handleError(w, r, err)
		return
	}
}

// DeleteComment godoc
// @Summary Delete comment with ID
// @Tags comments
// @Produce json
// @Param commentid path int true "Comment ID"
// @Success 200
// @Router /comments/{commentid} [delete]
func (m *Mux) DeleteComment(w http.ResponseWriter, r *http.Request) {
	cid, err := getHTTPRouterParamID(r, "commentid")
	if err != nil {
		handleError(w, r, err)
		return
	}
	if err := m.service.DeleteComment(cid); err != nil {
		handleError(w, r, err)
		return
	}
}

// ServeHTTP allows Mux to implement a http.Handler interface
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}
