package controllers

import "github.com/gin-gonic/gin"
import  "github.com/Rashad-Muntar/my-go-rest/config"
import  "github.com/Rashad-Muntar/my-go-rest/models"


func CreateShoper(c *gin.Context){
	var shopper  models.Shopper
	c.BindJSON(shopper)
	config.DB.Create(shopper)
}
