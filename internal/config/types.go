package config

// AppConfig represents the fully merged configuration used at runtime.
type AppConfig struct {
	Site     SiteSettings     `json:"site"`
	Content  ContentSettings  `json:"content"`
	Database DatabaseSettings `json:"database"`
	Server   ServerSettings   `json:"server"`
	CLI      CLISettings      `json:"cli"`
}

// SiteSettings describes metadata needed to render the site shell.
type SiteSettings struct {
	Name        string        `json:"name"`
	BaseURL     string        `json:"base_url"`
	Description string        `json:"description"`
	Keywords    []string      `json:"keywords"`
	Theme       ThemeSettings `json:"theme"`
}

// ThemeSettings controls the active site theme and available variants.
type ThemeSettings struct {
	Default     string   `json:"default"`
	Variants    []string `json:"variants"`
	AllowSystem bool     `json:"allow_system"`
}

// ContentSettings contains content pipeline paths and filters.
type ContentSettings struct {
	Root       string   `json:"root"`
	StaticRoot string   `json:"static_root"`
	Drafts     string   `json:"drafts"`
	Ignore     []string `json:"ignore"`
}

// DatabaseSettings declares persistence driver configuration.
type DatabaseSettings struct {
	Driver        string `json:"driver"`
	DSN           string `json:"dsn"`
	Path          string `json:"path"`
	MigrationsDir string `json:"migrations_dir"`
	LogSQL        bool   `json:"log_sql"`
}

// ServerSettings defines HTTP serving details.
type ServerSettings struct {
	Host         string   `json:"host"`
	Port         int      `json:"port"`
	PublicURL    string   `json:"public_url"`
	TLS          bool     `json:"tls"`
	TLSCertPath  string   `json:"tls_cert_path"`
	TLSKeyPath   string   `json:"tls_key_path"`
	AllowOrigins []string `json:"allow_origins"`
}

// CLISettings controls CLI level defaults.
type CLISettings struct {
	DefaultProfile string `json:"default_profile"`
	OutputFormat   string `json:"output_format"`
	Color          bool   `json:"color"`
	LogLevel       string `json:"log_level"`
}

// ProfileConfig is the resolved configuration for a named profile.
type ProfileConfig struct {
	Name    string    `json:"name"`
	Extends []string  `json:"extends"`
	Config  AppConfig `json:"config"`
}

// Bundle contains the base configuration and all resolved profiles.
type Bundle struct {
	Base     AppConfig                  `json:"base"`
	Profiles map[string]ProfileConfig   `json:"profiles"`
	Sources  []string                   `json:"sources"`
}

// DefaultConfig returns the built-in defaults used as a baseline.
func DefaultConfig() AppConfig {
	return AppConfig{
		Site: SiteSettings{
			Name:        "pynezz.dev",
			BaseURL:     "https://pynezz.dev",
			Description: "pynezz.dev publishing stack",
			Keywords:    []string{"pynezz", "go", "cms"},
			Theme: ThemeSettings{
				Default:     "catppuccin-mocha",
				Variants:    []string{"catppuccin-latte", "catppuccin-frappe", "catppuccin-mocha"},
				AllowSystem: true,
			},
		},
		Content: ContentSettings{
			Root:       "content",
			StaticRoot: "pynezz/public",
			Drafts:     "content/drafts",
			Ignore:     []string{".git", ".DS_Store", "node_modules"},
		},
		Database: DatabaseSettings{
			Driver:        "sqlite",
			Path:          "pynezz.db",
			DSN:           "file:pynezz.db?_foreign_keys=on",
			MigrationsDir: "migrations",
			LogSQL:        false,
		},
		Server: ServerSettings{
			Host:      "127.0.0.1",
			Port:      8080,
			PublicURL: "http://127.0.0.1:8080",
			TLS:       false,
		},
		CLI: CLISettings{
			DefaultProfile: "dev",
			OutputFormat:   "text",
			Color:          true,
			LogLevel:       "info",
		},
	}
}
