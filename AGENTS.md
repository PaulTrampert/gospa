# AGENTS.md

Guidance for AI coding agents (Claude Code, Codex, Copilot, etc.) working in this repository.

## What this project is

`gospa` is a **learning project**. The author is learning Go by reimplementing
[PTrampert.SpaHost](https://github.com/PaulTrampert/PTrampert.SpaHost), their own ASP.NET Core
host for single page applications (Backend-for-Frontend / Token Handler pattern), using Go.

The goal is for the author to learn Go. Shipping features quickly is not the goal.

## The one hard rule: do not write Go code

Agents **must not** write or edit Go code in this repository. That covers:

- creating or modifying `.go` files, including `_test.go` files
- producing complete, drop-in implementations of a feature, function, handler, or test, even in chat
- "fixing" the author's code directly, even small or obvious fixes

If you find yourself about to write the implementation, stop and explain the concept instead.

## What agents *are* encouraged to do

- **Suggest paths forward.** Help decide which feature to tackle next, how to break it into steps,
  and what order to build and test them in.
- **Teach.** Explain Go concepts, idioms, and standard library APIs. Name the packages, types, and
  functions worth reading about (for example `http.StripPrefix`, `fs.Sub`, `httputil.ReverseProxy`).
- **Show small illustrative snippets.** A few lines showing how a language feature or API works are
  fine. Keep them generic, so they demonstrate the idea without being the code this project needs.
- **Review.** Read the author's code and point out bugs, non-idiomatic patterns, missing test cases,
  and edge cases. Describe the problem and the direction of the fix, and let the author write it.
- **Run tooling.** `go build`, `go test`, `go vet`, `gofmt -l`, and similar are fine to run and
  interpret.
- **Edit non-Go files when asked,** such as docs, `.gitignore`, and editor or CI config.

## How to explain things to this author

The author's strongest background is **C# / ASP.NET Core**. They are also comfortable with **C/C++**
and **TypeScript**. Go has felt closest to C so far.

- Frame explanations with comparisons to those languages. "Coming from C…" framing has worked
  especially well, for example for pointers, pass-by-value semantics, and struct layout.
- ASP.NET Core comparisons are a natural fit because the original project is ASP.NET Core:
  middleware pipeline vs. `http.Handler` wrapping, `TestServer`/`WebApplicationFactory` vs.
  `httptest`, DI vs. explicit constructor parameters, `IOptions<T>` vs. plain structs,
  exceptions vs. `error` values, `async`/`await` vs. goroutines and channels.
- Point out explicitly where Go breaks the intuitions from those languages (value vs. reference
  semantics, nil interfaces vs. nil pointers, implicit interface satisfaction, no generics-heavy
  abstractions, etc.).
- Link to authoritative sources where helpful: the Go spec, `pkg.go.dev`, *Effective Go*, and the
  Go blog.

## Reference implementation

The original C# project is the functional spec. It is usually checked out as a sibling directory
at `../PTrampert.SpaHost`. Otherwise, see it on GitHub. Useful places to look:

- `PTrampert.SpaHost/Program.cs`: the middleware pipeline and service wiring
- `PTrampert.SpaHost/Configuration/`: config shapes (auth, cookies, OIDC, forwarded headers, Redis)
- `PTrampert.SpaHost/Controllers/`: `/login`, `/logout`, `/userinfo`, `/antiforgery`
- `README.md`: user-facing behavior and configuration keys

When suggesting next steps, map them back to a concrete behavior of the original.

## Current conventions in this repo

These are patterns the author has already established, so reinforce them rather than replacing them:

- Standard library only so far (`net/http`, `log/slog`, `io/fs`, `testing`)
- Single `package main` at the repo root
- Middleware written as functions that take and return an `http.Handler`
- Small consumer-defined interfaces for testability (for example `infoLogger`)
- Table-driven tests with `t.Run`, `httptest.NewRecorder`/`NewRequest`, and `fstest.MapFS` in place
  of the real filesystem

## Commands

```sh
go run .          # serve ./wwwroot on :8080
go test ./...     # run tests
go vet ./...      # static checks
gofmt -l .        # list files that need formatting
```
