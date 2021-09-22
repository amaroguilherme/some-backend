package mixins

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"some-backend/app/resources"
	"some-backend/app/routes"

	"github.com/gin-gonic/gin"
)

type Assertions struct {
	t TestingT
}

type TestingT interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Logf(format string, args ...interface{})
}

func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

var items []resources.Item
var responseRecorder = httptest.NewRecorder()

func TearDown() {
	items = []resources.Item{}
	responseRecorder = httptest.NewRecorder()
}

func GivenItemParams(_id string, _title string) {
	item := resources.Item{
		ID:    _id,
		Title: _title,
	}
	items = append(items, item)
	resources.Items = items
}

func WhenAccessRoute(route string) {
	r := gin.Default()
	r.GET(route, routes.GetItems)

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println("Couldn't create request: %v\n", err)
	}

	r.ServeHTTP(responseRecorder, req)
}

func (a *Assertions) AssertHttpStatus200(t TestingT) {
	if responseRecorder.Code == http.StatusOK {
		a.t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, responseRecorder.Code)
	} else {
		a.t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, responseRecorder.Code)
	}

	TearDown()
}

func (a *Assertions) AssertResponseIs(t TestingT, expected string) {
	if responseRecorder.Body.String() == expected {
		a.t.Logf("Expected to get response %d is same as %d\n", expected, responseRecorder.Body.String())
	} else {
		a.t.Fatalf("Expected to get response %d but instead got %d\n", expected, responseRecorder.Body.String())
	}

	TearDown()
}
