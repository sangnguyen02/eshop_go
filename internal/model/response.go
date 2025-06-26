package model

type ErrorResponse struct {
	Message string `json:"message" example:"invalid request"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"valid request"`
}
