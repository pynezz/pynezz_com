package config

// document represents a single TOML document with optional sections.
type document struct {
	Site     *siteFragment              `toml:"site"`
	Content  *contentFragment           `toml:"content"`
	Database *databaseFragment          `toml:"database"`
	Server   *serverFragment            `toml:"server"`
	CLI      *cliFragment               `toml:"cli"`
	Profiles map[string]profileFragment `toml:"profiles"`
}

func (d *document) merge(src document) {
	if src.Site != nil {
		if d.Site == nil {
			d.Site = src.Site.clone()
		} else {
			d.Site.merge(src.Site)
		}
	}
	if src.Content != nil {
		if d.Content == nil {
			d.Content = src.Content.clone()
		} else {
			d.Content.merge(src.Content)
		}
	}
	if src.Database != nil {
		if d.Database == nil {
			d.Database = src.Database.clone()
		} else {
			d.Database.merge(src.Database)
		}
	}
	if src.Server != nil {
		if d.Server == nil {
			d.Server = src.Server.clone()
		} else {
			d.Server.merge(src.Server)
		}
	}
	if src.CLI != nil {
		if d.CLI == nil {
			d.CLI = src.CLI.clone()
		} else {
			d.CLI.merge(src.CLI)
		}
	}
}

type siteFragment struct {
	Name        *string         `toml:"name"`
	BaseURL     *string         `toml:"base_url"`
	Description *string         `toml:"description"`
	Keywords    *[]string       `toml:"keywords"`
	Theme       *themeFragment  `toml:"theme"`
}

func (f *siteFragment) merge(src *siteFragment) {
	if src == nil {
		return
	}
	if src.Name != nil {
		f.Name = src.Name
	}
	if src.BaseURL != nil {
		f.BaseURL = src.BaseURL
	}
	if src.Description != nil {
		f.Description = src.Description
	}
	if src.Keywords != nil {
		f.Keywords = cloneStringSlicePointer(src.Keywords)
	}
	if src.Theme != nil {
		if f.Theme == nil {
			f.Theme = src.Theme.clone()
		} else {
			f.Theme.merge(src.Theme)
		}
	}
}

func (f *siteFragment) clone() *siteFragment {
	if f == nil {
		return nil
	}
	out := *f
	if f.Keywords != nil {
		out.Keywords = cloneStringSlicePointer(f.Keywords)
	}
	if f.Theme != nil {
		out.Theme = f.Theme.clone()
	}
	return &out
}

type themeFragment struct {
	Default     *string   `toml:"default"`
	Variants    *[]string `toml:"variants"`
	AllowSystem *bool     `toml:"allow_system"`
}

func (f *themeFragment) merge(src *themeFragment) {
	if src == nil {
		return
	}
	if src.Default != nil {
		f.Default = src.Default
	}
	if src.AllowSystem != nil {
		f.AllowSystem = src.AllowSystem
	}
	if src.Variants != nil {
		f.Variants = cloneStringSlicePointer(src.Variants)
	}
}

func (f *themeFragment) clone() *themeFragment {
	if f == nil {
		return nil
	}
	out := *f
	if f.Variants != nil {
		out.Variants = cloneStringSlicePointer(f.Variants)
	}
	return &out
}

type contentFragment struct {
	Root       *string   `toml:"root"`
	StaticRoot *string   `toml:"static_root"`
	Drafts     *string   `toml:"drafts"`
	Ignore     *[]string `toml:"ignore"`
}

func (f *contentFragment) merge(src *contentFragment) {
	if src == nil {
		return
	}
	if src.Root != nil {
		f.Root = src.Root
	}
	if src.StaticRoot != nil {
		f.StaticRoot = src.StaticRoot
	}
	if src.Drafts != nil {
		f.Drafts = src.Drafts
	}
	if src.Ignore != nil {
		f.Ignore = cloneStringSlicePointer(src.Ignore)
	}
}

func (f *contentFragment) clone() *contentFragment {
	if f == nil {
		return nil
	}
	out := *f
	if f.Ignore != nil {
		out.Ignore = cloneStringSlicePointer(f.Ignore)
	}
	return &out
}

type databaseFragment struct {
	Driver        *string `toml:"driver"`
	DSN           *string `toml:"dsn"`
	Path          *string `toml:"path"`
	MigrationsDir *string `toml:"migrations_dir"`
	LogSQL        *bool   `toml:"log_sql"`
}

