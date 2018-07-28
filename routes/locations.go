package routes

//import "fmt"
import "net/http"
import "github.com/gin-gonic/gin"

import "../models"
import "../db"

func LocationsRouteFind(c *gin.Context) {
  c.String(http.StatusOK, "Hello /locations")
}

func LocationsRouteGetByID(c *gin.Context) {
  id := c.Param("id")

  c.String(http.StatusOK, "Hello /locations/%s", id)
}

func LocationsRouteGetMediasByLocID(c *gin.Context) {
  id := c.Param("id")
  sort := c.Query("sort")
  n := c.DefaultQuery("n", "20")

  var medias []models.Media
  var db = db.GetDB()

  db.Table("media").Where("location_id = ?", id).Order(sort).Limit(n).Find(&medias)

  c.JSON(http.StatusOK, medias)
}
