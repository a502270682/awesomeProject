package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

var (
	endpoint = ""
	accessKeyId = ""
	accessKeySecret = ""
	bucketName = ""
	filePath = ""
	fileName = ""
)

func upload() error{
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return errors.Wrap(err, "failed to new oss client")
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return errors.Wrap(err, "failed to get bucket")
	}
	fd, err := os.Open(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}
	defer fd.Close()
	err = bucket.PutObject(fmt.Sprintf("%s_%d_%s", bucketName, time.Now().Unix(), fileName), fd)
	if err != nil {
		return errors.Wrap(err, "failed to PutObject")
	}
	return nil
}

// 追加上传
func appendUpload() {
	// 创建OSSClient实例。
	client, err := oss.New("<yourEndpoint>", "<yourAccessKeyId>", "<yourAccessKeySecret>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("<yourBucketName>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	var nextPos int64 = 0
	// 第一次追加上传的位置是0，返回值为下一次追加的位置。后续追加的位置是追加前文件的长度。
	// <yourObjectName>填写不包含Bucket名称在内的Object的完整路径，例如example/test.txt。
	nextPos, err = bucket.AppendObject("<yourObjectName>", strings.NewReader("YourObjectAppendValue1"), nextPos)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 如果不是第一次追加上传，可以通过bucket.GetObjectDetailedMeta方法或上次追加返回值的X-Oss-Next-Append-Position的属性，获取追加位置。
	//props, err := bucket.GetObjectDetailedMeta("<yourObjectName>")
	//if err != nil {
	//    fmt.Println("Error:", err)
	//    os.Exit(-1)
	//}
	//nextPos, err = strconv.ParseInt(props.Get("X-Oss-Next-Append-Position"), 10, 64)
	//if err != nil {
	//    fmt.Println("Error:", err)
	//    os.Exit(-1)
	//}

	// 第二次追加上传。
	nextPos, err = bucket.AppendObject("<yourObjectName>", strings.NewReader("YourObjectAppendValue2"), nextPos)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 您可以进行多次追加上传操作。
}

func checkPointUpload() {
	// 创建OSSClient实例。
	client, err := oss.New("<yourEndpoint>", "<yourAccessKeyId>", "<yourAccessKeySecret>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("<yourBucketName>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 设置分片大小为100 KB，指定分片上传并发数为3，并开启断点续传上传。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"为filePath，100*1024为partSize。
	err = bucket.UploadFile("<yourObjectName>", "LocalFile", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

// 分片上传
// https://help.aliyun.com/document_detail/88604.html?spm=a2c4g.11186623.6.1391.5e885d88RYYf7u


