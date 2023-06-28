package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func GetEmptyString() string {

	return ""
}

// IsIntArrString  是否是整形|多个整形字符串
func IsIntArrString(str string) bool {

	regexStr := `^\d+(,\d+)*$`
	var isExist bool = false
	match, err := regexp.Compile(regexStr)
	if err == nil {
		isExist = match.MatchString(str)
	}
	return isExist
}

// IsIntString  是否是整形字符串
func IsIntString(str string) bool {

	regexStr := `^\d+$`
	var isExist = false
	match, err := regexp.Compile(regexStr)
	if err == nil {
		isExist = match.MatchString(str)
	}
	return isExist
}

func GetIntArrFromString(str string) []int {

	arrStr := strings.Split(str, ",")
	var arrInt []int
	for _, v := range arrStr {
		intVal, err := strconv.Atoi(v)
		if err == nil {
			arrInt = append(arrInt, intVal)
		}
	}
	return arrInt
}

func Md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func DeleteExtraSpace(s string) string {
	// 删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)      // 替换tab为空格
	regstr := "\\s{2,}"                          // 两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             // 编译正则表达式
	s2 := make([]byte, len(s1))                  // 定义字符数组切片
	copy(s2, s1)                                 // 将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) // 在字符串中搜索
	for len(spc_index) > 0 {                     // 找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) // 删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            // 继续在字符串中搜索
	}
	return string(s2)
}

// RandomString return fixed length random string
func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// Substr => php substr("hello world", start, length);
func Substr(str string, start int, length int) string {
	return str[start : start+length]
}

// MbSubstr => PHP mb_substr
func MbSubstr(str string, start int, length int) string {
	return string([]rune(str)[start : start+length])
}

// MbStrpos => PHP mb_strpos
func MbStrpos(haystack, needle string) int {
	index := strings.Index(haystack, needle)
	if index == -1 || index == 0 {
		return index
	}
	pos := 0
	total := 0
	reader := strings.NewReader(haystack)
	for {
		_, size, err := reader.ReadRune()
		if err != nil {
			return -1
		}
		total += size
		pos++
		// got it
		if total == index {
			return pos
		}
	}
}

// MbStrlen => PHP mb_strlen
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}
