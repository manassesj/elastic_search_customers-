package utils

import (
	"strings"
	"time"
)

func BuildIndex(fill string, startDate string, endDate string) string {
	tStartDate, _ := time.Parse("2006-01-02", startDate)
	tEndDate, _ := time.Parse("2006-01-02", endDate)
	var indexs []string

	for ok := true; ok; ok = tStartDate != tEndDate {
		indexs = append(indexs, fill+tStartDate.Format("2006-01-02"))
		tStartDate = tStartDate.Add(time.Hour * 24)
	}
	indexs = append(indexs, fill+tStartDate.Format("2006-01-02"))
	return strings.Join(indexs, ",")
}
