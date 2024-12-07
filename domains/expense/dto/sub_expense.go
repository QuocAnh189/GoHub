package dto

type SubExpense struct {
	ID        string  `json:"id"`
	ExpenseId string  `json:"expenseId"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
}

type CreatedSubExpenseReq struct {
	ExpenseId string  `json:"expenseId"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
}

type UpdateSubExpenseReq struct {
	ID        string  `json:"id"`
	ExpenseId string  `json:"expenseId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
