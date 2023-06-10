package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sprectza/note-taking-app.git/service"
)

func NewRouter(userService service.UserService, noteService service.NoteService) (*gin.Engine, error) {
	userHandler := NewUserHandler(userService)
	noteHandler := NewNoteHandler(noteService)

	authMiddleware, err := userHandler.AuthMiddleware()
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/signup", userHandler.RegisterUser)

	noteRoutes := r.Group("/notes")
	noteRoutes.Use(authMiddleware.MiddlewareFunc())
	{
		noteRoutes.GET("/list", noteHandler.ListNotes)
		noteRoutes.POST("/create", noteHandler.CreateNote)
		noteRoutes.DELETE("/delete/:id", noteHandler.DeleteNote)
	}

	return r, nil
}
