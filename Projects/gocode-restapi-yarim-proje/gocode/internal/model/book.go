package model

import (
	"net/http"
	"time"
)

type (
	Book struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Route struct {
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	Pagination struct {
		Limit        int `json:"limit"`
		Offset       int `json:"offset"`
		TotalRecords int `json:"totalRecords"`
	}

	PaginatedResponse struct {
		Data       interface{} `json:"data,omitempty"`
		Pagination *Pagination `json:"pagination,omitempty"`
	}

	Item struct {
		Value interface{} `json:"value"`
	}

	ResponseBody struct {
		Data *[]Item `json:"data,omitempty"`
	}

	ResponseSingleBody struct {
		Data *Item `json:"data,omitempty"`
	}
)
