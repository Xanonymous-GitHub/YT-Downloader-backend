package api

import (
	"crypto/sha1"
	"fmt"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"io"
	"io/ioutil"
	"time"
)

func WriteToFile(b []byte) string {
	h := sha1.New()
	dt := time.Now().String()
	_, err := io.WriteString(h, dt)
	errorHandler.Handler("api.WriteToFile => _, err := io.WriteString(h, dt)", err)
	path := fmt.Sprintf("./%x.ytdr", h.Sum(nil))
	err = ioutil.WriteFile(path, b, 0644)
	errorHandler.Handler("api.WriteToFile => err = ioutil.WriteFile(path, b, 0644)", err)
	return path
}
