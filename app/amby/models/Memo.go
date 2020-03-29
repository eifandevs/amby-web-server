package models

type MemoItem struct {
    Title string `json:"title"`
    Content string `json:"content"`
}

type Memo struct {
    ID int `gorm:"type:int unsigned;not null;primary_key;auto_increment:false"`
    Token string `gorm:"type:varchar(255);not null;primary_key"`
    Title string
    Content string `gorm:"type:varchar(1000)"`
}

type MemoResponse struct {
    Item  []MemoItem `json:"items"`
}

func GetMemo() MemoResponse {
    items := []MemoItem{MemoItem{Title: "1", Content: "1"}, MemoItem{Title: "2", Content: "2"}}
    
	return MemoResponse{Item: items}
}