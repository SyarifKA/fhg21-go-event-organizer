package lib

type Response struct {
	Success bool        `json:"succes"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}