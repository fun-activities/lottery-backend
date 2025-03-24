package entity

const (
	Page     uint32 = 1
	PageSize uint32 = 10
)

type PaginationQuery struct {
	Page     uint32 `form:"page"`
	PageSize uint32 `form:"page_size"`
}

type SearchQuery struct {
	Search string `form:"search"`
}

type PrizeQuery struct {
	PaginationQuery
	SearchQuery
	PrizeId string `form:"prize_id"`
	Status  uint32 `form:"status"`
}

type PrizeUpdateStatus struct {
	PrizeId string `form:"prize_id"`
	Status  uint32 `form:"status"`
}
