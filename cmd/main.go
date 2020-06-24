package main

import (
	"fmt"
	"test_mercado_libre/domain/data"
	"test_mercado_libre/usecase/coupon"
)

func main() {
	res := coupon.Calculate(data.Items, 500.00)
	fmt.Println(res)
}
