package coupon

import (
	"testing"
	"test_mercado_libre/domain/data"
)

func TestCalculate(t *testing.T) {
	itemsCoupon, _ := Calculate(data.Items, 500.00)

	if len(itemsCoupon) == len(data.Items) {
		t.Errorf("Calculate() = %d; want len 4", len(itemsCoupon))
	}
}
