package routes

import (
	"net/http"
	"some-backend/app/resources"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, resources.Items)
}

func GetItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, _item := range resources.Items {
		if _item.ID == id {
			c.IndentedJSON(http.StatusOK, _item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}
