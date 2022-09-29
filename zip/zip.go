package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetZip(rw io.Writer, targetFile []string) error {
	//在流中创建一个 zipwriter
	zipwriter := zip.NewWriter(rw)
	//关闭zipwriter
	defer zipwriter.Close()
	//循环写入图片
	for idx, f := range targetFile {
		iowriter, err := zipwriter.Create(fmt.Sprintf("%d.jpg", idx))
		if err != nil {
			return err
		}
		var content []byte
		resp, err := http.Get(f)
		if err == nil {
			content, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				content = []byte("")
			}
			resp.Body.Close()
		}
		iowriter.Write(content)
	}
	return nil
}
