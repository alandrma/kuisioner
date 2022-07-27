package main

import (
	"kuisioner/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	_ = r.Run(":8585")
}
