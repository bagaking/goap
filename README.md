[![Build Status](https://travis-ci.com/bagaking/goap.svg?branch=master)](https://travis-ci.com/bagaking/goap)

# GOAP

GOlang Algorithm Pattern library

> version: 0.1

## Usage

Copy the entire required code into your topic code

### g

- MaxInt
- MinInt
- G.min(a, b int) int
- G.max(a, b,int) int
- G.minIn(arr ...int) int
- G.maxIn(arr ...int) int
- G.reduce(arr []int, reducer func (prev, v int) int, initVal int) int
- G.mapping(arr []int, mapper func (v int) int) []int
- G.mappingInPlace(arr []int, mapper func (v int) int) []int

### bitslice

- ConvertToBitSlice(nums []int) BitSlice
- bitSlice.Search(v int, l, r int) int
- bitSlice.BitSectLeft(v int, l, r int) int
- bitSlice.BitSectRight(v int, l, r int) int

### utils

- BitPow(base, power int) int
- BitPowMod(base, power, mod int) int


