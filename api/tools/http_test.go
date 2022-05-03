package tools

import (
	"fmt"
	"testing"
)

func TestHttpDownloadFile(t *testing.T) {
	fmt.Println(HttpDownloadFile("https://cdn.xiaoyou66.com/image/live2d.zip", "avatar.zip"))
}
