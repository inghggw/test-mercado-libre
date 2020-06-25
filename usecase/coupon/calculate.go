package coupon

import (
	"sort"
	"test_mercado_libre/domain/model"
)

type kv struct {
	Key   string
	Value float64
}

// Calculate get items apply cuopon by amount
func Calculate(items map[string]float64, amount float64) ([]model.List, float64) {
	var lists []model.List
	var sum float64
	var total float64

	itemsSort := sortItems(items)

	for _, row := range itemsSort {
		sum += row.Value

		if sum >= amount {
			break
		}

		lists = append(lists, model.List{row.Key})
		total += row.Value
	}

	return lists, total
}

func sortItems(items map[string]float64) []kv {
	var ss []kv

	for k, v := range items {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	return ss
}
