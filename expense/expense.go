package expense

import "time"

type Expense struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
}

func AddExpense(description, category string, amount float64) *Expense {
	data.MaxId++
	data.Expenses = append(data.Expenses, &Expense{
		Id:          data.MaxId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
		Category:    category,
		Amount:      amount,
	})

	return data.Expenses[len(data.Expenses)-1]
}
