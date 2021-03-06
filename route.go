package main

import (
	"net/http"
	"reflect"
	"strings"
	"sunrise/webcfg/controllers"
)

// var RouteMap = map[string]interface{}{
// 	"/server/": &controllers.Server{&controllers.Controller{}},
// 	"/alert/":  &controllers.Alert{&controllers.Controller{}},
// }

// func NewRoute() map[string]func(http.ResponseWriter, *http.Request) {
// 	route := make(map[string]func(http.ResponseWriter, *http.Request))
// 	for addr, inter := range RouteMap {
// 		// route[addr] = createRoute(inter)
// 		obj := reflect.ValueOf(inter)

// 		log.Println(addr, obj.InterfaceData())

// 	}
// 	return route
// }

// func createRoute(route interface{}) func(http.ResponseWriter, *http.Request) {
// 	// switch () {
// 	// 	case
// 	// }
// 	routeFunc := func(w http.ResponseWriter, r *http.Request) {
// 		client := controllers.NewController(w, r, cfg)
// 		route.(Controller) = client
// 		url := strings.Trim(r.URL.Path, "/")
// 		parts := strings.Split(url, "/")
// 		inMethod := strings.Title(url)
// 		if len(parts) >= 2 {
// 			inMethod = strings.Title(parts[1])
// 		}

// 		method := reflect.ValueOf({client}).MethodByName(inMethod)
// 		if !method.IsValid() {
// 			client.Error()
// 			return
// 		}
// 		method.Call(nil)
// 	}
// 	return routeFunc
// }

// Configure the routing
var RouteMap = map[string]func(http.ResponseWriter, *http.Request){
	"/server/": ServerRoute,
	"/alert/":  AlertRoute,
}

func ServerRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")             //返回数据格式是json

	client := controllers.NewController(w, r, cfg)
	url := strings.Trim(r.URL.Path, "/")
	// log.Println("url", url)
	parts := strings.Split(url, "/")
	inMethod := strings.Title(url)
	if len(parts) >= 2 {
		inMethod = strings.Title(parts[1])
	}

	controller := reflect.ValueOf(&controllers.Server{Controller: client})
	method := controller.MethodByName(inMethod)
	if !method.IsValid() {
		client.Error()
		return
	}

	method.Call(nil)
}

func AlertRoute(w http.ResponseWriter, r *http.Request) {
	client := controllers.NewController(w, r, cfg)
	url := strings.Trim(r.URL.Path, "/")
	// log.Println("url", url)
	parts := strings.Split(url, "/")
	inMethod := strings.Title(url)
	if len(parts) >= 2 {
		inMethod = strings.Title(parts[1])
	}

	controller := reflect.ValueOf(&controllers.Alert{Controller: client})
	method := controller.MethodByName(inMethod)
	if !method.IsValid() {
		client.Error()
		return
	}

	method.Call(nil)
}
