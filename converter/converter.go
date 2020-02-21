package converter

import (
	"encoding/hex"
	"github.com/Xanonymous-GitHub/YT-Downloader-backend/errorHandler"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"strconv"
)

type specialCharFinder struct {
	slash bool
	u     bool
}

func HttpHexNumberToSimpleText(path string) (result []byte) {
	var enc = simplifiedchinese.GBK
	file, err := os.Open(path)
	defer file.Close()
	errorHandler.Handler("converter.HttpHexNumberToSimpleText => file, err := os.Open(path)", err)
	r := transform.NewReader(file, enc.NewDecoder())
	buf := make([]byte, 1)
	tmpBuf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		errorHandler.Handler("converter.HttpHexNumberToSimpleText => n, err := r.Read(buf)", err)
		if n > 0 {
			words := string(buf[:n])
			if words == "%" {
				_, err := r.Read(tmpBuf)
				if err == io.EOF {
					break
				}
				errorHandler.Handler("converter.HttpHexNumberToSimpleText => _, err := r.Read(tmpBuf)", err)
				decoded, err := hex.DecodeString(string(tmpBuf[:]))
				errorHandler.Handler("converter.HttpHexNumberToSimpleText => decoded, err := hex.DecodeString(string(tmpBuf[:]))", err)
				result = append(result, decoded...)
			} else {
				result = append(result, buf[:n]...)
			}
		}
	}
	return
}

func DecodeUTF16(b []byte) (result []byte) {
	bLength := len(b)
	scf := specialCharFinder{false, false}
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
			scf = specialCharFinder{false, false}
			i += 3
			continue
		}
		result = append(result, b[i])
	}
	return
}
