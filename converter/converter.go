package converter

import (
	"encoding/hex"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"strconv"
)

type specialCharFinder struct {
	slash   bool
	u       bool
	percent int
}

func HttpHexNumberToSimpleText(b []byte) (result []byte) {
	scf := specialCharFinder{false, false, 0}
	tmpStr := make([]byte, 2)
	for _, data := range b {
		if data == '%' {
			scf.percent++
			continue
		}
		if data != '%' && scf.percent != 0 {
			tmpStr = append(tmpStr, data)
			scf.percent++
			if scf.percent == 3 {
				decoded, err := hex.DecodeString(string(tmpStr))
				errorHandler.Handler("converter.HttpHexNumberToSimpleText => decoded, err := hex.DecodeString(string(tmpStr))", err)
				result = append(result, decoded...)
				scf.percent = 0
				tmpStr = []byte{}
			}
			continue
		}
		scf.percent = 0
		tmpStr = []byte{}
		result = append(result, data)
	}
	return
}

func DecodeUTF16(b []byte) (result []byte) {
	bLength := len(b)
	scf := specialCharFinder{false, false, 0}
	for i := 0; i < bLength; i++ {
		if b[i] == '\\' {
			scf.slash = true
			continue
		}
		if b[i] == 'u' && scf.slash && !scf.u {
			scf.u = true
			continue
		}
		if i+3 < bLength && b[i]-'0' < 10 && b[i+1]-'0' < 10 && b[i+2]-'0' < 10 && b[i+3]-'0' < 10 && scf.u {
			n, err := strconv.ParseInt(string(b[i:i+4]), 16, 32)
			errorHandler.Handler("converter.DecodeUTF16 => n, err := strconv.ParseInt(string(b[i:i+4]), 16, 32)", err)
			result = append(result, byte(n))
			scf = specialCharFinder{false, false, 0}
			i += 3
			continue
		}
		result = append(result, b[i])
	}
	return
}
