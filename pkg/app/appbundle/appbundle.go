// The content of this file is intended to be used with go's -ldflags option to inject version control information.
// DO NOT EDIT THE VARIABLE NAMES UNLESS YOU KNOW WHAT YOU ARE DOING.

package appbundle

import "time"

var (
	// Version is the SemVer version of the binary.
	Version = "v0.0.0"

	// BuildTimeString is the time at which the application was built, in RFC3339 format.
	BuildTimeString = "1970-01-01T00:00:00Z"
)

func BuildTime() (t time.Time, err error) {
	return time.Parse(time.RFC3339, BuildTimeString)
}
