/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"fmt"

	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var (
	description string
	category    string
	amount      float64
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加新的消费记录",
	Long: `添加新的消费记录到账本中。
	
该命令会创建一条新的消费记录，包括消费描述、消费类型和金额，并将其保存到本地文件。

使用方法：
  expense-tracker add --description "午餐" --category "餐饮" --amount 35.5
  
必需参数：
  --description 消费描述（例如："购买书籍"）
  --amount 消费金额（例如：99.99）

可选参数：
  --category 消费类型（例如："学习用品"）

操作示例：
  expense-tracker add --description "咖啡" --amount 28.00
  expense-tracker add --description "地铁通勤" --category "交通" --amount 7.00

添加成功后会显示新记录的唯一ID，可用于后续查询或修改操作。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if amount <= 0 {
			return fmt.Errorf("消费金额必须为正数")
		}

		e := expense.Add(description, category, amount)
		if err := expense.SaveFile(); err != nil {
			return fmt.Errorf("添加消费记录失败: %v", err)
		}

		fmt.Printf("添加消费记录成功 (ID: %d)\n", e.Id)

		// 检查是否超预算
		month := int(e.CreatedAt.Month())
		budget, totalSpent, exceeded, err := expense.CheckBudget(month)
		if err == nil && exceeded {
			overAmount := totalSpent - budget
			fmt.Printf("⚠️  警告：本月已超预算 ¥%.2f\n", overAmount)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "", "消费描述")
	addCmd.MarkFlagRequired("description")

	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "消费金额")
	addCmd.MarkFlagRequired("amount")

	addCmd.Flags().StringVarP(&category, "category", "c", "默认", "消费类型")
}
