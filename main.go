package main

import (
	"fmt"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/converter"
)

func main() {
	result := converter.HttpHexNumberToSimpleText("/Users/xanonymous/Downloads/yt_info")
	fmt.Printf("%s", result)
}
