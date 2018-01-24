package putil

func reverse(arr []uint8) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		buffer := arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = buffer
	}
}

// convert baseStr to base-10 representation
func StrConv(baseStr []byte) int64 {
	var result int64 = 0
	var mul int64 = 1
	reverse(baseStr)
	for _, baseChar := range baseStr {
		result += mul * int64(baseChar-'0')
		mul *= 10
	}
	return result
}

func BaseConv(base, b10 int64) []uint8 {
	var accum []uint8
	for b10 > 0 {
		accum = append(accum, uint8(b10%base))
		b10 /= base
	}
	reverse([]byte(accum))
	return accum
}

func serial(start byte, n uint8) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = start + byte(i)
	}
	return string(result)
}

var encToks = serial('0', 10) + serial('a', 26)

func EncodeToks(tokens []uint8) string {
	output := make([]byte, len(tokens))
	for i, tok := range tokens {
		if i < 10 {
			output[i] = encToks[tok]
		}
	}
	return string(output)
}
