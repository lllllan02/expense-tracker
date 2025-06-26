/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/lllllan02/expense-tracker/expense"
	"github.com/spf13/cobra"
)

var exportMonth int
var exportFilename string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "导出消费记录到CSV文件",
	Long: `将消费记录导出到CSV文件，方便在其他软件中查看和分析。
	
该命令可以将所有消费记录或指定月份的记录导出为CSV格式，
文件包含ID、创建时间、消费描述、消费类型和消费金额等信息。

使用方法：
  expense-tracker export --file expenses.csv
  expense-tracker export --file august.csv --month 8
  
参数：
  --file 输出文件名（必需）
  --month 指定月份（1-12），不指定则导出所有记录

操作示例：
  # 导出所有消费记录
  expense-tracker export --file all_expenses.csv
  
  # 导出8月份的消费记录
  expense-tracker export --file august_expenses.csv --month 8
  
  # 导出12月份的消费记录
  expense-tracker export --file december_expenses.csv --month 12`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if exportFilename == "" {
			return fmt.Errorf("必须指定输出文件名")
		}

		if exportMonth < 0 || exportMonth > 12 {
			return fmt.Errorf("月份必须在 1-12 之间")
		}

		// 如果没有指定文件名扩展名，自动添加.csv
		if filepath.Ext(exportFilename) == "" {
			exportFilename += ".csv"
		}

		err := expense.ExportToCSV(exportFilename, exportMonth)
		if err != nil {
			return fmt.Errorf("导出失败: %v", err)
		}

		if exportMonth == 0 {
			fmt.Printf("已成功导出所有消费记录到：%s\n", exportFilename)
		} else {
			monthName := time.Month(exportMonth).String()
			fmt.Printf("已成功导出%s消费记录到：%s\n", monthName, exportFilename)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&exportFilename, "file", "f", "", "输出文件名")
	exportCmd.MarkFlagRequired("file")

	exportCmd.Flags().IntVarP(&exportMonth, "month", "m", 0, "指定月份（1-12）")
}
