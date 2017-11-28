package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sunrisedo/conf"
)

// Control public interface
type Controller struct {
	response http.ResponseWriter
	request  *http.Request
	cfg      *conf.Config
}

type Result struct {
	Status string      `json:"status,omitempty"` //success | failure
	Data   interface{} `json:"data,omitempty"`
}

func NewController(w http.ResponseWriter, r *http.Request, c *conf.Config) *Controller {
	return &Controller{w, r, c}
}

func (c *Controller) RequestStruct(i interface{}) {
	data, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		log.Printf("read body error:%v", err)
		return
	}

	if len(data) == 0 {
		log.Println("data is nil.")
		return
	}

	if err := json.Unmarshal(data, i); err != nil {
		log.Printf("json to struct error:%v", err)
	}
}

func (c *Controller) ResultJson(status string, i interface{}) {
	b, err := json.Marshal(&Result{status, i})
	if err != nil {
		log.Printf("result to json error:%v", err)
		return
	}
	c.response.Write(b)
}

func (c *Controller) Error() {
	c.response.Write([]byte("404 page not found"))
}
