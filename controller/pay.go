package controller

import (
	"furst/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaysController struct{}

type ByCategory struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount"`
}

type ByDay struct {
	Day        int          `json:"day"`
	Amount     int64        `json:"amount"`
	ByCategory []ByCategory `json:"by_category"`
}

type Ranking struct {
	Content string `json:"content"`
	Amount  int64  `json:"amount"`
}

type Response struct {
	ByDay   []ByDay   `json:"by_day"`
	Ranking []Ranking `json:"ranking"`
}

func (PaysController) GetPaysByYM(c *gin.Context) {
	year, _ := strconv.Atoi(c.Query("year"))
	month, _ := strconv.Atoi(c.Query("month"))
	mci, _ := strconv.Atoi(c.Query("main_category_id"))

	repo := repository.PayRepository{}
	cateRepo := repository.SubCategoryRepository{}
	pays := repo.GetPaysByYM(year, month, mci)

	jst, _ := time.LoadLocation("Asia/Tokyo")
	endDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, jst).AddDate(0, 1, -1).Day()

	res := Response{}

	for i := 1; i <= endDay; i++ {
		byDay := ByDay{
			Day: i,
		}
		res.ByDay = append(res.ByDay, byDay)
	}

	for i := 0; i < len(res.ByDay); i++ {
		var categoryIDs []uint
		categoryIDs = nil

		for _, pay := range pays {
			if (pay.Date.Day() != res.ByDay[i].Day) || pay.Transfer || !pay.CountIn {
				continue
			}

			res.ByDay[i].Amount += pay.Amount

			category, _ := cateRepo.FindByID(pay.SubCategoryID)
			if contains(categoryIDs, category.ID) {
				for j, byCategory := range res.ByDay[i].ByCategory {
					if byCategory.Name == category.Name {
						res.ByDay[i].ByCategory[j].Amount += pay.Amount
						break
					}
				}
			} else {
				byCategory := ByCategory{
					Name:   category.Name,
					Amount: pay.Amount,
				}
				res.ByDay[i].ByCategory = append(res.ByDay[i].ByCategory, byCategory)
				categoryIDs = append(categoryIDs, category.ID)
			}
		}
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    res,
	})
}

func contains(s []uint, i uint) bool {
	for _, v := range s {
		if i == v {
			return true
		}
	}
	return false
}

func pick(s []interface{}, f func(interface{}) bool) interface{} {
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	return nil
}
