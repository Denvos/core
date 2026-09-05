package buildmetadata

type BuildMetadata struct {
	BuildTime string
	Commit    string
	Version   string
}

func New(buildTime, commit, version string) *BuildMetadata {
	return &BuildMetadata{
		BuildTime: buildTime,
		Commit:    commit,
		Version:   version,
	}
}

func (b *BuildMetadata) String() string {
	return b.Version
}
