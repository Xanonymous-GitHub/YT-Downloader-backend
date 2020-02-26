package converter

import "strings"

func MakeUrlList(videoInfo string) (urlList []string) {
	for getPosition(videoInfo) != -1 {
		videoUrl, changedVideoInfo := wholeUrl(getPosition(videoInfo), videoInfo)
		urlList = append(urlList, videoUrl)
		videoInfo = changedVideoInfo
	}
	return
}

func getPosition(videoInfo string) int {
	return strings.Index(videoInfo, "videoplayback")
}

func wholeUrl(position int, videoInfo string) (videoUrl string, changedVideoInfo string) {
	videoUrl += string(videoInfo[position])
	left, right := position, position
	for ; videoInfo[left] != '"'; left -= 1 {
	}
	for ; videoInfo[right] != '"'; right += 1 {
	}
	videoUrl = videoInfo[left+1 : right]
	changedVideoInfo = videoInfo[right+1:]
	return
}
