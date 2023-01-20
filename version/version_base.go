package version

var (
	GitCommit   string
	GitDescribe string

	// The compilation date. This will be filled in by the compiler.
	BuildDate string

	Version           = "1.13.0"
	VersionPrerelease = "dev1"
	VersionMetadata   = ""
)
