package controllers

import "github.com/gin-gonic/gin"

type OrderController struct {
}

type Search struct {
	Cid  int    `json:"cid"`
	Name string `json:"name"`
}

func (o OrderController) GetList(c *gin.Context) {
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "chen")

	//param := make(map[string]interface{})
	//err := c.BindJSON(&param)

	search := Search{}
	err := c.BindJSON(&search)
	if err != nil {
		ReturnError(c, 10001, gin.H{"error": err.Error()})
		return
	}
	ReturnSuccess(c, 0, search.Name, search.Cid, 1)
}
