package pagination

import "math"

type QueryParams struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
}

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalData   int64 `json:"total_data"`
	TotalPages  int   `json:"total_pages"`
}

func (p *QueryParams) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p *QueryParams) GetLimit() int {
	if p.Limit <= 0 || p.Limit > 100 {
		return 10
	}
	return p.Limit
}

func (p *QueryParams) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func BuildMeta(total int64, page, limit int) Meta {
	return Meta{
		CurrentPage: page,
		PerPage:     limit,
		TotalData:   total,
		TotalPages:  int(math.Ceil(float64(total) / float64(limit))),
	}
}
