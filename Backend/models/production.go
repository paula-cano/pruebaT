package models

type Production struct {
	Items    []Item `json:"items"`
	NextPage string `json:"next_page"`
}