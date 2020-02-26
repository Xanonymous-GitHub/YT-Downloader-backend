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
	prePath := api.WriteToFile(resp)
	prePreResult := converter.HttpHexNumberToSimpleText(prePath)
	api.Recycle(prePath)
	path := api.WriteToFile(prePreResult)
	preResult := converter.HttpHexNumberToSimpleText(path)
	result := converter.DecodeUTF16(preResult)
	//fmt.Printf("%s", string(result))
	log.Println(converter.MakeUrlList(string(result)))
	api.Recycle(path)
}

func main() {
	flag.Parse()
	startServices(flag.Arg(0))
}
