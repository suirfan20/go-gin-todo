package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/suirfan20/go-gin-todo/internal/todo"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	svc := todo.NewService()
	h := NewHandler(svc)
	r := gin.New()
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET("/v1/todos", h.List)
	r.POST("/v1/todos", h.Create)
	return r
}

func TestHealthz(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
}

func TestCreateTodo(t *testing.T) {
	r := setupRouter()
	body := `{"title":"learn ci/cd"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/todos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201, got %d; body=%s", w.Code, w.Body.String())
	}
}
