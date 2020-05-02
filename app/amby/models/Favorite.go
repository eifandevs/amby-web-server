package models

import (
    "github.com/eifandevs/amby/repo"
    "github.com/thoas/go-funk"
)

type FavoriteInfo struct {
    ID int `json:"id"`
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    ID int `gorm:"type:int unsigned;not null;primary_key;auto_increment:false"`
    Token string `gorm:"type:varchar(255);not null;primary_key"`
    Title string
    Url string
}

type GetFavoriteResponse struct {
    BaseResponse
    Items  []FavoriteInfo `json:"data"`
}

type PostFavoriteRequest struct {
    Items  []FavoriteInfo `json:"data"`
}

type DeleteFavoriteRequest struct {
    Items  []FavoriteInfo `json:"data"`
}

func GetFavorite(accessToken string) GetFavoriteResponse {
    db := repo.Connect("development")
    defer db.Close()

    favorites := []Favorite{}
    if err := db.Where("token = ?", accessToken).Find(&favorites).Error; err != nil {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
    }

    items := funk.Map(favorites, func(favorite Favorite) FavoriteInfo {
        return FavoriteInfo{ID: favorite.ID, Title: favorite.Title, Url: favorite.Url}
    })
    
    if castedItems, ok := items.([]FavoriteInfo); ok {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: castedItems}
    } else {
        panic("cannot cast favorite item.")
    }
}

func PostFavorite(accessToken string, request PostFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        if err := db.Create(&Favorite{ID: item.ID, Token: accessToken, Title: item.Title, Url: item.Url}).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}

func DeleteFavorite(accessToken string, request DeleteFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        deletingRecord := Favorite{}
        deletingRecord.ID = item.ID
        deletingRecord.Token = accessToken
        db.First(&deletingRecord)
        if err := db.Unscoped().Delete(&deletingRecord).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}