package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sunrise/webcfg/datas"
	"time"
)

// Create your own business code

type Server struct {
	*Controller
}
type SymbolinitAsk struct {
	Device string `json:"Device"` //WEB|APP|TRADE|MARKET|ROBOT|DEPTH
}

func (c *Server) Products() {
	c.ResultPage("index")
}

func (c *Server) Login() {

	if c.request.FormValue("acc") != "Admin" || c.request.FormValue("pwd") != "Admin12345" {
		c.ResultJson(101, "Acc or Pwd error.")
		return
	}
	user := &http.Cookie{
		Name:    "UID",
		Value:   "1",
		Expires: time.Now().Add(20 * time.Minute),
	}

	http.SetCookie(c.response, user)
	c.ResultPage("upload")
}

//上传文件
func (c *Server) Upload() {

	if v, err := c.request.Cookie("UID"); err != nil || v == nil {
		c.ResultJson(102, "Please login in.")
		return
	}

	if c.request.Method == "GET" {
		c.ResultPage("upload", fmt.Sprintf("Error:%d Method is post.", 201))
		return
	}
	if err := c.request.ParseForm(); err != nil {
		c.ResultPage("upload", fmt.Sprintf("Error:520 %s", err.Error()))
		return
	}

	if c.request.ContentLength < 300 {
		c.ResultPage("upload", fmt.Sprintf("Error:%d Can not find file.", 202))
		return
	}
	file, handle, err := c.request.FormFile("file")
	defer file.Close()
	if err != nil {
		c.ResultPage("upload", 520, err.Error())
		return
	}
	fileName := fmt.Sprintf("./%s/%s", c.conf.Read("file", "uploadpath"), handle.Filename)
	os.Remove(fileName)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		c.ResultPage("upload", fmt.Sprintf("Error:520 %s", err.Error()))
		return
	}
	io.Copy(f, file)
	switch handle.Filename {
	case "2800_GOODS.xlsx":
		if err := datas.NewGoodsData(c.conf.Read("file", "uploadpath")); err != nil {
			c.ResultPage("upload", fmt.Sprintf("Error:520 %s", err.Error()))
			return
		}
	}

	c.ResultPage("upload", time.Now().Format("2006-01-02 15:04:05")+" SUCCEE")
}

//获取配置
func (c *Server) Symbolinit() {
	var ask SymbolinitAsk
	if query := c.RequestStruct(&ask); query != nil {
		ask.Device = query.Get("Device")
	}
	//WEB|APP|TRADE|MARKET|ROBOT|DEPTH
	switch ask.Device {
	case "WEB":
		if data := datas.RWebGoods.Get(); data != nil {
			c.ResultJson(0, data)
		}
		return
	case "APP":
		if data := datas.RAppGoods.Get(); data != nil {
			c.ResultJson(0, data)
		}
		return
	case "TRADE":
		if data := datas.RTradeGoods.Get(); data != nil {
			c.ResultJson(0, data)
		}
		return
	case "TEST":
		if data := datas.RTradeGoods.Get(); data != nil {
			c.ResultPage("upload", data)
		}
		return
	}

	// log.Println("goods", &datas.InitWeb{product, market, currency})
	c.ResultJson(520, "Init data is nil.")
}
