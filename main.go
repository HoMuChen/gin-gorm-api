package main

import "fmt"
import "github.com/gin-gonic/gin"

import "./db"
import "./configs"
import "./routes"

func main() {
  db.Init()
  defer db.Close()

  router := gin.Default()

  router.GET("/", routes.IndexRoute)

  media := router.Group("/media")
  {
    media.GET("/",     routes.MediaRouteFind)
    media.GET("/:id",  routes.MediaRouteGetByID)
  }

  locations := router.Group("/locations")
  {
    locations.GET("/",            routes.LocationsRouteFind)
    locations.GET("/:id",         routes.LocationsRouteGetByID)
    locations.GET("/:id/medias",  routes.LocationsRouteGetMediasByLocID)
  }

  router.Run( fmt.Sprintf(":%s", configs.API_PORT) )
}
