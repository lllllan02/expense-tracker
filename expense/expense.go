package expense

import (
	"fmt"
	"time"

	"github.com/mattn/go-runewidth"
)

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
	data.Expenses = append(data.Expenses, Expense{
		Id:          data.MaxId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: description,
		Category:    category,
		Amount:      amount,
	})

	return &data.Expenses[len(data.Expenses)-1]
}

type Expenses []Expense

func (es Expenses) Print() {
	widths := []int{
		runewidth.StringWidth("ID"),
		runewidth.StringWidth("创建时间"),
		runewidth.StringWidth("消费描述"),
		runewidth.StringWidth("消费类型"),
		runewidth.StringWidth("消费金额"),
	}

	// 计算每列需要的长度
	for _, e := range es {
		widths[0] = max(widths[0], runewidth.StringWidth(fmt.Sprint(e.Id)))
		widths[1] = max(widths[1], runewidth.StringWidth(e.CreatedAt.Format(time.DateOnly)))
		widths[2] = max(widths[2], runewidth.StringWidth(e.Description))
		widths[3] = max(widths[3], runewidth.StringWidth(e.Category))
		widths[4] = max(widths[4], runewidth.StringWidth(fmt.Sprintf("%.2f", e.Amount)))
	}

	id := runewidth.FillRight("ID", widths[0]+2)
	date := runewidth.FillRight("创建时间", widths[1]+2)
	description := runewidth.FillRight("消费描述", widths[2]+2)
	category := runewidth.FillRight("消费类型", widths[3]+2)
	amount := runewidth.FillRight("消费金额", widths[4]+2)
	fmt.Printf("%s%s%s%s%s\n", id, date, description, category, amount)

	for _, e := range es {
		id := runewidth.FillRight(fmt.Sprint(e.Id), widths[0]+2)
		date := runewidth.FillRight(e.CreatedAt.Format(time.DateOnly), widths[1]+2)
		description := runewidth.FillRight(e.Description, widths[2]+2)
		category := runewidth.FillRight(e.Category, widths[3]+2)
		amount := runewidth.FillRight(fmt.Sprintf("%.2f", e.Amount), widths[4]+2)
		fmt.Printf("%s%s%s%s%s\n", id, date, description, category, amount)
	}
}
