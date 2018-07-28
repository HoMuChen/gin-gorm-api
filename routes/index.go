package routes

import "net/http"
import "github.com/gin-gonic/gin"

func IndexRoute(c *gin.Context) {
  c.String(http.StatusOK, "Hello %s", "World")
}
