package router

import (
  "github.com/eifandevs/amby/controllers"
  "github.com/labstack/echo"
)

func Init() *echo.Echo {
	
  e := echo.New()
  
  api := e.Group("/api")
  {
    api.GET("/accesstoken", controllers.GetAccessToken())
    api.GET("/favorite", controllers.GetFavorite())
  }

  return e
}