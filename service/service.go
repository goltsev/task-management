package service

import (
	"sort"

	"github.com/goltsev/task-management/entities"
)

type Service struct {
	db Database
}

func NewService(r Database) *Service {
	return &Service{
		db: r,
	}
}

func (s *Service) CreateProject(name, desc string) (int, error) {
	p, err := entities.NewProject(name, desc)
	if err != nil {
		return 0, err
	}
	return s.db.CreateProject(p)
}

func (s *Service) GetProject(id int) (*entities.Project, error) {
	return s.db.ReadProject(id)
}

func (s *Service) GetProjectList() ([]*entities.Project, error) {
	list, err := s.db.ListProjects()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list, nil
}

func (s *Service) UpdateProject(p *entities.Project) error {
	if err := p.Validate(); err != nil {
		return err
	}
	return s.db.UpdateProject(p)
}

func (s *Service) UpdateProjectInfo(id int, name string, desc string) error {
	p, err := s.db.ReadProject(id)
	if err != nil {
		return err
	}
	p.Name = name
	p.Description = desc
	if err := s.UpdateProject(p); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteProject(id int) error {
	return s.db.DeleteProject(id)
}

func (s *Service) CreateColumn(pid int, name string) (int, error) {
	p, err := s.GetProject(pid)
	if err != nil {
		return 0, err
	}
	col, err := entities.NewColumn(name)
	if err != nil {
		return 0, err
	}
	if err := p.AddColumn(col); err != nil {
		return 0, err
	}
	if err := s.UpdateProject(p); err != nil {
		return 0, err
	}
	return col.ID, nil
}

func (s *Service) GetColumn(cid int) (*entities.Column, error) {
	c, err := s.db.ReadColumn(cid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Service) GetColumnList(pid int) ([]*entities.Column, error) {
	p, err := s.db.ReadProject(pid)
	if err != nil {
		return nil, err
	}
	return p.SortColumnList(), nil
}

func (s *Service) UpdateColumn(c *entities.Column) error {
	if err := c.Validate(); err != nil {
		return err
	}
	return s.db.UpdateColumn(c)
}

func (s *Service) UpdateColumnInfo(id int, name string) error {
	c, err := s.GetColumn(id)
	if err != nil {
		return err
	}
	c.Name = name
	return s.UpdateColumn(c)
}

func (s *Service) MoveColumn(cid int, pos int) error {
	c, err := s.GetColumn(cid)
	if err != nil {
		return err
	}
	p, err := s.GetProject(c.ProjectID)
	if err != nil {
		return err
	}
	if err := p.MoveColumn(c.Position, pos); err != nil {
		return err
	}
	return s.UpdateProject(p)
}

func (s *Service) DeleteColumn(cid int) error {
	c, err := s.GetColumn(cid)
	if err != nil {
		return err
	}
	p, err := s.GetProject(c.ProjectID)
	if err != nil {
		return err
	}
	if err := p.DeleteColumn(c.Position); err != nil {
		return err
	}
	if err := s.db.DeleteColumn(cid); err != nil {
		return err
	}
	if err := s.UpdateProject(p); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetProjectTaskList(pid int) ([]*entities.Task, error) {
	p, err := s.GetProject(pid)
	if err != nil {
		return nil, err
	}
	tasks := make([]*entities.Task, 0, 1)
	clist := p.SortColumnList()
	for _, c := range clist {
		tasks = append(tasks, c.Tasks...)
	}
	return tasks, nil
}

func (s *Service) GetTaskList(cid int) ([]*entities.Task, error) {
	c, err := s.GetColumn(cid)
	if err != nil {
		return nil, err
	}
	return c.SortTaskList(), nil
}

func (s *Service) CreateTask(cid int, name string, desc string) (int, error) {
	t, err := entities.NewTask(name, desc)
	if err != nil {
		return 0, err
	}
	c, err := s.GetColumn(cid)
	if err != nil {
		return 0, err
	}
	c.AddTask(t)
	if err := s.db.UpdateColumn(c); err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (s *Service) GetTask(tid int) (*entities.Task, error) {
	t, err := s.db.ReadTask(tid)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *Service) UpdateTask(t *entities.Task) error {
	if err := t.Validate(); err != nil {
		return err
	}
	if err := s.db.UpdateTask(t); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateTaskInfo(tid int, name string, desc string) error {
	t, err := s.db.ReadTask(tid)
	if err != nil {
		return err
	}
	t.Name = name
	t.Description = desc
	if err := s.UpdateTask(t); err != nil {
		return err
	}
	return nil
}

// move either "up" or "down"
func (s *Service) ChangeTaskPriority(tid int, prio int) error {
	t, err := s.GetTask(tid)
	if err != nil {
		return err
	}
	c, err := s.GetColumn(t.ColumnID)
	if err != nil {
		return err
	}

	if err := c.MoveTask(t.Priority, prio); err != nil {
		return err
	}
	if err := s.db.UpdateColumn(c); err != nil {
		return err
	}
	return nil
}

// MoveTaskTo moves task to another
// move task either "left" or "right"
func (s *Service) MoveTaskTo(tid int, colpos int) error {
	t, err := s.db.ReadTask(tid)
	if err != nil {
		return err
	}
	c, err := s.db.ReadColumn(t.ColumnID)
	if err != nil {
		return err
	}
	p, err := s.db.ReadProject(c.ProjectID)
	if err != nil {
		return err
	}
	err = p.MoveTask(c.Position, colpos, t.Priority)
	if err != nil {
		return err
	}
	if err := s.db.UpdateProject(p); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteTask(tid int) error {
	if err := s.db.DeleteTask(tid); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCommentList(tid int) ([]*entities.Comment, error) {
	t, err := s.GetTask(tid)
	if err != nil {
		return nil, err
	}
	return t.SortCommentList(), nil
}

func (s *Service) CreateComment(tid int, text string) (int, error) {
	c, err := entities.NewComment(text)
	if err != nil {
		return 0, err
	}
	t, err := s.GetTask(tid)
	if err != nil {
		return 0, err
	}
	t.AddComment(c)
	if err := s.UpdateTask(t); err != nil {
		return 0, err
	}
	return c.ID, nil
}

func (s *Service) GetComment(cid int) (*entities.Comment, error) {
	c, err := s.db.ReadComment(cid)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Service) UpdateComment(c *entities.Comment) error {
	if err := c.Validate(); err != nil {
		return err
	}
	if err := s.db.UpdateComment(c); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateCommentInfo(cid int, text string) error {
	c, err := s.db.ReadComment(cid)
	if err != nil {
		return err
	}
	c.Text = text
	if err := s.UpdateComment(c); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteComment(cid int) error {
	if err := s.db.DeleteComment(cid); err != nil {
		return err
	}
	return nil
}
