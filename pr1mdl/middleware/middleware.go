package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Definisi warna untuk status code dan metode
var (
	reset   = "\033[0m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	green   = "\033[32m"
	red     = "\033[31m"
	blue    = "\033[34m"
)

// LogWriter custom ResponseWriter untuk menangkap status code
type LogWriter struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader menangani status code yang dikirimkan ke client
func (lw *LogWriter) WriteHeader(statusCode int) {
	lw.StatusCode = statusCode
	lw.ResponseWriter.WriteHeader(statusCode)
}

// Logger adalah middleware yang mencatat log request dengan warna
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Membungkus ResponseWriter dengan LogWriter untuk menangkap status code
		lw := &LogWriter{ResponseWriter: w, StatusCode: http.StatusOK}
		next.ServeHTTP(lw, r)

		// Menentukan warna
		statusColor := getStatusColor(lw.StatusCode)
		methodColor := getMethodColor(r.Method)

		logStart := fmt.Sprintf("%s%s%s", cyan, start.Format(time.RFC1123), reset)
		logMethod := fmt.Sprintf("%s%s%s", methodColor, r.Method, reset)
		logURLPath := fmt.Sprintf("%s%s%s", yellow, r.URL.Path, reset)
		logProto := fmt.Sprintf("%s%s%s", blue, r.Proto, reset)
		logStatusCode := fmt.Sprintf("%s%d%s", statusColor, lw.StatusCode, reset)
		logDuration := fmt.Sprintf("%s%s%s", magenta, time.Since(start).String(), reset)

		fmt.Printf("[ %s ] %s %s %s %s %s\n", logStart, logMethod, logURLPath, logProto, logStatusCode, logDuration)
	})
}

func getStatusColor(statusCode int) string {
	switch {
	case statusCode < 200:
		return reset
	case statusCode < 300:
		return green
	case statusCode < 400:
		return yellow
	case statusCode < 500:
		return red
	default:
		return cyan
	}
}

func getMethodColor(method string) string {
	switch method {
	case http.MethodPost:
		return blue
	case http.MethodDelete:
		return red
	case http.MethodGet:
		return green
	case http.MethodPut:
		return cyan
	default:
		return reset
	}
}
