package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetAccessToken() echo.HandlerFunc {
  return func(c echo.Context) error {
    // TODO: ユーザー情報の登録、アクセストークンの発行
    accesstoken := models.GetAccessToken()
    return c.JSON(http.StatusOK, accesstoken)
  }
}
