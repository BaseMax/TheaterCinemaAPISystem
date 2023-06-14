package test

import (
	"TheaterCinemaAPISystem/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestmoviesController(t *testing.T) {
	r := gin.Default()

	r.GET("/movies", controller.Movies_index)
	r.Run()
	req, _ := http.NewRequest("GET", "/movies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Request failed with status: %v", w.Code)
	}

	expected := `{"status":"true"}`
	if w.Body.String() != expected {
		t.Errorf("Unexpected response body: got %v want %v", w.Body.String(), expected)
	}
}
func TestmoviesPOstController(t *testing.T) {
	r := gin.Default()

	r.POST("/categories", controller.Category_create)
	r.Run()
	req, _ := http.NewRequest("POst", "/categories", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Request failed with status: %v", w.Code)
	}

	expected := `{"status":"true"}`
	if w.Body.String() != expected {
		t.Errorf("Unexpected response body: got %v want %v", w.Body.String(), expected)
	}
}