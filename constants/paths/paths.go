package paths

import (
	"os"
	"path/filepath"
)

var (
	HomeDir, _   = os.UserHomeDir()
	CacheDir, _  = os.UserCacheDir()
	ConfigDir, _ = os.UserConfigDir()
	DataDir, _   = os.UserConfigDir()
	TempDir      = os.TempDir()
)

const (
	Separator     = string(os.PathSeparator)
	ListSeparator = string(os.PathListSeparator)
)

var (
	DenvosHome     = filepath.Join(HomeDir, ".denvos")
	DenvosCache    = filepath.Join(CacheDir, "denvos")
	DenvosConfig   = filepath.Join(ConfigDir, "denvos")
	DenvosData     = filepath.Join(DataDir, "denvos")
	DenvosLogs     = filepath.Join(DenvosHome, "logs")
	DenvosPlugins  = filepath.Join(DenvosHome, "plugins")
	DenvosTemplates = filepath.Join(DenvosHome, "templates")
)
