package test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
基准测试
go test -bench=. -benchmem -run=none -benchtime=3s

参考资料：https://juejin.cn/post/6963919796115079176
 */

func BenchmarkSprintf(b *testing.B) {
	num:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		fmt.Sprintf("%d",num)
	}
}

func BenchmarkFormat(b *testing.B){
	num:=int64(10)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		strconv.FormatInt(num,10)
	}
}

func BenchmarkItoa(b *testing.B){
	num:=10
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		strconv.Itoa(num)
	}
}
