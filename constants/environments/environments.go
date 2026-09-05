package environments

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
	Test        = "test"
	Local       = "local"
)

var All = []string{
	Development,
	Staging,
	Production,
	Test,
	Local,
}

func IsValid(env string) bool {
	for _, e := range All {
		if e == env {
			return true
		}
	}
	return false
}
