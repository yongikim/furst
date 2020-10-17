package repository

import (
	"furst/model"
	"time"
)

type PayRepository struct{}

func (PayRepository) SetPay(pay *model.Pay) error {
	result := db.Create(&pay)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (PayRepository) GetPayList() []model.Pay {
	pays := make([]model.Pay, 0)
	result := db.Limit(10).Find(&pays)
	if result.Error != nil {
		panic(result.Error)
	}
	return pays
}

func (PayRepository) UpdatePay(newPay *model.Pay) error {
	result := db.Save(&newPay)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (PayRepository) DeletePay(id int) error {
	result := db.Delete(&model.Pay{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (PayRepository) FindByMf(mf string) bool {
	pay := model.Pay{}
	result := db.Where("mf = ?", mf).First(&pay)
	if result.Error != nil {
		return false
	}
	return true
}

func (PayRepository) GetPaysByYM(year int, month int, mci int) []model.Pay {
	pays := make([]model.Pay, 0)
	jst, _ := time.LoadLocation("Asia/Tokyo")
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, jst)
	result := db.Where("date >= ?", startOfMonth).Where("main_category_id = ?", mci).Find(&pays)
	if result.Error != nil {
		panic(result.Error)
	}
	return pays
}
