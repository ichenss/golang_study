package main

import (
	"fmt"
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	err := r.Run()
	if err != nil {
		fmt.Println("r.Run() err: ", err)
		return
	}
}
