package wallet_api_calls

import (
	"log"
	"net/http"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// get users IP
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	InfoLogger.Println("Requesters ip:", r.RemoteAddr)
	return r.RemoteAddr
}
