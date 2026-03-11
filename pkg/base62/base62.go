package base62

import (
	"math"
	"strings"
)

// 62进制转换模块
// 0-9 a-z A-Z,正好表示了数字0-61

// const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

// 为了避免被人恶意请求，可以打乱上述字符

var (
	base62Str string
)

// MustInit 使用base62包必须调用该函数完成初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	base62Str = bs
}

// To62String 十进制数转为62进制字符串
func To62String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	var bl []byte
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}
	// 反转。除k取余法
	return string(reverse(bl))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// String2Int 62进制字符串转为十进制数
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// 从右往左遍历
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(base) * uint64(strings.Index(base62Str, string(b)))
	}
	return seq
}
