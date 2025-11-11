package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suirfan20/go-gin-todo/internal/todo"
)

type Handler struct {
	svc *todo.Service
}

func NewHandler(svc *todo.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) List(c *gin.Context) {
	items := h.svc.List()
	c.JSON(http.StatusOK, items)
}

type createReq struct {
	Title string `json:"title" binding:"required,min=1,max=140"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := h.svc.Create(req.Title)
	c.JSON(http.StatusCreated, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	ok := h.svc.Delete(id)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
