package expense

import (
	"encoding/json"
	"log"
	"os"
	"slices"
)

var data Data

type Data struct {
	MaxId    int                `json:"max_id"`
	Expenses []Expense          `json:"expenses"`
	Budgets  map[string]float64 `json:"budgets"` // 格式: "08" -> 预算金额
}

func init() {
	file, err := os.OpenFile("expense.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("open expense.json file err: %v", err)
	}
	defer file.Close()

	decodeer := json.NewDecoder(file)
	decodeer.Decode(&data)
}

func SaveFile() error {
	bytes, _ := json.MarshalIndent(data, "", " ")
	return os.WriteFile("expense.json", bytes, 0644)
}

func GetData() *Data {
	return &data
}

func List(category []string) Expenses {
	expenses := make([]Expense, 0, len(data.Expenses))

	for _, expense := range data.Expenses {
		if len(category) == 0 || slices.Contains(category, expense.Category) {
			expenses = append(expenses, expense)
		}
	}

	return expenses
}
