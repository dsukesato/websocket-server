package models

// クライアントからは JSON 形式で受け取る
type Message struct {
	Message string `json:message`
}
