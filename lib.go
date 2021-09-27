package ci

import (
	"os"
)

type CI string

const (
	CIGitHubActions CI = "GitHubActions"
	CIUnknown       CI = "Unknown"
)

func IsCI() (CI, bool) {
	if v := os.Getenv("GITHUB_ACTIONS"); v == "true" {
		return CIGitHubActions, true
	}
	return CIUnknown, false
}
