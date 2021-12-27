package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

type Sss struct {
	Auto int
}
func main() {
	group, _ := errgroup.WithContext(context.Background())
	ss := make([]*Sss, 0)
	ss = append(ss, &Sss{
		Auto: 1,
	})
	ss = append(ss, &Sss{
		Auto: 2,
	})
	for idx := range ss {
		index := idx
		group.Go(func() error {
			fmt.Printf("start to execute the %d gorouting\n", ss[index].Auto)
			time.Sleep(time.Duration(ss[index].Auto) * time.Second)
			//if index%2 == 0 {
			//	return fmt.Errorf("something has failed on grouting:%d", index)
			//}
			ss[index].Auto = index
			fmt.Printf("gorouting:%d end\n", ss[idx])
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
	for _, s := range ss {
		fmt.Println(s)
	}
}
