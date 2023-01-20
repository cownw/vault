package version

type VersionInfo struct {
	Revision          string `json:"revision,omitempty"`
	Version           string `json:"version,omitempty"`
	VersionPrerelease string `json:"version_prerelease,omitempty"`
	VersionMetadata   string `json:"version_metadata,omitempty"`
	BuildDate         string `json:"build_date,omitempty"`
}

func GetVersion() *VersionInfo {
	ver := Version
	rel := VersionPrerelease
	md := VersionMetadata
	if GitDescribe != "" {
		ver = GitDescribe
	}
	if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
		rel = "dev"
	}

	return &VersionInfo{
		Revision:          GitCommit,
		Version:           ver,
		VersionPrerelease: rel,
		VersionMetadata:   md,
		BuildDate:         BuildDate,
	}
}
