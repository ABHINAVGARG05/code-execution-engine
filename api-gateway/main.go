package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ABHINAVGARG05/code-execution-engine/api-gateway/router"
)

func main() {
	r := router.SetupRouter()
	fmt.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}