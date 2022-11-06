package urn

const ErrorBase = Base + ":errors:"

func Error(name string) string {
	return ErrorBase + name
}
