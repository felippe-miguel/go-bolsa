package entity

type OrderQueue struct {
	Orders []*Order
}

// Less
func (orderQueue *OrderQueue) Less(i, j int) bool {
	return orderQueue.Orders[i].Price < orderQueue.Orders[j].Price
}

// Swap
func (orderQueue *OrderQueue) Swap(i, j int) {
	orderQueue.Orders[i], orderQueue.Orders[j] = orderQueue.Orders[j], orderQueue.Orders[i]
}

// Len
func (orderQueue *OrderQueue) Len() int {
	return len(orderQueue.Orders)
}

// Push
func (orderQueue *OrderQueue) Push(x interface{}) {
	orderQueue.Orders = append(orderQueue.Orders, x.(*Order))
}

// Pop
func (orderQueue *OrderQueue) Pop() interface{} {
	old := orderQueue.Orders
	n := len(old)
	item := old[n-1]
	orderQueue.Orders = old[0 : n-1]

	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
