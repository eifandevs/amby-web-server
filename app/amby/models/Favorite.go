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
    Item  []FavoriteItem
}

type FavoriteResponse struct {
    Item  []FavoriteItem `json:"items"`
}

func GetFavorite() FavoriteResponse {
    items := []FavoriteItem{FavoriteItem{Title: "1", Url: "1"}, FavoriteItem{Title: "2", Url: "2"}}
    
	return FavoriteResponse{Item: items}
}

func PostFavorite(favorite Favorite) CommonResponse {
    log.Println("post favorite: ", favorite)
	return CommonResponse{Result: "OK", ErrorCode: ""}
}