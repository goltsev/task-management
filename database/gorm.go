package database

import (
	"github.com/goltsev/task-management/entities"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func CreateGormDB(db *gorm.DB) *GormDB {
	return &GormDB{
		db: db,
	}
}

func (db *GormDB) CreateProject(p *entities.Project) (int, error) {
	if err := db.db.Create(p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (db *GormDB) ReadProject(id int) (*entities.Project, error) {
	p := &entities.Project{
		ID: id,
	}
	if err := db.db.Preload("Columns.Tasks.Comments").First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (db *GormDB) UpdateProject(p *entities.Project) error {
	if err := db.db.Save(p).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) DeleteProject(id int) error {
	if err := db.db.Delete(&entities.Project{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) ListProjects() ([]*entities.Project, error) {
	var projects []*entities.Project
	if err := db.db.Preload("Columns.Tasks.Comments").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (db *GormDB) CreateColumn(p *entities.Column) (int, error) {
	if err := db.db.Create(p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (db *GormDB) ReadColumn(id int) (*entities.Column, error) {
	p := &entities.Column{
		ID: id,
	}
	if err := db.db.Preload("Tasks.Comments").First(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

func (db *GormDB) UpdateColumn(p *entities.Column) error {
	if err := db.db.Save(p).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) DeleteColumn(id int) error {
	if err := db.db.Delete(&entities.Column{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) CreateTask(t *entities.Task) (int, error) {
	if err := db.db.Create(t).Error; err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (db *GormDB) ReadTask(tid int) (*entities.Task, error) {
	t := &entities.Task{
		ID: tid,
	}
	if err := db.db.Preload("Comments").First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (db *GormDB) UpdateTask(t *entities.Task) error {
	if err := db.db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) DeleteTask(id int) error {
	if err := db.db.Delete(&entities.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) CreateComment(t *entities.Comment) (int, error) {
	if err := db.db.Create(t).Error; err != nil {
		return 0, err
	}
	return t.ID, nil
}

func (db *GormDB) ReadComment(tid int) (*entities.Comment, error) {
	t := &entities.Comment{}
	if err := db.db.First(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (db *GormDB) UpdateComment(p *entities.Comment) error {
	if err := db.db.Save(p).Error; err != nil {
		return err
	}
	return nil
}

func (db *GormDB) DeleteComment(id int) error {
	if err := db.db.Delete(&entities.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}
