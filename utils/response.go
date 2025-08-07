package utils

type ResponseWithMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
