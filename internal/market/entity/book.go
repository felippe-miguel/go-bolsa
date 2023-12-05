package entity

import "sync"

type Book struct {
	Order         []*Order
	Transactions  []*Transaction
	OrdersChan    chan *Order
	OrdersOutChan chan *Order
	Wg            *sync.WaitGroup
}

func NewBook(ordersChan chan *Order, ordersOutChan chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transactions:  []*Transaction{},
		OrdersChan:    ordersChan,
		OrdersOutChan: ordersOutChan,
		Wg:            wg,
	}
}
