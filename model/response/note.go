package model_response

type Note struct {
	IdNote       int    `json:"id_note"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Date_created string `json:"date_created"`
	Date_updated string `json:"date_updated"`
}
