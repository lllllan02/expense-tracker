/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var ids []int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除指定 id 的消费记录",
	Long: `删除指定 id 的消费记录。

该命令会根据提供的消费记录 id 永久删除对应记录，并更新存储文件。
删除操作不可逆，请谨慎使用。

使用方法：
  expense-tracker delete --id 123          # 删除 ID 为 123 的记录
  expense-tracker delete --id 456 --id 789 # 同时删除多个记录

参数说明：
  -i --id 消费记录的唯一标识（必填）`,
	RunE: func(cmd *cobra.Command, args []string) error {
		expense.Delete(ids)
		expense.SaveFile()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntSliceVarP(&ids, "id", "i", nil, "待删除的消费 id")
	deleteCmd.MarkFlagRequired("id")
}
