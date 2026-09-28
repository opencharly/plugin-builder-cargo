# plugin-builder-cargo

The `cargo` builder for OpenCharly — the Rust/Cargo package builder as a plugin
`builder:` word.

A candy that ships a `Cargo.toml` is built by this builder: the host selects it
by **detection** (the candy's `Cargo.toml`), never by an authored
`external_builder:`. The provider serves the build-time inline step
(`RUN cargo install --path /ctx`) and the deploy-time teardown IR (binary
uninstall, when the binaries are known from the `Cargo.toml` `[[bin]]` entries).

## What it provides

| Capability | Surface |
|---|---|
| `builder:cargo` | the `cargo` builder word — the inline Cargo builder |

The builder authors no `plugin_input` (it is triggered by detection, not an
authored field); its self-contained `#CargoBuilderInput` (`schema/cargo.cue`) is
empty by design and ships so the schema travels with the plugin.

## How to use it

Compose the plugin candy in a box's `candy:` list:

```yaml
- '@github.com/opencharly/plugin-builder-cargo/candy/plugin-builder-cargo:<tag>'
```

Then a candy with a `Cargo.toml` is built through this builder.

## Layout

- `candy/plugin-builder-cargo/` — the plugin module: `plugin.go`,
  `schema/cargo.cue`, `cmd/serve/main.go`.
- `charly.yml` — the root project manifest (`discover: candy`).
- `.github/workflows/tag-on-merge.yml` — CalVer tag + `CHANGELOG/` on merge.

## Related

- Owning skill: `/charly-image:image` — box/builder configuration and the box
  dependency graph. This candy carries no `skill:` entity of its own; the gap is
  tracked in
  [opencharly/opencharly#291](https://github.com/opencharly/opencharly/issues/291).
- `/charly-internals:plugin` — the plugin/provider model, including the
  `builder` provider class.
- [`opencharly/charly`](https://github.com/opencharly/charly) — the charly CLI.
