package main

import (
	"net/http"
	"reflect"
	"strings"
	"sunrise/webcfg/controllers"
)

//Configure the routing
var RouteMap = map[string]func(http.ResponseWriter, *http.Request){
	"/admin/": AdminRoute,
}


func AdminRoute(w http.ResponseWriter, r *http.Request) {
	client := controllers.NewController(w, r, cfg)
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		client.Error()
		return
	}
	controller := reflect.ValueOf(&controllers.Admin{client})
	method := controller.MethodByName(strings.Title(parts[1]))
	if !method.IsValid() {
		client.Error()
		return
	}

	method.Call(nil)
}
