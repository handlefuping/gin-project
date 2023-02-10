package util

import (
	"math/rand"
	"strings"
	"time"
)

// LONG_LETTERS 这里不能使用[]byte存储时乱码
var LONG_LETTERS = strings.Split("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

func RandStr() string {
	str := ""
	ran := rand.New(rand.NewSource(time.Now().UnixNano())) // 生成一个rand
	for i := 0; i < 32; i++ {
		str += LONG_LETTERS[ran.Intn(len(LONG_LETTERS))]
	}
	return str
}