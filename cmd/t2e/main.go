package main

import "log/slog"

var (
	version = "0.0.0-dev"
	commit  = "none"    //nolint:gochecknoglobals // バージョン管理で使う
	date    = "unknown" //nolint:gochecknoglobals // バージョン管理で使う
	builtBy = "go"      //nolint:gochecknoglobals // バージョン管理で使う
)

func main() {
	slog.Info("Hello, World!", "version", version, "commit", commit, "date", date, "builtBy", builtBy)
}
