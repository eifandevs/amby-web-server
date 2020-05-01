package controllers

import (
	"net/http"

	"github.com/eifandevs/amby/models"
	"github.com/labstack/echo"
)

func GetFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil})
		}

		favorites := models.GetFavorite(accessToken)

		return c.JSON(http.StatusOK, favorites)
	}
}

func PostFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		post := new(models.PostFavoriteRequest)
		if err := c.Bind(post); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Result: "NG", ErrorCode: ""})
		}

		response := models.PostFavorite(accessToken, *post)
		return c.JSON(http.StatusOK, response)
	}
}

func DeleteFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		delete := new(models.DeleteFavoriteRequest)
		if err := c.Bind(delete); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Result: "NG", ErrorCode: ""})
		}

		response := models.DeleteFavorite(accessToken, *delete)
		return c.JSON(http.StatusOK, response)
	}
}
