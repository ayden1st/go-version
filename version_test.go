package version

import (
	"runtime"
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	// Freeze time
	expectedTime := time.Now().Format(time.UnixDate)

	GoVersion = runtime.Version()

	tests := []struct {
		name       string
		input      string
		version    string
		revision   string
		buildDate  string
		wantPrefix string
	}{
		{
			name:       "normal case",
			input:      "myapp",
			version:    "1.2.3",
			revision:   "abc123",
			buildDate:  "2025-01-01T00:00:00Z",
			wantPrefix: "myapp, version: 1.2.3, revision: abc123, build date: 2025-01-01T00:00:00Z, " + GoVersion,
		},
		{
			name:       "empty name",
			input:      "",
			version:    "1.0.0",
			revision:   "def456",
			buildDate:  "2025-06-01T12:00:00Z",
			wantPrefix: ", version: 1.0.0, revision: def456, build date: 2025-06-01T12:00:00Z, " + GoVersion,
		},
		{
			name:       "empty version",
			input:      "myapp",
			version:    "",
			revision:   "abc123",
			buildDate:  "2025-01-01",
			wantPrefix: "myapp, version: , revision: abc123, build date: 2025-01-01, " + GoVersion,
		},
		{
			name:       "empty revision",
			input:      "myapp",
			version:    "1.2.3",
			revision:   "",
			buildDate:  "2025-01-01",
			wantPrefix: "myapp, version: 1.2.3, revision: , build date: 2025-01-01, " + GoVersion,
		},
		{
			name:       "empty build date (uses current time)",
			input:      "myapp",
			version:    "1.2.3",
			revision:   "abc123",
			buildDate:  "",
			wantPrefix: "myapp, version: 1.2.3, revision: abc123, build date: " + expectedTime + ", " + GoVersion,
		},
		{
			name:       "all empty",
			input:      "",
			version:    "",
			revision:   "",
			buildDate:  "",
			wantPrefix: ", version: , revision: , build date: " + expectedTime + ", " + GoVersion,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock global vars
			Version = tt.version
			Revision = tt.revision
			BuildDate = tt.buildDate

			// Capture result
			got := Info(tt.input)

			// Assert
			if got != tt.wantPrefix {
				t.Errorf("Info(%q) = %q, want prefix %q", tt.input, got, tt.wantPrefix)
			}

			// Extra: ensure it contains the input name (if non-empty)
			if tt.input != "" && got[:len(tt.input)] != tt.input {
				t.Errorf("Info(%q): output must start with %q", tt.input, tt.input)
			}
		})
	}
}

// Benchmark for performance
func BenchmarkInfo(b *testing.B) {
	Version = "1.2.3"
	Revision = "abc123"
	BuildDate = "2025-01-01"
	GoVersion = runtime.Version()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Info("myapp")
	}
}
