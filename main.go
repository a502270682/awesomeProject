package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
	}
}

func test(a Input) *Input {

	fmt.Printf("%x\n", &a)
	a.E = 4
	return &a
}

func main() {
	//res, _ := Get("http://www.baidu.com/")
	//fmt.Println(config_util.GetAllConfig("test"))

}

func getmap() map[int64][]*string {
	var ret map[int64][]*string
	fmt.Println(ret[1])
	return ret
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func fi(A Input) {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			fmt.Println(err)
		}
	}()
	se(A)
}

func se(A Input) {
	fmt.Println(A.C.NextEvent)
}

/*
	ff := func(ctx context.Context,input *Input) (*Output, error){
		return nil, nil
	}
	tt := reflect.TypeOf(ff)
	a := &Input{map[string]string{}, 2}
*/

type Input struct {
	Ctx map[string]string
	E   int
	C   *Output
}

type Output struct {
	NextEvent    string
	TemplateData map[string]string
}

type Action func(ctx context.Context, input *Input) (*Output, error)

func Get(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, _ := http.NewRequest("GET", url, body)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil || res == nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
