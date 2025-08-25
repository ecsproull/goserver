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
		u := user.(*models.DbUser)
		com := comment.(*models.DbComment)

		isOwner := u.Username == com.Name || u.Email == com.Email
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
		u, okUser := user.(*models.DbUser)
		b, okBlog := blog.(*models.DbBlog)

		if !okUser || !okBlog {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user or blog information"})
			c.Abort()
			return
		}

		isOwner := u.Username == b.AuthorID
		isAdmin := u.Role == "Admin"

		if !isOwner && !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. You can only modify your own blog posts or must be an Admin."})
			c.Abort()
			return
		}
		c.Next()
	}
}
