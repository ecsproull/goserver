package handlers

import (
	"goserver/internal/models"
	"goserver/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct{}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}

func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := services.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

func (h *BlogHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	blog, err := services.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog models.DbBlog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := services.SaveBlog(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"id":      id,
	})
}

func (h *BlogHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
		"id":      id,
	})
}
