package main

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/database"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/middleware"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/controller/users"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/controller/genders"
)

func main() {
	PORT := os.Getenv("PORT")
	database.Connect()

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.POST("/signup", users.SignUp)
	r.POST("/signin", users.SignIn)

	router := r.Group("/")
	router.Use(middleware.Authorizes())
	{
		router.PUT("/user/:id", users.Update)
		router.GET("/users", users.GetAll)
		router.GET("/user/:id", users.Get)
		router.DELETE("/user/:id", users.Delete)
	}

	r.GET("/genders", genders.GetAll)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	r.Run(":" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}