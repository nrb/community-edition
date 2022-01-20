package manifest

type ManifestReadWriter interface {
	// ReadManifest takes a filePath to a TCE manifest YAML file, and returns the TanzuBuild struct.
	// If the file cannot be found or read, ReadManifest will return an error.
	ReadManifest(filePath string) (*TanzuBuild, error)

	// WriteManifest will write a Tanzu Build manifest to a YAML file.
	// If the file cannot be written, the function will return an error.
	WriteManifest(build *TanzuBuild, filePath string) error
}
