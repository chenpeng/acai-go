package dto

import "time"

type MoneyRecordItemDto struct {
	Id                 int64     `json:"id"`
	ClassificationCode int32     `json:"classification_code"`
	ClassificationName string    `json:"classification_name"`
	Money              float64   `json:"money"`
	Type               int32     `json:"type"`
	Remark             string    `json:"remark"`
	RecordDateTime     string    `json:"record_date_time"`
	PicUrl             string    `json:"pic_url"`
	DeleteFlag         bool      `json:"delete_flag"`
	CreateUserId       int64     `json:"create_user_id"`
	CreateUserName     string    `json:"create_user_name"`
	CreateDateTime     time.Time `json:"create_date_time"`
	UpdateUserId       int64     `json:"update_user_id"`
	UpdateUserName     string    `json:"update_user_name"`
	UpdateDateTime     time.Time `json:"update_date_time"`
}
