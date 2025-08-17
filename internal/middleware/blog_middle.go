package middleware

import (
	"goserver/internal/models"
	"goserver/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyBlogExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("blogId")
		if id == "" {
			id = c.Param("id")
		}
		blog, err := services.GetBlogByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if blog == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
			c.Abort()
			return
		}
		c.Set("blog", blog)
		c.Next()
	}
}

func VerifyCommentExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("commentId")
		comment, err := services.GetBlogCommentByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if comment == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			c.Abort()
			return
		}
		c.Set("comment", comment)
		c.Next()
	}
}

func VerifyCommentOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user") // Set by your auth middleware
		comment, _ := c.Get("comment")
		u := user.(*models.User)
		com := comment.(*models.Comment)

		isOwner := u.Name == com.CommenterName || u.Email == com.CommenterEmail
		isAdmin := u.Role == "Admin"

		if !isOwner && !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. You can only modify your own comments or must be an admin."})
			c.Abort()
			return
		}
		c.Next()
	}
}

func VerifyBlogOwnership() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		blog, _ := c.Get("blog")
		u := user.(*models.User)
		b := blog.(*models.Blog)

		isOwner := u.Name == b.OwnerName
		isAdmin := u.Role == "Admin"

		if !isOwner && !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. You can only modify your own blog posts or must be an Admin."})
			c.Abort()
			return
		}
		c.Next()
	}
}
