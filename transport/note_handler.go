package transport

import (
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sprectza/note-taking-app.git/repository"
	"github.com/sprectza/note-taking-app.git/service"
)

type NoteInput struct {
	Content string `json:"content" binding:"required"`
}

type NoteHandler struct {
	noteService service.NoteService
}

func NewNoteHandler(noteService service.NoteService) *NoteHandler {
	return &NoteHandler{noteService: noteService}
}

func (h *NoteHandler) ListNotes(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["id"].(float64))

	notes, err := h.noteService.GetNotesByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"notes": notes})
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["id"].(float64))

	var noteInput NoteInput
	if err := c.ShouldBindJSON(&noteInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	note := repository.Note{
		UserID:  userID,
		Content: noteInput.Content,
	}

	err := h.noteService.CreateNote(note)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "note created successfully"})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := uint(claims["id"].(float64))

	noteIDStr := c.Param("id")

	noteID, err := strconv.ParseUint(noteIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	note, err := h.noteService.GetNoteByID(uint(noteID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if note.UserID != userID {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	err = h.noteService.DeleteNoteByID(uint(noteID))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "note deleted successfully"})
}
