package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoaderMergesBaseConfig(t *testing.T) {
	dir := t.TempDir()

	content := `
[site]
name = "Pynezz Labs"
base_url = "https://labs.pynezz.dev"
description = "Labs"

[content]
root = "notes"
ignore = ["drafts"]

[server]
port = 9090
public_url = "https://labs.pynezz.dev"

[cli]
output_format = "json"
`
	if err := os.WriteFile(filepath.Join(dir, "base.toml"), []byte(strings.TrimSpace(content)), 0o600); err != nil {
		t.Fatalf("write base config: %v", err)
	}

	loader := NewLoader(dir)
	bundle, err := loader.Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	if bundle.Base.Site.Name != "Pynezz Labs" {
		t.Fatalf("expected site.name to be overridden, got %q", bundle.Base.Site.Name)
	}
	if bundle.Base.Content.Root != "notes" {
		t.Fatalf("expected content.root to be notes, got %q", bundle.Base.Content.Root)
	}
	if bundle.Base.Server.Port != 9090 {
		t.Fatalf("expected server.port to be 9090, got %d", bundle.Base.Server.Port)
	}
	if bundle.Base.CLI.OutputFormat != "json" {
		t.Fatalf("expected cli.output_format to be json, got %q", bundle.Base.CLI.OutputFormat)
	}
	if len(bundle.Sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(bundle.Sources))
	}
}

func TestLoaderResolvesProfiles(t *testing.T) {
	dir := t.TempDir()

	baseContent := `
[profiles.dev.content]
root = "content/dev"

[profiles.dev.server]
port = 8081

[profiles.prod]
extends = ["dev"]

[profiles.prod.server]
tls = true
tls_cert_path = "certs/prod.crt"
tls_key_path = "certs/prod.key"
public_url = "https://pynezz.dev"
`
	if err := os.WriteFile(filepath.Join(dir, "profiles.toml"), []byte(strings.TrimSpace(baseContent)), 0o600); err != nil {
		t.Fatalf("write profiles config: %v", err)
	}

	loader := NewLoader(dir)
	bundle, err := loader.Load()
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}

	dev, ok := bundle.Profiles["dev"]
	if !ok {
		t.Fatalf("expected dev profile to be resolved")
	}
	if dev.Config.Content.Root != "content/dev" {
		t.Fatalf("expected dev content root override, got %q", dev.Config.Content.Root)
	}
	if dev.Config.Server.Port != 8081 {
		t.Fatalf("expected dev server port override, got %d", dev.Config.Server.Port)
	}

	prod, ok := bundle.Profiles["prod"]
	if !ok {
		t.Fatalf("expected prod profile to be resolved")
	}
	if !prod.Config.Server.TLS {
		t.Fatalf("expected prod to enable TLS")
	}
	if prod.Config.Server.PublicURL != "https://pynezz.dev" {
		t.Fatalf("expected prod public url override, got %q", prod.Config.Server.PublicURL)
	}
	if len(prod.Extends) != 1 || prod.Extends[0] != "dev" {
		t.Fatalf("expected prod to extend dev, got %v", prod.Extends)
	}
	if prod.Config.Content.Root != "content/dev" {
		t.Fatalf("expected prod to inherit content root from dev, got %q", prod.Config.Content.Root)
	}
}

func TestLoaderValidationError(t *testing.T) {
	dir := t.TempDir()

	content := `
[site]
name = ""
base_url = "://invalid"

[database]
driver = "unknown"
`
	if err := os.WriteFile(filepath.Join(dir, "invalid.toml"), []byte(strings.TrimSpace(content)), 0o600); err != nil {
		t.Fatalf("write invalid config: %v", err)
	}

	loader := NewLoader(dir)
	_, err := loader.Load()
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}
	if !strings.Contains(err.Error(), "config validation failed") {
		t.Fatalf("expected validation error description, got %v", err)
	}
}
