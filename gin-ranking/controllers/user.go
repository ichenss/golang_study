package controllers

import "github.com/gin-gonic/gin"

type User struct {
}

func (u User) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "user info", 1)
}

func (u User) GetList(c *gin.Context) {
	ReturnError(c, 4004, "没有相关信息")
}
