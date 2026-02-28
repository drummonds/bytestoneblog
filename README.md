# Bytestone blog


This is the thoughts of my random walk to develop bytestone.

## Content Structure

- `content/posts/` — Blog posts (appear in timeline automatically)
- `content/AFP/` — Accounting for Programmers section
- `content/slides/` — Slide decks

## Tags

Tags are standard Hugo taxonomies set in frontmatter:

```yaml
tags: ["accounting", "go"]
```

Tags in use: `ABNF`, `EBNF`, `AI`, `Accounts`, `Humphrey`, `accounting`, `bytestone`, `cli`, `colour`, `color`, `cooking`, `datacolor`, `food`, `go`, `gokrazy`, `gophoto`, `history`, `hmrc`, `htmx`, `jupyter`, `lamotrek`, `negative numbers`, `purpose`, `python`, `recipe`, `slides`, `tartan`, `tax`.

## AFP Content in Blog Timeline

AFP pages don't appear in the main blog list by default. To include an AFP page in the blog timeline, add `show_in_timeline: true` to its frontmatter:

```yaml
---
title: Your Title
date: 2026-02-28
show_in_timeline: true
---
```

These pages appear with a blue **AFP** badge in the blog list. The templates merge regular posts with AFP pages that have this flag via `union`.