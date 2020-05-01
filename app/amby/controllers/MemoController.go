package controllers

import (
	"net/http"

	"github.com/eifandevs/amby/models"
	"github.com/labstack/echo"
)

func GetMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.GetMemoResponse{BaseResponse: models.BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil})
		}

		memos := models.GetMemo(accessToken)

		return c.JSON(http.StatusOK, memos)
	}
}

func PostMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		post := new(models.PostMemoRequest)
		if err := c.Bind(post); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Result: "NG", ErrorCode: ""})
		}

		response := models.PostMemo(accessToken, *post)
		return c.JSON(http.StatusOK, response)
	}
}

func DeleteMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		delete := new(models.DeleteMemoRequest)
		if err := c.Bind(delete); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("User-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Result: "NG", ErrorCode: ""})
		}

		response := models.DeleteMemo(accessToken, *delete)
		return c.JSON(http.StatusOK, response)
	}
}