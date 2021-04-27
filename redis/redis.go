package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

func init() {
	pool = NewPool("localhost:6379")
}
var (
	pool *redis.Pool
)

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func main() {
	conn:= pool.Get()
	v,err:=redis.Strings(conn.Do("keys","mini_data_switch*"))
	if err!=nil{
		println(err.Error())
		return
	}
	onShopIds := make([]string, 0)
	for _, this := range v {
		turn, err := redis.String(conn.Do("get", this))
		if err != nil {
			fmt.Println(err)
			continue
		}
		if turn == "on" {
			onShopIds = append(onShopIds, strings.TrimPrefix(this, "mini_data_switch:"))
		}
	}
	fmt.Println(onShopIds)
}



