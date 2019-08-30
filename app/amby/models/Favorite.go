package models

type Favorite struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func GetFavorite() Favorite {
	return Favorite{Id: 1, Name: "favorite"}
}