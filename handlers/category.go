package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type CategoryInput struct {
	Name   string `json:"name" binding:"required"`
	BookId int    `json:"book_id" binding:"required"`
}

type CategoryEdit struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateCategory(c *gin.Context) {
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

	var book models.Book
	if err := s.db.Where("id = ? and user_id = ?", categoryInput.BookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	bookId, valid := c.GetQuery("book")

	if !valid {
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := s.db.Preload("Categories").Where("id = ? and user_id = ?", bookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	if err := s.db.Joins("Book").Where("categories.id = ? and book.user_id = ?", categoryId, user.ID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (s *Server) EditCategory(c *gin.Context) {

	var categoryEdit CategoryEdit
	if err := c.ShouldBindJSON(&categoryEdit); err != nil {
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
	if err := s.db.Joins("Book").Where("categories.id = ? and book.user_id = ?", categoryId, user.ID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	category.Name = categoryEdit.Name
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
	if err := s.db.Joins("Book").Where("categories.id = ? and book.user_id = ?", categoryId, user.ID).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Delete(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
