package test

import (
	"TheaterCinemaAPISystem/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestcategoriesController(t *testing.T) {
	r := gin.Default()

	r.GET("/categories", controller.Category_index)
	r.Run()
	req, _ := http.NewRequest("GET", "/categories", nil)
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
