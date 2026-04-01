package testing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func respond(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	headers := map[string]string{}
	for k, v := range r.Header {
		headers[k] = v[0]
	}
	respond(w, r, map[string]any{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": headers,
		"body":    string(body),
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(405)
			return
		}
		echoHandler(w, r)
	})
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			return
		}
		echoHandler(w, r)
	})
	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(405)
			return
		}
		echoHandler(w, r)
	})
	mux.HandleFunc("/patch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			w.WriteHeader(405)
			return
		}
		echoHandler(w, r)
	})
	mux.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(405)
			return
		}
		echoHandler(w, r)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		echoHandler(w, r)
	})

	fmt.Println("Listening on :8081")
	http.ListenAndServe(":8081", mux)
}
