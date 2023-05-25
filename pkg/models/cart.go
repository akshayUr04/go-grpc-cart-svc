package models

type Cart struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	ProductId int64
	Total     int64
	UserId    int64
	Qty       int64
}
