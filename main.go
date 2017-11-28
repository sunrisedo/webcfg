package main

import (
	"log"
	"net/http"
	"runtime"

	"github.com/sunrisedo/conf"
)

var (
	cfg = conf.NewConfig("init.conf")
)

func init() {
	// 初始化配置文件
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("init data start...")
	CreateDir(cfg.Read("file", "uploadpath"))
	log.Println("init data finish.")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//注册路由
	log.Println("init route start...")
	for addr, controller := range RouteMap {
		http.HandleFunc(addr, controller)
	}
	log.Println("init route finish.")

	log.Println("listen port", cfg.Read("server", "port"))
	http.Handle("/", http.FileServer(http.Dir("webroot")))
	log.Println(http.ListenAndServe(cfg.Read("server", "port"), nil))
}

// func main() {
// var server *http.Server = &http.Server{
// 	Addr:           ":8080",
// 	Handler:        &customHandler{},
// 	ReadTimeout:    10 * time.Second,
// 	WriteTimeout:   10 * time.Second,
// 	MaxHeaderBytes: 1 << 20,
// }
// server.ListenAndServe()

// select {}
// }

// type customHandler struct {
// }

// func (cb *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("customHandler!!")
// 	w.Write([]byte("customHandler!!"))
// }

// func (cb *customHandler) Test(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Test!!")
// 	w.Write([]byte("Test!!"))
// }
