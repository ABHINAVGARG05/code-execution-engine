package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

)

type CodePayload struct {
	Language string `json:language`
	Code string `json:"code"`
}

func ResolveExecutor(lang string) string {
	switch lang {
	case "c":
		return "http://localhost:5001/run-c"
	case "cpp":
		return "http://executor-cpp:5002/run"
	case "java":
		return "http://executor-java:5003/run"
	case "python":
		return "http://executor-python:5004/run"
	case "go":
		return "http://executor-go:5005/run"
	default:
		return ""
	}
}

func ForwardCode(url, code string, language string) (*http.Response, error) {
	payload := CodePayload{
		Code: code,
		Language: language,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return http.Post(url, "application/json", bytes.NewReader(body))
}
