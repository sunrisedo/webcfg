package datas

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hugozhu/godingtalk"
	"github.com/tealeg/xlsx"
)

func NewDingConn(corpid, corpsecret string) {
	ConnDing = godingtalk.NewDingTalkClient(corpid, corpsecret)
}

func NewGoodsData(fileName string) error {
	xflie, err := xlsx.OpenFile(fmt.Sprintf("./%s/2800_GOODS.xlsx", fileName))
	if err != nil {
		return err
	}
	if len(xflie.Sheets) < 2 {
		return errors.New("Xlsx lack sheet.")
	}

	prodSheet := xflie.Sheets[0]
	currSheet := xflie.Sheets[1]
	if len(prodSheet.Rows) < 2 {
		return errors.New("Product lack row.")
	}
	if len(currSheet.Rows) < 2 {
		return errors.New("Currency lack row.")
	}

	prodRows := prodSheet.Rows[1:]
	currRows := currSheet.Rows[1:]
	goods := NewGoodsValue()
	for _, row := range prodRows {
		rid, err := row.Cells[0].String()
		if err == nil && strings.TrimSpace(rid) != "" {
			mid, _ := row.Cells[1].String()
			priceFixed, _ := row.Cells[2].String()
			volFixed, _ := row.Cells[3].String()
			id, _ := row.Cells[4].String()
			mPrecise, _ := row.Cells[5].String()
			rPrecise, _ := row.Cells[6].String()
			munlimited, _ := row.Cells[7].String()
			runlimited, _ := row.Cells[8].String()
			soldUnits, _ := row.Cells[9].String()
			buyCommission, _ := row.Cells[10].String()
			sellCommission, _ := row.Cells[11].String()
			side, _ := row.Cells[12].String()
			method, _ := row.Cells[13].String()
			closed, _ := row.Cells[14].String()
			unitPrice, _ := row.Cells[15].String()

			var obj Product
			obj.RID = rid
			obj.MID = mid
			obj.Coin = obj.RID
			obj.Code = obj.RID + obj.MID
			obj.PriceFixed, _ = strconv.Atoi(priceFixed)
			obj.VolFixed, _ = strconv.Atoi(volFixed)
			obj.ID, _ = strconv.Atoi(id)
			obj.MPrecise, _ = strconv.Atoi(mPrecise)
			obj.RPrecise, _ = strconv.Atoi(rPrecise)
			obj.Munlimited = StringToFloat(munlimited)
			obj.Runlimited = StringToFloat(runlimited)
			obj.SoldUnits = StringToFloat(soldUnits)
			obj.BuyCommission, _ = strconv.Atoi(buyCommission)
			obj.SellCommission, _ = strconv.Atoi(sellCommission)
			obj.Side = side
			obj.Method = method
			obj.Closed = closed
			obj.UnitPrice = StringToFloat(unitPrice)
			goods.Prods.PushBack(obj)
		}
	}

	for _, row := range currRows {
		code, err := row.Cells[0].String()
		if err == nil && strings.TrimSpace(code) != "" {
			value, _ := row.Cells[1].String()
			symbol, _ := row.Cells[2].String()
			minTrade, _ := row.Cells[3].String()
			minCash, _ := row.Cells[4].String()
			cashFree, _ := row.Cells[5].String()

			var obj Currency
			obj.Code = code
			obj.Value, _ = strconv.Atoi(value)
			obj.Symbol = symbol
			obj.MinTrade = StringToFloat(minTrade)
			obj.MinCash = StringToFloat(minCash)
			obj.CashFree = StringToFloat(cashFree)
			goods.Currs.PushBack(obj)
		}
	}

	RGoods.Set(goods)
	goods.SaveWebGoods()
	goods.SaveAppGoods()
	goods.SaveTradeGoods()
	return nil
}

type GoodsValue struct {
	Prods *list.List
	Currs *list.List
}

func NewGoodsValue() *GoodsValue {
	return &GoodsValue{list.New(), list.New()}
}
func (goods *GoodsValue) SaveWebGoods() {
	product := make(map[string]*ProdSimple)
	market := make(map[string]*MarketSimple)
	var markcnt, prodcnt int
	var old string
	for e := goods.Prods.Front(); e != nil; e = e.Next() {
		v := e.Value.(Product)
		prodcnt++
		var prod ProdSimple
		prod.Sort = prodcnt
		prod.Code = v.RID + v.MID
		prod.Name = v.RID + "/" + v.MID
		prod.PriceFixed = v.PriceFixed
		prod.VolFixed = v.VolFixed
		product[prod.Code] = &prod

		if v.MID != old {
			markcnt++
			var mark MarketSimple
			mark.Sort = markcnt
			mark.Code = v.MID
			mark.Name = v.MID
			market[mark.Code] = &mark
			old = v.MID
		}
	}

	currency := make(map[string]*Currency)
	for e := goods.Currs.Front(); e != nil; e = e.Next() {
		v := e.Value.(Currency)
		currency[v.Code] = &v
	}
	RWebGoods.Set(&RamGoodsValue{market, product, currency})
}
func (goods *GoodsValue) SaveAppGoods() {

	cnt := -1
	var old string
	var market []*MarketSimple
	for e := goods.Prods.Front(); e != nil; e = e.Next() {
		v := e.Value.(Product)
		if v.MID != old {
			var mark MarketSimple
			mark.Code = v.MID
			market = append(market, &mark)
			old = v.MID
			cnt++
		}
		var prod ProdSimple
		prod.Code = v.RID + v.MID
		prod.Name = v.RID + "/" + v.MID
		prod.PriceFixed = v.PriceFixed
		prod.VolFixed = v.VolFixed
		market[cnt].Product = append(market[cnt].Product, prod)
	}

	// var currency []*Currency
	// for e := goods.Currs.Front(); e != nil; e = e.Next() {
	// 	v := e.Value.(Currency)
	// 	currency = append(currency, &v)
	// }

	currency := make(map[string]*Currency)
	for e := goods.Currs.Front(); e != nil; e = e.Next() {
		v := e.Value.(Currency)
		currency[v.Code] = &v
	}
	RAppGoods.Set(&RamGoodsValue{market, nil, currency})
}
func (goods *GoodsValue) SaveTradeGoods() {

	product := make(map[string]*Product)
	for e := goods.Prods.Front(); e != nil; e = e.Next() {
		v := e.Value.(Product)
		product[v.Code] = &v
	}

	currency := make(map[string]int)
	for e := goods.Currs.Front(); e != nil; e = e.Next() {
		v := e.Value.(Currency)
		currency[v.Code] = v.Value
	}
	RTradeGoods.Set(&RamGoodsValue{nil, product, currency})
}
