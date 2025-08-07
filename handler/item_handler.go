package handler

import (
	"net/http"
	"strconv"

	"github.com/dickysetiawan031000/go-backend/dto/item"
	"github.com/dickysetiawan031000/go-backend/mapper"
	"github.com/dickysetiawan031000/go-backend/usecase"
	"github.com/dickysetiawan031000/go-backend/utils"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	Usecase usecase.ItemUseCase
}

func NewItemHandler(r *gin.RouterGroup, uc usecase.ItemUseCase) {
	handler := &ItemHandler{Usecase: uc}

	r.POST("/items", handler.Create)
	r.GET("/items", handler.GetAll)
	r.GET("/items/:id", handler.GetByID)
	r.PUT("/items/:id", handler.Update)
	r.DELETE("/items/:id", handler.Delete)
}

func (h *ItemHandler) Create(c *gin.Context) {
	var input item.CreateItemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := h.Usecase.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, utils.ResponseWithMessage{
		Message: "item created successfully",
		Data:    mapper.ToItemResponse(newItem),
	})
}

func (h *ItemHandler) GetAll(c *gin.Context) {
	items, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "success",
		Data:    mapper.ToItemResponses(items),
	})
}

func (h *ItemHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	item, err := h.Usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "success",
		Data:    mapper.ToItemResponse(item),
	})
}

func (h *ItemHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var input item.UpdateItemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.Usecase.Update(uint(id), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "item updated successfully",
		Data:    mapper.ToItemResponse(updated),
	})
}

func (h *ItemHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if err := h.Usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWithMessage{
		Message: "item deleted successfully",
	})
}
