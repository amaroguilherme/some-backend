package web

import (
	"some-backend/app/mixins"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestAssert200Status(t *testing.T) {
	assert := mixins.New(t) //HOW TO IMPROVE THIS?

	mixins.GivenItemParams("1", "Gui")

	mixins.WhenAccessRoute("/items")

	assert.AssertHttpStatus200(t)

}

func TestAssertResponseSingleItem(t *testing.T) {
	assert := mixins.New(t)

	mixins.GivenItemParams("1", "Gui")

	mixins.WhenAccessRoute("/items")

	assert.AssertResponseIs(t, `[
    {
        "ID": "1",
        "Title": "Gui"
    }
]`)

}
