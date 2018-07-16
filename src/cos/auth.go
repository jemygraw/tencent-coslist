package cos

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

//@doc http://www.qcloud.com/doc/product/227/%E7%AD%BE%E5%90%8D%E7%AE%97%E6%B3%95

func CreateMultiUseAuthToken(appId, secretId, secretKey, bucket string) (token string) {
	now := time.Now()
	ts := now.Unix()
	expire := now.Add(time.Second * 1800).Unix()
	rand := rand.Intn(1000000000)
	pattern := fmt.Sprintf("a=%s&k=%s&e=%d&t=%d&r=%d&f=&b=%s", appId, secretId, expire, ts, rand, bucket)

	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(pattern))
	signTmp := h.Sum(nil)

	signTmp = append(signTmp, []byte(pattern)...)
	token = base64.StdEncoding.EncodeToString(signTmp)
	return
}
