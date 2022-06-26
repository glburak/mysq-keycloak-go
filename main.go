package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm-mysql/controller"
	"gorm-mysql/middleware"
	"github.com/gin-contrib/cors"

)

func main() {
	r := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge:   60 * 60 * 24,Path: "/", HttpOnly: true, Domain: "localhost",Secure: false })
	r.Use(sessions.Sessions("SESSIONID22", store),cors.Default())
	//r.Use(cors.Default())
	r.GET("/user",controller.User)
	r.POST("/login",controller.Login)
	r.GET("/logout", controller.Logout)
	auth := r.Group("/auth")
	auth.Use(middleware.Authentication())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})
	}
	r.Run(":8088")
}
