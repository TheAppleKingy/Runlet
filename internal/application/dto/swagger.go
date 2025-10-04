package dto

var NoBodyResponse any = nil

// Error is response body for 4xx statuses
type ErrorBody struct {
	Error string `json:"error" example:"error msg"`
}

// OkBody response body for 2xx statuses for undefined body values
type OkBody struct {
	Detail string `json:"detail" example:"detail msg"`
}
