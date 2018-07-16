package cos

import (
	"os"
	"testing"
)

func TestMultiUseToken(t *testing.T) {
	appId := os.Getenv("appId")
	secretId := os.Getenv("secretId")
	secretKey := os.Getenv("secretKey")
	bucket := os.Getenv("bucket")
	token := CreateMultiUseAuthToken(appId, secretId, secretKey, bucket)
	t.Log(token)
}
