package models

import (
	"github.com/eifandevs/amby/repo"
	"github.com/thoas/go-funk"
)

type MemoItem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Memo struct {
	ID      int    `gorm:"type:int unsigned;not null;primary_key;auto_increment:false"`
	Token   string `gorm:"type:varchar(255);not null;primary_key"`
	Title   string
	Content string `gorm:"type:varchar(1000)"`
}

type GetMemoResponse struct {
	BaseResponse
	Items []MemoItem `json:"data"`
}

type PostMemoRequest struct {
	Items []MemoItem `json:"data"`
}

type DeleteMemoRequest struct {
	Items []MemoItem `json:"data"`
}

func GetMemo(accessToken string) GetMemoResponse {
	db := repo.Connect("development")
	defer db.Close()

	memos := []Memo{}
	if err := db.Where("token = ?", accessToken).Find(&memos).Error; err != nil {
		return GetMemoResponse{BaseResponse: BaseResponse{Result: "NG", ErrorCode: ""}, Items: nil}
	}

	items := funk.Map(memos, func(memo Memo) MemoItem {
		return MemoItem{ID: memo.ID, Title: memo.Title, Content: memo.Content}
	})

	if castedItems, ok := items.([]MemoItem); ok {
		return GetMemoResponse{BaseResponse: BaseResponse{Result: "OK", ErrorCode: ""}, Items: castedItems}
	} else {
		panic("cannot cast memo item.")
	}
}

func PostMemo(accessToken string, request PostMemoRequest) BaseResponse {
	db := repo.Connect("development")
	defer db.Close()

	for _, item := range request.Items {
		if err := db.Create(&Memo{ID: item.ID, Token: accessToken, Title: item.Title, Content: item.Content}).Error; err != nil {
			return BaseResponse{Result: "NG", ErrorCode: ""}
		}
	}

	return BaseResponse{Result: "OK", ErrorCode: ""}
}

func DeleteMemo(accessToken string, request DeleteMemoRequest) BaseResponse {
	db := repo.Connect("development")
	defer db.Close()

	for _, item := range request.Items {
		deletingRecord := Memo{}
		deletingRecord.ID = item.ID
		deletingRecord.Token = accessToken
		db.First(&deletingRecord)
		if err := db.Unscoped().Delete(&deletingRecord).Error; err != nil {
			return BaseResponse{Result: "NG", ErrorCode: ""}
		}
	}

	return BaseResponse{Result: "OK", ErrorCode: ""}
}
