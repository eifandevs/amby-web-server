package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetFavorite() echo.HandlerFunc {
  return func(c echo.Context) error {
    favorites := models.GetFavorite()
    return c.JSON(http.StatusOK, favorites)
  }
}

func PostFavorite() echo.HandlerFunc {
  return func(c echo.Context) error {
    favorites := models.PostFavorite()
    return c.JSON(http.StatusOK, favorites)
  }
}
