package orderbook

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestOrderBookHeap(t *testing.T, db *gorm.DB) {
	ob := NewOrderBook()

	ob.HandleIncomingOrder(&Order{Price: 0.6, Shares: 5, UserID: 1, Timestamp: time.Now(), Yes: true}, db)
	ob.HandleIncomingOrder(&Order{Price: 0.3, Shares: 5, UserID: 1, Timestamp: time.Now(), Yes: false}, db)

	t.Errorf("top of yesOrder : %v", ob.YesOrders.Peek().Price)
	t.Errorf("top of noOrder : %v", ob.NoOrders.Peek().Price)

}
