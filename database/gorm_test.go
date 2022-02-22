package database_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goltsev/task-management/database"
	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dsn = "host=192.168.0.3 user=postgres password=root port=5432 dbname=testpostgres"

	gormdb *gorm.DB
)

func prepareDB() error {
	if gormdb != nil {
		return nil
	}
	var err error
	var conf *gorm.Config
	conf = &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		FullSaveAssociations: true,
	}
	gormdb, err = gorm.Open(postgres.Open(dsn), conf)
	if err != nil {
		return err
	}
	prj := &entities.Project{}
	col := &entities.Column{}
	tsk := &entities.Task{}
	cmt := &entities.Comment{}
	gormdb.Migrator().DropTable(prj)
	gormdb.Migrator().DropTable(col)
	gormdb.Migrator().DropTable(tsk)
	gormdb.Migrator().DropTable(cmt)
	if err := gormdb.AutoMigrate(prj, col, tsk, cmt); err != nil {
		return err
	}
	return nil
}

func TestCreateProject(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)
	len := 5
	for i := 0; i < len; i++ {
		name, _ := helpers.GenerateString(10)
		p, err := entities.NewProject(name, "")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := db.CreateProject(p); err != nil {
			t.Fatal(err)
		}
	}
	ps := make([]*entities.Project, len)
	if err := gormdb.Find(&ps).Error; err != nil {
		t.Fatal(err)
	}

}

func TestReadProject(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)
	name, _ := helpers.GenerateString(10)
	p, err := entities.NewProject(name, "")
	if err != nil {
		t.Fatal(err)
	}
	if err := gormdb.Create(p).Error; err != nil {
		t.Fatal(err)
	}
	np, err := db.ReadProject(p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != np.Name {
		t.Errorf("expected:\n%v\ngot:\n%v\n", p, np)
	}
}

func TestUpdateProjectNameDesc(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)
	name, _ := helpers.GenerateString(10)
	p, err := entities.NewProject(name, "")
	if err != nil {
		t.Fatal(err)
	}
	id, err := db.CreateProject(p)
	if err != nil {
		t.Fatal(err)
	}
	p.Name = p.Name + "ADDTEXT"
	p.Description = p.Description + "ADDTEXT"
	if err := db.UpdateProject(p); err != nil {
		t.Fatal(err)
	}
	np, err := db.ReadProject(id)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != np.Name {
		t.Errorf("expected:\n%v\ngot:\n%v\n", p, np)
	}
}

func TestUpdateProjectAddColumn(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)

	pname, _ := helpers.GenerateString(10)
	cname := "new column"
	p, err := entities.NewProject(pname, "")
	if err != nil {
		t.Fatal(err)
	}
	id, err := db.CreateProject(p)
	if err != nil {
		t.Fatal(err)
	}
	col, _ := entities.NewColumn(cname)
	p.AddColumn(col)
	if err := db.UpdateProject(p); err != nil {
		t.Fatal(err)
	}
	np, err := db.ReadProject(id)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(np.Columns); i++ {
		if np.Columns[i].Name != p.Columns[i].Name {
			t.Fatalf("project was not updated properly:\nexpected:\n%v\ngot:\n%v\n",
				helpers.PrintProject(p), helpers.PrintProject(np))
		}
	}
}

func TestDeleteProject(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)
	name, _ := helpers.GenerateString(10)
	p, err := entities.NewProject(name, "DELETE")
	if err != nil {
		t.Fatal(err)
	}
	id, err := db.CreateProject(p)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.DeleteProject(id); err != nil {
		t.Fatal(err)
	}
	_, err = db.ReadProject(id)
	if err == nil || err != gorm.ErrRecordNotFound {
		t.Errorf("expected error: %v, got nothing", gorm.ErrRecordNotFound)
	}
}

func TestReadTask(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}

	p, _ := entities.NewProject("project", "")
	c, _ := p.GetColumnPosition(0)
	expected, _ := entities.NewTask("task", "")
	expected.Comments = []*entities.Comment{}
	c.AddTask(expected)
	if err := gormdb.Create(p).Error; err != nil {
		t.Fatal(err)
	}

	db := database.CreateGormDB(gormdb)
	got, err := db.ReadTask(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nexpected:\n%v\ngot:\n%v\n", expected, got)
	}
}

func TestUpdateProjectDeleteColumn(t *testing.T) {
	if err := prepareDB(); err != nil {
		t.Fatal(err)
	}
	db := database.CreateGormDB(gormdb)

	p, _ := entities.NewProject("project", "")
	id, err := db.CreateProject(p)
	if err != nil {
		t.Fatal(err)
	}
	pos := 0
	// c0, _ := p.GetColumnPosition(pos)
	c1, _ := entities.NewColumn("c1")
	p.AddColumn(c1)
	if err := db.UpdateProject(p); err != nil {
		t.Fatal(err)
	}
	if err := p.DeleteColumn(pos); err != nil {
		t.Fatal(err)
	}
	fmt.Println("AFTER DELETION")
	if err := db.UpdateProject(p); err != nil {
		t.Fatal(err)
	}
	// for _, c := range p.Columns {
	// 	db.UpdateColumn(c)
	// }
	// db.DeleteColumn(c0.ID)
	np, err := db.ReadProject(id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(helpers.PrintProject(p))
	fmt.Println(helpers.PrintProject(np))
}
