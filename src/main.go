package main

import (
	"bufio"
	"cos"
	"flag"
	"fmt"
	"os"
)

func main() {
	var appId string
	var bucket string
	var secretId string
	var secretKey string
	var result string
	var dir string
	var prefix string

	flag.StringVar(&appId, "appId", "", "app id")
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.StringVar(&secretId, "secretId", "", "secret id")
	flag.StringVar(&secretKey, "secretKey", "", "secret key")
	flag.StringVar(&dir, "dir", "/", "dir")
	flag.StringVar(&prefix, "prefix", "", "prefix")
	flag.StringVar(&result, "result", "", "result file")

	flag.Parse()

	if appId == "" {
		fmt.Println("Err: no appId specified")
		return
	}

	if bucket == "" {
		fmt.Println("Err: no bucket specified")
		return
	}

	if secretId == "" {
		fmt.Println("Err: no secretId specified")
		return
	}

	if secretKey == "" {
		fmt.Println("Err: no secretKey specified")
		return
	}

	resultFp, openErr := os.Create(result)
	if openErr != nil {
		fmt.Println("Open result file error,", openErr)
		return
	}
	defer resultFp.Close()
	bWriter := bufio.NewWriter(resultFp)
	defer bWriter.Flush()

	cos.ListBucket(appId, bucket, secretId, secretKey, dir, prefix, bWriter)
}
