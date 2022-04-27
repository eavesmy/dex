package utils

import "golang.org/x/crypto/sha3"

func LeftPadding(b []byte, length int) []byte {
	i := len(b)
	for ; i < length; i++ {
		b = append([]byte{0}, b...)
	}
	// b = append([]byte("0x"), b...)

	return b
}

func SplitBytes(b []byte, length int) [][]byte {
	ret := [][]byte{}
	for i := 0; i < len(b); i += length {
		ret = append(ret, b[i:i+length])
	}
	return ret
}

func ShaMethod(method string) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(method))
	return hash.Sum(nil)[:4]
}
