package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func md5file() {
	path := "../../static/version"
	testFile := path + "/IHaoyueManager_2015061502.apk"
	log.Println(testFile)
	file, inerr := os.Open(testFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		fmt.Printf("%x", md5h.Sum([]byte(""))) //md5
	}
}
