package serve

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func Serve(args ...string) {
	fmt.Println("Hello from the serve package!")
}

func Help(args ...string) string {
	return "Help for serve module: [usage instructions]"
}

func Execute(args ...string) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
