package main

import (
	"github.com/faizalwicak/ars-note-be/handlers"
	"github.com/faizalwicak/ars-note-be/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router := r.Group("/api")
	router.POST("/register", server.Register)
	router.POST("/login", server.Login)

	authorized := r.Group("/api")
	authorized.Use(middleware.JwtAuthMiddleware())

	authorized.POST("/books", server.CreateBook)
	authorized.GET("/books", server.ListBook)
	authorized.GET("/book/:bookId", server.GetBook)
	authorized.DELETE("/book/:bookId", server.DeleteBook)
	authorized.PUT("/book/:bookId", server.EditBook)

	authorized.POST("/book/:bookId/categories", server.CreateCategory)
	authorized.GET("/book/:bookId/categories", server.ListCategory)
	authorized.GET("/category/:categoryId", server.GetCategory)
	authorized.PUT("/category/:categoryId", server.EditCategory)
	authorized.DELETE("/category/:categoryId", server.DeleteCategory)

	authorized.POST("/book/:bookId/locations", server.CreateLocation)
	authorized.GET("/book/:bookId/locations", server.ListLocation)
	authorized.GET("/location/:locationId", server.GetLocation)
	authorized.PUT("/location/:locationId", server.EditLocation)
	authorized.DELETE("/location/:locationId", server.DeleteLocation)

	// authorized.GET("/groceries", server.GetGroceries)
	// authorized.POST("/grocery", server.PostGrocery)

	return r
}
