package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateBook(c *gin.Context) {
	var bookInput BookInput

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := models.Book{Name: bookInput.Name}
	book.UserId = user.ID

	if err := s.db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (s *Server) ListBook(c *gin.Context) {
	user, err := utils.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var books []models.Book
	if err := s.db.Where("user_id = ?", user.ID).Find(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (s *Server) GetBook(c *gin.Context) {

	bookId := c.Param("bookId")

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := s.db.Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (s *Server) EditBook(c *gin.Context) {

	bookId := c.Param("bookId")

	var bookInput BookInput

	if err := c.ShouldBindJSON(&bookInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := s.db.Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	book.Name = bookInput.Name
	s.db.Save(&book)

	c.JSON(http.StatusOK, book)
}

func (s *Server) DeleteBook(c *gin.Context) {

	bookId := c.Param("bookId")

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := s.db.Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	if err := s.db.Where("id = ?", bookId).Delete(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
