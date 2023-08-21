package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/faizalwicak/ars-note-be/models"
	"github.com/faizalwicak/ars-note-be/utils"
	"github.com/gin-gonic/gin"
)

type Date time.Time

var _ json.Unmarshaler = &Date{}

func (mt *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = Date(t)
	return nil
}

type TransactionInput struct {
	Date        Date   `json:"date" binding:"required"`
	Value       int    `json:"value" binding:"required"`
	Description string `json:"description"`
	BookId      uint   `json:"book_id" binding:"required"`
	CategoryId  *uint  `json:"category_id,omitempty"`
	LocationId  *uint  `json:"location_id,omitempty"`
}

type TransactionEdit struct {
	Date        Date   `json:"date" binding:"required"`
	Value       int    `json:"value" binding:"required"`
	Description string `json:"description"`
	CategoryId  *uint  `json:"category_id"`
	LocationId  *uint  `json:"location_id"`
}

func (s *Server) CreateTransaction(c *gin.Context) {
	var transactionInput TransactionInput

	if err := c.ShouldBindJSON(&transactionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	if err := s.db.Where("id = ? and user_id = ?", transactionInput.BookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if transactionInput.CategoryId != nil {
		var category models.Category
		if err := s.db.Where("id = ? and book_id = ?", transactionInput.CategoryId, transactionInput.BookId).First(&category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
	}

	if transactionInput.LocationId != nil {
		var location models.Location
		if err := s.db.Where("id = ? and book_id = ?", transactionInput.LocationId, transactionInput.BookId).First(&location).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
			return
		}
	}

	transaction := models.Transaction{Date: time.Time(transactionInput.Date), Value: transactionInput.Value, Description: transactionInput.Description, BookId: book.ID, LocationId: transactionInput.LocationId, CategoryId: transactionInput.CategoryId}

	if err := s.db.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (s *Server) ListTransaction(c *gin.Context) {
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
	if err := s.db.Preload("Transactions").Where("id = ? and user_id = ?", bookId, user.ID).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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

	transactionId := c.Param("transactionId")
	var transaction models.Transaction
	if err := s.db.Joins("Book").Where("transactions.id = ? and Book.user_id = ?", transactionId, user.ID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (s *Server) EditTransaction(c *gin.Context) {

	var transactionEdit TransactionEdit

	if err := c.ShouldBindJSON(&transactionEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId := c.Param("transactionId")
	var transaction models.Transaction
	if err := s.db.Joins("Book").Where("transactions.id = ? and Book.user_id = ?", transactionId, user.ID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if transactionEdit.CategoryId != nil {
		var category models.Category
		if err := s.db.Where("id = ? and book_id = ?", transactionEdit.CategoryId, transaction.BookId).First(&category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}
	}

	if transactionEdit.LocationId != nil {
		var location models.Location
		if err := s.db.Where("id = ? and book_id = ?", transactionEdit.LocationId, transaction.BookId).First(&location).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
			return
		}
	}

	transaction.Date = time.Time(transactionEdit.Date)
	transaction.Value = transactionEdit.Value
	transaction.Description = transactionEdit.Description
	transaction.CategoryId = transactionEdit.CategoryId
	transaction.LocationId = transactionEdit.LocationId

	s.db.Save(&transaction)

	c.JSON(http.StatusOK, transaction)
}

func (s *Server) DeleteTransaction(c *gin.Context) {

	user, err := utils.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId := c.Param("transactionId")
	var transaction models.Transaction
	if err := s.db.Joins("Book").Where("transactions.id = ? and Book.user_id = ?", transactionId, user.ID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := s.db.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
