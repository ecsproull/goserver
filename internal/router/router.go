package router

import (
	"goserver/internal/config"
	"goserver/internal/handlers"
	"goserver/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	cfg := config.Load()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.FrontendURL)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Debug middleware
	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		log.Printf("Origin: %s", c.Request.Header.Get("Origin"))
		c.Next()
		log.Printf("Response headers: %+v", c.Writer.Header())
	})

	router.Use(middleware.Logger())

	api := router.Group("/api/v1")
	{
		authHandler := handlers.NewAuthHandler()
		apiRoutes := api.Group("/auth")
		{
			apiRoutes.POST("/login", authHandler.Login)
			apiRoutes.POST("/signup", authHandler.Signup)
			apiRoutes.POST("/logout", authHandler.Logout)
			//apiRoutes.POST("/refresh", authHandler.RefreshToken)
			//apiRoutes.POST("/resend-verification", authHandler.ResendVerificationEmail)
		}

		blogHandler := handlers.NewBlogHandler()
		blogRoutes := api.Group("/blog")
		{
			blogRoutes.GET("/", blogHandler.GetAll)
			blogRoutes.GET("/:id", middleware.RequireAuth(), middleware.VerifyBlogExists(), middleware.VerifyBlogOwnership(), blogHandler.GetByID)
			blogRoutes.POST("/", middleware.RequireAuth(), middleware.RequireRole("Creator", "Admin"), blogHandler.Create)
			blogRoutes.DELETE("/:id", middleware.RequireAuth(), middleware.VerifyBlogExists(), middleware.VerifyBlogOwnership(), blogHandler.Delete)
		}

		commentHandler := handlers.NewCommentHandler()
		commentRoutes := api.Group("/comments")
		{
			commentRoutes.GET("/:blogId", commentHandler.GetByBlogID)
			commentRoutes.POST("/:blogId", middleware.RequireAuth(), middleware.RequireRole("Commentor", "Creator", "Admin"), commentHandler.Create)
			commentRoutes.PUT("/:blogId/:id", middleware.RequireAuth(), middleware.RequireRole("Creator", "Admin"), commentHandler.Update)
			commentRoutes.DELETE("/:blogId/:id", middleware.RequireAuth(), middleware.RequireRole("Creator", "Admin"), commentHandler.Delete)
		}

		userHandler := handlers.NewUserHandler()
		userRoutes := api.Group("/users")
		{
			userRoutes.GET("/", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.GetAll)
			userRoutes.GET("/:id", userHandler.GetByID)
			userRoutes.POST("/", userHandler.Create)
			userRoutes.POST("/verify-email/", userHandler.VerifyEmail)
			userRoutes.PUT("/:id", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.Update)
			userRoutes.DELETE("/:id", middleware.RequireAuth(), middleware.RequireRole("Admin"), userHandler.Delete)
		}

		placeHandler := handlers.NewPlaceHandler()
		placeRoutes := router.Group("/api/v1/places")
		{
			placeRoutes.GET("/", placeHandler.GetPlaces)
			placeRoutes.POST("/", placeHandler.SavePlace)
			placeRoutes.PUT("/:id", placeHandler.UpdatePlace)
			placeRoutes.DELETE("/:id", placeHandler.DeletePlace)
		}

		fileHandler := handlers.NewFileHandler()
		fileRoutes := router.Group("/api/v1/files")
		{
			fileRoutes.GET("/structure", fileHandler.GetStructure)
			fileRoutes.GET("/", fileHandler.GetYearMakes)
			fileRoutes.GET("/:yearMake", fileHandler.GetModels)
			fileRoutes.GET("/:yearMake/:model", fileHandler.GetFiles)
			fileRoutes.GET("/:yearMake/:model/download/:fileName", fileHandler.DownloadFile)
			fileRoutes.GET("/:yearMake/:model/download/:fileName/:parentDir", fileHandler.DownloadFile)
			fileRoutes.GET("/:yearMake/:model/download-directory/:dirName", fileHandler.DownloadDirectory)
			fileRoutes.GET("/:yearMake/:model/download-all", fileHandler.DownloadAll)
			fileRoutes.POST("/:yearMake/:model/download-selected", fileHandler.DownloadSelected)
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}
