package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetAccessToken() echo.HandlerFunc {
  return func(c echo.Context) error {
    accesstoken := models.GetAccessToken()
    return c.JSON(http.StatusOK, accesstoken)
  }
}
