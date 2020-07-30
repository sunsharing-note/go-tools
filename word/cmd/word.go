package cmd

import (
	"log"
	"strings"
	"github.com/sunsharing-note/go-tools/word/internal/word"
	"github.com/spf13/cobra"
)

// 定义目前单词所支持的转换模式枚举值
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下线线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

var str string
var mode int8

// 对具体的单词子命进行设置和集成
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转换为大写",
	"2: 全部单词转换为小写",
	"3：下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"5: 驼峰单词转换为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command,arg []string) {
		var content  string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.UnderscoreToLowerCamelCase(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果： %s",content)
	},
}

// 命令行参数的设置和初始化
func init()  {
	wordCmd.Flags().StringVarP(&str, "str", "s", "","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0,"请输入单词转换模式")
}