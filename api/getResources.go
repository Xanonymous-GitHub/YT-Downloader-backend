package api

import (
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"io/ioutil"
	"net/http"
)

type Header struct {
	Host      string
	UserAgent string
}

func Request(url string, queryId map[string]string, method string, header Header) []byte {
	req, err := http.NewRequest(method, url, nil)
	errorHandler.Handler("api.Request => req, err := http.NewRequest(method, url, nil)", err)
	req.Header.Set("Host", header.Host)
	req.Header.Set("User-Agent", header.UserAgent)
	query := req.URL.Query()
	for key, value := range queryId {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp == nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("http resp error!\n")
	}
	errorHandler.Handler("api.Request => resp, err := client.Do(req)", err)
	body, err := ioutil.ReadAll(resp.Body)
	errorHandler.Handler("api.Request => body, err := ioutil.ReadAll(resp.Body)", err)
	return body
}
