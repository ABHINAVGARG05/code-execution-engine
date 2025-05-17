package utils

import (
	"bytes"
	"net/http"
	"strings"
)

func ResolveExecutor(lang string) string {
	switch lang {
	case "c":
		return "http://executor-c:5004/run"
	default:
		return ""
	}
}

func ForwardCode(url string, code string) (*http.Response, error) {
	escaped := strings.ReplaceAll(code, `"`, `\"`)
	payload := []byte(`{"code":"` + escaped + `"}`)
	return http.Post(url, "application/json", bytes.NewBuffer(payload))
}