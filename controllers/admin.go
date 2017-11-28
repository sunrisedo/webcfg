package controllers

import (
	"fmt"
	"io"
	"os"

	"github.com/tealeg/xlsx"
)

// Create your own business code

type Admin struct {
	*Controller
}

//发送信息
func (c *Admin) Upload() {
	c.request.ParseForm()
	if c.request.Method == "GET" {
		c.ResultJson("failure", "Method is post.")
		return
	}
	file, handle, err := c.request.FormFile("file")
	defer file.Close()
	if err != nil {
		c.ResultJson("failure", err.Error())
		return
	}
	fileName := fmt.Sprintf("./%s/%s", c.cfg.Read("file", "uploadpath"), handle.Filename)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		c.ResultJson("failure", err.Error())
		return
	}
	io.Copy(f, file)
	xflie, err := xlsx.OpenFile(fileName)
	if err != nil {
		c.ResultJson("failure", err.Error())
		return
	}
	for _, sheet := range xflie.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text, _ := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
	c.ResultJson("success", nil)
}

func (c *Admin) Symbolinit() {

	var parent []interface{}
	for _, pname := range []string{"BTC", "ETH", "USDT"} {
		var child []interface{}
		for _, cname := range []string{"OMG", "BCC", "NEO"} {
			child = append(child, &Coin{cname, 2, 100, nil})
		}
		parent = append(parent, &Coin{Name: pname, List: child})
	}
	c.ResultJson("success", parent)
}
