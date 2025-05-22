package web

import "github.com/freightcms/carriers/models"

type GetCarriersRequest struct {
	// Page should be the page number starting at 0 that is being requested for data
	Page int `json:"page" xml:"page" query:"page"`
	// Limit should the max number of records that can be fetched in the request
	Limit int `json:"limit" xml:"limit" query:"limit"`
}

type GetCarriersResponse struct {
	Total    int64             `json:"total" xml:"total"`
	Carriers []*models.Carrier `json:"carriers" xml:"carriers"`
}
