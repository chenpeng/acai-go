package dto

type MoneyRecordChartDto struct {
	ClassificationCode int    `json:"classification_code"`
	ClassificationName string `json:"classification_name"`
	Money              string `json:"money"`
}
