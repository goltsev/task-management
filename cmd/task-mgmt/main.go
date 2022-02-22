package main

import (
	"log"
)

// @title Task Management API
// @BasePath /api/v1
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
