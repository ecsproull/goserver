package main

import (
	"log"
	"net/http"

	"goserver/internal/config"
	"goserver/internal/database"
	"goserver/internal/handlers"
	"goserver/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()
	database.InitMongo(cfg.DatabaseURL)

	// Initialize Gin router
	router := gin.Default()

	// Setup CORS
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3001")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		log.Printf("CORS headers set for origin: %s", origin)

		if c.Request.Method == "OPTIONS" {
			log.Printf("Handling OPTIONS preflight request")
			log.Printf("Requested headers: %s", c.Request.Header.Get("Access-Control-Request-Headers"))
			log.Printf("Requested method: %s", c.Request.Header.Get("Access-Control-Request-Method"))
			log.Printf("All request headers: %+v", c.Request.Header)
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Add a debug middleware to see what's happening
	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Origin: %s", c.Request.Header.Get("Origin"))
		c.Next()
		log.Printf("Response headers: %+v", c.Writer.Header())
	})

	// Add middleware
	router.Use(middleware.Logger())

	// API routes
	api := router.Group("/api/v1")
	{
		authHandler := handlers.NewAuthHandler()
		apiRoutes := api.Group("/auth")
		{
			apiRoutes.POST("/login", authHandler.Login)
			apiRoutes.POST("/signup", authHandler.Signup)
			apiRoutes.POST("/logout", authHandler.Logout)
			apiRoutes.POST("/refresh", authHandler.RefreshToken)
			apiRoutes.POST("/resend-verification", authHandler.ResendVerificationEmail)
		}

		// Blog routes
		blogHandler := handlers.NewBlogHandler()
		blogRoutes := api.Group("/blog")
		{
			blogRoutes.GET("/", blogHandler.GetAll)
			blogRoutes.GET("/:id", blogHandler.GetByID)
			blogRoutes.POST("/", middleware.RequireAuth(), blogHandler.Create)
			blogRoutes.DELETE("/:id", middleware.RequireAuth(), blogHandler.Delete)
		}

		// Comment routes
		commentHandler := handlers.NewCommentHandler()
		commentRoutes := api.Group("/comments")
		{
			commentRoutes.GET("/blog/:blogId", commentHandler.GetByBlogID)
			commentRoutes.POST("/blog/:blogId", middleware.RequireAuth(), middleware.RequireRole("Commentor", "Creator", "Admin"), commentHandler.Create)
			commentRoutes.PUT("/blog/:blogId/:id", middleware.RequireAuth(), middleware.RequireRole("Creator", "Admin"), commentHandler.Update)
			commentRoutes.DELETE("/blog/:blogId/:id", middleware.RequireAuth(), middleware.RequireRole("Creator", "Admin"), commentHandler.Delete)
		}

		userHandler := handlers.NewUserHandler()
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.GetAll)
			userRoutes.GET("/:id", userHandler.GetByID)
			userRoutes.POST("", userHandler.Create)
			userRoutes.PUT("/:id", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.Update)
			userRoutes.DELETE("/:id", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.Delete)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	port := cfg.Port
	if port == "" {
		port = "3003"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
