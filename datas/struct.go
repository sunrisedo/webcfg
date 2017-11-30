package datas

import (
	"sync"

	"github.com/hugozhu/godingtalk"
)

var (
	RGoods      = NewGoods()
	ConnDing    = new(godingtalk.DingTalkClient)
	RWebGoods   = NewRamGoods()
	RAppGoods   = NewRamGoods()
	RTradeGoods = NewRamGoods()
)

type Product struct {
	ID             int
	Code           string
	MID            string
	RID            string
	Coin           string
	PriceFixed     int
	VolFixed       int
	MPrecise       int
	RPrecise       int
	Munlimited     float64
	Runlimited     float64
	SoldUnits      float64
	BuyCommission  int
	SellCommission int
	Side           string
	Method         string
	Closed         string
	UnitPrice      float64
}

type Currency struct {
	Code     string  `json:"Code,omitempty"`
	Value    int     `json:"ID,omitempty"`
	Symbol   string  `json:"Symbol,omitempty"`
	MinTrade float64 `json:"MinTrade,omitempty"`
	MinCash  float64 `json:"MinCash,omitempty"`
	CashFree float64 `json:"CashFree,omitempty"`
}
type ProdSimple struct {
	Sort       int `json:"Sort,omitempty"`
	Code       string
	Name       string
	PriceFixed int
	VolFixed   int
}

type MarketSimple struct {
	Sort    int          `json:"Sort,omitempty"`
	Code    string       `json:"Code,omitempty"`
	Name    string       `json:"Name,omitempty"`
	Product []ProdSimple `json:"product,omitempty"`
}

/************************************************/
type Goods struct {
	value *GoodsValue
	lock  *sync.RWMutex
}

func NewGoods() *Goods {
	return &Goods{NewGoodsValue(), new(sync.RWMutex)}
}
func (r *Goods) Get() *GoodsValue {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.value
}
func (r *Goods) Set(value *GoodsValue) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.value = value
}

/************************************************/
type RamGoods struct {
	value *RamGoodsValue
	lock  *sync.RWMutex
}

type RamGoodsValue struct {
	Market   interface{} `json:"Market,omitempty"`
	Product  interface{} `json:"Product,omitempty"`
	Currency interface{} `json:"Currency,omitempty"`
}

func NewRamGoods() *RamGoods {
	return &RamGoods{new(RamGoodsValue), new(sync.RWMutex)}
}
func (r *RamGoods) Get() *RamGoodsValue {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.value
}
func (r *RamGoods) Set(value *RamGoodsValue) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.value = value
}

// /************************************************/
// type AppGoods struct {
// 	value *GoodsValue
// 	lock  *sync.RWMutex
// }

// func NewAppGoods() *AppGoods {
// 	return &AppGoods{new(GoodsValue), new(sync.RWMutex)}
// }
// func (r *AppGoods) Get() *GoodsValue {
// 	r.lock.RLock()
// 	defer r.lock.RUnlock()
// 	return r.value
// }
// func (r *AppGoods) Set(value *GoodsValue) {
// 	r.lock.Lock()
// 	defer r.lock.Unlock()
// 	r.value = value
// }

// /************************************************/
// type TradeGoods struct {
// 	value *GoodsValue
// 	lock  *sync.RWMutex
// }

// func NewTradeGoods() *TradeGoods {
// 	return &TradeGoods{new(GoodsValue), new(sync.RWMutex)}
// }
// func (r *TradeGoods) Get() *GoodsValue {
// 	r.lock.RLock()
// 	defer r.lock.RUnlock()
// 	return r.value
// }
// func (r *TradeGoods) Set(value *GoodsValue) {
// 	r.lock.Lock()
// 	defer r.lock.Unlock()
// 	r.value = value
// }
