package web

import (
	"some-backend/app/routes"

	"github.com/gin-gonic/gin"
)

func CreateApp() {
	router := gin.Default()
	router.GET("/items", routes.GetItems)
	router.GET("/items/:id", routes.GetItemByID)

	router.Run("localhost:8080")
}
