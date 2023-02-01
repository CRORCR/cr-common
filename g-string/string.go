package g_string

import (
	"errors"
	"strings"
)

type StringClass struct {
}

var String = StringClass{}

func (this *StringClass) Desensitize(str string) string {
	index := strings.Index(str, `@`)
	if index == -1 {
		return this.DesensitizeMobile(str)
	} else {
		return this.DesensitizeEmail(str)
	}
}

/*
*
>7        前3后4中间4个*
<=7 && >4 前2后2中间4个*
<=4       报错
*/
func (this *StringClass) DesensitizeMobile(str string) string {
	result := ``
	length := len(str)
	if length > 7 {
		result = str[:3] + `****` + str[length-4:]
	} else if length <= 7 && length > 4 {
		result = str[:2] + `****` + str[length-2:]
	} else {
		panic(errors.New(`mobile length is too small`))
	}
	return result
}

/*
*
@前字符串长度>3   前4 中4个* 后@后面所有
@前字符串长度<=3  前@前面所有 中4个* 后@后面所有
*/
func (this *StringClass) DesensitizeEmail(str string) string {
	result := ``
	index := strings.Index(str, `@`)
	if index == -1 {
		panic(errors.New(`not email`))
	}
	preAt := str[:index]
	if len(preAt) > 3 {
		result = str[:4] + `****` + str[index:]
	} else {
		result = preAt + `****` + str[index:]
	}
	return result
}

func (this *StringClass) RemoveLast(str string, num int) string {
	return str[:len(str)-num]
}

func (this *StringClass) RemoveFirst(str string, num int) string {
	return str[num:]
}

func (this *StringClass) Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func (this *StringClass) ReplaceAll(str string, oldStr string, newStr string) (result string) {
	return strings.Replace(str, oldStr, newStr, -1)
}

func (this *StringClass) SpanLeft(str string, length int, fillChar string) string {
	if len(str) > length {
		panic(errors.New(`length is too small`))
	}
	if len(fillChar) != 1 {
		panic(errors.New(`length of fillChar must be 1`))
	}
	result := ``
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result + str
}

func (this *StringClass) SpanRight(str string, length int, fillChar string) string {
	if len(str) > length {
		panic(errors.New(`length is too small`))
	}
	if len(fillChar) != 1 {
		panic(errors.New(`length of fillChar must be 1`))
	}
	result := str
	for i := 0; i < length-len(str); i++ {
		result += fillChar
	}
	return result
}

func (this *StringClass) StartWith(str string, substr string) bool {
	return strings.HasPrefix(str, substr)
}

func (this *StringClass) EndWith(str string, substr string) bool {
	return strings.HasSuffix(str, substr)
}
