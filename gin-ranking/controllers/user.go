package controllers

import (
	"gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
}

func (u User) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)
}

func (u User) GetList(c *gin.Context) {
	//logger.Write("日志信息", "user")
	ReturnError(c, 4004, "获取用户列表失败")
}
