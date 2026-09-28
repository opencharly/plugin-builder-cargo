# AGENTS.md — plugin-builder-cargo

Standalone out-of-tree plugin repo serving the `cargo` builder word
(`builder:cargo`). The plugin is a Go module at `candy/plugin-builder-cargo/`
(module path `github.com/opencharly/plugin-builder-cargo/candy/plugin-builder-cargo`);
the root `charly.yml` only declares `discover: candy` so the repo is a project
and its candy is scanned.

Canonical files:

- `candy/plugin-builder-cargo/charly.yml` — the `plugin-builder-cargo:` candy
  entity (`plugin:` block, `plan:` check).
- `candy/plugin-builder-cargo/plugin.go` — the builder provider
  (`Invoke(OpResolve)` → `kit.BuilderResolve`, `OpCollectContext`, `OpReverse`)
  and `NewProvider()`/`NewMeta()`.
- `candy/plugin-builder-cargo/schema/cargo.cue` — the self-contained
  `#CargoBuilderInput`.
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.
- `README.md` — user overview only; never agent guidance.

## Load these skills first (R0)

- `/charly-image:image` — box/builder configuration and the builder vocabulary.
- `/charly-internals:plugin` — the plugin authoring reference: the `plugin:`
  block, the `builder` provider class, the per-plugin CUE-schema contract.
- `/charly-internals:git-workflow` — before any git/PR action.

## Build / validate / test

- `go build ./...` in `candy/plugin-builder-cargo/` — compile the plugin module.
- `go test ./...` in `candy/plugin-builder-cargo/` — the plugin's Go tests.
- `charly box validate` at the repo root — the structural check (the candy +
  `plugin:` block, CUE schema).
- The merge gate is the **org-wide** `charly/pr-validator` (required check
  `validate / validate`, defined in `opencharly/.github`); this repo has **no**
  per-repo candy gate.
- The changed path is exercised by any box build composing a candy with a
  `Cargo.toml`.

## Modify this repo

- Edit the `plugin-builder-cargo:` candy entity, the Go source, and
  `schema/cargo.cue` **together** — the schema is the single source for the
  builder's served declaration surface.
- The builder is selected by DETECTION (a candy's `Cargo.toml`), never by an
  authored `external_builder:`.

## Landing

- PR-only. Every change lands through a pull request; the org-required
  `charly/pr-validator` validates the diff and body and arms native auto-merge on
  PASS. Direct pushes to `main` are blocked.
- History lives in `CHANGELOG/` (written by `tag-on-merge` at merge time); the PR
  body IS the changelog.
- The authoritative rulebook is the umbrella `AGENTS.md` in
  `opencharly/opencharly` and `charly/AGENTS.md` in the charly repo. Do not
  restate its rules here.
