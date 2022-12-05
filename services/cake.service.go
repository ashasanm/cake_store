package services

import (
	"cake_store/entities"
	"cake_store/repositories"

	"github.com/gin-gonic/gin"
)

func CreateCake(c *gin.Context, cake entities.Cake) error {
	err := repositories.CreateCake(c, cake)
	if err != nil {
		return err
	}
	return nil
}

func GetCakeDetail(cakeId int) (entities.Cake, error) {
	result, err := repositories.GetOneCake(cakeId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetCakes() ([]entities.Cake, error) {
	result, err := repositories.GetAllCake()
	if err != nil {
		return result, err
	}

	return result, nil
}

func UpdateCake(c *gin.Context, cakeId int, cake entities.Cake) error {
	err := repositories.UpdateCake(c, cakeId, cake)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCake(c *gin.Context, cakeId int) error {
	err := repositories.DeleteCake(c, cakeId)
	if err != nil {
		return err
	}
	return nil
}
