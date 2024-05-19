package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proj-go-5/accounts/pkg/authorization"
)

type AdminMiddlewareAdapter struct {
	authService *authorization.Service
}

func NewAdminMiddlewareAdapter(authService *authorization.Service) *AdminMiddlewareAdapter {
	return &AdminMiddlewareAdapter{authService: authService}
}

func (a *AdminMiddlewareAdapter) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := a.authService.AdminMiddleware(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Writer = w.(gin.ResponseWriter)
			c.Next()
		})

		w := &responseRecorder{ResponseWriter: c.Writer, Context: c}
		r := c.Request
		handler(w, r)

		if w.statusCode == http.StatusUnauthorized || w.statusCode == http.StatusBadRequest {
			c.AbortWithStatus(w.statusCode)
		} else {
			c.Next()
		}
	}
}

type responseRecorder struct {
	gin.ResponseWriter
	Context    *gin.Context
	statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
