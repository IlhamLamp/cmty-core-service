package dto

type PaginationMeta struct {
	TotalRecords int64 `json:"total_records"`
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	NextPage     *int  `json:"next_page,omitempty"`
	PrevPage     *int  `json:"prev_page,omitempty"`
}
