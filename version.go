package version

import (
	"fmt"
	"runtime"
	"time"
)

// STEP 1: Determinate the required values

// VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
// COMMIT_HASH="$(git rev-parse --short HEAD)"
// BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

// STEP 2: Build the ldflags

// LDFLAGS=(
//   "-X '${PACKAGE}/version.Version=${VERSION}'"
//   "-X '${PACKAGE}/version.Revision=${COMMIT_HASH}'"
//   "-X '${PACKAGE}/version.BuildDate=${BUILD_TIMESTAMP}'"
// )

// STEP 3: Actual Go build process

// go build -ldflags="${LDFLAGS[*]}"

// Build information. Populated at build-time.
var (
	Version   = "dev"
	Revision  = "none"
	BuildDate = ""
	GoVersion = runtime.Version()
)

// Info returns version, branch and revision information.
func Info(name string) string {
	var date = BuildDate
	if len(date) == 0 {
		date = time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
	}
	return fmt.Sprintf("%s, version: %s, revision: %s, build date: %s, %s", name, Version, Revision, date, GoVersion)
}
