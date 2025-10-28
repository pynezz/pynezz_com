package templates

const cpsHeaders = `"Content-Security-Policy", "default-src 'self'; script-src 'self' 'nonce-"+nonce+"'; style-src 'self'; img-src 'self' *.github.com; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; frame-src 'none'; base-uri 'self'; block-all-mixed-content; upgrade-insecure-requests;")`

var _ = cpsHeaders
