package cmd

import (
	"strings"
	"testing"
)

func TestGorillize(t *testing.T) {
	original := "あ"
	gorillized := gorillize([]string{"あ"})
	humanized := humanize(strings.Split(gorillized, " "))
	if original != humanized {
		t.Fatalf("unexpected result. want: %s, got: %s", original, humanized)
	}
}
