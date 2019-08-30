package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

type Favorite struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

func GetFavorite() echo.HandlerFunc {
  return func(c echo.Context) error {
    favorites := models.GetFavorite()
    return c.JSON(http.StatusOK, favorites)
  }
}
