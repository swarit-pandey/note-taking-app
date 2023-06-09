package repository

import "github.com/jinzhu/gorm"

type NoteRepository interface {
	CreateNote(note Note) error
	GetNotesByUserID(userID uint) ([]Note, error)
	DeleteNoteByID(id uint) error
	GetNoteByID(id uint) (Note, error)
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) CreateNote(note Note) error {
	return r.db.Create(&note).Error
}

func (r *noteRepository) GetNotesByUserID(userID uint) ([]Note, error) {
	var notes []Note
	err := r.db.Where("user_id = ?", userID).Find(&notes).Error
	return notes, err
}

func (r *noteRepository) DeleteNoteByID(id uint) error {
	return r.db.Where("id = ?", id).Delete(Note{}).Error
}

func (r *noteRepository) GetNoteByID(id uint) (Note, error) {
	var note Note
	err := r.db.First(&note, id).Error
	return note, err
}
