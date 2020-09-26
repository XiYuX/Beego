package models

type User struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
