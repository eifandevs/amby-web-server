package models

import (
    "log"
)

type FavoriteItem struct {
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    Item  []FavoriteItem `json:"items"`
}

func GetFavorite() Favorite {
    items := []FavoriteItem{FavoriteItem{Title: "1", Url: "1"}, FavoriteItem{Title: "2", Url: "2"}}
    
	return Favorite{Item: items}
}

func PostFavorite(favorite Favorite) Favorite {
    log.Println("post favorite: ", favorite)
    items := []FavoriteItem{FavoriteItem{Title: "3", Url: "3"}, FavoriteItem{Title: "4", Url: "4"}}
    
	return Favorite{Item: items}
}