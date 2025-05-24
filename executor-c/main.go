package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ABHINAVGARG05/code-execution-engine/executor-lib"
)

type CodeRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

type CodeExecutionJob struct {
	Code   string 
	Config executor.ExecutionConfig
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func handleCodeExecution(w http.ResponseWriter, r *http.Request) {
	log.Println("HI");
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

func startRedisWorker() {
	for {
		result, err := rdb.BRPop(ctx, 0*time.Second, "code_execution_queue").Result()
		if err != nil {
			log.Println("Redis BRPop error:", err)
			continue
		}
		if len(result) < 2 {
			continue
		}

		var job CodeExecutionJob
		if err := json.Unmarshal([]byte(result[1]), &job); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		output, err := executor.RunCode(job.Code, job.Config)
		if err != nil {
			log.Printf("Execution Error: %v\nOutput:\n%s\n", err, output)
		} else {
			log.Printf("Execution Output:\n%s\n", output)
		}
	}
}

func main() {
	go startRedisWorker() // Worker starts in background

	http.HandleFunc("/run-c", handleCodeExecution)
	log.Println("C Executor running on port 5001...")
	if err := http.ListenAndServe(":5001", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
