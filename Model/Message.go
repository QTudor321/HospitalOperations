package model
type Message struct{
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Title string `json:"title"`
	Content string `json:"content"`
}