package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_mercado_libre/domain/model"
	"test_mercado_libre/usecase/coupon"
)

// ItemsCoupon get items apply coupon by amount
func ItemsCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.ItemsPostRequest{}
		c.Bind(&requestBody)

		results, _ := coupon.GetItemCoupon(&requestBody)

		if results.Total == 0 {
			c.JSON(http.StatusNotFound, results)
		} else {
			c.JSON(http.StatusOK, results)
		}
	}
}
