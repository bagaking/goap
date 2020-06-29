package bitslice

import (
	"sort"
)

type BitSlice []int


func (nums *BitSlice) tidyRange(l, r int) (int, int) {
	if l < 0 {
		l = 0
	}
	if r < 0 {
		n := len(*nums)
		r = (n + r % n) % n
	}
	return l, r
}

// ConvertToBitSlice creates a BitSlice from a normal slice.
//   Note that the bitSlice shares memory with the original slice, and is therefore a mutable method, but also modifies the original slice.
//   If used in an immutable scenario, a copied array would be passed instead.
func ConvertToBitSlice(nums []int) BitSlice {
	sort.Ints(nums)
	return BitSlice(nums)
}

// Search returns the index of the given value, -1 if not found, and the time complexity is O(lgn).
func (nums BitSlice) Search(v int, l, r int) int {
	l, r = nums.tidyRange(l, r)
	for l <= r {
		m := l + (r-l) >> 2
		val := nums[m]
		if val == v {
			return m
		} else if val > v {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// BitSectLeft searches for and returns the insertion position of the given value, ensuring that the BitSect is still ordered after the insertion.
//   Inserting from this position guarantees that the bitSect is still ordered after the insertion.
//   If the same value is already present in the BitSect, the FIRST insertion position is returned.
func (nums BitSlice) BitSectLeft(v int, l, r int) int {
	l, r = nums.tidyRange(l, r)
	for l < r {
		m := l + (r-l) >> 1
		vMid := nums[m]
		if v <= vMid {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

// BitSectRight searches for and returns the insertion position of the given value, ensuring that the BitSect is still ordered after the insertion.
//   Inserting from this position guarantees that the bitSect is still ordered after the insertion.
//   If the same value is already present in the BitSect, the LAST insertion position is returned.
func (nums BitSlice) BitSectRight(v int, l, r int) int {
	l, r = nums.tidyRange(l, r)
	for l < r {
		m := l + (r-l+1) >> 1
		vMid := nums[m]
		if v >= vMid {
			l = m
		} else {
			r = m - 1
		}
	}
	return l + 1
}