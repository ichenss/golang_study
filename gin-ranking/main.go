package main

import (
	"fmt"
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("r.Run() err: ", err)
		return
	}
}
