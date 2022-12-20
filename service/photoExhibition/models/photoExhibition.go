package models

import "time"

type PhotoExhibition struct {
	Id        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Title     string     `json:"title"`
	SubTitle  string     `json:"sub_title"`
	Des       string     `json:"des"`
	Cover     string     `json:"cover"`
	UserId    uint       `json:"user_id"`
	Status    int        `json:"status" gorm:"default:1"` // 1待审核 2审核通过 3未通过审核
	RejectRes string     `json:"reject_res"`              // 驳回原因
}

func (PhotoExhibition) TableName() string {
	return "photoExhibition"
}
