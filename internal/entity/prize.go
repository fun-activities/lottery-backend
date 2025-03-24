package entity

type PrizeType int8

// '实物', '虚拟', '优惠券'
const (
	PrizeGoodsType   PrizeType = 1
	PrizeVirtualType PrizeType = 2
	PrizeCouponType  PrizeType = 3
)

type Status uint8

const (
	OnlineStatus  Status = 1
	OfflineStatus Status = 2
	TestStatus    Status = 3
)

type Prize struct {
	Id         string    `json:"id"   db:"id"`
	PrizeId    string    `json:"prize_id"   db:"prize_id"`
	Name       string    `json:"name"   db:"name"`
	Type       PrizeType `json:"type"   db:"type"`
	Quantity   int32     `json:"quantity"   db:"quantity"`
	ImageUrl   string    `json:"image_url"   db:"image_url"`
	Price      float32   `json:"price"   db:"price"`
	TotalStock int32     `json:"total_stock"   db:"total_stock"`
	Status     Status    `json:"status"   db:"status"`
	Ext        string    `json:"ext"   db:"ext"`
}
