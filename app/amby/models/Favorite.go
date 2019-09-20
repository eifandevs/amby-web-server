package models

type FavoriteItem struct {
    Title string `json:"title"`
    Url string `json:"url"`
}

type FavoriteResponse struct {
    Item  []FavoriteItem `json:"items"`
}

func GetFavorite() FavoriteResponse {
    items := []FavoriteItem{FavoriteItem{Title: "1", Url: "1"}, FavoriteItem{Title: "2", Url: "2"}}
    
	return FavoriteResponse{Item: items}
}