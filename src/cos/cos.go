package cos

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//@doc http://www.qcloud.com/doc/product/227/%E7%9B%AE%E5%BD%95%E5%88%97%E8%A1%A8
const (
	//appId, bucket, dir ,prefix
	LIST_BUCKET_API_URL = "http://web.file.myqcloud.com/files/v1/%s/%s%s%s"
)

const (
	//int
	MAX_NUM         = 199
	MAX_RETRY_TIMES = 10
)

const (
	PATTERN_BOTH = "eListBoth"
	PATTERN_DIR  = "eListDirOnly"
	PATTERN_FILE = "eListFileOnly"
)

type ListResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    ListItem `json:"data"`
}

type ListItem struct {
	Context   string         `json:"context"`
	HasMore   bool           `json:"has_more"`
	DirCount  int            `json:"dircount"`
	FileCount int            `json:"filecount"`
	Infos     []ListItemInfo `json:"infos"`
}

type ListItemInfo struct {
	Name      string `json:"name"`
	BizAttr   string `json:"bit_attr,omitempty"`
	FileSize  int    `json:"filesize,omitempty"`
	FileLen   int    `json:"filelen,omitempty"`
	Sha       string `json:"sha,omitempty"`
	Ctime     string `json:"ctime"`
	Mtime     string `json:"mtime"`
	AccessUrl string `json:"access_url,omitempty"`
	SrcUrl    string `json:"source_url,omitempty"`
}

func ListBucket(appId, bucket, secretId, secretKey, dir, prefix string, bWriter *bufio.Writer) {
	if dir != "" {
		if !strings.HasPrefix(dir, "/") {
			dir = fmt.Sprintf("/%s", dir)
		}

		if !strings.HasSuffix(dir, "/") {
			dir = fmt.Sprintf("%s/", dir)
		}
	}

	fmt.Printf("Listing dir of `%s` ...\n", dir)

	var apiUrlBase = fmt.Sprintf(LIST_BUCKET_API_URL, appId, bucket, dir, prefix)
	var respErr error
	var resp *http.Response
	var context string

	for {
		apiUrl := fmt.Sprintf("%s?op=list&num=%d&context=%s&pattern=%s&order=%d",
			apiUrlBase, MAX_NUM, context, PATTERN_BOTH, 0)
		req, reqErr := http.NewRequest("GET", apiUrl, nil)
		if reqErr != nil {
			fmt.Println("New request error,", reqErr)
			break
		}

		token := CreateMultiUseAuthToken(appId, secretId, secretKey, bucket)
		req.Header.Add("Authorization", token)

		for i := 0; i < MAX_RETRY_TIMES; i++ {
			if i != 0 {
				fmt.Println("Retrying times", i)
			}

			resp, respErr = http.DefaultClient.Do(req)
			if respErr != nil {
				fmt.Println("Fire request error,", respErr)
				continue
			}

			decoder := json.NewDecoder(resp.Body)
			listResp := ListResponse{}
			if decodeErr := decoder.Decode(&listResp); decodeErr != nil {
				fmt.Println("Parse response error,", decodeErr)
				time.After(time.Second * 1)
				continue
			}
			resp.Body.Close()

			for _, item := range listResp.Data.Infos {
				//write item
				if item.SrcUrl != "" {
					//it is a file
					data := fmt.Sprintf("%s\t%s\t%d", item.SrcUrl, item.AccessUrl, item.FileSize)
					bWriter.WriteString(data + "\n")
				} else {
					//it is a dir
					newDir := fmt.Sprintf("%s%s", dir, item.Name)
					ListBucket(appId, bucket, secretId, secretKey, newDir, "", bWriter)
				}
			}

			if listResp.Data.HasMore {
				context = listResp.Data.Context
			} else {
				fmt.Printf("All list of `%s` done!\n", dir)
				return
			}

			break
		}

		if respErr != nil {
			fmt.Printf("Last request of `%s` faild too many times, quit!\n", dir)
			break
		}
	}
}
