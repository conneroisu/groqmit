package git

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
)

// Diff returns the diff between the current branch and the index.
func Diff(ctx context.Context) (string, error) {
	var w, ew bytes.Buffer
	cmd := exec.CommandContext(
		ctx,
		"git",
		"diff",
		"--cached",
	)
	cmd.Stdout = &w
	cmd.Stderr = &ew
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	if ew.Len() > 0 {
		return "", errors.New(ew.String())
	}
	return w.String(), nil
}
