package help

func StrToByteSlice(str string) (result []byte) {
	result = make([]byte, len(str))
	for i, v := range str {
		result[i] = byte(v)
	}
	return
}
