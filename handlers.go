package main

import (
	"net/http"
	"os"
)

func (c *Config) homeHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "./index.html")
}

func (c *Config) serveFilesHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := "." + r.URL.Path
			// If the requested path is a file , serve the file, eg when html request for the css file
			if f, err := os.Stat(path); err == nil && !f.IsDir() {
				http.ServeFile(w, r, path)
				return
			}

			next.ServeHTTP(w, r)

		})

	}
}
