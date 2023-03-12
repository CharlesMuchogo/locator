package main

import (
	"fmt"
	"main.go/location"
	"main.go/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", user.Register)
	public.POST("/login", user.Login)
	public.POST("/location", location.UpdateLocation)
	public.GET("/location", location.GetLocation)

	r.Run(":8000")
	fmt.Println("Server started on port 8080")
}
