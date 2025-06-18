/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var categoryFilter []string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出消费记录",
	Long: `列出所有或指定分类的消费记录。

该命令会以表格形式展示消费记录，并支持按分类筛选。
未指定分类时，默认显示所有记录。

使用方法：
  expense-tracker list                     # 显示所有消费记录
  expense-tracker list -c 餐饮 -c 交通      # 显示多个分类的消费
  expense-tracker list --category 购物     # 只显示购物类消费

参数说明：
  -c, --category 消费类型筛选（可重复使用）

输出字段说明：
  ID        - 消费记录唯一标识
  日期      - 消费发生时间
  分类      - 消费所属类别
  描述      - 消费具体描述
  金额      - 消费金额（单位：元）`,
	Run: func(cmd *cobra.Command, args []string) {
		expense.List(categoryFilter).Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringArrayVarP(&categoryFilter, "category", "c", nil, "消费类型筛选")
}
