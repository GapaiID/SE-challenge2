package dto

type Pagination struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"page_size"`
}

type PaginationParams struct {
	Current  int `json:"current" query:"current"`
	PageSize int `json:"page_size" query:"page_size"`
}

func (p *PaginationParams) SetDefaultPageSize(pageSize int) {
	if p.PageSize == 0 {
		p.PageSize = pageSize
	}
}
