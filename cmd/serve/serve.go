package serve

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Serve(args ...string) {
	fmt.Println("Hello from the serve package!")
}

func Help(args ...string) string {
	return "Help for serve module: [usage instructions]"
}

func Execute(args ...string) {
	var t = template.Must(template.ParseFiles("index.html.tmpl"))

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
