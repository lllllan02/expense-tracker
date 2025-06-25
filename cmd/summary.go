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

var summaryMonth int

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "显示消费统计信息",
	Long: `显示消费记录的统计信息，包括总支出金额和各类别分布。
	
该命令可以显示所有消费记录的总计，也可以按月份筛选显示特定月份的统计信息。

使用方法：
  expense-tracker summary
  expense-tracker summary --month 8
  
参数：
  --month 指定月份（1-12），不指定则显示所有记录

操作示例：
  expense-tracker summary
  # 总支出：¥20.00
  
  expense-tracker summary --month 8
  # 8月份总支出：¥20.00
  
  expense-tracker summary --month 12
  # 12月份总支出：¥150.00`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if summaryMonth < 0 || summaryMonth > 12 {
			return fmt.Errorf("月份必须在 1-12 之间")
		}

		total, categoryTotals := expense.Summary(summaryMonth)

		if summaryMonth == 0 {
			fmt.Printf("总支出：¥%.2f\n", total)
		} else {
			monthName := time.Month(summaryMonth).String()
			fmt.Printf("%s总支出：¥%.2f\n", monthName, total)
		}

		// 显示每个类别的统计
		if len(categoryTotals) > 0 {
			fmt.Println("\n按类别统计：")
			for category, amount := range categoryTotals {
				fmt.Printf("  %s：¥%.2f\n", category, amount)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVarP(&summaryMonth, "month", "m", 0, "指定月份（1-12）")
}
