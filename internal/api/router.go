package api

import (
	stdhttp "net/http"

	"github.com/gin-gonic/gin"
	todosvc "github.com/suirfan20/go-gin-todo/internal/todo"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	// healthz
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(stdhttp.StatusOK, gin.H{"status": "ok"})
	})

	// TODO endpoints
	svc := todosvc.NewService()
	h := NewHandler(svc)

	v1 := r.Group("/v1")
	{
		v1.GET("/todos", h.List)
		v1.POST("/todos", h.Create)
		v1.DELETE("/todos/:id", h.Delete)
	}

	return r
}
