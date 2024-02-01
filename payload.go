package gg

type GetAllResponse[T any] struct {
	FirstIndex int   `json:"firstIndex"`
	Total      int64 `json:"total"`
	Data       []T   `json:"data"`
}

type GetResponse[T any] struct {
	Data T `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
