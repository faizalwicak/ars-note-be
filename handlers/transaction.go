package handlers

import (
	"net/http"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	Name string `json:"name" binding:"required"`
}

func (s *Server) CreateTransaction(c *gin.Context) {
	var TransactionInput TransactionInput

	if err := c.ShouldBindJSON(&TransactionInput); err != nil {
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

	Transaction := models.Transaction{Name: TransactionInput.Name, BookId: book.ID}

	if err := s.db.Create(&Transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, Transaction)
}

func (s *Server) ListTransaction(c *gin.Context) {
	bookId := c.Param("bookId")
	var book models.Book
	if err := s.db.Preload("Transactions").Where("id = ?", bookId).First(&book).Error; err != nil {
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

	c.JSON(http.StatusOK, book.Transactions)
}

func (s *Server) GetTransaction(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	TransactionId := c.Param("transactionId")
	var Transaction models.Transaction
	if err := s.db.Preload("Book").Where("id = ?", TransactionId).First(&Transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Transaction.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	c.JSON(http.StatusOK, Transaction)
}

func (s *Server) EditTransaction(c *gin.Context) {

	var TransactionInput TransactionInput

	if err := c.ShouldBindJSON(&TransactionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	TransactionId := c.Param("transactionId")
	var Transaction models.Transaction
	if err := s.db.Preload("Book").Where("id = ?", TransactionId).First(&Transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Transaction.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	Transaction.Name = TransactionInput.Name
	s.db.Save(&Transaction)

	c.JSON(http.StatusOK, Transaction)
}

func (s *Server) DeleteTransaction(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	TransactionId := c.Param("TransactionId")
	var Transaction models.Transaction
	if err := s.db.Preload("Book").Where("id = ?", TransactionId).First(&Transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if Transaction.Book.UserId != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You not allowed to access this book"})
		return
	}

	if err := s.db.Delete(&Transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
