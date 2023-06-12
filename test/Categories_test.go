package test

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"TheaterCinemaAPISystem/controller"
	"TheaterCinemaAPISystem/initializers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func init()  {

	initializers.LoadEnvVarables()
	initializers.ConectToDb()
	
	}
  func TestCategoryIndex(t *testing.T) {
    // Create a new HTTP request
    req, err := http.NewRequest("GET", "/categories", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new HTTP response recorder
    rr := httptest.NewRecorder()

    // Create a new gin Engine instance
    r := gin.Default()

    // Register the Category_index function as the handler for the GET /categories route
    r.GET("/categories", controller.Category_index)

    // Call the handler function by passing the request and response recorder
    r.ServeHTTP(rr, req)

    // Check that the HTTP status code is 200 OK
    assert.Equal(t, http.StatusOK, rr.Code)

    // Check that the response body contains the expected JSON data
    expected := `{
        
        "status": true
    }`
    assert.JSONEq(t, expected, rr.Body.String())
}
