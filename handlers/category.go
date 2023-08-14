package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type CategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateCategory(c *gin.Context) {
	var categoryInput CategoryInput

	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookId := c.Param("bookId")
	var book models.Book
	if err := s.db.Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	category := models.Category{Name: categoryInput.Name, BookId: book.ID}

	if err := s.db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (s *Server) ListCategory(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.Book
	if err := s.db.Preload("Categories").Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	c.JSON(http.StatusOK, book.Categories)
}

func (s *Server) GetCategory(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryId := c.Param("categoryId")
	var category models.Category
	if err := s.db.Preload("Book").Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if category.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (s *Server) EditCategory(c *gin.Context) {

	var categoryInput CategoryInput

	if err := c.ShouldBindJSON(&categoryInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryId := c.Param("categoryId")
	var category models.Category
	if err := s.db.Preload("Book").Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if category.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	category.Name = categoryInput.Name
	s.db.Save(&category)

	c.JSON(http.StatusOK, category)
}

func (s *Server) DeleteCategory(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryId := c.Param("categoryId")
	var category models.Category
	if err := s.db.Preload("Book").Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if category.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	if err := s.db.Delete(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
