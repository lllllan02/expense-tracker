/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var (
	updateId          int
	updateDescription string
	updateCategory    string
	updateAmount      float64
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新现有的消费记录",
	Long: `更新现有的消费记录信息。
	
该命令会根据指定的ID更新消费记录的描述、类型和金额，并自动更新修改时间。
至少需要提供一个要更新的字段（描述、类型或金额）。

使用方法：
  expense-tracker update --id 1 --description "晚餐" --category "餐饮" --amount 45.5
  
必需参数：
  --id 要更新的消费记录ID（例如：1）

可选参数（至少需要提供一个）：
  --description 新的消费描述（例如："购买书籍"）
  --amount 新的消费金额（例如：99.99）
  --category 新的消费类型（例如："学习用品"）

操作示例：
  expense-tracker update --id 1 --description "星巴克咖啡"
  expense-tracker update --id 2 --amount 8.00
  expense-tracker update --id 3 --category "交通"
  expense-tracker update --id 4 --description "地铁通勤" --category "交通" --amount 8.00

更新成功后会显示更新后的记录信息。如果指定的ID不存在，会显示错误信息。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if updateId <= 0 {
			return fmt.Errorf("消费记录ID必须为正整数")
		}

		if updateAmount < 0 {
			return fmt.Errorf("消费金额不能为负数")
		}

		// 获取原始记录
		originalExpense := expense.GetById(updateId)
		if originalExpense == nil {
			return fmt.Errorf("未找到ID为 %d 的消费记录", updateId)
		}

		// 直接修改指针指向的数据
		if updateDescription != "" {
			originalExpense.Description = updateDescription
		}
		if updateCategory != "" {
			originalExpense.Category = updateCategory
		}
		if updateAmount != 0 {
			originalExpense.Amount = updateAmount
		}
		originalExpense.UpdatedAt = time.Now()

		if err := expense.SaveFile(); err != nil {
			return fmt.Errorf("保存文件失败: %v", err)
		}

		fmt.Printf("更新消费记录成功 (ID: %d)\n", originalExpense.Id)
		fmt.Printf("描述: %s\n", originalExpense.Description)
		fmt.Printf("类型: %s\n", originalExpense.Category)
		fmt.Printf("金额: %.2f\n", originalExpense.Amount)
		fmt.Printf("更新时间: %s\n", originalExpense.UpdatedAt.Format("2006-01-02 15:04:05"))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateId, "id", "i", 0, "要更新的消费记录ID")
	updateCmd.MarkFlagRequired("id")

	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "新的消费描述")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0, "新的消费金额")
	updateCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "新的消费类型")
	updateCmd.MarkFlagsOneRequired("description", "amount", "category")
}
