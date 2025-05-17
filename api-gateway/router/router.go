package router

import (
	"net/http"
	"github.com/ABHINAVGARG05/code-execution-engine/api-gateway/handler"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handler.ExecuteCode)
	return mux
}