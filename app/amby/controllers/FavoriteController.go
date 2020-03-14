package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetHandler() echo.HandlerFunc {
  return func(c echo.Context) error {
    userToken := c.Request().Header.Get("User-Token")
    if userToken == "" {
      return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil})
    }

    favorites := models.GetFavorite(userToken)

    return c.JSON(http.StatusOK, favorites)
  }
}

func PostHandler() echo.HandlerFunc {
  return func(c echo.Context) error {

    post := new(models.PostFavoriteRequest)
    if err := c.Bind(post); err != nil {
        return err
    }

    userToken := c.Request().Header.Get("User-Token")
    if userToken == "" {
      return c.JSON(http.StatusOK, models.BaseResponse{Result: "NG", ErrorCode: ""})
    }

    response := models.PostFavorite(userToken, *post)
    return c.JSON(http.StatusOK, response)
  }
}

func DeleteHandler() echo.HandlerFunc {
  return func(c echo.Context) error {

    delete := new(models.DeleteFavoriteRequest)
    if err := c.Bind(delete); err != nil {
        return err
    }

    response := models.DeleteFavorite(*delete)
    return c.JSON(http.StatusOK, response)
  }
}