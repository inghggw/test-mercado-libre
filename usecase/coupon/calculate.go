package coupon

import ("sort")

// List with item coupon
type List struct {
	ItemID string
}

type kv struct {
	Key   string
	Value float64
}

// Calculate get items apply cuopon by amount
func Calculate(items map[string]float64, amount float64) []List {
	var lists []List
	var sum float64

	itemsSort := sortItems(items)

	for _, row := range itemsSort {
		sum += row.Value

		if sum >= amount {
			break
		}

		lists = append(lists, List{row.Key})
	}

	return lists
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
