/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./templates/**/*.templ",
    "./templates/**/*.go",
    "./internal/parser/*.go"
  ],
  theme: {
    extend: {},
  },
  fontFamily: {
    mono: ['"Hack Nerd Font Mono"', 'Menlo', 'Monaco', 'Consolas', 'Liberation Mono', 'Courier New', 'monospace'],
    sans: ['"Inter"', 'system-ui', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'Helvetica Neue', 'Arial', 'Noto Sans', 'sans-serif']
  },
  plugins: [require("@catppuccin/tailwindcss")({
    // prefix: "ctp",
    defaultFlavour: "mocha"
  })],
}
