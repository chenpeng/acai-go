package dto

import "container/list"

type MoneyRecordDto struct {
	DateStr                string    `json:"dateStr"`
	WeekStr                string    `json:"weekStr"`
	DayPayMoney            float64   `json:"dayPayMoney"`
	DayIncomeMoney         float64   `json:"dayIncomeMoney"`
	MoneyRecordItemDtoList list.List `json:"moneyRecordItemDtoList"`
}
