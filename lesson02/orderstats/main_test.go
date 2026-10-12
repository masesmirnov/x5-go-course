package main

import "testing"

func TestStats(t *testing.T) {
	minAmount, maxAmount, sum, avg, ok := stats([]int64{12000, 7500, 31000, 4500})
	if !ok {
		t.Errorf("ok = %v, want true", ok)
	}
	if minAmount != 4500 {
		t.Errorf("min = %d, want 4500", minAmount)
	}
	if maxAmount != 31000 {
		t.Errorf("max = %d, want 31000", maxAmount)
	}
	if sum != 55000 {
		t.Errorf("sum = %d, want 55000", sum)
	}
	if avg != 13750 {
		t.Errorf("avg = %v, want 13750", avg)
	}
}

func TestStatsOneOrder(t *testing.T) {
	minAmount, maxAmount, _, _, ok := stats([]int64{3_000_000_000})
	if !ok {
		t.Errorf("ok = %v, want true", ok)
	}
	if minAmount != 3_000_000_000 {
		t.Errorf("min = %d, want 3000000000", minAmount)
	}
	if maxAmount != 3_000_000_000 {
		t.Errorf("max = %d, want 3000000000", maxAmount)
	}
}

func TestStatsFractionalAvg(t *testing.T) {
	_, _, _, avg, _ := stats([]int64{1, 2})
	if avg != 1.5 {
		t.Errorf("avg = %v, want 1.5", avg)
	}
}

func TestStatsEmpty(t *testing.T) {
	_, _, _, _, ok := stats([]int64{})
	if ok {
		t.Errorf("ok = %v, want false", ok)
	}
}
