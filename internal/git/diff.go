package git

import (
	"bytes"
	"errors"
	"os/exec"
)

// GetDiff returns the diff between the current branch and the index.
func GetDiff() (string, error) {
	var w, ew bytes.Buffer
	cmd := exec.Command(
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
