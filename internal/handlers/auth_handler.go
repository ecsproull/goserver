package handlers

import (
	"net/http"

	"goserver/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginData struct {
		Name     string `json:"user_name"`
		Password string `json:"user_password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	user, validationErrors, _ := services.LoginUser(loginData.Name, loginData.Password)
	if len(validationErrors) > 0 || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	accessToken, err := services.GenerateAccessToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

func (h *AuthHandler) Signup(c *gin.Context) {
	// TODO: Register user, send verification email
	c.JSON(http.StatusOK, gin.H{"message": "Signup endpoint"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// Nothing to do here. Token gets deleted on client-side.
	c.JSON(http.StatusOK, gin.H{"message": "Logout endpoint"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// TODO: Issue new JWT if refresh token is valid
	c.JSON(http.StatusOK, gin.H{"message": "Refresh token endpoint"})
}

func (h *AuthHandler) ResendVerificationEmail(c *gin.Context) {
	// TODO: Resend verification email
	c.JSON(http.StatusOK, gin.H{"message": "Resend verification endpoint"})
}
