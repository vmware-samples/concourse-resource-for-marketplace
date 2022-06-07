// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func (i *Input) GetProductJSON() ([]byte, error) {
	args := []string{"product", "get", "--output", "json", "--product", i.Source.ProductSlug, "--product-version", i.Version.VersionNumber}
	getVersions := MakeMkpcliCommand(i, args...)
	results, err := getVersions.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil, fmt.Errorf("failed to run mkpcli %s:\n%s\n%w", strings.Join(args, " "), string(exitErr.Stderr), err)
		}
		return nil, fmt.Errorf("failed to run mkpcli %s: %w", strings.Join(args, " "), err)
	}
	return results, nil
}

func (i *Input) GetVersions() ([]*Version, error) {
	args := []string{"product", "list-versions", "--output", "json", "--product", i.Source.ProductSlug}
	getVersions := MakeMkpcliCommand(i, args...)
	results, err := getVersions.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil, fmt.Errorf("failed to run mkpcli %s:\n%s\n%w", strings.Join(args, " "), string(exitErr.Stderr), err)
		}
		return nil, fmt.Errorf("failed to run mkpcli %s: %w", strings.Join(args, " "), err)
	}

	var versions []*Version
	err = json.Unmarshal(results, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to parse output for mkpcli %s: %w", strings.Join(args, " "), err)
	}

	Sort(versions)
	return GetOnlySince(versions, i.Version), nil
}

func (i *Input) DownloadAsset() error {
	return nil
}
