/* ================================================================
 * ┌ GO.A.P - golang algorithm pattern [/](github.com/bagaking/goap)
 * └ usage: copy the hole segment into your conquest code
 *
 * Copyright (c) 2020 kinghand@foxmail.com
 * ================================================================
*/
package goap

/*** !!! GO.A.P START !!! ***/
/*** ref: const ***/

// min: find out the min value between two int32 values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max: find out the max value between two int32 values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// minIn: to use this method, you should ensure the inputs are not empty
func minIn(arr ...int) int {
	minVal := MaxInt
	for _, v := range arr {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// maxIn: to use this method, you should ensure the inputs are not empty
func maxIn(arr ...int) int {
	maxVal := MinInt
	for _, v := range arr {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

/*** !!! GO.A.P END !!! ***/