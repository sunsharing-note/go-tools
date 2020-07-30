package word

import (
	"strings"
	"unicode"
)

// 1.单词全部转为小写/大写
//把单词全部转为小写或大写的实现方法非常简单，直接调用标准库strings中的ToLower方法或ToUpper方法进行转换即可，
//其原生方法的作用就是把单词转为小写或大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 2.下画线单词转大写驼峰单词把以下画线方式命名的单词转为大写驼峰单词的主体逻辑是，把下画线替换为空格字符，
//然后再把所有字符修改为其对应的首字母大写形式，最后把空格字符替换为空，就可以确保各个部分返回的是首字母大写的字符串了

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 3.下画线单词转小写驼峰单词把以下画线方式命名的单词转为小写驼峰单词的主体逻辑是，直接复用大写驼峰单词的转换方法，
//然后对其首字母进行处理。在该方法中，我们直接将字符串的第一位取出来，然后利用unicode.ToLower进行转换

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 4.驼峰单词转下画线单词
//这里直接使用govalidator库提供的转换方法把大写或小写驼峰单词转为下画线单词。主体逻辑是在将字符转为小写的同时添加下画线。比较特殊的一点在于，
//如果当前字符不是小写字母、下画线或数字，那么在处理时将对segment置空，保证其每一段的区间转换都是正确的
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}