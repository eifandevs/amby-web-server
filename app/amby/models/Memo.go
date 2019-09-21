package models

type MemoItem struct {
    Title string `json:"title"`
    content string `json:"content"`
}

type MemoResponse struct {
    Item  []MemoItem `json:"items"`
}

func GetMemo() MemoResponse {
    items := []MemoItem{MemoItem{Title: "1", Content: "1"}, MemoItem{Title: "2", Content: "2"}}
    
	return MemoResponse{Item: items}
}