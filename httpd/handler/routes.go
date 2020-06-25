package handler

import (
	"github.com/gin-gonic/gin"
)

// Routes api
func Routes() *gin.Engine {

	r := gin.Default()

	r.POST("/coupon", ItemsCoupon())

	return r
}
