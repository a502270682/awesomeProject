package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"runtime"
)

type Cent int64

func main() {
	//a := map[string]interface{}{
	//	"a": "b",
	//}
	//
	//tta, _ := json.Marshal(a)
	//
	//e := Event{
	//	UserId:     1,
	//	Properties: tta,
	//}
	//oriE, _ := json.Marshal(e)
	//var ret Event
	//err := json.Unmarshal(oriE, &ret)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ret)
	//_ = script_cmd.Command.Execute()
}

//base编码
func base64EncodeStr(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

//base解码
func base64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "error"
	}
	//b, err := base64.StdEncoding.DecodeString(string(a))
	//if err != nil {
	//	return "error"
	//}
	return string(a)
}

func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}

func Decrypt(encrypted string, key []byte) (string, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	src, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(src))
	var block cipher.Block
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, src)
	return string(decrypted), nil
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
