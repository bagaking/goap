package goap

import "testing"

func TestMin(t *testing.T) {
	if v := G.min(1, 2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := G.min(-1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.min(1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.min(MaxInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.min(MinInt, -2); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}
}

func TestMax(t *testing.T) {
	if v := G.max(1, 2); v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}

	if v := G.max(-1, -2); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := G.max(1, -2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := G.max(MaxInt, -2); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := G.max(MinInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}
}

func TestMinIn(t *testing.T) {
	if v := G.minIn(1, 2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := G.minIn(-1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.minIn(1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.minIn(MaxInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.minIn(MinInt, -2); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}

	if v := G.minIn(); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := G.minIn(-1, 0, 1, 2, 3); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := G.minIn([]int{-1, 0, 1, 2, 3}...); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}
}

func TestMaxIn(t *testing.T) {
	if v := G.maxIn(1, 2); v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}

	if v := G.maxIn(-1, -2); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := G.maxIn(1, -2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := G.maxIn(MaxInt, -2); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := G.maxIn(MinInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := G.maxIn(); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}

	if v := G.maxIn(-1, 0, 1, 2, 3); v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}

	if v := G.maxIn([]int{-1, 0, 1, 2, 3}...); v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
}

func TestReduce(t *testing.T) {
	if v := G.reduce([]int{-2, -1, 0, 1, 2}, func(a, b int) int { return a + b }, 0); v != 0 {
		t.Errorf("WA: expect 0, got %d\n", v)
	}
	if v := G.reduce([]int{-2, -1, 0, 1, 3}, func(a, b int) int { return a + b }, 0); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}
	if v := G.reduce([]int{-2, -1, 0, 1, 3}, func(a, b int) int { return a + b }, 2); v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
	if v := G.reduce([]int{-2, -1, 1, 2}, func(a, b int) int { return a * b }, 2); v != 8 {
		t.Errorf("WA: expect 8, got %d\n", v)
	}
}

func TestMapping(t *testing.T) {
	req := []int{-2, -1, 0, 1, 2}
	rsp := G.mapping(req, func(a int) int { return a + 2 })
	if v := rsp[0]; v != 0 {
		t.Errorf("WA: expect 0, got %d\n", v)
	}
	if v := rsp[1]; v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}
	if v := rsp[2]; v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}
	if v := rsp[3]; v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
	if v := rsp[4]; v != 4 {
		t.Errorf("WA: expect 4, got %d\n", v)
	}
	if v := req[0]; v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}
	if v := req[1]; v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}
	if v := req[2]; v != 0 {
		t.Errorf("WA: expect 0, got %d\n", v)
	}
	if v := req[3]; v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}
	if v := req[4]; v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}
}

func TestMappingInPlace(t *testing.T) {
	req := []int{-2, -1, 0, 1, 2}
	rsp := G.mappingInPlace(req, func(a int) int { return a + 2 })
	if v := rsp[0]; v != 0 {
		t.Errorf("WA: expect 0, got %d\n", v)
	}
	if v := rsp[1]; v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}
	if v := rsp[2]; v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}
	if v := rsp[3]; v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
	if v := rsp[4]; v != 4 {
		t.Errorf("WA: expect 4, got %d\n", v)
	}
	if v := req[0]; v != 0 {
		t.Errorf("WA: expect 0, got %d\n", v)
	}
	if v := req[1]; v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}
	if v := req[2]; v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}
	if v := req[3]; v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
	if v := req[4]; v != 4 {
		t.Errorf("WA: expect 4, got %d\n", v)
	}
}

