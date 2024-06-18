package serve

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var usage func(...string) = func(args ...string) {
	fmt.Println(`Usage: serve [options]
	Options:
		--help			Print this help message
		--port, -p	Specify the port to listen on

		Example:
		serve --port 8080

		Visit http://localhost:8080 in your browser to see the webapp.`)
}


func Serve(args ...string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			data := struct {
				Title string
				Body  string
			}{
				Title: "Hello, World!",
				Body:  "This is a test page.",
			}

			t.ExecuteTemplate(w, "index.html.tmpl", data)
		})
}

func Help(args ...string) string {
	return fmt.Sprintf("Help for serve module: %s", usage(args...) )
}

func Execute(args ...string) {
	fmt.Println("Hello from the serve package!")

	var t = template.Must(template.ParseFiles("index.html.tmpl"))
	var port string

	// Some args parsing
	if len(args) < 1 {̈́
		fmt.Println("Please provide a command.")
		usage(args...)
		return
	}

	for i, arg := range args[:1] {
		if arg == "--help" {
			usage(args...)
			return
		}

		if arg == "--port" || arg == "-p" {
			if len(args) < 2 {
				fmt.Println("Please provide a port number.")
				usage(args...)
				return
			}

			port := args[i+1]

			fmt.Println("Listening on port", port)
			return
		}
	}


	log.Fatal(http.ListenAndServe(":"+port, nil))
}
