package controllers

import (
	"net/http"

	models "github.com/go-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var Users []models.User

	db.Find(&Users)

	c.JSON(http.StatusOK, gin.H{"data": Users})

}
