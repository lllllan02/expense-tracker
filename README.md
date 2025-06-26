# 费用追踪器 (Expense Tracker)

一个简单而强大的命令行费用追踪工具，帮助您管理个人财务支出。

[![Go Version](https://img.shields.io/badge/Go-1.21.13+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 项目简介

这是一个基于 [roadmap.sh 费用追踪器项目](https://roadmap.sh/projects/expense-tracker) 实现的 Go 语言版本。该工具提供了完整的费用管理功能，包括添加、删除、更新、查看和统计消费记录，支持预算管理和数据导出。

## ✨ 主要功能

### 核心功能
- ✅ **添加消费记录** - 记录消费描述、类型和金额
- ✅ **查看消费记录** - 以表格形式展示所有记录
- ✅ **更新消费记录** - 修改现有记录的信息
- ✅ **删除消费记录** - 删除指定的消费记录
- ✅ **消费统计** - 查看总支出和分类统计
- ✅ **月度统计** - 查看指定月份的消费情况

### 高级功能
- 🎯 **预算管理** - 设置月度预算并监控超支情况
- 📊 **分类管理** - 按消费类型分类管理
- 📁 **数据导出** - 导出消费记录到 CSV 文件
- ⚠️ **超预算警告** - 自动提醒预算超支

## 🚀 快速开始

### 环境要求
- Go 1.21.13 或更高版本

### 安装

1. **克隆项目**
   ```bash
   git clone https://github.com/lllllan02/expense-tracker.git
   cd expense-tracker
   ```

2. **安装依赖**
   ```bash
   go mod tidy
   ```

3. **编译项目**
   ```bash
   go build -o expense-tracker
   ```

4. **运行程序**
   ```bash
   ./expense-tracker --help
   ```

## 📖 使用指南

### 基本命令

#### 添加消费记录
```bash
# 添加基本消费记录
expense-tracker add --description "午餐" --amount 35.5

# 添加带分类的消费记录
expense-tracker add --description "购买书籍" --category "学习用品" --amount 99.99

# 使用简写参数
expense-tracker add -d "咖啡" -c "餐饮" -a 28.00
```

#### 查看消费记录
```bash
# 查看所有消费记录
expense-tracker list

# 按分类筛选
expense-tracker list --category "餐饮"
expense-tracker list -c "交通" -c "购物"
```

#### 更新消费记录
```bash
# 更新消费描述
expense-tracker update --id 1 --description "星巴克咖啡"

# 更新消费金额
expense-tracker update --id 2 --amount 45.50

# 更新多个字段
expense-tracker update --id 3 --description "地铁通勤" --category "交通" --amount 8.00
```

#### 删除消费记录
```bash
# 删除单个记录
expense-tracker delete --id 1

# 删除多个记录
expense-tracker delete --id 2 --id 3
```

#### 查看统计信息
```bash
# 查看总支出
expense-tracker summary

# 查看指定月份支出
expense-tracker summary --month 8
```

### 预算管理

#### 设置预算
```bash
# 设置8月预算为1000元
expense-tracker budget --month 8 --amount 1000

# 查看8月预算
expense-tracker budget --month 8

# 查看所有月份预算
expense-tracker budget
```

#### 预算监控
当添加消费记录时，如果超出预算会自动显示警告：
```
⚠️  警告：本月已超预算 ¥50.00
```

### 数据导出

#### 导出到 CSV
```bash
# 导出所有消费记录
expense-tracker export --file all_expenses.csv

# 导出指定月份记录
expense-tracker export --file august_expenses.csv --month 8
```

## 📊 输出示例

### 消费记录列表
```
ID    创建时间    消费描述    消费类型    消费金额  
1     2024-01-15  午餐       餐饮        35.50    
2     2024-01-15  购买书籍   学习用品    99.99    
3     2024-01-16  地铁通勤   交通        8.00     
```

### 统计信息
```
总支出：¥143.49

按类别统计：
  餐饮：¥35.50
  学习用品：¥99.99
  交通：¥8.00
```

### 预算信息
```
8月总支出：¥150.00
预算：¥100.00
⚠️  超预算：¥50.00
```

## 🛠️ 技术实现

### 项目结构
```
expense-tracker/
├── cmd/                    # 命令行命令
│   ├── add.go             # 添加消费记录
│   ├── list.go            # 列出消费记录
│   ├── update.go          # 更新消费记录
│   ├── delete.go          # 删除消费记录
│   ├── summary.go         # 消费统计
│   ├── budget.go          # 预算管理
│   ├── export.go          # 数据导出
│   └── root.go            # 根命令
├── expense/               # 核心业务逻辑
│   ├── expense.go         # 消费记录模型和操作
│   └── file.go            # 文件操作
├── main.go                # 程序入口
└── go.mod                 # 依赖管理
```

### 依赖库
- [Cobra](https://github.com/spf13/cobra) - 命令行框架
- [go-runewidth](https://github.com/mattn/go-runewidth) - 中文字符宽度计算

### 数据存储
- 使用 JSON 格式存储消费记录和预算信息
- 数据文件：`expenses.json`

## 🔧 开发

### 构建
```bash
# 开发构建
go build

# 发布构建
go build -ldflags="-s -w" -o expense-tracker
```

### 测试
```bash
# 运行测试
go test ./...
```

## 📝 许可证

本项目基于 MIT 许可证开源 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📚 参考

本项目基于 [roadmap.sh 费用追踪器项目](https://roadmap.sh/projects/expense-tracker) 实现，感谢原项目的设计思路。

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 [Issue](https://github.com/lllllan02/expense-tracker/issues)
- 发送邮件：[your-email@example.com]

---

⭐ 如果这个项目对您有帮助，请给它一个星标！ 