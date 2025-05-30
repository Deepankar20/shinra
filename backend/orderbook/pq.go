package orderbook

import "container/heap"

type PriorityQueue []*Order

func (h PriorityQueue) Len() int {
	return len(h)
}

func (h PriorityQueue) Less(i, j int) bool {
	if h[i].Price == h[j].Price {
		return h[i].Timestamp.Before(h[j].Timestamp) // FIFO for same price
	}
	return h[i].Price > h[j].Price
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(*Order))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PriorityQueue) Peek() *Order {
	h2 := *h

	return h2[0]
}

func NewPriorityQueue() *PriorityQueue {
	h := &PriorityQueue{}
	heap.Init(h)

	return h
}

func PushOrder(pq *PriorityQueue, order *Order) {
	heap.Push(pq, order)
}

func PopOrder(pq *PriorityQueue) *Order {
	return heap.Pop(pq).(*Order)
}

func PeekTop(pq *PriorityQueue) *Order {
	return pq.Peek()
}
