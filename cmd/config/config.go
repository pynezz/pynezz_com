package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	cfg "github.com/pynezz/pynezz_com/internal/config"
	"github.com/pynezz/pynezz_com/internal/runtime"
	"github.com/pynezz/pynezzentials/ansi"
)

var errHelpRequested = errors.New("help requested")

type showOptions struct {
	profile string
	format  string
}

func usage() string {
	return `Usage: config <command> [options]

Commands:
    show            Display the resolved configuration (default)
    profiles        List available profiles

show options:
    --profile, -p   Override the profile to display (use "base" for defaults)
    --format, -f    Output format: text (default) or json
    --json          Shortcut for --format json
    --text          Shortcut for --format text

Examples:
    pynezz config
    pynezz config show --profile prod --format json
    pynezz config profiles
`
}

func Help(args ...string) string {
	return usage()
}

func Execute(args ...string) {
	if len(args) == 0 || args[0] == "show" {
		offset := 0
		if len(args) > 0 && args[0] == "show" {
			offset = 1
		}
		opts, err := parseShowOptions(args[offset:])
		if err != nil {
			if errors.Is(err, errHelpRequested) {
				fmt.Print(usage())
				return
			}
			ansi.PrintError(err.Error())
			fmt.Print(usage())
			return
		}
		if err := runShow(opts); err != nil {
			ansi.PrintError(err.Error())
		}
		return
	}

	switch args[0] {
	case "profiles":
		if len(args) > 1 {
			ansi.PrintWarning("profiles command does not take additional arguments")
			fmt.Print(usage())
			return
		}
		runProfiles()
	case "help", "--help", "-h":
		fmt.Print(usage())
	default:
		ansi.PrintWarning("unknown config command: " + args[0])
		fmt.Print(usage())
	}
}

func parseShowOptions(args []string) (showOptions, error) {
	opts := showOptions{
		format: "text",
	}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--help", "-h":
			return opts, errHelpRequested
		case "--profile", "-p":
			if i+1 >= len(args) {
				return opts, errors.New("missing value for --profile")
			}
			i++
			opts.profile = strings.TrimSpace(args[i])
		case "--format", "-f":
			if i+1 >= len(args) {
				return opts, errors.New("missing value for --format")
			}
			i++
			opts.format = strings.ToLower(strings.TrimSpace(args[i]))
		case "--json":
			opts.format = "json"
		case "--text":
			opts.format = "text"
		default:
			return opts, fmt.Errorf("unknown option %q", arg)
		}
	}

	switch opts.format {
	case "text", "json":
	default:
		return opts, fmt.Errorf("unsupported format %q", opts.format)
	}

	return opts, nil
}

func runShow(opts showOptions) error {
	env := runtime.Current()

	cfg, profileName, active := resolveConfig(opts.profile, env)
	if profileName == "" {
		return fmt.Errorf("unknown profile %q", opts.profile)
	}

	switch opts.format {
	case "text":
		renderText(cfg, profileName, active, env.Bundle.Sources)
	case "json":
		payload, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal config: %w", err)
		}
		fmt.Println(string(payload))
	}

	return nil
}

func runProfiles() {
	env := runtime.Current()

	fmt.Println("Available profiles:")

	baseLabel := "base"
	if env.ActiveProfile == "" {
		fmt.Printf("  * %s (active)\n", baseLabel)
	} else {
		fmt.Printf("    %s\n", baseLabel)
	}

	names := make([]string, 0, len(env.Bundle.Profiles))
	for name := range env.Bundle.Profiles {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		profile := env.Bundle.Profiles[name]
		activeMarker := " "
		if name == env.ActiveProfile {
			activeMarker = "*"
		}
		if len(profile.Extends) > 0 {
			fmt.Printf("  %s %s (extends %s)\n", activeMarker, name, strings.Join(profile.Extends, ", "))
		} else {
			fmt.Printf("  %s %s\n", activeMarker, name)
		}
	}
}

func resolveConfig(profile string, env runtime.Environment) (cfg.AppConfig, string, bool) {
	switch strings.TrimSpace(profile) {
	case "":
		if env.ActiveProfile == "" {
			return env.Bundle.Base, "base", true
		}
		return env.Active, env.ActiveProfile, true
	case "base":
		return env.Bundle.Base, "base", env.ActiveProfile == ""
	default:
		profileConfig, ok := env.Bundle.Profiles[profile]
		if !ok {
			return cfg.AppConfig{}, "", false
		}
		return profileConfig.Config, profile, profile == env.ActiveProfile
	}
}

func renderText(cfg cfg.AppConfig, profile string, active bool, sources []string) {
	header := fmt.Sprintf("Profile: %s", profile)
	if active {
		header += " (active)"
	}
	fmt.Println(header)
	if len(sources) > 0 {
		fmt.Println("Sources: " + strings.Join(sources, ", "))
	}

	fmt.Println("\nSite")
	fmt.Println("  name:        " + cfg.Site.Name)
	fmt.Println("  base_url:    " + cfg.Site.BaseURL)
	fmt.Println("  description: " + cfg.Site.Description)
	if len(cfg.Site.Keywords) > 0 {
		fmt.Println("  keywords:    " + strings.Join(cfg.Site.Keywords, ", "))
	}
	fmt.Println("  theme.default:     " + cfg.Site.Theme.Default)
	if len(cfg.Site.Theme.Variants) > 0 {
		fmt.Println("  theme.variants:    " + strings.Join(cfg.Site.Theme.Variants, ", "))
	}
	fmt.Printf("  theme.allow_system: %t\n", cfg.Site.Theme.AllowSystem)

	fmt.Println("\nContent")
	fmt.Println("  root:        " + cfg.Content.Root)
	fmt.Println("  static_root: " + cfg.Content.StaticRoot)
	fmt.Println("  drafts:      " + cfg.Content.Drafts)
	if len(cfg.Content.Ignore) > 0 {
		fmt.Println("  ignore:      " + strings.Join(cfg.Content.Ignore, ", "))
	}

	fmt.Println("\nDatabase")
	fmt.Println("  driver:        " + cfg.Database.Driver)
	fmt.Println("  dsn:           " + cfg.Database.DSN)
	fmt.Println("  path:          " + cfg.Database.Path)
	fmt.Println("  migrations:    " + cfg.Database.MigrationsDir)
	fmt.Printf("  log_sql:       %t\n", cfg.Database.LogSQL)

	fmt.Println("\nServer")
	fmt.Println("  host:          " + cfg.Server.Host)
	fmt.Printf("  port:          %d\n", cfg.Server.Port)
	fmt.Println("  public_url:    " + cfg.Server.PublicURL)
	fmt.Printf("  tls:           %t\n", cfg.Server.TLS)
	if cfg.Server.TLSCertPath != "" {
		fmt.Println("  tls_cert_path: " + cfg.Server.TLSCertPath)
	}
	if cfg.Server.TLSKeyPath != "" {
		fmt.Println("  tls_key_path:  " + cfg.Server.TLSKeyPath)
	}
	if len(cfg.Server.AllowOrigins) > 0 {
		fmt.Println("  allow_origins: " + strings.Join(cfg.Server.AllowOrigins, ", "))
	}

	fmt.Println("\nCLI")
	fmt.Println("  default_profile: " + cfg.CLI.DefaultProfile)
	fmt.Println("  output_format:   " + cfg.CLI.OutputFormat)
	fmt.Printf("  color:           %t\n", cfg.CLI.Color)
	fmt.Println("  log_level:       " + cfg.CLI.LogLevel)
}
