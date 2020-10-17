package controller

import (
	"furst/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaysController struct {}

func (PaysController) GetPaysByYM(c *gin.Context){
  year, _ := strconv.Atoi(c.Query("year"))
  month, _ := strconv.Atoi(c.Query("month"))
  mci, _ := strconv.Atoi(c.Query("main_category_id"))

  repo := repository.PayRepository{}
  pays := repo.GetPaysByYM(year, month, mci)
  c.JSONP(http.StatusOK, gin.H{
    "message": "ok",
    "data": pays,
  })
}
