package goap

/* ================================================================
 * ┌ GOAP - golang algorithm pattern [/](github.com/bagaking/goap)
 * └ usage: copy the hole segment into your conquest code
 *
 * !!! GOAP START !!!
 *
 * ================================================================
*/

type Goap int
const G Goap = 0
const MaxInt = 1 << 31 - 1
const MinInt = -1 << 31

// min: find out the min value between two int32 values
func (g Goap) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max: find out the max value between two int32 values
func (g Goap) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minIn: to use this method, you should ensure the inputs are not empty
func (g Goap) minIn(arr ...int) int {
	min := MaxInt
	for _, v := range arr {
		min = g.min(min, v)
	}
	return min
}

// maxIn: to use this method, you should ensure the inputs are not empty
func (g Goap) maxIn(arr ...int) int {
	max := MinInt
	for _, v := range arr {
		max = g.max(max, v)
	}
	return max
}

// mapping: apply reduce method to an slice
func (g Goap) reduce(arr []int, reducer func (prev, v int) int, initVal int) int {
	for _, v := range arr {
		initVal = reducer(initVal, v)
	}
	return initVal
}

// mapping: apply mapping method to an slice, immutable
func (g Goap) mapping(arr []int, mapper func (v int) int) []int {
	ret := make([]int, len(arr))
	for i, v := range arr {
		ret[i] = mapper(v)
	}
	return ret
}

// mapping: apply mapping method to an slice, mutable
func (g Goap) mappingInPlace(arr []int, mapper func (v int) int) []int {
	for i, v := range arr {
		arr[i] = mapper(v)
	}
	return arr
}


/* ================================================================
 *
 * !!! GOAP END !!!
 *
 * ================================================================
 */