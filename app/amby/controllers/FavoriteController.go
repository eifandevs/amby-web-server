package controllers

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/eifandevs/amby/models"
)

func GetHandler() echo.HandlerFunc {
  return func(c echo.Context) error {
    favorites := models.GetFavorite()
    return c.JSON(http.StatusOK, favorites)
  }
}

func PostHandler() echo.HandlerFunc {
  return func(c echo.Context) error {

    post := new(models.Favorite)
    if err := c.Bind(post); err != nil {
        return err
    }

    response := models.PostFavorite(*post)
    return c.JSON(http.StatusOK, response)
  }
}
