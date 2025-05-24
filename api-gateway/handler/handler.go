package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ABHINAVGARG05/code-execution-engine/api-gateway/utils"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Language == "" || req.Code == "" {
		http.Error(w, "Missing code or language", http.StatusBadRequest)
		return
	}
	log.Println("Before");
	err := utils.EnqueueCodeJob(req.Code, req.Language)
	if err != nil {
		http.Error(w, "Failed to enqueue code job: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("After")

	w.WriteHeader(http.StatusAccepted)
	log.Println("Code execution job enqueued successfully");
	w.Write([]byte("Code execution job enqueued successfully"))
}
