package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, templatePaths string) {

	t, err := template.ParseFiles(templatePaths)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error parsing html to template. %v", err))
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error exucuting template. %v", err))
		return
	}
}
