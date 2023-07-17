package models

type Pasar struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Address string `json:"address"`
}
