package utils

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
	for i := 0; i < length; i += 64 {
		ret = append(ret, b[i:i+64])
	}
	return ret
}
