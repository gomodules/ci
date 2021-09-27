package ci_test

import (
	"gomodules.xyz/ci"
	"testing"
)

func TestIsCI(t *testing.T) {
	provider, isCI := ci.IsCI()
	if !isCI {
		t.Fatalf("IsCI() got = %v, want %v", isCI, true)
	}
	if provider != ci.CIGitHubActions {
		t.Errorf("IsCI() provider got = %v, want %v", provider, ci.CIGitHubActions)
	}
}
