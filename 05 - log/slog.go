package main

import (
	"log/slog"
	"os"
)

func main() {
	// plain text
	slog.Info("Hello, World!")
	slog.Default().With("key", "value").Info("Hello, World!")

	// json
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "x-itau-correlation", "id-01010")
}

// 2023/07/30 16:05:40 INFO Hello, World!
// 2023/07/30 16:05:40 INFO Hello, World! key=value
// {"time":"2023-07-30T16:05:40.866412+07:00","level":"INFO","msg":"hello","user":"user1","age":10}
