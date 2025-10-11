package version

var (
	// version is the version of the application. Set via ldflags.
	version = "dev"
	// commit is the git commit hash. Set via ldflags.
	commit = "unknown"
)

// Version returns the application version.
func Version() string {
	return version
}

// Commit returns the git commit hash.
func Commit() string {
	return commit
}

// FullVersion returns the full version string including commit.
func FullVersion() string {
	return version + " (" + commit + ")"
}
