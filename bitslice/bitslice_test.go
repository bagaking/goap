package bitslice

import "testing"

func TestBitSliceTidyRange(t *testing.T) {
	bs := ConvertToBitSlice([]int{ 1, 2, 3})
	l, r := bs.tidyRange(-1, -1)
	if l != 0 {
		t.Errorf("WA: expect 0, got %d", l)
	}
	if r != 2 {
		t.Errorf("WA: expect 2, got %d", r)
	}
}

func TestBitSliceSearch(t *testing.T) {
	bs := ConvertToBitSlice([]int{ 2,2,1,2,3,4,5,6,7,9})
	ind := 0

	if ind = bs.Search(1, 0, -1); ind != 0 {
		t.Errorf("WA: expect 0, got %d", ind)
	}

	if ind = bs.Search(2, 0, -1); ind != 1 && ind != 2 && ind != 3 {
		t.Errorf("WA: expect 1~3, got %d", ind)
	}

	if ind = bs.Search(3, 0, -1); ind != 4 {
		t.Errorf("WA: expect 4, got %d", ind)
	}

	if ind = bs.Search(5, 0, -1); ind != 6 {
		t.Errorf("WA: expect 6, got %d", ind)
	}

	if ind = bs.Search(9, 0, -1); ind != 9 {
		t.Errorf("WA: expect 9, got %d", ind)
	}

	if ind = bs.Search(10, 0, -1); ind != -1 {
		t.Errorf("WA: expect -1, got %d", ind)
	}
}

func TestBitSliceSectLeft(t *testing.T) {
	ind := 0

	if ind = ConvertToBitSlice([]int{ 1, 5, 5, 5, 5, 6}).BitSectLeft(5, 0, -1); ind != 1 {
		t.Errorf("WA: expect 1, got %d", ind)
	}

	if ind = ConvertToBitSlice([]int{ 5, 5, 5, 5, 6}).BitSectLeft(5, 0, -1); ind != 0 {
		t.Errorf("WA: expect 0, got %d", ind)
	}

	if ind = ConvertToBitSlice([]int{ 5, 5, 5, 5, 6}).BitSectLeft(5, 1, -1); ind != 1 {
		t.Errorf("WA: expect 1, got %d", ind)
	}
}

func TestBitSliceSectRight(t *testing.T) {
	ind := 0

	if ind = ConvertToBitSlice([]int{ 1, 5, 5, 5, 5, 6}).BitSectRight(5, 0, -1); ind != 5 {
		t.Errorf("WA: expect 5, got %d", ind)
	}

	if ind = ConvertToBitSlice([]int{ 1, 5, 5, 5, 5}).BitSectRight(5, 0, -1); ind != 5 {
		t.Errorf("WA: expect 5, got %d", ind)
	}
}