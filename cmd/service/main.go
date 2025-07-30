package main

import (
	"fmt"
	"net/http"

	getstats "example.com/interview/internal/handlers/get_stats"
	recordstat "example.com/interview/internal/handlers/record_stat"
	"example.com/interview/internal/logging"
	"example.com/interview/internal/storage"
)

func main() {
	storage := storage.NewInMemoryStorage()
	logger := &logging.ConsoleLogger{}

	mux := http.NewServeMux()
	mux.Handle("/record_stat", recordstat.New(storage, logger))
	mux.Handle("/get_stats", getstats.New(storage, logger))

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
