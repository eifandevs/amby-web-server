package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func LoginHandler() echo.HandlerFunc {
  return func(c echo.Context) error {
    post := new(models.PostUserRequest)
		if err := c.Bind(post); err != nil {
      return err
    }
    
    // ユーザー情報の登録
    user, err := models.CreateUser((*post).Item)
    if err != nil {
      return err
    }

    return c.JSON(http.StatusOK, user)
  }
}
