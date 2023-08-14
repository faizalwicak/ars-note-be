package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

// func (s *Server) GetGroceries(c *gin.Context) {
// 	var groceries []models.Grocery

// 	if err := s.db.Find(&groceries).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, groceries)
// }

func (s *Server) GetGroceries(c *gin.Context) {
	user, err := utils.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Groceries})
}

func (s *Server) PostGrocery(c *gin.Context) {
	var newGrocery NewGrocery

	if err := c.ShouldBindJSON(&newGrocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	grocery := models.Grocery{Name: newGrocery.Name, Quantity: newGrocery.Quantity}
	grocery.UserId = user.ID

	if err := s.db.Create(&grocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grocery)
}
