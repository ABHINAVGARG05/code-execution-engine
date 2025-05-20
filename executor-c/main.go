package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ABHINAVGARG05/code-execution-engine/executor-lib"
)

type CodeRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func handleCodeExecution(w http.ResponseWriter, r *http.Request) {
	// log.Print("hi") --> For Debugging
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CodeRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	// log.Printf("Body: %s",body); --> For Debugging
	// log.Printf("Language: %s",req.Language); --> For Debugging
	config, ok := executor.LanguageConfigs[req.Language]
	if !ok {
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	output, err := executor.RunCode(req.Code, config)
	if err != nil {
		http.Error(w, "Execution error: "+err.Error()+"\n"+output, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(output))
}

func main() {
	http.HandleFunc("/run-c", handleCodeExecution)

	log.Println("C Executor running on port 5001...")
	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
