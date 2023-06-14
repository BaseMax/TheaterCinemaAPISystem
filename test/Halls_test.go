package test

import (
	"TheaterCinemaAPISystem/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHallsController(t *testing.T) {
	r := gin.Default()

	r.GET("/Halls", controller.Hall_index)
	r.Run()
	req, _ := http.NewRequest("GET", "/Halls", nil)
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
func TestHallsPOstController(t *testing.T) {
	r := gin.Default()

	r.POST("/Halls", controller.Hall_create)
	r.Run()
	req, _ := http.NewRequest("POst", "/Halls", nil)
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