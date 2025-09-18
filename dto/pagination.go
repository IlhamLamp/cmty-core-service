package dto

type Paginatable interface {
	GetPage() int
	GetLimit() int
	SetPage(int)
	SetLimit(int)
}
type PaginationMeta struct {
	TotalRecords int64 `json:"total_records"`
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	NextPage     *int  `json:"next_page,omitempty"`
	PrevPage     *int  `json:"prev_page,omitempty"`
}

func (f *CoreFilter) GetPage() int   { return f.Page }
func (f *CoreFilter) GetLimit() int  { return f.Limit }
func (f *CoreFilter) SetPage(p int)  { f.Page = p }
func (f *CoreFilter) SetLimit(l int) { f.Limit = l }

func (f *MemberFilter) GetPage() int   { return f.Page }
func (f *MemberFilter) GetLimit() int  { return f.Limit }
func (f *MemberFilter) SetPage(p int)  { f.Page = p }
func (f *MemberFilter) SetLimit(l int) { f.Limit = l }
