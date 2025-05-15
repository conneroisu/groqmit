package git

import (
	"bytes"
	"context"
	"os/exec"
)

// Root returns the root directory of the git repository.
func Root(ctx context.Context) (string, error) {
	var w bytes.Buffer
	// git rev-parse --show-toplevel
	cmd := exec.CommandContext(
		ctx,
		"git",
		"rev-parse",
		"--show-toplevel",
	)
	cmd.Stdout = &w
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return w.String(), nil
}
