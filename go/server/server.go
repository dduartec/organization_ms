package server

import (
	"net/http"
	"time"

	"goji.io"
)

//New conexion
func New(mux *goji.Mux, serverAddres string) *http.Server {
	srv := &http.Server{
		Addr:         serverAddres,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return srv
}
