package models

type JsonResponse struct {
	Status  interface{} `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
