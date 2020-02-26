package converter

import "strings"

func MakeUrlList(s string) (result []string) {
	for getPosition(s) != -1 {
		r, l := wholeUrl(getPosition(s), s)
		result = append(result, r)
		s = l
	}
	return
}

func getPosition(s string) int {
	return strings.Index(s, "videoplayback")
}

func wholeUrl(p int, s string) (result string, l string) {
	result += string(s[p])
	q := p - 1
	for ; s[q] != '"'; q -= 1 {
	}
	result = s[q+1:p] + result
	q = p + 1
	for ; s[q] != '"'; q++ {
	}
	result += s[p+1 : q]
	l = s[q+1:]
	return
}
