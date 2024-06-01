package domain

import (
	"time"
)

type Company struct {
	CompanyID   int64     `json:"company_id"`
	LogoUrl     string    `json:"logo_url"`
	UserID      int64     `json:"user_id"`
	CompanyName string    `json:"company_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type CompaniesFetch struct {
	Page          uint      `json:"page"`
	NumberOfPages uint      `json:"number_of_pages"`
	Total         uint      `json:"total"`
	Data          []Company `json:"data"`
}

type CompaniesResponse struct {
	StatusCode    uint      `json:"status_code"`
	Message       string    `json:"message"`
	Page          uint      `json:"page"`
	NumberOfPages uint      `json:"number_of_pages"`
	Total         uint      `json:"total"`
	Data          []Company `json:"data"`
}

type CompanyResponse struct {
	StatusCode uint    `json:"status_code"`
	Message    string  `json:"message"`
	Data       Company `json:"data"`
}
