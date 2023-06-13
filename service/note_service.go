package service

import "github.com/sprectza/note-taking-app.git/repository"

type NoteService interface {
	CreateNote(note repository.Note) error
	GetNotesByUserID(userID uint) ([]repository.NoteResponse, error)
	DeleteNoteByID(id uint) error
	GetNoteByID(id uint) (repository.Note, error)
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{noteRepository: noteRepository}
}

func (s *noteService) CreateNote(note repository.Note) error {
	return s.noteRepository.CreateNote(note)
}

func (s *noteService) GetNotesByUserID(userID uint) ([]repository.NoteResponse, error) {
	return s.noteRepository.GetNotesByUserID(userID)
}

func (s *noteService) DeleteNoteByID(id uint) error {
	return s.noteRepository.DeleteNoteByID(id)
}

func (s *noteService) GetNoteByID(id uint) (repository.Note, error) {
	return s.noteRepository.GetNoteByID(id)
}
