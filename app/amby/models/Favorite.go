package models

import (
    "log"
    "github.com/jinzhu/gorm"
)

type FavoriteItem struct {
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    gorm.Model
    Token string 
    Items  []FavoriteItem
}

type GetFavoriteResponse struct {
    BaseResponse
    Items  []FavoriteItem `json:"data"`
}

type PostFavoriteRequest struct {
    Items  []FavoriteItem `json:"data"`
}

func GetFavorite() GetFavoriteResponse {
    items := []FavoriteItem{FavoriteItem{Title: "1", Url: "1"}, FavoriteItem{Title: "2", Url: "2"}}
    
	return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: items}
}

func PostFavorite(request PostFavoriteRequest) BaseResponse {
    log.Println("post favorite: ", request.Items)
	return BaseResponse{Result: "OK", ErrorCode: ""}
}