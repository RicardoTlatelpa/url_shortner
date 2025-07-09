# url_shortner

## API layer
- Stateless HTTP server in Go
- Endpoints:
    - POST /shorten -> Accepts JSON: { "url": "https://eaxmple.com" }, returns: {"short_url": "https://short.ly/abc123"}
- GET /{id} -> Redirects to original URL or 404

## ID Generation
- Use uniqueidgen Snowflake-inspired library
- Ensures distributed-safe unique IDs, sortable by time
- Encoded to Base62 (to shorten)

## Storage
- Maps id -> original URL
- Options:
    - MVP: in-memory map[string]string
    - Scalable: Redis or SQL DB

## Cache
- Redis to cache hot short URLs for faster redirects

## Logging & Monitoring
- Request logs, metrics, errors
