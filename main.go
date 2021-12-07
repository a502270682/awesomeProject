package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"runtime"
)

type Cent int64

func main() {
	getGongjijin()

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

func getGongjijin() {
	ratio := 0.9
	a := 214.0
	b := 480.0
	c := 3360.0
	sum := 0.0
	jiao := 35
	aTime := 16
	bTime := 5
	//cTime := jiao - aTime - bTime
	for i := jiao; i > 0; i-- {
		if i >= jiao-aTime {
			sum += a * ratio * float64(i)
		} else if i >= jiao-aTime-bTime {
			sum += b * ratio * float64(i)
		} else {
			sum += c * ratio * float64(i)
		}
	}
	fmt.Println("交的月数：", jiao, "贷款额：", sum)
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
