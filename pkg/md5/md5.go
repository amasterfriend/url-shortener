package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum 对传入的参数求md5值
// 通过加密散列函数，产生出一个128位（16字节）的散列值（通常表现为一个32个的16进制数字，每个16进制代表4字节）
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)) // 32个十六进制数
}
