package models

type Locations struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	Lat   string `json:"lat"`
	Long  string `json:"long"`
}
