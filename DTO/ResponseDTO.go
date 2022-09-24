package dto

type Response struct {
	Status  int         `json:"status"`
	Message []string    `json:"message"`
	Entity  interface{} `json:"entity"`
}
