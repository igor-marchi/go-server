package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  0 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
