package putil

var encToks = serial('0', 10) + serial('a', 26)

// ======== Public ========

// StrConv ... convert baseStr to base-10 representation
func StrConv(baseStr []byte) int64 {
	var result int64
	var mul int64 = 1
	reverse(baseStr)
	for _, baseChar := range baseStr {
		result += mul * int64(baseChar-'0')
		mul *= 10
	}
	return result
}

// BaseConv ... convert b10 base 10 value to base
func BaseConv(base, b10 int64) []uint8 {
	var accum []uint8
	for b10 > 0 {
		tok := uint8(b10 % base)
		accum = append(accum, tok)
		b10 /= base
	}
	reverse([]byte(accum))
	return accum
}

// EncodeToks ... convert base-n value to its string representation
func EncodeToks(tokens []uint8) string {
	output := make([]byte, len(tokens))
	for i, tok := range tokens {
		output[i] = encToks[tok]
	}
	return string(output)
}

// ======== Private ========

// reverse byte string in place
func reverse(arr []byte) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		buffer := arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = buffer
	}
}

// return a string starting with byte string from start to start + n - 1
func serial(start byte, n uint8) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = start + byte(i)
	}
	return string(result)
}
