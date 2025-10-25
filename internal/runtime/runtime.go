package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pynezz/pynezz_com/internal/config"
)

// Environment wraps the resolved configuration for the current process.
type Environment struct {
	Bundle         config.Bundle
	Active         config.AppConfig
	ActiveProfile  string
	ProfileFromEnv bool
	Paths          []string
}

var current Environment

// Bootstrap resolves configuration using filesystem + env overrides.
func Bootstrap() (Environment, error) {
	paths := defaultPaths()
	if override := strings.TrimSpace(os.Getenv("PYNEZZ_CONFIG_PATHS")); override != "" {
		paths = parsePathList(override)
	}

	loader := config.NewLoader(paths...)
	bundle, err := loader.Load()
	if err != nil {
		return Environment{}, err
	}

	requestedProfile := strings.TrimSpace(os.Getenv("PYNEZZ_PROFILE"))
	profileFromEnv := requestedProfile != ""
	if requestedProfile == "" {
		requestedProfile = strings.TrimSpace(bundle.Base.CLI.DefaultProfile)
	}

	active := bundle.Base
	activeProfile := ""

	if requestedProfile != "" {
		if p, ok := bundle.Profiles[requestedProfile]; ok {
			active = p.Config
			activeProfile = requestedProfile
		} else if profileFromEnv {
			return Environment{}, fmt.Errorf("config: requested profile %q not found", requestedProfile)
		}
	}

	env := Environment{
		Bundle:         bundle,
		Active:         active,
		ActiveProfile:  activeProfile,
		ProfileFromEnv: profileFromEnv,
		Paths:          paths,
	}
	current = env
	return env, nil
}

// Current exposes the cached runtime environment. Bootstrap must be called first.
func Current() Environment {
	return current
}

func parsePathList(paths string) []string {
	segments := strings.Split(paths, string(os.PathListSeparator))
	out := make([]string, 0, len(segments))
	seen := map[string]struct{}{}
	for _, segment := range segments {
		trim := strings.TrimSpace(segment)
		if trim == "" {
			continue
		}
		clean := filepath.Clean(trim)
		if _, ok := seen[clean]; ok {
			continue
		}
		seen[clean] = struct{}{}
		out = append(out, clean)
	}
	return out
}

func defaultPaths() []string {
	return []string{"config"}
}

