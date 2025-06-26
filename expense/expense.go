package expense

import (
	"fmt"
	"slices"
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

func Add(description, category string, amount float64) *Expense {
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

func Delete(ids []int) {
	deleted := make([]Expense, 0, len(data.Expenses))
	for _, expense := range data.Expenses {
		if !slices.Contains(ids, expense.Id) {
			deleted = append(deleted, expense)
		}
	}
	data.Expenses = deleted
}

func GetById(id int) *Expense {
	for i := range data.Expenses {
		if data.Expenses[i].Id == id {
			return &data.Expenses[i]
		}
	}
	return nil
}

func Summary(month int) (float64, map[string]float64) {
	total := 0.0
	categoryTotals := make(map[string]float64)

	for _, expense := range data.Expenses {
		if month == 0 || int(expense.CreatedAt.Month()) == month {
			total += expense.Amount
			categoryTotals[expense.Category] += expense.Amount
		}
	}

	return total, categoryTotals
}

// SetBudget 设置指定月份的预算
func SetBudget(month int, amount float64) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("月份必须在 1-12 之间")
	}
	if amount < 0 {
		return fmt.Errorf("预算金额不能为负数")
	}

	if data.Budgets == nil {
		data.Budgets = make(map[string]float64)
	}

	key := fmt.Sprintf("%02d", month)
	data.Budgets[key] = amount
	return nil
}

// GetBudget 获取指定月份的预算
func GetBudget(month int) (float64, bool) {
	if data.Budgets == nil {
		return 0, false
	}

	key := fmt.Sprintf("%02d", month)
	amount, exists := data.Budgets[key]
	return amount, exists
}

// CheckBudget 检查指定月份是否超预算
func CheckBudget(month int) (float64, float64, bool, error) {
	budget, exists := GetBudget(month)
	if !exists {
		return 0, 0, false, fmt.Errorf("未设置 %d月 的预算", month)
	}

	total, _ := Summary(month)
	exceeded := total > budget

	return budget, total, exceeded, nil
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
