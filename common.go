package utils

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/davecgh/go-spew/spew"
	"time"
)

func Uniqid(prefix ...string) string {
	var p = ""
	for k, v := range prefix {
		if k == 1 {
			p = v
		}
	}
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", p, now.Unix(), now.UnixNano()%0x100000)
}

func GenerateOrderSn() string {
	spew.Dump(TimeSecondStr())
	return TimeSecondStr() + snowflakeId()
}

func snowflakeId() string {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return ""
	}
	id := n.Generate()
	return id.String()
}

func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && IsASCIIUpper(s[i+1]) {
			continue
		}
		if IsASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if IsASCIIUpper(c) {
			c ^= ' '
		}
		t = append(t, c)

		for i+1 < len(s) && IsASCIIUpper(s[i+1]) {
			i++
			t = append(t, '_')
			t = append(t, bytes.ToLower([]byte{s[i]})[0])
		}
	}
	return string(t)
}

func IsASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// If 模拟三元表达式
// Go不支持运算符重载，因此需要先将 a<b 在函数外转换为 bool 条件
// Go不支持泛型，只能用 interface{} 模拟
// 返回的类型安全需要用户自己保证，.(type) 的类型必须匹配
// a, b := 2, 3
// max := If(a > b, a, b).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
