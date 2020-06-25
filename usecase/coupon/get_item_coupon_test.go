package coupon

import (
	"os"
	"test_mercado_libre/domain/model"
	"testing"
)

func TestGetItemCoupon(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	items := model.ItemsPostRequest{
		ItemIDs: []string{
			"MCO492245833",
			"MCO492245834",
			"MCO492245835",
		},
		Amount: 150000.60,
	}

	itemCoupon, _ := GetItemCoupon(&items)

	if len(itemCoupon.ItemIDs) <= 0 {
		t.Errorf("GetItemCoupon() = %v; want items", len(itemCoupon.ItemIDs))
	}
}

func TestGetItemCouponNoExist(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	items := model.ItemsPostRequest{
		ItemIDs: []string{
			"MLA1",
			"MLA2",
		},
		Amount: 50000.50,
	}

	itemCoupon, _ := GetItemCoupon(&items)

	if len(itemCoupon.ItemIDs) > 0 {
		t.Errorf("GetItemCoupon() = %v; want items not found", len(itemCoupon.ItemIDs))
	}
}

func TestGetItemCouponAmountInsuficient(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	items := model.ItemsPostRequest{
		ItemIDs: []string{
			"MCO492245833",
			"MCO492245834",
		},
		Amount: 500.80,
	}

	itemCoupon, _ := GetItemCoupon(&items)

	if itemCoupon.Total > 0 {
		t.Errorf("GetItemCoupon() = %v; want 0", itemCoupon.Total)
	}
}
