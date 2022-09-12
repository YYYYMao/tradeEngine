package order

import (
	"exam/model"
	"sync"
	"testing"
	"time"
)

func TestTradingAndUpdateOrderConcurrent(t *testing.T) {
	model.TradeList = make([]model.Trade, 0)
	model.Db.OrderMap = make(map[string]model.Order)
	model.Db.PendingOrderMap = make(map[string]map[string][]model.Order)
	fakeOrders := []Param{
		{
			Type:     0,
			Quantity: 10,
			Price:    1,
		},
		{
			Type:     0,
			Quantity: 5,
			Price:    10,
		},
		{
			Type:     0,
			Quantity: 2,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 3,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 4,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 3,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 10,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 10,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 2,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 1,
			Price:    1,
		},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(fakeOrders))
	for _, v := range fakeOrders {
		go func(v Param) {
			defer wg.Done()
			TradingAndUpdateOrder(v)
		}(v)
	}
	wg.Wait()

	model.PrintTrade()
	model.PrintAllOrder()
}

func TestTradingAndUpdateOrder(t *testing.T) {
	model.TradeList = make([]model.Trade, 0)
	model.Db.OrderMap = make(map[string]model.Order)
	model.Db.PendingOrderMap = make(map[string]map[string][]model.Order)
	fakeOrders := []Param{
		{
			Type:     0,
			Quantity: 10,
			Price:    1,
		},
		{
			Type:     0,
			Quantity: 5,
			Price:    10,
		},
		{
			Type:     0,
			Quantity: 2,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 3,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 4,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 3,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 10,
			Price:    10,
		},
		{
			Type:     1,
			Quantity: 10,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 2,
			Price:    1,
		},
		{
			Type:     1,
			Quantity: 1,
			Price:    1,
		},
	}

	for _, v := range fakeOrders {
		TradingAndUpdateOrder(v)
		time.Sleep(1000 * time.Millisecond)
	}

	model.PrintTrade()
	model.PrintAllOrder()
}
