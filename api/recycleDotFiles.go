package api

import (
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"os"
)

func Recycle(path string) {
	err := os.Remove(path)
	errorHandler.Handler("api.recycle => err := os.Remove(path)", err)
}
