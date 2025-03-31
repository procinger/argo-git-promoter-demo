package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/encoding/javaproperties"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var backendURL atomic.Value

func initConfigWatcher() {
	codec := &javaproperties.Codec{}

	codecRegistry := viper.NewCodecRegistry()
	codecRegistry.RegisterCodec("properties", codec)

	v := viper.NewWithOptions(
		viper.WithCodecRegistry(codecRegistry),
	)

	v.SetConfigType("properties")
	v.SetConfigName("labels")
	v.AddConfigPath("/etc/pod-labels")

	v.SetDefault("backend_url", "backend")

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("[frontend] config file not found or unreadable, using default: %v", err)
	} else {
		log.Printf("[frontend] using config file: %s", v.ConfigFileUsed())
	}

	backendURL.Store(v.GetString("backend_url"))
	log.Printf("[frontend] initial backend URL: %s", backendURL.Load())

	v.OnConfigChange(func(e fsnotify.Event) {
		newURL := v.GetString("backend_url")
		backendURL.Store(newURL)
		log.Printf("[frontend] backend URL updated to: %s", newURL)
	})
	v.WatchConfig()
}

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

func fetchBackendVersion() (string, map[string][]string) {
	client := http.Client{Timeout: 2 * time.Second}
	backendURLString := strings.Trim(backendURL.Load().(string), "\"")
	resp, err := client.Get(fmt.Sprintf("http://%s/version", backendURLString))
	if err != nil {
		log.Printf("failed to contact backend: %v", err)
		return "unreachable", resp.Header
	}
	defer resp.Body.Close()

	var versionInfo VersionInfo
	err = json.NewDecoder(resp.Body).Decode(&versionInfo)
	if err != nil {
		log.Printf("failed to decode backend response: %v", err)
		return "invalid response", resp.Header
	}

	return versionInfo.Version, resp.Header
}

type PageData struct {
	FrontendVersion string
	BackendVersion  string
	RequestHeaders  map[string][]string
	BackendHeaders  map[string][]string
}

func versionPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request to %s", r.Method, r.URL.Path)
	backendVersion, backendHeaders := fetchBackendVersion()

	data := PageData{
		FrontendVersion: getVersion(),
		BackendVersion:  backendVersion,
		RequestHeaders:  r.Header,
		BackendHeaders:  backendHeaders,
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
		<div class="card shadow-sm mb-4">
			<div class="card-body">
				<h1 class="card-title text-center mb-4">Service Versions</h1>
				<ul class="list-group list-group-flush">
					<li class="list-group-item"><strong>Frontend Version:</strong> {{.FrontendVersion}}</li>
					<li class="list-group-item"><strong>Backend Version:</strong> {{.BackendVersion}}</li>
				</ul>
			</div>
		</div>
		<div class="card shadow-sm mb-4">
			<div class="card-body">
				<h2 class="card-title">Request Headers</h2>
				<ul class="list-group list-group-flush">
					{{range $key, $vals := .RequestHeaders}}
						<li class="list-group-item"><strong>{{$key}}:</strong> {{range $i, $v := $vals}}{{if $i}}, {{end}}{{$v}}{{end}}</li>
					{{end}}
				</ul>
			</div>
		</div>
		<div class="card shadow-sm">
			<div class="card-body">
				<h2 class="card-title">Backend Response Headers</h2>
				<ul class="list-group list-group-flush">
					{{range $key, $vals := .BackendHeaders}}
						<li class="list-group-item"><strong>{{$key}}:</strong> {{range $i, $v := $vals}}{{if $i}}, {{end}}{{$v}}{{end}}</li>
					{{end}}
				</ul>
			</div>
		</div>
	</div>
</body>
</html>`

	t, err := template.New("version").Parse(tmpl)
	if err != nil {
		log.Printf("[frontend] failed to parse template: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Printf("[frontend] failed to execute template: %v", err)
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

func main() {
	initConfigWatcher()
	http.HandleFunc("/", versionPageHandler)
	print("Frontend listening on :8080\n")
	http.ListenAndServe(":8080", nil)
}
