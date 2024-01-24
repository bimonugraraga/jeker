package services

import (
	"errors"

	g "github.com/bdn/jeker/db"
	"github.com/bdn/jeker/models"
)

var (
	errCategoryNotFound   = errors.New("Category Not Found")
	errFailedToCreateBook = errors.New("Failed To Create Book")
	errBookNotFound       = errors.New("Book Not Found")
)

func CreateBookService(params models.Book) (models.Book, error) {
	//Check If Category Exist
	var newBook models.Book
	var err error
	var db = g.GetConn()
	var targetCategory models.Category
	result := db.Where("id = ?", params.CategoryId).Find(&targetCategory)
	err = result.Error
	if err != nil || targetCategory.ID == 0 {
		return newBook, errCategoryNotFound
	}
	params.Category.ID = targetCategory.ID
	params.Category.Name = targetCategory.Name
	params.User.ID = params.UserId
	params.Status = models.BookStatus.BookStatusDraft
	result = db.Model(&newBook).Joins("Category").Joins("User").Create(&params)
	err = result.Error
	if err != nil {
		return newBook, errFailedToCreateBook
	}
	return params, nil
}

func UpdateBookService(params models.Book) (models.Book, error) {
	var targetBook models.Book
	var err error
	var db = g.GetConn()
	result := db.Where("books.id = ?", params.ID).Where("user_id = ?", params.UserId).Joins("Category").Find(&targetBook)
	err = result.Error
	if err != nil || targetBook.ID == 0 {
		return targetBook, errBookNotFound
	}
	result = db.Model(&targetBook).Updates(params)
	err = result.Error
	if err != nil {
		return targetBook, errBookNotFound
	}
	params.CategoryId = targetBook.CategoryId
	params.Category.Name = targetBook.Category.Name
	return params, nil
}
