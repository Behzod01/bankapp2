package stats

import "github.com/Behzod01/bank/pkg/bank/types"

//Avg
func Avg(payments []types.Payment) types.Money {
	var middle, count types.Money
	for _, v := range payments {
		middle += v.Amount
		count++		
	}
	middle /= count
	return middle
}
