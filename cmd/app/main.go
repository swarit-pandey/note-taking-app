package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sprectza/note-taking-app.git/repository"
	"github.com/sprectza/note-taking-app.git/service"
	"github.com/sprectza/note-taking-app.git/transport"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnectionString := os.Getenv("DB_CONNECTION_STRING_NOTEAPP")

	db, err := gorm.Open("postgres", dbConnectionString)
	if err != nil {
		log.Print("Error connecting to database")
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&repository.User{}, &repository.Note{})
	userRepo := repository.NewUserRepository(db)
	noteRepo := repository.NewNoteRepository(db)

	userService := service.NewUserService(userRepo)
	noteService := service.NewNoteService(noteRepo)

	r, err := transport.NewRouter(userService, noteService)
	if err != nil {
		log.Print("Error creating router")
		panic(err)
	}

	r.Run(":8080")
}
