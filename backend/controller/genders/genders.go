package genders

import (
	"net/http"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/database"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	db := database.DB
	var genders []entity.Genders
	db.Find(&genders)
	c.JSON(http.StatusOK, &genders)
}