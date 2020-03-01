package models

import (
    "log"
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

func GetFavorite() GetFavoriteResponse {
    items := []FavoriteItem{FavoriteItem{Title: "1", Url: "1"}, FavoriteItem{Title: "2", Url: "2"}}
    
	return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: items}
}

func PostFavorite(request PostFavoriteRequest) BaseResponse {
    log.Println("post favorite: ", request.Items)

    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        db.Create(&Favorite{Token: "1111", Title: item.Title, Url: item.Url})
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}

func DeleteFavorite(request DeleteFavoriteRequest) BaseResponse {
    log.Println("delete favorite: ", request.Items)

    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        db.Unscoped().Delete(&Favorite{Token: "1111", Title: item.Title, Url: item.Url})
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}