package main

// var obj = map[string]map[string]ChildInfo{
// 	"BTC": map[string]ChildInfo{
// 		"ETH": ChildInfo{2, 100},
// 		"BCC": ChildInfo{2, 100},
// 		"NEO": ChildInfo{2, 100},
// 	},
// 	"ETH": map[string]ChildInfo{
// 		"BCC": ChildInfo{2, 100},
// 		"NEO": ChildInfo{2, 100},
// 	},
// }
type ChildInfo struct {
	Prec       int
	DepositMax int
}

type Coin struct {
	Name       string      `json:"Name,omitempty"`
	Prec       int         `json:"Prec,omitempty"`
	DepositMax int         `json:"DepositMax,omitempty"`
	List       interface{} `json:"List,omitempty"`
}
