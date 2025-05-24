module github.com/ABHINAVGARG05/code-execution-engine/executor-c

go 1.24.2

require github.com/ABHINAVGARG05/code-execution-engine v0.0.0-20250518125046-6fb82df2b1a9

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/redis/go-redis/v9 v9.8.0 // indirect
)

replace github.com/ABHINAVGARG05/code-execution-engine/executor-lib => ../executor-lib
