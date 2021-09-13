package types

//Money  представляет собой денежную сумму в минимальных единицах(центы, копейки, и т.д)
type Money int64

//Category  представляет сабой категорию в которой был совершен платёж(авто, рестораны и т.д)
type Category string

//Payment представляет  информацию о платеже  
type Payment struct {
	ID int
	Amount Money
	Category Category
}
