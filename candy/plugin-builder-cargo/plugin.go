// Package buildercargo is the charly plugin serving the `cargo` builder. Every leg — the build-time
// OpResolve and the deploy-time OpCollectContext / OpReverse — is the ONE shared
// sdk.InvokeBuilder dispatch over sdk/kit, the same for every builder (opencharly/sdk#608); this
// package is only the composable selection point. Dual-placement by construction: the SAME
// NewProvider()/NewMeta() compile INTO charly in-process when listed in compiled_plugins, or
// cmd/serve serves them OUT-OF-PROCESS over go-plugin gRPC when they are not — placement is
// invisible above the registry.
package buildercargo

import (
	"context"
	"embed"

	"github.com/opencharly/sdk"
	pb "github.com/opencharly/spec/proto"
)

//go:embed schema/*.cue
var schemaFS embed.FS

// builderWord is the reserved builder word this plugin serves.
const builderWord = "cargo"

// NewProvider returns the buildercargo provider.
func NewProvider() pb.ProviderServer { return &provider{} }

// NewMeta advertises the builder:cargo capability + the plugin's self-contained CUE schema
// (via sdk.NewMeta → BuildCapabilities), compiled standalone and failing loudly if broken.
func NewMeta() pb.PluginMetaServer {
	return sdk.NewMeta("2026.182.0300",
		[]sdk.ProvidedCapability{{Class: "builder", Word: builderWord, InputDef: "#CargoBuilderInput"}},
		schemaFS)
}

type provider struct{ pb.UnimplementedProviderServer }

// Invoke serves every builder leg through the one shared dispatch.
func (provider) Invoke(_ context.Context, req *pb.InvokeRequest) (*pb.InvokeReply, error) {
	return sdk.InvokeBuilder(builderWord, req)
}
