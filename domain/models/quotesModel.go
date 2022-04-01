package models

type Quote struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
