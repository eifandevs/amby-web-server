package models

import (
    "github.com/eifandevs/amby/repo"
    "github.com/thoas/go-funk"
    "github.com/jinzhu/gorm"
)

type FavoriteInfo struct {
    FID int `json:"fid"`
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    gorm.Model
    FID int `gorm:"type:int unsigned;not null;unique;primary_key;auto_increment:false"`
    UserID uint
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

func GetFavorite(userID uint) GetFavoriteResponse {
    db := repo.Connect("development")
    defer db.Close()

    favorites := []Favorite{}
    if err := db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
    }

    items := funk.Map(favorites, func(favorite Favorite) FavoriteInfo {
        return FavoriteInfo{FID: favorite.FID, Title: favorite.Title, Url: favorite.Url}
    })
    
    if castedItems, ok := items.([]FavoriteInfo); ok {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: castedItems}
    } else {
        panic("cannot cast favorite item.")
    }
}

func PostFavorite(userID uint, request PostFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        if err := db.Create(&Favorite{FID: item.FID, UserID: userID, Title: item.Title, Url: item.Url}).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}

func DeleteFavorite(userID uint, request DeleteFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Items {
        deletingRecord := Favorite{}
        deletingRecord.FID = item.FID
        deletingRecord.UserID = userID
        db.First(&deletingRecord)
        if err := db.Unscoped().Delete(&deletingRecord).Error; err != nil {
            return BaseResponse{Result: "NG", ErrorCode: ""}
        }
    }

	return BaseResponse{Result: "OK", ErrorCode: ""}
}