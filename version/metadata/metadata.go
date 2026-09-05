package metadata

type Metadata struct {
	BuildTime string
	Commit    string
	GoVersion string
	OS        string
	Arch      string
}

var Global = &Metadata{}

func Set(buildTime, commit, goVersion, os, arch string) {
	Global.BuildTime = buildTime
	Global.Commit = commit
	Global.GoVersion = goVersion
	Global.OS = os
	Global.Arch = arch
}
