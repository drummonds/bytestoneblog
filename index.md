# Bytestone Blog — Documentation

Tooling and infrastructure documentation for the [Bytestone Blog](https://www.bytestone.uk/).

## Commands

- `cmd/server` — Go file server for serving the Hugo site
- `cmd/afp2book` — Generate AFP (Accounting for Programmers) book from blog content

## Build

| Task | Description |
|---|---|
| `task public:build` | Build the Hugo blog into `public/` |
| `task docs:build` | Build this documentation into `docs/` |
| `task serve` | Run Hugo dev server |
| `task server` | Build Hugo + run Go file server locally |

## Deployment

Two statichost sites:

| Site | Content | Directory |
|---|---|---|
| `blog-bytestone` | Public blog (www.bytestone.uk) | `public/` |
| `h3-bytestoneblog` | This documentation | `docs/` |

## Links

| | |
|---|---|
| Documentation | https://h3-bytestoneblog.statichost.page/ |
| Blog | https://www.bytestone.uk/ |
| Source (Codeberg) | https://codeberg.org/hum3/bytestoneblog |
| Mirror (GitHub) | https://github.com/drummonds/bytestoneblog |
