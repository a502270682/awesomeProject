package decorator

import (
	"io"
	"log"
	"net/http"
	"time"
)

/*
装饰器模式
属于：结构型模式
优点
可以通过一种动态的方式来扩展一个对象的功能。
可以使用多个具体装饰类来装饰同一对象，增加其功能。
具体组件类与具体装饰类可以独立变化，符合"开闭原则"。
缺点：
对于多次装饰的对象，易于出错，排错也很困难。
对于产生很多具体装饰类，增加系统的复杂度以及理解成本。
适合场景：
需要给一个对象增加功能，这些功能可以动态地撤销。
需要给一批兄弟类增加或者改装功能。
*/

// Step1: 实现一个HTTP Server
// Step2: 实现一个HTTP Handler
// Step3: 实现中间件的功能：
//    (1)、实现HTTP中间件记录请求的URL、方法。
//    (2)、实现HTTP中间件记录请求的网络地址。
//    (3)、实现HTTP中间件记录请求的耗时。

// 记录请求的URL和方法
func method(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("实现HTTP中间件记录请求的URL和方法：%s, %s", req.URL, req.Method)
		next.ServeHTTP(resp, req)
	}
}

// 记录请求的网络地址
func address(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("实现HTTP中间件记录请求的网络地址：%s", req.RemoteAddr)
		next.ServeHTTP(resp, req)
	}
}

// 记录请求的耗时
func cost(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(resp, req)
		duration := time.Since(start)
		log.Printf("实现HTTP中间件记录请求的耗时: %v", duration)
	}
}

func HelloHandler(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "hello world")
}
