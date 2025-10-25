package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
)

// Loader reads TOML configuration snippets and produces a merged bundle.
type Loader struct {
	paths []string
}

// NewLoader constructs a loader that will search the provided paths in order.
func NewLoader(paths ...string) *Loader {
	cleaned := make([]string, 0, len(paths))
	seen := map[string]struct{}{}
	for _, p := range paths {
		if p == "" {
			continue
		}
		cp := filepath.Clean(p)
		if _, ok := seen[cp]; ok {
			continue
		}
		seen[cp] = struct{}{}
		cleaned = append(cleaned, cp)
	}
	return &Loader{paths: cleaned}
}

// Load parses every TOML file found under the configured paths and returns the merged bundle.
func (l *Loader) Load() (Bundle, error) {
	var bundle Bundle
	bundle.Base = DefaultConfig()
	bundle.Profiles = map[string]ProfileConfig{}

	files, err := l.collectFiles()
	if err != nil {
		return bundle, err
	}

	if len(files) == 0 {
		return bundle, nil
	}

	partials := partialBundle{
		base:     document{},
		profiles: map[string]*profileFragment{},
	}

	for _, file := range files {
		data, readErr := os.ReadFile(file)
		if readErr != nil {
			return bundle, fmt.Errorf("config: read %s: %w", file, readErr)
		}

		var doc document
		if err := toml.Unmarshal(data, &doc); err != nil {
			return bundle, fmt.Errorf("config: parse %s: %w", file, err)
		}

		partials.merge(doc)
		bundle.Sources = append(bundle.Sources, file)
	}

	bundle.Base = applyDocument(bundle.Base, partials.base)

	resolver := newProfileResolver(bundle.Base, partials.profiles)
	for name := range partials.profiles {
		cfg, err := resolver.resolve(name)
		if err != nil {
			return bundle, err
		}
		frag := partials.profiles[name]
		bundle.Profiles[name] = ProfileConfig{
			Name:    name,
			Extends: cloneStrings(frag.Extends),
			Config:  cfg,
		}
	}

	var validationErrors []error
	if err := validateConfig("base", bundle.Base); err != nil {
		validationErrors = append(validationErrors, err)
	}

	for name, profile := range bundle.Profiles {
		if err := validateConfig("profile:"+name, profile.Config); err != nil {
			validationErrors = append(validationErrors, err)
		}
	}

	if len(validationErrors) > 0 {
		return bundle, errors.Join(validationErrors...)
	}

	return bundle, nil
}

func (l *Loader) collectFiles() ([]string, error) {
	if len(l.paths) == 0 {
		return nil, nil
	}

	var files []string
	seen := map[string]struct{}{}

	for _, root := range l.paths {
		info, err := os.Stat(root)
		if err != nil {
			return nil, fmt.Errorf("config: stat %s: %w", root, err)
		}

		if info.IsDir() {
			err = filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}
				if d.IsDir() {
					return nil
				}
				if !isTOML(path) {
					return nil
				}
				clean := filepath.Clean(path)
				if _, ok := seen[clean]; ok {
					return nil
				}
				seen[clean] = struct{}{}
				files = append(files, clean)
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("config: walk %s: %w", root, err)
			}
		} else {
			if !isTOML(root) {
				continue
			}
			clean := filepath.Clean(root)
			if _, ok := seen[clean]; ok {
				continue
			}
			seen[clean] = struct{}{}
			files = append(files, clean)
		}
	}

	sort.Strings(files)
	return files, nil
}

func isTOML(path string) bool {
	return strings.EqualFold(filepath.Ext(path), ".toml")
}

type partialBundle struct {
	base     document
	profiles map[string]*profileFragment
}

func (pb *partialBundle) merge(doc document) {
	pb.base.merge(doc)
	for name, frag := range doc.Profiles {
		existing := pb.profiles[name]
		if existing == nil {
			pb.profiles[name] = cloneProfileFragment(frag)
			continue
		}
		existing.merge(frag)
	}
}

type profileResolver struct {
	base      AppConfig
	fragments map[string]*profileFragment
	resolved  map[string]AppConfig
	inFlight  map[string]bool
}

func newProfileResolver(base AppConfig, fragments map[string]*profileFragment) *profileResolver {
	return &profileResolver{
		base:      base,
		fragments: fragments,
		resolved:  map[string]AppConfig{},
		inFlight:  map[string]bool{},
	}
}

func (r *profileResolver) resolve(name string) (AppConfig, error) {
	if cfg, ok := r.resolved[name]; ok {
		return cloneAppConfig(cfg), nil
	}

	frag, ok := r.fragments[name]
	if !ok {
		return AppConfig{}, fmt.Errorf("config: unknown profile %q", name)
	}

	if r.inFlight[name] {
		return AppConfig{}, fmt.Errorf("config: cyclic profile inheritance at %q", name)
	}

	r.inFlight[name] = true

	var cfg AppConfig
	if len(frag.Extends) == 0 {
		cfg = cloneAppConfig(r.base)
	} else {
		if len(frag.Extends) > 1 {
			return AppConfig{}, fmt.Errorf("config: profile %q extends multiple parents (%v) which is not supported yet", name, frag.Extends)
		}
		parent := frag.Extends[0]
		parentCfg, err := r.resolve(parent)
		if err != nil {
			return AppConfig{}, err
		}
		cfg = cloneAppConfig(parentCfg)
	}

	cfg = applyProfile(cfg, frag)

	r.inFlight[name] = false
	r.resolved[name] = cloneAppConfig(cfg)
	return cfg, nil
}
