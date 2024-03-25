package controllers

import "github.com/gin-gonic/gin"

type User struct {
}

func (u User) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}

func (u User) GetList(c *gin.Context) {
	ReturnError(c, 4004, "没有相关信息")
}
