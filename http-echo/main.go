package main

import (
	"encoding/json"
	"io"
	"libs/version"
	"log"
	"net/http"
	"os"
)

func main() {
	version.ShowVersionDetect("1.0.0")
	if len(os.Args) < 2 {
		log.Fatal("usage: http-echo <host:port>")
	}
	addr := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 读取 body
		var body map[string]any
		raw, err := io.ReadAll(r.Body)
		if err == nil && len(raw) > 0 {
			json.Unmarshal(raw, &body)
		}
		if body == nil {
			body = map[string]any{}
		}

		// 构建返回 JSON
		resp := map[string]any{
			"method":  r.Method,
			"host":    r.Host,
			"path":    r.URL.Path,
			"query":   r.URL.Query(),
			"headers": r.Header,
			"data":    body,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
