package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ABHINAVGARG05/code-execution-engine/api-gateway/utils"
)

type CodeRequest struct {
	Language string `json:"language_id"`
	Code     string `json:"code"`
}

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	var req CodeRequest
	log.Println("HI")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	log.Printf("Request: %s",req.Language)
	targetURL := utils.ResolveExecutor(req.Language)
	if targetURL == "" {
		http.Error(w, "Unsupported language: API-Gateway", http.StatusBadRequest)
		return
	}

	resp, err := utils.ForwardCode(targetURL, req.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	log.Printf("response: ",resp.Body)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
