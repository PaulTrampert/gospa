# gospa

A Go reimplementation of [PTrampert.SpaHost](https://github.com/PaulTrampert/PTrampert.SpaHost),
a host for single page applications that implements the
[Backend for Frontend](https://learn.microsoft.com/en-us/azure/architecture/patterns/backends-for-frontends)
and Token Handler patterns.

This is primarily a **learning project**. I'm using it to learn Go by rebuilding something I
already know well in C#. Expect the code to change as I learn better ways of doing things.

## Running

Requires Go 1.27+.

```sh
go run .
```

Serves the contents of `./wwwroot` on `http://localhost:8080`.

## Testing

```sh
go test ./...
```

## Roadmap

These features come from the original SpaHost and still need to be ported:

- [x] Serve static files from `wwwroot`
- [x] Request logging
- [ ] SPA fallback: unknown routes serve `index.html`
- [ ] Layered configuration: JSON files, environment variables (`foo__bar`), `/run/secrets/*`, and
      command-line flags
- [ ] Forwarded headers support for running behind a reverse proxy
- [ ] OpenID Connect login with cookie-based sessions (`/login`, `/logout`)
- [ ] `/userinfo` endpoint returning the current user's claims
- [ ] Antiforgery protection (`/antiforgery`, `X-XSRF-Token` header)
- [ ] Optionally require authentication for static files
- [ ] API proxy under `/api/` that forwards and refreshes access tokens
- [ ] WebSocket support through the proxy
- [ ] Shared session encryption keys (Redis) for running multiple replicas
- [ ] Docker base image

## A note for AI agents

See [AGENTS.md](AGENTS.md). In short, agents may guide and teach here, but they don't write the Go
code.
