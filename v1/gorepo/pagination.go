package gorepo

type PaginationResponse[t any] struct {
	Data           []*t  `json:"data"`
	PagesRemaining int64 `json:"pages_remaining"`
	CurrentPage    int64 `json:"current_page"`
	CurrentSize    int64 `json:"current_size"`
}

type PaginationInput struct {
	Page int
	Size int
}

func PAG(page int, size int) *PaginationInput {
	return &PaginationInput{
		Page: page,
		Size: size,
	}
}
