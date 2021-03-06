package converter

import (
	"encoding/json"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"log"
	"strings"
)

func getPosition(videoInfo string, searchBy string) int {
	return strings.Index(videoInfo, searchBy)
}

func MakeLists(videoInfo string, keyWord string) (urlList []string, propList []map[string]interface{}, validRequest bool) {
	validRequest = false
	for getPosition(videoInfo, keyWord) != -1 {
		validRequest = true
		videoUrl, changedVideoInfo := wholeUrl(getPosition(videoInfo, keyWord), videoInfo)
		urlList = append(urlList, videoUrl)
		videoInfo = changedVideoInfo
		prop, changedVideoInfo := properties(0, videoInfo)
		videoProps := make(map[string]interface{})
		err := json.Unmarshal([]byte(prop), &videoProps)
		if err != nil {
			log.Println(prop)
		}
		errorHandler.Handler("converter.MakeLists => err := json.Unmarshal([]byte(prop), videoProps)", err)
		propList = append(propList, videoProps)
		videoInfo = changedVideoInfo
	}
	return
}

func wholeUrl(position int, videoInfo string) (videoUrl string, changedVideoInfo string) {
	left, right := position, position
	for ; videoInfo[left] != '"'; left -= 1 {
	}
	for ; videoInfo[right] != '"'; right += 1 {
	}
	videoUrl = videoInfo[left+1 : right]
	changedVideoInfo = videoInfo[right+1:]
	return
}

func properties(position int, videoInfo string) (prop string, changedVideoInfo string) {
	head := position + 1
	ql := struct {
		tmpStorage []int //save the pos to be added to storage.
		qStatus    bool  //true:already in a quotation mark scope; false: out of the scope.
		active     bool  //true:record the byte; false: stop record.
	}{[]int{}, false, false}
	for ; videoInfo[position] != ']'; position++ {
		if position != 0 && videoInfo[position] == ',' && videoInfo[position-1] == '}' && videoInfo[position+1] == '{' {
			break
		}
		if ql.active {
			if videoInfo[position] == '"' {
				if position+1 < len(videoInfo) && findWords([]string{",", "}", ":"}, string(videoInfo[position+1])) {
					ql.active = !ql.active
					for _, k := range ql.tmpStorage {
						videoInfo = videoInfo[:k] + "'" + videoInfo[k+1:]
					}
					ql.tmpStorage = []int{}
					continue
				}
				ql.tmpStorage = append(ql.tmpStorage, position)
			}
		} else if videoInfo[position] == '"' && (findWords([]string{"{", ":", ","}, string(videoInfo[position-1])) || (position == 1)) {
			ql.active = !ql.active
		}
	}
	prop = "{" + videoInfo[head:position]
	v := struct {
		active       bool
		hasNotNumber bool
		headPos      int
	}{false, false, 0}
	for i, k := range prop {
		if v.active {
			if k == ',' {
				if v.hasNotNumber {
					prop = prop[:v.headPos] + "\"" + prop[v.headPos:i] + "\"" + prop[i:]
				}
				v.active = false
				v.hasNotNumber = false
			} else if findWords([]string{"\"", "{", "}", "[", "]"}, string(k)) {
				v.active = false
				v.hasNotNumber = false
			} else if (i == v.headPos && k == '0') || !findWords([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}, string(k)) {
				v.hasNotNumber = true
			}
		} else if k == ':' {
			v.active = true
			v.headPos = i + 1
		}
	}
	prop = finalFix("tr\"ue", "true", prop)
	prop = finalFix("\"\":", "\":", prop)
	for i := 0; i < 32; i++ {
		prop = finalFix(string(uint8(i)), "", prop)
	}
	changedVideoInfo = videoInfo[position+1:]
	return
}

func findWords(findList []string, target string) bool {
	for _, k := range findList {
		if k == target {
			return true
		}
	}
	return false
}

func strangeQuotation(pos int, data string, key string, move int) (result string) {
	result = data[:pos] + key + data[pos+move:]
	return
}

func finalFix(searchBy, key, data string) string {
	for {
		pos := getPosition(data, searchBy)
		if pos == -1 {
			break
		}
		data = strangeQuotation(pos, data, key, len(searchBy))
	}
	return data
}
