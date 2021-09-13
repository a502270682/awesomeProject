package main

import (
	"fmt"
	"github.com/thoas/go-funk"
)

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}


func main(){
	f := &Foo{
		ID:        1,
		FirstName: "Gilles",
		LastName:  "Fabio",
		Age:       70,
	}

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       80,
	}

	results := []*Foo{f, b}

	mapping := funk.ToMap(results, "ID") // map[int]*Foo{1: f, 2: b}

	for key, this := range mapping.(map[int]*Foo){
		fmt.Println(key, ":", this)
	}

	// to map
	r := funk.Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	}) // []int{2, 4}
	fmt.Println(r)

	// reduce
	// Using operation runes. '+' and '*' only supported.
	fmt.Println(funk.Reduce([]int{1, 2, 3, 4}, '+', float64(0))) // 10
	fmt.Println(funk.Reduce([]int{1, 2, 3, 4}, '*', 1)) // 24

	// Using accumulator function
	fmt.Println(funk.Reduce([]int{1, 2, 3, 4}, func(acc float64, num int) float64 {
		return acc + float64(num)
	}, float64(0))) // 10

	fmt.Println(funk.Reduce([]int{1, 2, 3, 4}, func(acc string, num int) string {
		return acc + fmt.Sprint(num)
	}, "") ) // "1234"

	// 提取键值对中的值
	funk.Values(map[string]int{"one": 1, "two": 2}) // []string{1, 2} (iteration order is not guaranteed)

	foo := &Foo{
		ID:        1,
		FirstName: "Dark",
		LastName:  "Vador",
		Age:       30,
	}

	funk.Values(foo) // []interface{}{1, "Dark", "Vador", 30} (iteration order is not guaranteed)
}