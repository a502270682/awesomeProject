package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
	}
}

func main(){
	res, _ := Get("http://www.baidu.com/")
	fmt.Println(string(res))
}

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
		Timeout: 3*time.Second,
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