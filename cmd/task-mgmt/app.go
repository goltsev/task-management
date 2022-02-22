package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/goltsev/task-management/database"
	"github.com/goltsev/task-management/entities"
	"github.com/goltsev/task-management/handlers"
	"github.com/goltsev/task-management/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/goltsev/task-management/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func run() error {
	port := ":" + os.Getenv("PORT")
	dsn := os.Getenv("DATABASE_URL")
	root := os.Getenv("ROOT_URL")
	migrate := os.Getenv("MIGRATE")

	gormdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		FullSaveAssociations: true,
	})
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	if migrate == "true" {
		if err := gormdb.AutoMigrate(
			&entities.Project{},
			&entities.Column{},
			&entities.Task{},
			&entities.Comment{},
		); err != nil {
			return fmt.Errorf("error while migrating: %w", err)
		}
	}
	db := database.CreateGormDB(gormdb)
	svc := service.NewService(db)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", handlers.NewMux(svc))
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL(root+"/swagger/doc.json"),
	))
	log.Println("starting server on", port)
	return http.ListenAndServe(port, mux)
}