func (f *databaseFragment) merge(src *databaseFragment) {
	if src == nil {
		return
	}
	if src.Driver != nil {
		f.Driver = src.Driver
	}
	if src.DSN != nil {
		f.DSN = src.DSN
	}
	if src.Path != nil {
		f.Path = src.Path
	}
	if src.MigrationsDir != nil {
		f.MigrationsDir = src.MigrationsDir
	}
	if src.LogSQL != nil {
		f.LogSQL = src.LogSQL
	}
}

func (f *databaseFragment) clone() *databaseFragment {
	if f == nil {
		return nil
	}
	out := *f
	return &out
}

type serverFragment struct {
	Host         *string   `toml:"host"`
	Port         *int      `toml:"port"`
	PublicURL    *string   `toml:"public_url"`
	TLS          *bool     `toml:"tls"`
	TLSCertPath  *string   `toml:"tls_cert_path"`
	TLSKeyPath   *string   `toml:"tls_key_path"`
	AllowOrigins *[]string `toml:"allow_origins"`
}

func (f *serverFragment) merge(src *serverFragment) {
	if src == nil {
		return
	}
	if src.Host != nil {
		f.Host = src.Host
	}
	if src.Port != nil {
		f.Port = src.Port
	}
	if src.PublicURL != nil {
		f.PublicURL = src.PublicURL
	}
	if src.TLS != nil {
		f.TLS = src.TLS
	}
	if src.TLSCertPath != nil {
		f.TLSCertPath = src.TLSCertPath
	}
	if src.TLSKeyPath != nil {
		f.TLSKeyPath = src.TLSKeyPath
	}
	if src.AllowOrigins != nil {
		f.AllowOrigins = cloneStringSlicePointer(src.AllowOrigins)
	}
}

func (f *serverFragment) clone() *serverFragment {
	if f == nil {
		return nil
	}
	out := *f
	if f.AllowOrigins != nil {
		out.AllowOrigins = cloneStringSlicePointer(f.AllowOrigins)
	}
	return &out
}

type cliFragment struct {
	DefaultProfile *string `toml:"default_profile"`
	OutputFormat   *string `toml:"output_format"`
	Color          *bool   `toml:"color"`
	LogLevel       *string `toml:"log_level"`
}

func (f *cliFragment) merge(src *cliFragment) {
	if src == nil {
		return
	}
	if src.DefaultProfile != nil {
		f.DefaultProfile = src.DefaultProfile
	}
	if src.OutputFormat != nil {
		f.OutputFormat = src.OutputFormat
	}
	if src.Color != nil {
		f.Color = src.Color
	}
	if src.LogLevel != nil {
		f.LogLevel = src.LogLevel
	}
}

func (f *cliFragment) clone() *cliFragment {
	if f == nil {
		return nil
	}
	out := *f
	return &out
}

type profileFragment struct {
	Extends []string         `toml:"extends"`
	Site    *siteFragment    `toml:"site"`
	Content *contentFragment `toml:"content"`
	Database *databaseFragment `toml:"database"`
	Server  *serverFragment  `toml:"server"`
}

func (f *profileFragment) merge(src profileFragment) {
	if len(src.Extends) > 0 {
		f.Extends = cloneStrings(src.Extends)
	}
	if src.Site != nil {
		if f.Site == nil {
			f.Site = src.Site.clone()
		} else {
			f.Site.merge(src.Site)
		}
	}
	if src.Content != nil {
		if f.Content == nil {
			f.Content = src.Content.clone()
		} else {
			f.Content.merge(src.Content)
		}
	}
	if src.Database != nil {
		if f.Database == nil {
			f.Database = src.Database.clone()
		} else {
			f.Database.merge(src.Database)
		}
	}
	if src.Server != nil {
		if f.Server == nil {
			f.Server = src.Server.clone()
		} else {
			f.Server.merge(src.Server)
		}
	}
}

func cloneProfileFragment(src profileFragment) *profileFragment {
	out := profileFragment{
		Extends: cloneStrings(src.Extends),
	}
	if src.Site != nil {
		out.Site = src.Site.clone()
	}
	if src.Content != nil {
		out.Content = src.Content.clone()
	}
	if src.Database != nil {
		out.Database = src.Database.clone()
	}
	if src.Server != nil {
		out.Server = src.Server.clone()
	}
	return &out
}

func cloneStringSlicePointer(src *[]string) *[]string {
	if src == nil {
		return nil
	}
	cloned := cloneStrings(*src)
	return &cloned
}

func cloneStrings(in []string) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	return out
}

