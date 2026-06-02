package utils

import (
        "strings" // <--- 必须加上这一行
	// 使用别名 latexconv，确保不与 markdown.go 中的 converter 结构体冲突
	latexconv "github.com/burgr033/latex2unicode/pkg/converter"
)
func ConvertLatexToUnicode(input string) string {
	// 1. 先进行手动映射替换（把箭头转成 Unicode 字符）
	input = strings.ReplaceAll(input, "$\\rightarrow$", "→")
	input = strings.ReplaceAll(input, "$\\leftarrow$", "←")

        // 2. 清理边界：先去掉空格，再去掉首尾的 $
        // TrimSpace 会去掉字符串首尾的空白符（空格、换行、制表符）
        input = strings.TrimSpace(input)
        // Trim 会去掉字符串首尾的所有 $ 符号
        input = strings.Trim(input, "$")
        // 如果处理后里面还有空格（比如 "$ → $" 变成了 " → "），再去掉一次空格
        input = strings.TrimSpace(input)

	// 3. 创建实例
	c, err := latexconv.New()
	if err != nil {
		return input
	}

	// 4. 执行转换
	return c.ConvertString(input)
}
