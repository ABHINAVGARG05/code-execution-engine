package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/ABHINAVGARG05/code-execution-engine/api-gateway/utils"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	targetURL := utils.ResolveExecutor(req.Language)
	if targetURL == "" {
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	resp, err := utils.ForwardCode(targetURL, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}