/* ================================================================
 * ┌ GO.A.P - golang algorithm pattern [/](github.com/bagaking/goap)
 * └ usage: copy the hole segment into your conquest code
 *
 * Copyright (c) 2020 kinghand@foxmail.com
 * ================================================================
*/

package goap

/*** !!! GO.A.P START !!! ***/

// mapping: apply reduce method to an slice
func reduce(arr []int, reducer func (prev, v int) int, initVal int) int {
	for _, v := range arr {
		initVal = reducer(initVal, v)
	}
	return initVal
}

// mapping: apply mapping method to an slice, immutable
func mapping(arr []int, mapper func (v int) int) []int {
	ret := make([]int, len(arr))
	for i, v := range arr {
		ret[i] = mapper(v)
	}
	return ret
}

// mapping: apply mapping method to an slice, mutable
func mappingInPlace(arr []int, mapper func (v int) int) []int {
	for i, v := range arr {
		arr[i] = mapper(v)
	}
	return arr
}

/*** !!! GO.A.P END !!! ***/