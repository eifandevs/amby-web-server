package models

import(
    "github.com/eifandevs/amby/repo"
    // "github.com/thoas/go-funk"
	"github.com/jinzhu/gorm"
)

type AccessTokenItem struct {
    AccessToken string `json:"access_token"`
}

type AccessToken struct {
    gorm.Model
    VendorToken string
    AccessToken string
    Expire string
}

type UserInfo struct {
    Mail string `json:"mail"`
    VendorToken string `json:"vendor_token"`
}

type GetAccessTokenRequest struct {
    Item  UserInfo `json:"data"`
}

type GetAccessTokenResponse struct {
    BaseResponse
    Item  AccessTokenItem `json:"data"`
}

func GetAccessToken() GetAccessTokenResponse {
    db := repo.Connect("development")
    defer db.Close()

    return GetAccessTokenResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Item: AccessTokenItem{AccessToken: "aaaaaaa"}}

    //　ユーザー情報登録
    // accessTokens := []AccessToken{}
    // if err := db.Find(&accessTokens).Error; err != nil {
    //     return GetAccessTokenResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
    // }

    // items := funk.Map(accessTokens, func(accessToken AccessToken) AccessTokenItem {
    //     return AccessTokenItem{Token: accessToken.Token, Expire: accessToken.Expire}
    // })
    
    // if castedItems, ok := items.([]AccessTokenItem); ok {
    //     return GetAccessTokenResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: castedItems}
    // } else {
    //     panic("cannot cast accessToken item.")
    // }
}