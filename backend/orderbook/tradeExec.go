package orderbook

import (
	"github.com/Deepankar20/shinra/backend/models"
	"gorm.io/gorm"
)

func ExecuteTrade(ob *OrderBook, buyOrder *Order, sellOrder *Order, db *gorm.DB) error {

	sharesToTrade := min(buyOrder.Shares, sellOrder.Shares)

	yesAmountLocked := sharesToTrade * int(buyOrder.Price)
	noAmountLocked := sharesToTrade * int(sellOrder.Price)

	remainingYesShares := buyOrder.Shares - sharesToTrade
	remainingNoShares := sellOrder.Shares - sharesToTrade

	tradeYes := models.Trade{
		UserID:     buyOrder.UserID,
		EventID:    ob.EventID,
		Amount:     float64(yesAmountLocked),
		Prediction: "yes",
		Status:     "pending",
		Result:     "pending",
	}

	tradeNo := models.Trade{
		UserID:     sellOrder.UserID,
		EventID:    ob.EventID,
		Amount:     float64(noAmountLocked),
		Prediction: "no",
		Status:     "pending",
		Result:     "pending",
	}

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(tradeYes).Error; err != nil {
		tx.Rollback()

	}

	if err := tx.Create(tradeNo).Error; err != nil {
		tx.Rollback()
	}

	if remainingNoShares > 0 {
		orderToPush := sellOrder
		orderToPush.Shares = remainingNoShares
		PushOrder(ob.NoOrders, orderToPush)
	}

	if remainingYesShares > 0 {
		orderToPush := buyOrder
		orderToPush.Shares = remainingYesShares
		PushOrder(ob.YesOrders, orderToPush)
	}

	return nil
}
