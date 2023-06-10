package test

import (
	"fmt"
	"io"
	"io/ioutil"
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
func TestPingRoute(t *testing.T) {
	r := gin.Default()


  // Handlers for testing
  r.GET("/categories", controller.Category_index)

  // Setup the router
  mockResponse := `{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2023-06-05T09:01:02.023Z",
            "UpdatedAt": "2023-06-05T09:01:02.023Z",
            "DeletedAt": null,
            "title": "sad",
            "slug": "sad",
            "parint_id": 0
        },
        {
            "ID": 2,
            "CreatedAt": "2023-06-05T09:07:25.321Z",
            "UpdatedAt": "2023-06-05T16:43:56.454Z",
            "DeletedAt": null,
            "title": "دسته بندی شماره 2",
            "slug": "دسته_بندی_جدید",
            "parint_id": 0
        },
        {
            "ID": 3,
            "CreatedAt": "2023-06-05T09:07:39.331Z",
            "UpdatedAt": "2023-06-05T09:07:39.331Z",
            "DeletedAt": null,
            "title": "sad",
            "slug": "sad",
            "parint_id": 0
        },
        {
            "ID": 4,
            "CreatedAt": "2023-06-05T09:09:00.366Z",
            "UpdatedAt": "2023-06-05T09:09:00.366Z",
            "DeletedAt": null,
            "title": "",
            "slug": "",
            "parint_id": 0
        }
    ],
    "status": true
}`

  req, _ := http.NewRequest("GET", "/categories", nil)
  fmt.Print("hi",req)
  w := httptest.NewRecorder()
  r.ServeHTTP(w, req)
fmt.Println(w)
  responseData, _ := ioutil.ReadAll(w.Body)
  assert.Equal(t, mockResponse, string(responseData))
  assert.Equal(t, http.StatusOK, w.Code)


}