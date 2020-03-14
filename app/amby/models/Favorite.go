package models

import (
    "github.com/jinzhu/gorm"
    "github.com/eifandevs/amby/repo"
)

type FavoriteItem struct {
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    gorm.Model
    Token string
    Title string
    Url string
}

type GetFavoriteResponse struct {
    BaseResponse
    Items  []FavoriteItem `json:"data"`
}

type PostFavoriteRequest struct {
    Items  []FavoriteItem `json:"data"`
}

type DeleteFavoriteRequest struct {
    Items  []FavoriteItem `json:"data"`
}

func GetFavorite(userToken string) GetFavoriteResponse {
    db := repo.Connect("development")
    defer db.Close()

    favorites := []Favorite{}
    if err := db.Find(&favorites).Error; err != nil {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
    }

    // TODO: []Favorite -> []FavoriteItem
	return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: favorites}
}

func PostFavorite(userToken string, request PostFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        if err := db.Create(&Favorite{Token: userToken, Title: item.Title, Url: item.Url}).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}

func DeleteFavorite(request DeleteFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        if err := db.Unscoped().Delete(&Favorite{Token: "1111", Title: item.Title, Url: item.Url}).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}