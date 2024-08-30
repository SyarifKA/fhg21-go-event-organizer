package dtos

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name" form:"name"`
}
