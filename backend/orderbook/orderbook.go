package orderbook

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Price     float64
	Shares    int
	UserID    uint
	Timestamp time.Time
	Yes       bool
}

type OrderBook struct {
	NoOrders  *PriorityQueue
	YesOrders *PriorityQueue
	EventID   uint
}

func NewOrderBook() *OrderBook {

	yesPQ := NewPriorityQueue()
	noPQ := NewPriorityQueue()

	return &OrderBook{
		NoOrders:  noPQ,
		YesOrders: yesPQ,
	}
}

func (ob *OrderBook) HandleIncomingOrder(incomingOrder *Order, db *gorm.DB) {

	if incomingOrder.Yes {

		if ob.NoOrders.Len() > 0 {
			topNO := PeekTop(ob.NoOrders)

			if topNO.Price+incomingOrder.Price >= 1 {
				PopOrder(ob.NoOrders)
				ExecuteTrade(ob, incomingOrder, topNO, db)
			} else {

				PushOrder(ob.YesOrders, incomingOrder)

			}
		} else {
			PushOrder(ob.YesOrders, incomingOrder)
		}
	} else {

		if ob.YesOrders.Len() > 0 {
			topYes := PeekTop(ob.YesOrders)

			if topYes.Price+incomingOrder.Price >= 1 {
				PopOrder(ob.YesOrders)
				ExecuteTrade(ob, topYes, incomingOrder, db)
			} else {

				PushOrder(ob.NoOrders, incomingOrder)

			}
		} else {
			PushOrder(ob.NoOrders, incomingOrder)
		}

	}

}
