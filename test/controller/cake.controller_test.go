package controller_test

import (
	"cake_store/controller"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCakeDetail(t *testing.T) {
	// get cake  id
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.AddParam("id", "2")
	controller.GetCakeDetail(c)
	assert.Equal(t, 200, w.Code)
}

func TestGetCakes(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.GetCakes(c)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteCake(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.AddParam("id", "1")
	controller.DeleteCake(c)
	assert.Equal(t, 200, w.Code)
}
