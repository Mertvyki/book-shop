package core_http_response

type PaginatedResponse struct {
	Data      any    `json:"data"`
	Total     int    `json:"total"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	TotalPages int   `json:"total_pages"`
}

func NewPaginatedResponse(data any, total, page, limit int) PaginatedResponse {
	totalPages := total / limit
	if total%limit != 0 {
		totalPages++
	}

	return PaginatedResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}
