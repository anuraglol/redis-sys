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

func handleGetAll(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	res, err := executeCommand("GETALL", []string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rawArr, ok := res.([]interface{})
	if !ok {
		http.Error(w, "Invalid data format received", http.StatusInternalServerError)
		return
	}

	type Item struct {
		Value  string `json:"value"`
		Expiry string `json:"expiry"`
	}
	dataMap := make(map[string]Item)

	for i := 0; i < len(rawArr); i += 3 {
		if i+2 < len(rawArr) {
			key, _ := rawArr[i].(string)
			val, _ := rawArr[i+1].(string)
			ttl, _ := rawArr[i+2].(string)

			dataMap[key] = Item{
				Value:  val,
				Expiry: ttl,
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"data": dataMap})
}
