package main

import "fmt"

// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
type MyMap[K int | string, V float32 | float64] map[K]V

// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
var a MyMap[string, float64] = map[string]float64{
	"jack_score": 9.6,
	"bob_score":  8.4,
}

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

// 去除歧义
type NewType[T interface{ *int }] []T
type NewType2[T interface{ *int | *float64 }] []T

// ～：可以使用新定义的来实例化底层类型
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float interface {
	~float32 | ~float64
}

type Slice[T Int | Uint | Float] []T

var s Slice[int] // 正确

type MyInt int

var s2 Slice[MyInt] // MyInt底层类型是int，所以可以用于实例化

type MyMyInt MyInt

var s3 Slice[MyMyInt] // 正确。MyMyInt 虽然基于 MyInt ，但底层类型也是int，所以也能用于实例化

type MyFloat32 float32 // 正确
var s4 Slice[MyFloat32]

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
