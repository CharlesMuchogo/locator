package main

import (
	"fmt"
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

	r.Run(":8000")
	fmt.Println("Server started on port 8080")

}
