package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type LocationInput struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateLocation(c *gin.Context) {
	var LocationInput LocationInput

	if err := c.ShouldBindJSON(&LocationInput); err != nil {
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

	Location := models.Location{Name: LocationInput.Name, BookId: book.ID}

	if err := s.db.Create(&Location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Location)
}

func (s *Server) ListLocation(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.Book
	if err := s.db.Preload("Locations").Where("id = ?", bookId).First(&book).Error; err != nil {
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

	c.JSON(http.StatusOK, book.Locations)
}

func (s *Server) GetLocation(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	LocationId := c.Param("locationId")
	var Location models.Location
	if err := s.db.Preload("Book").Where("id = ?", LocationId).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Location.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	c.JSON(http.StatusOK, Location)
}

func (s *Server) EditLocation(c *gin.Context) {

	var LocationInput LocationInput

	if err := c.ShouldBindJSON(&LocationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	LocationId := c.Param("locationId")
	var Location models.Location
	if err := s.db.Preload("Book").Where("id = ?", LocationId).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Location.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	Location.Name = LocationInput.Name
	s.db.Save(&Location)

	c.JSON(http.StatusOK, Location)
}

func (s *Server) DeleteLocation(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	LocationId := c.Param("LocationId")
	var Location models.Location
	if err := s.db.Preload("Book").Where("id = ?", LocationId).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Location.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	if err := s.db.Delete(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
