package utils

// BitPow is a golang implementation of binary exponential algorithm with O(lgn) time complexity.
// You can use BitPow to exponentiate quickly when you are sure it will not overflow int32.
// If the operation is likely to exceed a range of int32, use BitPowMod instead.
func BitPow(base, power int) int {
	ret := 1
	for power > 0 {
		if power&1 > 0 {
			ret *= base
		}
		base *= base
		power >>= 1
	}
	return ret
}

// BitPowMod is a golang implementation of binary exponential algorithm with O(lgn) time complexity.
// Use the `mod` parameter for overflow processing, and the size of `mod` should not exceed an int32 value.
// The recommended `mod` value is 1e9+7.
func BitPowMod(base, power, mod int) int {
	var ret, b, m int64 = 1, int64(base), int64(mod)
	for power > 0 {
		if power & 1 > 0 {
			ret *= b
			ret %= m
		}
		base *= base
		base %= mod
		power >>= 1
	}
	return int(ret)
}