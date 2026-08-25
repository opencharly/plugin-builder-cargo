# plugin-builder-cargo

The `plugin-builder-cargo` plugin candy of the [opencharly/charly](https://github.com/opencharly/charly)
candy library, as a standalone repo (the candy de-submodule cutover, plugin
kind). The Go module lives at `candy/plugin-builder-cargo/` with module path
`github.com/opencharly/plugin-builder-cargo/candy/plugin-builder-cargo`; the charly resolver fetches this repo at the pinned tag and
the compiled-in wiring imports the module at that path.
