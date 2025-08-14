package handlers

import (
	"goserver/internal/models"
	"goserver/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentHandler struct{}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{}
}

func (h *CommentHandler) GetByBlogID(c *gin.Context) {
	blogID := c.Param("blogId")
	comments, err := services.GetCommentsByBlogID(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Create(c *gin.Context) {
	blogID := c.Param("blogId")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the BlogID field (assuming it's an ObjectID in your model)
	objID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	comment.BlogID = objID

	id, err := services.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"id":      id.Hex(),
		"blogId":  blogID,
	})
}

func (h *CommentHandler) Update(c *gin.Context) {
	blogID := c.Param("blogId")
	id := c.Param("id")

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateComment(blogID, id, updateData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
		"blogId":  blogID,
		"id":      id,
	})
}

func (h *CommentHandler) Delete(c *gin.Context) {
	blogID := c.Param("blogId")
	id := c.Param("id")

	err := services.DeleteComment(blogID, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
		"blogId":  blogID,
		"id":      id,
	})
}
