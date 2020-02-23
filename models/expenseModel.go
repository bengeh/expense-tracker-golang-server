package models

type Expense struct {
	Name  string `json:"name" bson:"name"`
	Price string `json:"price" bson:"price"`
}

// SetName receives a pointer to Foo so it can modify it.
func (expense *Expense) SetName(name string) {
    expense.Name = name
}

// SetName receives a pointer to Foo so it can modify it.
func (expense *Expense) SetPrice(price string) {
    expense.Price = price
}