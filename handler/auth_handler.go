package handler

import (
	"net/http"
	"strings"

	"github.com/dickysetiawan031000/go-backend/dto/auth"
	"github.com/dickysetiawan031000/go-backend/mapper"
	"github.com/dickysetiawan031000/go-backend/model"
	"github.com/dickysetiawan031000/go-backend/usecase"
	"github.com/dickysetiawan031000/go-backend/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Usecase usecase.AuthUseCase
}

func NewAuthHandler(r *gin.RouterGroup, usecase usecase.AuthUseCase) *AuthHandler {
	handler := &AuthHandler{Usecase: usecase}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	// NOTE: logout JANGAN di sini jika kamu pakai middleware di group terproteksi

	return handler
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.Usecase.Register(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseWithMessage{
		Message: "register success",
		Data:    mapper.ToUserResponse(createdUser),
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loggedInUser, err := h.Usecase.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(uint(loggedInUser.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "login success",
		Data: gin.H{
			"user":  mapper.ToUserResponse(loggedInUser),
			"token": token,
		},
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// validasi dulu tokennya
	_, err := utils.VerifyJWT(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// hapus dari whitelist
	utils.RemoveToken(tokenStr)

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "logout success",
	})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDVal.(uint)
	user, err := h.Usecase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "success",
		Data:    mapper.ToUserResponse(user),
	})
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var input auth.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := userIDVal.(uint)
	updatedUser, err := h.Usecase.UpdateProfile(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "profile updated successfully",
		Data:    mapper.ToUserResponse(updatedUser),
	})
}
