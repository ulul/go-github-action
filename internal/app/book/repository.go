package book

import "gorm.io/gorm"

type Repository interface {
	Get() ([]Book, error)
	FindByID(id string) (Book, error)
	Create(book Book) (Book, error)
	Delete(id string) error
	UpdateByID(id string, book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]Book, error) {
	var books []Book

	if e := r.db.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func (r *repository) FindByID(id string) (Book, error) {
	var book Book
	if e := r.db.Where("id = ?", id).First(&book).Error; e != nil {
		return book, e
	}
	return book, nil
}

func (r *repository) Create(book Book) (Book, error) {
	if e := r.db.Create(&book).Error; e != nil {
		return book, e
	}
	return book, nil
}

func (r *repository) Delete(id string) error {
	var book Book
	if e := r.db.Where("id = ?", id).Delete(&book).Error; e != nil {
		return e
	}
	return nil
}

func (r *repository) UpdateByID(id string, book Book) (Book, error) {
	if e := r.db.Where("id = ?", id).Updates(&book).Error; e != nil {
		return book, e
	}
	return book, nil
}
