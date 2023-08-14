package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type LocationInput struct {
	Name   string `json:"name" binding:"required"`
	BookId int    `json:"book_id" binding:"required"`
}

type LocationEdit struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateLocation(c *gin.Context) {
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

	var book models.Book
	if err := s.db.Where("id = ? and user_id = ?", LocationInput.BookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	if err := s.db.Preload("Locations").Where("id = ? and user_id = ?", bookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	if err := s.db.Joins("Book").Where("locations.id = ? and book.user_id = ?", LocationId, user.ID).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Location)
}

func (s *Server) EditLocation(c *gin.Context) {

	var locationEdit LocationEdit
	if err := c.ShouldBindJSON(&locationEdit); err != nil {
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
	if err := s.db.Joins("Book").Where("locations.id = ? and book.user_id = ?", LocationId, user.ID).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	Location.Name = locationEdit.Name
	s.db.Save(&Location)

	c.JSON(http.StatusOK, Location)
}

func (s *Server) DeleteLocation(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	locationId := c.Param("locationId")
	var Location models.Location
	if err := s.db.Joins("Book").Where("locations.id = ? and book.user_id = ?", locationId, user.ID).First(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Delete(&Location).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
