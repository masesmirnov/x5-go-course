package main

import "testing"

func TestCategory(t *testing.T) {
	name, ok := category(0)
	if name != "small" || !ok {
		t.Errorf("category(0) = %q, %v, want small, true", name, ok)
	}
	name, ok = category(9_999)
	if name != "small" || !ok {
		t.Errorf("category(9_999) = %q, %v, want small, true", name, ok)
	}
	name, ok = category(10_000)
	if name != "regular" || !ok {
		t.Errorf("category(10_000) = %q, %v, want regular, true", name, ok)
	}
	name, ok = category(49_999)
	if name != "regular" || !ok {
		t.Errorf("category(49_999) = %q, %v, want regular, true", name, ok)
	}
	name, ok = category(50_000)
	if name != "large" || !ok {
		t.Errorf("category(50_000) = %q, %v, want large, true", name, ok)
	}
}

func TestCategoryNegative(t *testing.T) {
	_, ok := category(-1)
	if ok {
		t.Errorf("category(-1) ok = %v, want false", ok)
	}
}
