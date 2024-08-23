package server

import (
	"fmt"
	"net/http"
	"os"
)

func Run() error {
	addr := fmt.Sprintf(":%s", defaultEnv("PORT", "8080"))

	fmt.Println("listen", addr)
	fmt.Println("env", defaultEnv("APP_ENV", "dev"))

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "ok")
	})
	return http.ListenAndServe(addr, nil)
}

func defaultEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}
