package routes

//import "fmt"
import "net/http"
import "github.com/gin-gonic/gin"

import "../models"
import "../db"

func MediaRouteFind(c *gin.Context) {
  location_id := c.Query("location_id")
  n := c.Query("n")

  var medias []models.Media
  var db = db.GetDB()

  db.Table("media").Where("location_id = ?", location_id).Limit(n).Find(&medias)

  c.JSON(http.StatusOK, medias)
}

func MediaRouteGetByID(c *gin.Context) {
  id := c.Param("id")

  var media models.Media
  var db = db.GetDB()

  db.Table("media").Where("id = ?", id).First(&media)

  c.JSON(http.StatusOK, media)
}
