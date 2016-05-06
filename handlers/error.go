package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// ErrorHandler handles the 404 page
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	log.Println("Serving error page")

	var t *template.Template

	w.WriteHeader(status)
	t = template.Must(template.New("error.html").Delims("[[", "]]").ParseFiles("templates/error.html"))

	t.Execute(w, map[string]interface{}{
		"StatusCode":           status,
		"google_analytics_key": googleAnalyticsKey,
	})
	return
}
