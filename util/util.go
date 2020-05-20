//File  : util.go
//Author: duanhaobin
//Date  : 2020/5/20

package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var alphabet = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM"
	randName := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range randName {
		randName[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(randName)
}
