package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

func fetchBackendVersion() string {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://localhost:8081" // Default fallback
	}

	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(fmt.Sprintf("%s/version", backendURL))
	if err != nil {
		log.Printf("failed to contact backend: %v", err)
		return "unreachable"
	}
	defer resp.Body.Close()

	var versionInfo VersionInfo
	err = json.NewDecoder(resp.Body).Decode(&versionInfo)
	if err != nil {
		log.Printf("failed to decode backend response: %v", err)
		return "invalid response"
	}

	return versionInfo.Version
}

type PageData struct {
	FrontendVersion string
	BackendVersion  string
}

func versionPageHandler(w http.ResponseWriter, r *http.Request) {
	backendVersion := fetchBackendVersion()

	data := PageData{
		FrontendVersion: getVersion(),
		BackendVersion:  backendVersion,
	}

	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Service Versions</title>
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
	<div class="container py-5">
		<div class="card shadow-sm">
			<div class="card-body">
				<h1 class="card-title text-center mb-4">Service Versions</h1>
				<ul class="list-group list-group-flush">
					<li class="list-group-item"><strong>Frontend Version:</strong> {{.FrontendVersion}}</li>
					<li class="list-group-item"><strong>Backend Version:</strong> {{.BackendVersion}}</li>
				</ul>
			</div>
		</div>
	</div>
</body>
</html>`
	t, _ := template.New("frontendVersion").Parse(tmpl)
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", versionPageHandler)
	print("Frontend listening on :8080\n")
	http.ListenAndServe(":8080", nil)
}
