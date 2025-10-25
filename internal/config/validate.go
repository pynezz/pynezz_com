package config

import (
	"fmt"
	"net/url"
	"strings"
)

type ValidationError struct {
	Scope  string
	Issues []string
}

func (e *ValidationError) Error() string {
	if len(e.Issues) == 0 {
		return fmt.Sprintf("config validation failed for %s", e.Scope)
	}
	return fmt.Sprintf("config validation failed for %s: %s", e.Scope, strings.Join(e.Issues, "; "))
}

func validateConfig(scope string, cfg AppConfig) error {
	var issues []string

	if strings.TrimSpace(cfg.Site.Name) == "" {
		issues = append(issues, "site.name must be set")
	}

	baseURL := strings.TrimSpace(cfg.Site.BaseURL)
	if baseURL == "" {
		issues = append(issues, "site.base_url must be set")
	} else {
		parsed, err := url.Parse(baseURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			issues = append(issues, "site.base_url must be a valid URL")
		}
	}

	if strings.TrimSpace(cfg.Content.Root) == "" {
		issues = append(issues, "content.root must be set")
	}

	if strings.TrimSpace(cfg.Database.Driver) == "" {
		issues = append(issues, "database.driver must be set")
	} else {
		switch cfg.Database.Driver {
		case "sqlite", "mysql", "postgres":
		default:
			issues = append(issues, fmt.Sprintf("database.driver %q is not supported", cfg.Database.Driver))
		}
	}

	if cfg.Database.Driver == "sqlite" && strings.TrimSpace(cfg.Database.Path) == "" {
		issues = append(issues, "database.path must be set when using sqlite")
	}

	if cfg.Server.Port <= 0 || cfg.Server.Port > 65535 {
		issues = append(issues, "server.port must be between 1 and 65535")
	}

	if cfg.Server.TLS {
		if strings.TrimSpace(cfg.Server.TLSCertPath) == "" {
			issues = append(issues, "server.tls_cert_path must be set when tls is true")
		}
		if strings.TrimSpace(cfg.Server.TLSKeyPath) == "" {
			issues = append(issues, "server.tls_key_path must be set when tls is true")
		}
	}

	if strings.TrimSpace(cfg.CLI.OutputFormat) == "" {
		issues = append(issues, "cli.output_format must be set")
	} else {
		switch cfg.CLI.OutputFormat {
		case "text", "json":
		default:
			issues = append(issues, fmt.Sprintf("cli.output_format %q is not supported", cfg.CLI.OutputFormat))
		}
	}

	if len(issues) > 0 {
		return &ValidationError{
			Scope:  scope,
			Issues: issues,
		}
	}

	return nil
}
