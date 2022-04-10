package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kislovandrew/task/models"
)

var item6_columns = []string{"id", "name", "metric"}
var item6 []models.Item6

var order, size, page, order_by string

func GetDBData(c *gin.Context) {

	if models.Err_connect != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to DB connection"})
		return
	}

	order = c.DefaultQuery("order", "metric")
	size = c.DefaultQuery("size", "5")
	page = c.DefaultQuery("page", "1")
	order_by = c.DefaultQuery("order_by", "desc")

	// проверка параметров для корректного запроса к БД
	s, err := strconv.Atoi(size)
	if s > 100 && s < 1 || err != nil {
		s = 5
	}
	p, err := strconv.Atoi(page)
	if p < 0 || err != nil {
		p = 0
	}
	contain := false
	for _, i := range item6_columns {
		if i == order {
			contain = true
			break
		}
	}
	if !contain {
		order = "metric"
	}
	if order_by != "asc" && order_by != "desc" {
		order_by = "desc"
	}

	models.DB.Order(order + " " + order_by).Limit(s).Offset((p - 1) * s).Find(&item6)

	c.JSON(http.StatusOK, gin.H{"result": item6})
}
