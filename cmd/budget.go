/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var budgetMonth int
var budgetAmount float64

// budgetCmd represents the budget command
var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "设置和查看月度预算",
	Long: `设置和查看指定月份的消费预算。
	
该命令可以设置特定月份的预算金额，也可以查看已设置的预算信息。
设置预算后，在查看消费统计时会显示预算使用情况和超预算警告。

使用方法：
  expense-tracker budget --month 8 --amount 1000
  expense-tracker budget --month 8
  expense-tracker budget
  
参数：
  --month 指定月份（1-12）
  --amount 预算金额

操作示例：
  # 设置8月预算为1000元
  expense-tracker budget --month 8 --amount 1000
  
  # 查看8月的预算
  expense-tracker budget --month 8
  
  # 查看所有月份的预算
  expense-tracker budget`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 检查参数绑定：要么都不传，要么都传
		monthFlag := cmd.Flags().Lookup("month")
		amountFlag := cmd.Flags().Lookup("amount")

		monthSet := monthFlag != nil && monthFlag.Changed
		amountSet := amountFlag != nil && amountFlag.Changed

		if monthSet != amountSet {
			return fmt.Errorf("--month 和 --amount 参数必须同时使用")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// 如果两个参数都没有设置，显示所有月份预算
		if budgetMonth == 0 && budgetAmount == 0 {
			return showAllBudgets()
		}

		// 如果两个参数都设置了，设置预算
		if budgetMonth < 1 || budgetMonth > 12 {
			return fmt.Errorf("月份必须在 1-12 之间")
		}

		err := expense.SetBudget(budgetMonth, budgetAmount)
		if err != nil {
			return fmt.Errorf("设置预算失败: %v", err)
		}

		if err := expense.SaveFile(); err != nil {
			return fmt.Errorf("保存文件失败: %v", err)
		}

		fmt.Printf("已设置 %d月 预算：¥%.2f\n", budgetMonth, budgetAmount)
		return nil
	},
}

// showAllBudgets 显示所有月份的预算
func showAllBudgets() error {
	if len(expense.GetData().Budgets) == 0 {
		fmt.Println("未设置任何月份预算")
		return nil
	}

	// 获取所有月份并排序
	months := make([]int, 0, len(expense.GetData().Budgets))
	for key := range expense.GetData().Budgets {
		if month, err := strconv.Atoi(key); err == nil {
			months = append(months, month)
		}
	}
	sort.Ints(months)

	fmt.Println("所有月份预算：")
	for _, month := range months {
		key := fmt.Sprintf("%02d", month)
		amount := expense.GetData().Budgets[key]
		fmt.Printf("  %d月：¥%.2f\n", month, amount)

		// 显示预算使用情况
		budget, totalSpent, exceeded, err := expense.CheckBudget(month)
		if err == nil {
			if exceeded {
				overAmount := totalSpent - budget
				fmt.Printf("    当前支出：¥%.2f，超预算：¥%.2f\n", totalSpent, overAmount)
			} else {
				remaining := budget - totalSpent
				fmt.Printf("    当前支出：¥%.2f，剩余：¥%.2f\n", totalSpent, remaining)
			}
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(budgetCmd)

	budgetCmd.Flags().IntVarP(&budgetMonth, "month", "m", 0, "指定月份（1-12）")
	budgetCmd.Flags().Float64VarP(&budgetAmount, "amount", "a", 0, "预算金额")
}
