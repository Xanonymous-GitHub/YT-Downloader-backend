package converter

import (
	"encoding/hex"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
)

func HttpHexNumberToSimpleText(path string) (result []byte) {
	var enc = simplifiedchinese.GBK
	file, err := os.Open(path)
	defer file.Close()
	errorHandler.Handler("", err)
	r := transform.NewReader(file, enc.NewDecoder())
	buf := make([]byte, 1)
	tmpBuf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			words := string(buf[:n])
			if words == "%" {
				_, err := r.Read(tmpBuf)
				if err == io.EOF {
					break
				}
				errorHandler.Handler("converter error", err)
				decoded, err := hex.DecodeString(string(tmpBuf[:]))
				errorHandler.Handler("decode error", err)
				result = append(result, decoded...)
			} else {
				result = append(result, buf[:n]...)
			}
		}
		if err == io.EOF {
			break
		}
		errorHandler.Handler("converter error", err)
	}
	return
}
