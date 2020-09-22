package types
// Money for money
type Money int64
// Category for card
type Category string
// Payment for structer
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}