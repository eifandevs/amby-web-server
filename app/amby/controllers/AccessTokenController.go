package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetAccessTokenHandler() echo.HandlerFunc {
  return func(c echo.Context) error {
    post := new(models.GetAccessTokenRequest)
		if err := c.Bind(post); err != nil {
      return err
    }
    
    // ユーザー情報の登録
    if err := models.CreateUser((*post).Item); err != nil {
      return err
    }

    // TODO: アクセストークンの発行
    accesstoken := models.GetAccessToken()
    return c.JSON(http.StatusOK, accesstoken)
  }
}
