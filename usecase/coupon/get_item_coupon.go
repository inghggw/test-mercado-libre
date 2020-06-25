package coupon

import (
	"test_mercado_libre/domain/model"
	"test_mercado_libre/service"
)

// GetItemCoupon item exist and calculate coupon
func GetItemCoupon(items *model.ItemsPostRequest) (*model.ItemsCoupon, error) {
	var itemsID []string
	itemsMap := make(map[string]float64)
	itemsCoupon := &model.ItemsCoupon{}

	for _, id := range items.ItemIDs {
		itemExist, _ := service.GetItemMercadoLibre(id)

		if itemExist.Status == "404" {
			itemsCoupon.Message = "El item " + id + " no existe"

			return itemsCoupon, nil
		}

		itemsMap[id] = itemExist.Price
	}

	res, total := Calculate(itemsMap, items.Amount)

	for _, row := range res {
		itemsID = append(itemsID, row.ItemID)
	}

	itemsCoupon.ItemIDs = itemsID
	itemsCoupon.Total = total

	if total == 0 {
		itemsCoupon.Message = "El monto del cupon no es suficiente para comprar mínimamente un item"
	}

	return itemsCoupon, nil
}
