package main

import (
	"github.com/Panmax/gin-template/interfaces/facade"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	leaveApi := api.Group("/leaves")
	facade.LeaveRouterRegister(leaveApi)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