func cloneAppConfig(cfg AppConfig) AppConfig {
	cfg.Site.Keywords = cloneStrings(cfg.Site.Keywords)
	cfg.Site.Theme.Variants = cloneStrings(cfg.Site.Theme.Variants)
	cfg.Content.Ignore = cloneStrings(cfg.Content.Ignore)
	cfg.Server.AllowOrigins = cloneStrings(cfg.Server.AllowOrigins)
	return cfg
}

func applyDocument(base AppConfig, doc document) AppConfig {
	base.Site = applySite(base.Site, doc.Site)
	base.Content = applyContent(base.Content, doc.Content)
	base.Database = applyDatabase(base.Database, doc.Database)
	base.Server = applyServer(base.Server, doc.Server)
	base.CLI = applyCLI(base.CLI, doc.CLI)
	return base
}

func applyProfile(base AppConfig, frag *profileFragment) AppConfig {
	if frag == nil {
		return base
	}
	base.Site = applySite(base.Site, frag.Site)
	base.Content = applyContent(base.Content, frag.Content)
	base.Database = applyDatabase(base.Database, frag.Database)
	base.Server = applyServer(base.Server, frag.Server)
	return base
}

func applySite(base SiteSettings, frag *siteFragment) SiteSettings {
	if frag == nil {
		return base
	}
	if frag.Name != nil {
		base.Name = *frag.Name
	}
	if frag.BaseURL != nil {
		base.BaseURL = *frag.BaseURL
	}
	if frag.Description != nil {
		base.Description = *frag.Description
	}
	if frag.Keywords != nil {
		base.Keywords = cloneStrings(*frag.Keywords)
	}
	if frag.Theme != nil {
		base.Theme = applyTheme(base.Theme, frag.Theme)
	}
	return base
}

func applyTheme(base ThemeSettings, frag *themeFragment) ThemeSettings {
	if frag == nil {
		return base
	}
	if frag.Default != nil {
		base.Default = *frag.Default
	}
	if frag.AllowSystem != nil {
		base.AllowSystem = *frag.AllowSystem
	}
	if frag.Variants != nil {
		base.Variants = cloneStrings(*frag.Variants)
	}
	return base
}

func applyContent(base ContentSettings, frag *contentFragment) ContentSettings {
	if frag == nil {
		return base
	}
	if frag.Root != nil {
		base.Root = *frag.Root
	}
	if frag.StaticRoot != nil {
		base.StaticRoot = *frag.StaticRoot
	}
	if frag.Drafts != nil {
		base.Drafts = *frag.Drafts
	}
	if frag.Ignore != nil {
		base.Ignore = cloneStrings(*frag.Ignore)
	}
	return base
}

func applyDatabase(base DatabaseSettings, frag *databaseFragment) DatabaseSettings {
	if frag == nil {
		return base
	}
	if frag.Driver != nil {
		base.Driver = *frag.Driver
	}
	if frag.DSN != nil {
		base.DSN = *frag.DSN
	}
	if frag.Path != nil {
		base.Path = *frag.Path
	}
	if frag.MigrationsDir != nil {
		base.MigrationsDir = *frag.MigrationsDir
	}
	if frag.LogSQL != nil {
		base.LogSQL = *frag.LogSQL
	}
	return base
}

func applyServer(base ServerSettings, frag *serverFragment) ServerSettings {
	if frag == nil {
		return base
	}
	if frag.Host != nil {
		base.Host = *frag.Host
	}
	if frag.Port != nil {
		base.Port = *frag.Port
	}
	if frag.PublicURL != nil {
		base.PublicURL = *frag.PublicURL
	}
	if frag.TLS != nil {
		base.TLS = *frag.TLS
	}
	if frag.TLSCertPath != nil {
		base.TLSCertPath = *frag.TLSCertPath
	}
	if frag.TLSKeyPath != nil {
		base.TLSKeyPath = *frag.TLSKeyPath
	}
	if frag.AllowOrigins != nil {
		base.AllowOrigins = cloneStrings(*frag.AllowOrigins)
	}
	return base
}

func applyCLI(base CLISettings, frag *cliFragment) CLISettings {
	if frag == nil {
		return base
	}
	if frag.DefaultProfile != nil {
		base.DefaultProfile = *frag.DefaultProfile
	}
	if frag.OutputFormat != nil {
		base.OutputFormat = *frag.OutputFormat
	}
	if frag.Color != nil {
		base.Color = *frag.Color
	}
	if frag.LogLevel != nil {
		base.LogLevel = *frag.LogLevel
	}
	return base
}

