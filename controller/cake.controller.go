package controller

import (
	"cake_store/entities"
	"cake_store/repositories"
	"cake_store/services"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCakeDetail(c *gin.Context) {
	// get cake  id
	cakeId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     "Cake id must be a number/integer",
			"status_code": http.StatusNotFound})
		return
	}

	// cake query from database
	cakeObj, err := services.GetCakeDetail(int(cakeId))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message":     "cake is not found",
				"status_code": http.StatusNotFound,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Internal Server Error",
				"status_code": http.StatusInternalServerError,
			})
			return
		}
	}
	c.JSON(http.StatusOK, cakeObj)
}

func CreateCake(c *gin.Context) {
	var cake entities.Cake
	c.BindJSON(&cake)
	err := services.CreateCake(c, cake)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Internal Server Error",
			"status_code": http.StatusInternalServerError,
			"error":       err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Cake successfully created!",
	})
}

func GetCakes(c *gin.Context) {
	results, err := repositories.GetAllCake()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_data": len(results),
		"data":       results,
	})
}

func UpdateCake(c *gin.Context) {
	// Parse cake id from uri param
	cakeId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     "Cake id must be a number/integer",
			"status_code": http.StatusNotFound})
		return
	}

	// Validate json data
	var cake entities.Cake
	err = c.ShouldBindJSON(&cake)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     "Bad Request",
			"status_code": http.StatusNotFound})
		return
	}

	// execute update logic
	err = services.UpdateCake(c, int(cakeId), cake)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Internal Server Error",
			"status_code": http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     "Cake successfully updated!"})
}

func DeleteCake(c *gin.Context) {
	cakeId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     "Cake id must be a number/integer",
			"status_code": http.StatusNotFound})
		return
	}

	err = services.DeleteCake(c, int(cakeId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     err.Error(),
			"status_code": http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Delete Success",
		"status_code": http.StatusOK})
}
