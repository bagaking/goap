package goap

import "testing"

func TestMin(t *testing.T) {
	if v := min(1, 2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := min(-1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := min(1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := min(MaxInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := min(MinInt, -2); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}
}

func TestMax(t *testing.T) {
	if v := max(1, 2); v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}

	if v := max(-1, -2); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := max(1, -2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := max(MaxInt, -2); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := max(MinInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}
}

func TestMinIn(t *testing.T) {
	if v := minIn(1, 2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := minIn(-1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := minIn(1, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := minIn(MaxInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := minIn(MinInt, -2); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}

	if v := minIn(); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := minIn(-1, 0, 1, 2, 3); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := minIn([]int{-1, 0, 1, 2, 3}...); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}
}

func TestMaxIn(t *testing.T) {
	if v := maxIn(1, 2); v != 2 {
		t.Errorf("WA: expect 2, got %d\n", v)
	}

	if v := maxIn(-1, -2); v != -1 {
		t.Errorf("WA: expect -1, got %d\n", v)
	}

	if v := maxIn(1, -2); v != 1 {
		t.Errorf("WA: expect 1, got %d\n", v)
	}

	if v := maxIn(MaxInt, -2); v != MaxInt {
		t.Errorf("WA: expect MaxInt, got %d\n", v)
	}

	if v := maxIn(MinInt, -2); v != -2 {
		t.Errorf("WA: expect -2, got %d\n", v)
	}

	if v := maxIn(); v != MinInt {
		t.Errorf("WA: expect MinInt, got %d\n", v)
	}

	if v := maxIn(-1, 0, 1, 2, 3); v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}

	if v := maxIn([]int{-1, 0, 1, 2, 3}...); v != 3 {
		t.Errorf("WA: expect 3, got %d\n", v)
	}
}


