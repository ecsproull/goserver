package handlers

import (
	"net/http"
	"os"
	"time"

	"goserver/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginData struct {
		UserName     string `json:"user_name"`
		UserPassword string `json:"user_password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	user, validationErrors, _ := services.LoginUser(loginData.UserName, loginData.UserPassword)
	if len(validationErrors) > 0 || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	userRole := user.Role
	claims := jwt.MapClaims{
		"user_name": user.UserName,
		"user":      user.ID,
		"role":      userRole,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	accessToken, err := token.SignedString([]byte(secret))
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
	// TODO: Invalidate token (if using token blacklist or similar)
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
