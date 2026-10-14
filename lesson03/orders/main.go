package main

import (
	"errors"
	"fmt"
)

type OrderItem struct {
	Name  string
	Price int64
}

type Order struct {
	ID     int
	Status string
	Items  []OrderItem
}

func NewOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("order id must be positive")
	}
	return Order{ID: id, Status: "new"}, nil
}

func AddItem(order *Order, item OrderItem) error {
	if order == nil {
		return errors.New("order is nil")
	}
	if item.Name == "" {
		return errors.New("item name is empty")
	}
	if item.Price < 0 {
		return errors.New("item price is negative")
	}
	order.Items = append(order.Items, item)
	return nil
}

func CalculateTotal(order Order) int64 {
	var total int64
	for _, item := range order.Items {
		total += item.Price
	}
	return total
}

func statusNames() map[string]string {
	names := make(map[string]string, 2)
	names["new"] = "новый"
	names["paid"] = "оплачен"
	return names
}

func main() {
	order, err := NewOrder(1)
	if err != nil {
		fmt.Println("не удалось создать заказ:", err)
		return
	}
	items := []OrderItem{
		{Name: "молоко", Price: 8_990},
		{Name: "хлеб", Price: 5_490},
		{Name: "сыр", Price: 45_000},
	}
	for _, item := range items {
		err = AddItem(&order, item)
		if err != nil {
			fmt.Println("не удалось добавить позицию:", err)
			return
		}
	}
	fmt.Printf("заказ %d: %s\n", order.ID, statusNames()[order.Status])
	fmt.Println("позиций:", len(order.Items))
	fmt.Println("сумма:", CalculateTotal(order), "коп.")
}
