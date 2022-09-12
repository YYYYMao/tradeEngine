package model

import (
	"exam/utils"
	"fmt"

	"github.com/google/uuid"
)

type Trade struct {
	ID        uuid.UUID
	BuyerId   uuid.UUID
	SellerId  uuid.UUID
	Quantity  float64
	Price     float64
	CreatedAt int64
}

var TradeList []Trade

func AddTrade(buyer, seller Order) {
	TradeList = append(TradeList, Trade{
		ID:        utils.GetID(),
		BuyerId:   buyer.ID,
		SellerId:  seller.ID,
		Quantity:  seller.Quantity,
		Price:     buyer.Price,
		CreatedAt: utils.GetTimestamp(),
	})
}

func PrintTrade() {
	for _, v := range TradeList {
		fmt.Println("===============")
		fmt.Println("TRADE ID: ", v.ID)
		fmt.Println("Buyer ID: ", v.BuyerId, "Seller ID: ", v.SellerId)
		fmt.Println("Quantity: ", v.Quantity, "Price: ", v.Price, "Created: ", v.CreatedAt)
	}
}
