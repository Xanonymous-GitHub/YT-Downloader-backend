package main

import (
	"flag"
	"fmt"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/api"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/converter"
)

func startServices(id string) {
	defaultURL := "https://youtube.com/get_video_info"
	header := api.Header{
		Host:      "www.youtube.com",
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36",
	}
	queries := map[string]string{"video_id": id}
	resp := api.Request(defaultURL, queries, "GET", header)
	path := api.WriteToFile(resp)
	preResult := converter.HttpHexNumberToSimpleText(path)
	result := converter.DecodeUTF16(preResult)
	fmt.Printf("%s", string(result))
	api.Recycle(path)
}

func main() {
	flag.Parse()
	startServices(flag.Arg(0))
}
