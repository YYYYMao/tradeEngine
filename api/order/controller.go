package order

import (
	"encoding/json"
	"exam/model"
	"net/http"
)

type Param struct {
	Type     model.OrderType `json:"type"`
	Quantity float64         `json:"quantity"`
	Price    float64         `json:"price"`
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	var p Param

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if p.Type != model.Sell && p.Type != model.Buy {
		http.Error(w, "BAD REQUEST", http.StatusBadRequest)
		return
	}
	TradingAndUpdateOrder(p)
	w.WriteHeader(http.StatusOK)
}
