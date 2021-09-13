package stats

import (
	"fmt"

	"github.com/Behzod01/bank/pkg/bank/types"
)

func ExampleAvg() {
	cards := []types.Payment {
		{
			ID: 1234,
			Amount: 300_00,
			Category: "Кредит",
		},
		{
			ID: 2222,
			Amount: 10_000,
			Category: "Телефон",
		},
		{
			ID: 4613,
			Amount: 800_00,
			Category: "TV",
		},
	}
	middle := Avg(cards)
	fmt.Println(middle) 
	//Output: 40000

}