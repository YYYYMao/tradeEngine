package model

import (
	"exam/utils"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type OrderType int
type StatusType int

const (
	Sell OrderType = iota
	Buy
)

const (
	Cancel StatusType = iota
	Pending
	Completed
)

type Order struct {
	ID               uuid.UUID
	Type             OrderType
	OriginalQuantity float64
	Quantity         float64
	Price            float64
	CreatedAt        int64
	Status           StatusType
}

type DataStore struct {
	sync.Mutex
	OrderMap        map[string]Order
	PendingOrderMap map[string]map[string][]Order
}

func New() *DataStore {
	return &DataStore{
		OrderMap:        make(map[string]Order),
		PendingOrderMap: make(map[string]map[string][]Order),
	}
}

var Db *DataStore

func init() {
	Db = New()
}

func (ds *DataStore) AddOrder(order Order) {
	ds.Lock()
	defer ds.Unlock()

	ds.OrderMap[order.ID.String()] = order
}

func (ds *DataStore) GetOrder(id string) (Order, bool) {
	if v, ok := ds.OrderMap[id]; ok {
		return v, true
	}
	return Order{}, false
}

func (ds *DataStore) AddPendingOrder(order Order) {
	ds.Lock()
	defer ds.Unlock()
	if key := getPendingOrderKey(order); key != "" {
		if _, ok := ds.PendingOrderMap[key]; !ok {
			ds.PendingOrderMap[key] = make(map[string][]Order)
		}
		timestampKey := fmt.Sprintf("%d", order.CreatedAt)
		ds.PendingOrderMap[key][timestampKey] = append(ds.PendingOrderMap[key][timestampKey], order)
	}
}

func (ds *DataStore) SearchOrder(order Order) []Order {
	ds.Lock()
	defer ds.Unlock()

	matchingOrders := make([]Order, 0)
	quantity := float64(0)
	if key := getMatchingOrderKey(order); key != "" {
		if v, ok := ds.PendingOrderMap[key]; ok {
			for _, orders := range v {
				for _, _order := range orders {
					if targetOrder, ok := ds.GetOrder(_order.ID.String()); ok && targetOrder.Quantity > 0 {
						quantity += targetOrder.Quantity
						matchingOrders = append(matchingOrders, targetOrder)
					}
				}
				if quantity > order.Quantity {
					break
				}
			}
		}
	}
	return matchingOrders
}

func (ds *DataStore) RemoveOrder(order Order) {
	ds.Lock()
	defer ds.Unlock()

	order.Status = Cancel
	ds.OrderMap[order.ID.String()] = order
	ds.RemovePendingOrder(order)
}

func (ds *DataStore) RemovePendingOrder(order Order) {
	ds.Lock()
	defer ds.Unlock()

	if key := getPendingOrderKey(order); key != "" {
		if _, ok := ds.PendingOrderMap[key]; !ok {
			return
		}
		timestampKey := fmt.Sprintf("%d", order.CreatedAt)
		if _, ok := ds.PendingOrderMap[key][timestampKey]; !ok {
			return
		}
		removeIndex := -1
		for i, v := range ds.PendingOrderMap[key][timestampKey] {
			if v.ID == order.ID {
				removeIndex = i
				break
			}
		}
		if removeIndex != -1 {
			ds.PendingOrderMap[key][timestampKey] = append(ds.PendingOrderMap[key][timestampKey][:removeIndex], ds.PendingOrderMap[key][timestampKey][removeIndex+1:]...)
		}
	}
}

func (ds *DataStore) UpdateOrder(order Order) {
	if order.Quantity == 0 {
		order.Status = Completed
		ds.RemovePendingOrder(order)
	}
	ds.AddOrder(order)
}

func getPendingOrderKey(order Order) string {
	key := ""
	if order.Type == Sell {
		key = utils.GetHashkey(fmt.Sprintf("Sell#%f", order.Price))
	} else if order.Type == Buy {
		key = utils.GetHashkey(fmt.Sprintf("Buy#%f", order.Price))
	}
	return key
}

func getMatchingOrderKey(order Order) string {
	key := ""
	if order.Type == Sell {
		key = utils.GetHashkey(fmt.Sprintf("Buy#%f", order.Price))
	} else if order.Type == Buy {
		key = utils.GetHashkey(fmt.Sprintf("Sell#%f", order.Price))
	}
	return key
}

func PrintAllOrder() {
	for _, v := range Db.OrderMap {
		fmt.Println("===============")
		t := "Sell"
		if v.Type == Buy {
			t = "Buy"
		}
		s := "Pending"
		if v.Status == Completed {
			s = "Completed"
		}
		fmt.Println("Order ID: ", v.ID)
		fmt.Println("Type: ", t, "Original Quantity: ", v.OriginalQuantity, "Quantity: ", v.Quantity, " Price: ", v.Price)
		fmt.Println("Created: ", v.CreatedAt, "Status: ", s)
	}
}
