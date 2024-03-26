package controllers

import (
	"gin-ranking/pkg/logger"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u User) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)
}

func (u User) GetList(c *gin.Context) {
	logger.Write("日志信息", "user")
	ReturnError(c, 4004, "获取用户列表失败")
}
