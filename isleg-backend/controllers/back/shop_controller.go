package controllers

import "github.com/gin-gonic/gin"

func CreateShop(c *gin.Context) {

	ownerName := c.PostForm("owner_name")
	address := c.PostForm("address")
	phoneNumber := c.PostForm("phone_number")
	runningTime := c.PostForm("running_time")
	categories, _ := c.GetPostFormArray("category_id")

}
