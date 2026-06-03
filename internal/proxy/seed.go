package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleSeeding(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	for i := range 100 {
		_, err := executeCommand("set", []string{fmt.Sprintf("%d", i), fmt.Sprintf("%d", i+1)})
		if err != nil {
			log.Printf("oops failed on one front")
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"message": "ok"})
}
