package order

import (
	"exam/model"
	"exam/utils"
	"math"
)

func TradingAndUpdateOrder(order Param) {

	newOrder := model.Order{
		ID:               utils.GetID(),
		Type:             order.Type,
		Quantity:         order.Quantity,
		OriginalQuantity: order.Quantity,
		Price:            order.Price,
		CreatedAt:        utils.GetTimestamp(),
		Status:           model.Pending,
	}

	quantity := float64(0)
	maxQuantity := order.Quantity

	updateOrders := make([]model.Order, 0)
	updateTrades := make([]model.Order, 0)

	if matchingOrders := model.Db.SearchOrder(newOrder); len(matchingOrders) > 0 {
		for _, matchingOrder := range matchingOrders {
			if quantity < maxQuantity {
				needQuantity := maxQuantity - quantity
				minQuantity := math.Min(needQuantity, matchingOrder.Quantity)
				quantity += minQuantity
				matchingOrder.Quantity -= minQuantity
				updateOrders = append(updateOrders, matchingOrder)
				updateTrades = append(updateTrades, model.Order{
					ID:       matchingOrder.ID,
					Quantity: minQuantity,
				})
			} else if quantity == maxQuantity {
				break
			}
		}
	}
	newOrder.Quantity -= quantity
	updateOrders = append(updateOrders, newOrder)
	for _, v := range updateOrders {
		model.Db.UpdateOrder(v)
	}
	if newOrder.Quantity > 0 {
		model.Db.AddPendingOrder(newOrder)
	}
	for _, v := range updateTrades {
		model.AddTrade(newOrder, v)
	}
}
