package serve

import (
	"fmt"
	"strconv"

	"github.com/pynezz/pynezz_com/internal/helpers"
	"github.com/pynezz/pynezz_com/internal/runtime"
	"github.com/pynezz/pynezz_com/internal/server"
	"github.com/pynezz/pynezzentials/ansi"
)

type options struct {
	host string
	port int
}

func usage() string {
	env := runtime.Current()
	host := env.Active.Server.Host
	if host == "" {
		host = "127.0.0.1"
	}
	port := env.Active.Server.Port
	if port == 0 {
		port = 8080
	}
	return fmt.Sprintf(`Usage: serve [options]

  Serve the web server using configuration defaults (host %s, port %d).

Options:
    --help, -h      Print this help message
    --host          Override the listening host (default %s)
    --port, -p      Override the listening port (default %d)
`, host, port, host, port)
}

func Help(args ...string) string {
	return usage()
}

func Execute(args ...string) {
	env := runtime.Current()
	opts := options{
		host: env.Active.Server.Host,
		port: env.Active.Server.Port,
	}
	if opts.host == "" {
		opts.host = "127.0.0.1"
	}
	if opts.port == 0 {
		opts.port = 8080
	}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--help", "-h":
			fmt.Println(usage())
			return
		case "--host":
			if i+1 >= len(args) {
				helpers.Warning("Missing value for --host")
				fmt.Println(usage())
				return
			}
			i++
			opts.host = args[i]
		case "--port", "-p":
			if i+1 >= len(args) {
				helpers.Warning("Missing value for --port")
				fmt.Println(usage())
				return
			}
			i++
			port, err := strconv.Atoi(args[i])
			if err != nil || port <= 0 || port > 65535 {
				helpers.Warning("Invalid port supplied.")
				fmt.Println(usage())
				return
			}
			opts.port = port
		default:
			helpers.Warning(fmt.Sprintf("Unknown argument: %s", arg))
			fmt.Println(usage())
			return
		}
	}

	address := fmt.Sprintf("%s:%d", opts.host, opts.port)
	server.Serve(strconv.Itoa(opts.port))
	ansi.PrintInfo("Serving content on " + address)
	ansi.PrintInfo("Waiting for SIGINT (Ctrl+C) to shutdown...")
}
