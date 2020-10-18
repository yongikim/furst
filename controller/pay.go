package controller

import (
	"furst/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaysController struct{}

type ByCategory struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ByDay struct {
	Date       int          `json:"date"`
	Amount     int          `json:"amount"`
	ByCategory []ByCategory `json:"by_category"`
}

type Ranking struct {
	Content string `json:"content"`
	Amount  int    `json:"amount"`
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
	pays := repo.GetPaysByYM(year, month, mci)
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    pays,
	})
}
