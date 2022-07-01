package main

import "fmt"

func SumIntsOrFloats[V int64 | float64 | string](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}
func main() {
	tt := MySlice[int]{12, 123}
	fmt.Println(tt.Sum())
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[int64]([]int64{1, 2, 3}),
		SumIntsOrFloats[float64]([]float64{0.1, 0.2, 0.3}))
}
