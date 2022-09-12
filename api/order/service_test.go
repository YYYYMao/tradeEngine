package order

import (
	"exam/model"
	"fmt"
	"sync"
	"testing"
)

func TestTradingAndUpdateOrder(t *testing.T) {
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

	fmt.Println("===")
	model.PrintTrade()
	model.PrintAllOrder()
}
