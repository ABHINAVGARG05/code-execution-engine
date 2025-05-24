package utils

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/ABHINAVGARG05/code-execution-engine/executor-lib"
	"github.com/redis/go-redis/v9"
)

type CodePayload struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type CodeExecutionJob struct {
	Code   string
	Config executor.ExecutionConfig
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func EnqueueCodeJob(code string, language string) error {
	config, ok := executor.LanguageConfigs[language]
	if !ok {
		return errors.New("unsupported language")
	}

	job := CodeExecutionJob{
		Code:   code,
		Config: config,
	}

	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	log.Println("Before1");
	log.Println(data);
	result, err := rdb.BRPop(ctx, 5*time.Second, "code_execution_queue").Result()
	
		if err != nil {
			log.Println("Redis BRPop error:", err)
		}
		log.Println(result);
		if len(result) < 2 {
			log.Println("redis connected", result)
		}
		log.Println("Code enqued and redis connected");
	return rdb.LPush(ctx, "code_execution_queue", data).Err()
}
