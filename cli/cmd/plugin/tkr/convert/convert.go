package convert

import (
	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/tkr/manifest"
	tkrtypes "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkr/pkg/types"
)

func ManifestToTKR(manifest *manifest.TanzuBuild) tkrtypes.Bom {
	tkr := tkrtypes.Bom{}

	return tkr
}

func TKRToManifest(tkr tkrtypes.Bom) *manifest.TanzuBuild {
	m := manifest.TanzuBuild{}

	return &m
}
