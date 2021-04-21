package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	//go func() {
	//	for i:=0; i<1000; i ++ {
	//		fmt.Println(i)
	//		time.Sleep(100*time.Millisecond)
	//	}
	//}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}
