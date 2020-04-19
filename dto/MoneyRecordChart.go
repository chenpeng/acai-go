package dto

import "time"

type MoneyRecordChartDto struct {
	Date  time.Time `json:"date"`
	Money string    `json:"money"`
}
