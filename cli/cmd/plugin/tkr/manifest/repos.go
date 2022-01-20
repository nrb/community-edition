package manifest

// VerifyCrane checks for `crane` on the local $PATH.
// Not part of long term API, used because `crane` is an implementation detail at the moment
func VerifyCrane() error {
	return nil
}

type RepoCopier interface {
	// CopyAll processes a TanzuBuild manifest and copies the images and packages from one repository to another.
	CopyAll(manifest *TanzuBuild, destination string) error

	// DownloadPackage downloads a Carvel Package from an OCI registry.
	DownloadPackage(registry, outPath string) error

	// ExtractImages extracts image URI information from a Carvel Package.
	ExtractImages(pkg Package) []ContainerImage

	// CopyImages copies container images from one OCI registry to another.
	CopyImages(images []ContainerImage, destination string) error

	// CopyPackage copies an OCI package from one registry to another.
	// TODO(nrb): This is essentially the same as CopyImages, but because they use different fields to define the repository (Package.ImageBundleUri vs ContainerImage.Repository), we need different functions
	CopyPackages(pkgs []Package, destination string) error

	// RemapUri changes the registry information of a container image to the destination.
	// TODO(nrb): Should the ImageBundleUri field on Packages be changed to a ContainerImage for consistency?
	// We could do remapping for both with this function if Packages were changed.
	RemapUri(source string, destination string)
}
