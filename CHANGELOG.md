# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Complete Datastar integration for reactive frontend components
- Enhanced dashboard with real-time statistics and error handling
- Dynamic search and filter functionality for posts list
- Interactive tags page with search, sort, and filter capabilities
- Reactive login form with loading states and validation
- Enhanced register form with real-time password validation
- Modernized passkey authentication with Datastar state management
- Retry functionality for failed API calls
- Loading states across all interactive components
- Error handling with user-friendly feedback
- Local Datastar hosting (removed CDN dependency)

### Changed
- **BREAKING**: Migrated from vanilla JavaScript to Datastar for all frontend interactions
- Updated `.air.toml` to modern Air v1.52+ configuration format
- Improved navigation burger menu with proper signal scoping
- Enhanced API endpoints to return structured data for frontend consumption
- Updated static file serving paths for better organization
- Moved navigation state (`navOpen`) from body to header for proper scoping

### Enhanced
- Dashboard now includes:
  - Real-time statistics loading
  - Tab-based navigation (Overview, Posts, Settings)
  - Error handling with retry functionality
  - Loading indicators and user feedback
- Posts list features:
  - Real-time search functionality
  - Tag-based filtering
  - Refresh capability
  - Loading states
- Tags page includes:
  - Search functionality
  - Sort by name or count
  - Statistics display
  - Loading states
- Authentication forms:
  - Real-time validation feedback
  - Loading states during submission
  - Enhanced error messages
  - Success feedback with automatic redirects

### Fixed
- Resolved Datastar 404 errors by hosting library locally
- Fixed static file path mismatches (`/static/js/` → `/js/`)
- Corrected navigation signal scope issues
- Fixed Air configuration validation errors
- Resolved template syntax errors in register form

### Technical Improvements
- Updated to Datastar v1.0.0-beta.11
- Enhanced API endpoints with proper Datastar headers
- Improved error handling patterns across all components
- Better state management with reactive signals
- Consistent loading states and user feedback
- Optimized static asset serving

### Developer Experience
- Modern Air configuration with proper hot reloading
- Improved build process with `templ generate` integration
- Better error messages and debugging information
- Consistent patterns across all Datastar components

### Dependencies
- Updated to Templ v0.3.865
- Updated Go modules to latest stable versions
- Added local Datastar v1.0.0-beta.11
- Updated webauthn libraries to latest versions

### Migration Notes
For developers working on this codebase:

1. **Datastar Integration**: All frontend interactions now use Datastar attributes instead of vanilla JavaScript
2. **Signal Management**: Component state is managed through Datastar signals with proper scoping
3. **API Patterns**: All API endpoints now return data in Datastar-compatible format with proper headers
4. **Development Setup**: Use `air` command for hot reloading during development
5. **Static Assets**: Local hosting of Datastar eliminates external CDN dependencies

### Performance Improvements
- Eliminated external CDN dependencies
- Reduced JavaScript bundle size through Datastar's efficient reactivity
- Improved DOM update performance with Datastar's targeted updates
- Better caching of static assets

---

## Previous Versions

### [Pre-Datastar] - Legacy Version
- Initial Go + Templ + Tailwind implementation
- Basic Echo server setup
- Static HTML templates
- Vanilla JavaScript interactions
- CDN-based asset loading