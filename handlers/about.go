package handlers

import (
	"log"
	"net/http"
	"text/template"
)

// AboutHandler handles the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving about page")

	if r.URL.Path != "/about/" {
		http.Redirect(w, r, "/about/", http.StatusTemporaryRedirect)
		return
	}

	t := template.Must(template.New("about.html").Delims("[[", "]]").ParseFiles("templates/about.html"))
	t.Execute(w, map[string]interface{}{
		"google_analytics_key": googleAnalyticsKey,
	})
	return
}
