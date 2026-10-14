package main

import "testing"

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Errorf("NewOrder(1) err = %v, want nil", err)
	}
	if order.ID != 1 {
		t.Errorf("NewOrder(1) ID = %d, want 1", order.ID)
	}
	if order.Status != "new" {
		t.Errorf("NewOrder(1) Status = %q, want new", order.Status)
	}
	order, err = NewOrder(2)
	if err != nil {
		t.Errorf("NewOrder(2) err = %v, want nil", err)
	}
	if order.ID != 2 {
		t.Errorf("NewOrder(2) ID = %d, want 2", order.ID)
	}
}

func TestNewOrderNotPositiveID(t *testing.T) {
	_, err := NewOrder(0)
	if err == nil {
		t.Errorf("NewOrder(0) err = %v, want error", err)
	}
	_, err = NewOrder(-1)
	if err == nil {
		t.Errorf("NewOrder(-1) err = %v, want error", err)
	}
}

func TestCalculateTotalEmpty(t *testing.T) {
	order, _ := NewOrder(1)
	total := CalculateTotal(order)
	if total != 0 {
		t.Errorf("total = %d, want 0", total)
	}
}

func TestCalculateTotalTwoItems(t *testing.T) {
	order, _ := NewOrder(1)
	_ = AddItem(&order, OrderItem{Name: "молоко", Price: 8_990})
	_ = AddItem(&order, OrderItem{Name: "хлеб", Price: 5_490})
	total := CalculateTotal(order)
	if total != 14_480 {
		t.Errorf("total = %d, want 14480", total)
	}
}

func TestAddItemFreeItem(t *testing.T) {
	order, _ := NewOrder(1)
	err := AddItem(&order, OrderItem{Name: "пакет", Price: 0})
	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
	if len(order.Items) != 1 {
		t.Errorf("len(Items) = %d, want 1", len(order.Items))
	}
}

func TestAddItemErrors(t *testing.T) {
	err := AddItem(nil, OrderItem{Name: "молоко", Price: 8_990})
	if err == nil {
		t.Errorf("nil order err = %v, want error", err)
	}
	order, _ := NewOrder(1)
	err = AddItem(&order, OrderItem{Name: "", Price: 8_990})
	if err == nil {
		t.Errorf("empty name err = %v, want error", err)
	}
	err = AddItem(&order, OrderItem{Name: "молоко", Price: -1})
	if err == nil {
		t.Errorf("negative price err = %v, want error", err)
	}
	if len(order.Items) != 0 {
		t.Errorf("len(Items) = %d, want 0", len(order.Items))
	}
}

func TestOrderCopy(t *testing.T) {
	order, _ := NewOrder(1)
	_ = AddItem(&order, OrderItem{Name: "молоко", Price: 8_990})
	orderCopy := order
	orderCopy.Status = "paid"
	_ = AddItem(&orderCopy, OrderItem{Name: "хлеб", Price: 5_490})
	if order.Status != "new" {
		t.Errorf("Status = %q, want new", order.Status)
	}
	if len(order.Items) != 1 {
		t.Errorf("len(Items) = %d, want 1", len(order.Items))
	}
	total := CalculateTotal(order)
	if total != 8_990 {
		t.Errorf("total = %d, want 8990", total)
	}
}

func TestStatusNames(t *testing.T) {
	names := statusNames()
	if names["new"] != "новый" {
		t.Errorf("names[new] = %q, want новый", names["new"])
	}
	if names["paid"] != "оплачен" {
		t.Errorf("names[paid] = %q, want оплачен", names["paid"])
	}
}
