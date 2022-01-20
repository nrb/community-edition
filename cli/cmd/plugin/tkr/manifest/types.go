package manifest

// The types listed here are the manifest for TCE's included images.
// These will be translated into a TKR for consumption by the `tanzu` controllers.
// TODO: Focus on creation of package repos based on this
// TODO: Copy functionality (upstream -> projects harbor repo) Gets all of the target arches, and avoids writing our own build system
// . Crack them open to find upstream images
// Use a golden YAML file first
// Writing the YAML file by hand? Or via Go code?
// tanzu manifest add metapackage antrea
// Interface should know how to more than just translate fields
// Example: copy upstream images looks at the (Core|User)Packages to the target repo
// Example: create the package repo
// Example: push package repo

type TanzuBuild struct {
	// the version of this release and build
	Version string
	// packages to create a core package repo of
	// created repo will be uploaded to an OCI repo
	// uri of uploaded repo will be added to the generated TKR
	CorePackages []MetaPackage `yaml:"corePackages"`
	// packages to create a user-managed package repo of
	// created repo will be uploaded to an OCI repo
	// uri of uploaded repo will be added to the generated TKR
	UserPackages []MetaPackage `yaml:"userPackages"`
	// host images to add to the generated TKR
	HostImages []HostImage `yaml:"hostImages"`
	// k8s metadata added to the generated TKR
	KubernetesMeta `yaml:"kubernetesMeta"`
}
type MetaPackage struct {
	// name of the metapackage, e.g. contour
	Name string
	// each versioned package to make available
	Packages []Package
	// Embeds contents of package metadata:
	// https://carvel.dev/kapp-controller/docs/latest/packaging/#package-metadata
	PackageMetadata string `yaml:"packageMetadata"`
}
type Package struct {
	// fully qualified name of package, e.g. contour.dev.1.5.3
	Name string
	// location of the OCI bundle
	ImageBundleURI string `yaml:"imageBundleUri"`
	// Arbitrary set of options to add in the package CR
	Options map[string]interface{}
}
type HostImage struct {
	// name of the host image (ami-dfdal)
	Name string
	// the infrastructure provider it relates to (aws, azure, etc)
	Provider string
	// metadata relevant to the specific provider. (aws-region: us-west-2)
	Metadata map[string]string
}
type KubernetesMeta struct {
	// the version of Kubernetes this TKR represents
	Version string
	// the kubernetes components used in bootstrap
	// these should match what is inside the host images
	// otherwise kubeadm will pull down new images on each
	// host during bootstrap
	Components []ContainerImage
}
type ContainerImage struct { // the name of the container image
	Name string
	// the host, and url used to access the container image
	Repository string
	// the tag associated with the image
	Tag string
}
