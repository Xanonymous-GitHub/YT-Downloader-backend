package main

import (
	"flag"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/api"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/converter"
	"log"
)

func startServices(id string) {
	defaultURL := "https://youtube.com/get_video_info"
	header := api.Header{
		Host:      "www.youtube.com",
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36",
	}
	queries := map[string]string{"video_id": id}
	resp := api.Request(defaultURL, queries, "GET", header)
	result := converter.HttpHexNumberToSimpleText(resp)
	result = converter.HttpHexNumberToSimpleText(result)
	result = converter.DecodeUTF16(result)
	log.Println(converter.MakeUrlList(string(result)))
}

func main() {
	flag.Parse()
	startServices(flag.Arg(0))
}
