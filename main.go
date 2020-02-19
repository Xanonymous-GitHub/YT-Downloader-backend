package main

import (
	"fmt"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/converter"
)

func main() {
	preResult := converter.HttpHexNumberToSimpleText("/Users/xanonymous/Downloads/yt_info")
	result := converter.DecodeUTF16(preResult)
	fmt.Printf("%s", string(result))
}
