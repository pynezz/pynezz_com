# Pynezz.dev V2 Architecture Plan

## Scope & Mandate

- Revamp the Go-based CLI + web stack to deliver a production-ready content workflow.
- Keep CLI-driven authoring (no web editor) while exposing authenticated APIs/webhooks for automation.
- Ship a polished, responsive frontend with theming, Safari parity, and installable PWA experience.
- Provide composable TOML snippets for configuration and Podman-friendly deployment artifacts.

## Current Pain Points

- CMS commands are mostly stubs; publish pipeline (MD → HTML → DB) is unreliable.
- Parser lacks deterministic output, metadata validation, and structured error handling.
- Server couples presentation/data access and exposes limited APIs; auth story unfinished.
- UI is desktop-first, lacks theme control, manifest/service worker, and Safari compatibility fixes.
- Build/config story is ad-hoc: no external config format, no container image definition, Tailwind/templ steps undocumented.

## V2 Pillars

1. **Content Pipeline**
   - Declarative TOML config snippets (site, database, content roots); merge at startup.
   - Strongly typed parser with markdown linting, front-matter validation, slug management.
   - Deterministic CLI commands (`parse`, `publish`, `unpublish`, `list`, `delete`, `config`, `tags`, `status`) with structured output (text/JSON).
   - Transactional persistence: use GORM with migrations, hashed checksums, draft/published states.

2. **Platform/CLI**
   - `pynezz` CLI with modular subcommands (`cms`, `serve`, `config`, `version`) and shared flag parsing.
   - Configurable profiles (dev/preview/prod) + environment overrides.
   - CLI integration with API/webhook (e.g., `cms publish --push` triggers webhook).

3. **API & Automation**
   - REST(ish) JSON API under `/api/v2` with JWT + signed webhook secrets.
   - Endpoints: health, posts CRUD, tags, publish action, inbound markdown upload.
   - Webhook receiver for `publish` actions with replay protection (HMAC + expiry).
   - Rate limiting + structured logging middleware.

4. **Frontend Experience**
   - Component library built with templ + Tailwind.
   - Responsive layout strategy (mobile-first, flex/grid, accessible landmarks).
   - Theme system (light/dark + Catppuccin variants) with persisted preference & Safari-safe CSS.
   - PWA assets: manifest, service-worker caching, offline shell, install prompts.
   - Content rendering from database with incremental revalidation hooks.

5. **Operations & Tooling**
   - Podman `Containerfile` + `podman-compose` for app + assets build stages.
   - Makefile targets updated (`make dev`, `make build`, `make lint`, `make podman-run`).
   - GitHub-style hook examples (pre-commit `go fmt`, `templ generate`, Tailwind build).
   - Observability hooks (structured logs, health/readiness endpoints).

## Milestones / Deliverables

1. **Foundation**
   - ~~Finalize config schema; implement loader/validator package.~~
   - Refactor parser & models; ensure CLI commands operate end-to-end locally.
2. **Automation Layer**
   - Build `/api/v2` routes, JWT middleware, webhook HMAC signing.
   - Update CLI to leverage API (opt-in) and add JSON output mode.
3. **Frontend & PWA**
   - Rebuild templates, add theme switcher, ensure Tailwind CSS for Safari.
   - Add manifest, icons, service worker, offline caching strategy.
4. **Container & Docs**
   - Provide Podman definitions, updated Make targets, runbooks, README refresh.
   - Add smoke tests (Go + npm) and CI-ready scripts.

## Open Questions

- Do we need multi-language support (i18n) in V2?
- Should webhooks queue and retry or remain fire-and-forget?
- Preferred secret storage mechanism (env vars vs. files) for prod?

## Next Actions

- Flesh out CMS command UX with structured text/JSON output and config-driven defaults.
- Establish integration tests for parser → DB → templates.
- Drive server/middleware bootstrapping from the new configuration bundle (DB, auth, assets).

## Progress Log

- ✅ **Config baseline**: Added TOML-driven configuration loader with profile inheritance (`internal/config`) plus sample snippets in `config/`. CLI bootstrap respects `PYNEZZ_CONFIG_PATHS` and `PYNEZZ_PROFILE` overrides.
- ✅ **CLI tools**: Added `pynezz config` inspector (text/JSON) and profile listing; `serve` and CMS parsing now respect configuration paths/ignore rules.
