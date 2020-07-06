package main

import (
	"zhengrchan/ginessential/data"

	"github.com/gin-gonic/gin"
)

func main() {
	db := data.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())

}
