package tips

import (
	"fmt"
	"testing"
)

func NextNumber() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func TestBiBao(t *testing.T) {
	next := NextNumber()
	next()
	next()
	next()

	next1 := NextNumber()
	next1()
	NextNumber() //不会输出i
}

