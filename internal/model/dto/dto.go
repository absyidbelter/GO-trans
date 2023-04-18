package dto

import "time"

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ApiloginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ApiUserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewApiResponseSuccess(msg string, data any) ApiResponse {
	return ApiResponse{
		Status:  "Success",
		Message: msg,
		Data:    data,
	}
}

func NewApiResponseFailed(msg string) ApiResponse {
	return ApiResponse{
		Status:  "Failed",
		Message: msg,
	}
}
