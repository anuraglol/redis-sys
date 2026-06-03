package main

import (
	"encoding/json"
	"net/http"
)

func handleFlushAll(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res, _ := executeCommand("FLUSHALL", []string{})

	json.NewEncoder(w).Encode(map[string]interface{}{"message": res})
}
