package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handleGet(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := strings.TrimPrefix(req.URL.Path, "/get/")
	if key == "" || strings.Contains(key, "/") {
		http.Error(w, "Missing or invalid key", http.StatusBadRequest)
		return
	}

	res, err := executeCommand("GET", []string{key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Key not found"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": res})
}

func handleSet(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body SetRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.Key == "" || body.Value == "" {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	args := []string{body.Key, body.Value}
	if body.EX > 0 {
		args = append(args, "EX", fmt.Sprintf("%d", body.EX))
	}

	res, err := executeCommand("SET", args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": res})
}

func handleDel(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := strings.TrimPrefix(req.URL.Path, "/del/")
	if key == "" || strings.Contains(key, "/") {
		http.Error(w, "Missing or invalid key", http.StatusBadRequest)
		return
	}

	res, err := executeCommand("DEL", []string{key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"deleted_count": res})
}

func handleIncr(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := strings.TrimPrefix(req.URL.Path, "/incr/")
	if key == "" || strings.Contains(key, "/") {
		http.Error(w, "Missing or invalid key", http.StatusBadRequest)
		return
	}

	res, err := executeCommand("INCR", []string{key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": res})
}

func handleTtl(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := strings.TrimPrefix(req.URL.Path, "/ttl/")
	if key == "" || strings.Contains(key, "/") {
		http.Error(w, "Missing or invalid key", http.StatusBadRequest)
		return
	}

	res, err := executeCommand("TTL", []string{key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "ttl": res})
}

func handleStats(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := executeCommand("STATS", []string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"stats": res})
}
