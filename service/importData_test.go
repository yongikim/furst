package service

import (
	"testing"
	"time"
)

func TestImportData(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	year := now.Year()
	month := int(now.Month())

	service := ImportDataService{}
	_, err := service.Call(year, month)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
