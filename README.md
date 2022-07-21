# Marla.ONE Website

This is the codebase for my [Website](https://marla.one).

## Tech Stack

*This is the tech stack that I aim to use for the whole project.*

### Backend

- Golang v1.18 - used for static content and background tasks
- **Markdown** - website content
  - [goldmark](https://github.com/yuin/goldmark) - markdown parser
  - [goldmark-meta](https://github.com/yuin/goldmark-meta) - meta data from markdown
  - [bluemonday](https://github.com/microcosm-cc/bluemonday) - html sanitizer to keep custom web components from markdown
- [go-chi](https://github.com/go-chi/chi) - http server for static content and graphql api
- [gqlgen](https://github.com/99designs/gqlgen) - graphql server generator
- [ent](https://github.com/ent/ent) - entity framework
  - [PostgreSQL](https://www.postgresql.org/) - relational database
  - [pgx](https://github.com/jackc/pgx) - PostgreSQL driver
  - [entql](https://github.com/ent/contrib/entql) - graphql extension for ent
- [testify](https://github.com/stretchr/testify) - unit test framework
- [gocron](https://github.com/go-co-op/gocron) - for background tasks
- [cobra](https://github.com/spf13/cobra) - cli applications
- [zap](https://github.com/uber-go/zap) - B L A Z I N G _ F A S T logger
- [blocks](https://github.com/kataras/blocks/blob/main/blocks.go) - html template engine
- Social Media APIs
  - [YouTube](https://developers.google.com/youtube/v3/code_samples/go) - fetch videos
  - [Twitch](https://github.com/nicklaw5/helix) - latest VODs
  - [Instagram](https://github.com/yanatan16/golang-instagram) - posts from feed

### Frontend

- Web Components - used for dynamic content on the website
  - [Lit](https://lit.dev/) - Web Components lib
- [Vitest](https://vitest.dev/) - B L A Z I N G _ F A S T unit test framework
- [Vite](https://vitejs.dev/) lightning fast frontend tool (sadly not blazingly fast)
- [TailwindCSS](https://github.com/tailwindlabs/tailwindcss) - utility-first CSS framework
- [daisyUI](https://github.com/saadeghi/daisyui) - Tailwind CSS component library

Looks pretty over engineered at first, it probably is for the beginning, but I prefer to choose my tech stack myself instead of using a completely pre-built framework. That way I can replace or improve individual bottlenecks more easily and have more control for the future.

## Roadmap

### v0.1 - The Alpha

Basic website that just shows the legal notice needed for german law.

1. [ ] static http server
2. [ ] markdown parser (for now just the legal notice)
3. [ ] basic UI

### v1.0 - The Release

Fetch social media videos/posts from my accounts and show them on the website.

1. [ ] generate ent entities
2. [ ] graphql api for ent entities
3. [ ] fetch social media content from my channels and save it to the database
4. [ ] build UI components
5. [ ] show social media content in the UI

### v2.0 - The Upgrade

Add a blog to write cool stuff about tech and dogs.

1. [ ] markdown blog structure
2. [ ] blog UI

### v3.0

1. [ ] ???