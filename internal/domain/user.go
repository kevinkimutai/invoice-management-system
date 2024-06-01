package domain

import "time"

type User struct {
	UserID    int64
	Email     string
	CreatedAt time.Time
	Name      string
}

type UsersFetch struct {
	Page          uint   `json:"page"`
	NumberOfPages uint   `json:"number_of_pages"`
	Total         uint   `json:"total"`
	Data          []User `json:"data"`
}

type UsersResponse struct {
	StatusCode    uint   `json:"status_code"`
	Message       string `json:"message"`
	Page          uint   `json:"page"`
	NumberOfPages uint   `json:"number_of_pages"`
	Total         uint   `json:"total"`
	Data          []User `json:"data"`
}

type UserResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
	Data       User   `json:"data"`
}
