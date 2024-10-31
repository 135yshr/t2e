package cmd

import "fmt"

var (
	version = "0.0.0-dev"
	commit  = "none"    //nolint:gochecknoglobals // バージョン管理で使う
	date    = "unknown" //nolint:gochecknoglobals // バージョン管理で使う
	builtBy = "go"      //nolint:gochecknoglobals // バージョン管理で使う
)

const template = `
========================================
 __    ______
|  |_ |__    |.-----.
|   _||    __||  -__|
|____||______||_____|

========================================
  Version: %s
  Commit: %s
  Date: %s
  BuiltBy: %s
`

// Version returns the version string.
func Version() string {
	return fmt.Sprintf(template, version, commit, date, builtBy)
}
