package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

type VersionInfo struct {
	Version string `json:"version"`
}

func getVersion() string {
	data, err := os.ReadFile("version")
	if err != nil {
		log.Printf("failed to read version file: %v", err)
		return "unknown"
	}
	return strings.TrimSpace(string(data))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(VersionInfo{Version: getVersion()})
}

func main() {
	http.HandleFunc("/version", versionHandler)
	print("Backend listening on :8080\n")
	http.ListenAndServe(":8080", nil)
}
