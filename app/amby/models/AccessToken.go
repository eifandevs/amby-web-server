package models

// import(
//     "github.com/eifandevs/amby/repo"
// )

type AccessToken struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func GetAccessToken() AccessToken {
	return AccessToken{Id: 1, Name: "test"}
}