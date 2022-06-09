// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

//go:generate counterfeiter . MarketplaceCLI
type MarketplaceCLI interface {
	GetProductJSON() ([]byte, error)
	GetVersions() ([]*Version, error)
	DownloadAsset(folder string) error

	GetInputVersion() *Version
	GetInputSlug() string
}

func (i *Input) run(args []string) ([]byte, error) {
	command := MakeMkpcliCommand(i, args...)
	results, err := command.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil, fmt.Errorf("failed to run mkpcli %s:\n%s\n%w", strings.Join(args, " "), string(exitErr.Stderr), err)
		}
		return nil, fmt.Errorf("failed to run mkpcli %s: %w", strings.Join(args, " "), err)
	}
	return results, nil
}

func (i *Input) GetProductJSON() ([]byte, error) {
	args := []string{"product", "get", "--output", "json", "--product", i.Source.ProductSlug, "--product-version", i.Version.VersionNumber}
	return i.run(args)
}

func (i *Input) GetVersions() ([]*Version, error) {
	args := []string{"product", "list-versions", "--output", "json", "--product", i.Source.ProductSlug}
	results, err := i.run(args)
	if err != nil {
		return nil, err
	}

	var versions []*Version
	err = json.Unmarshal(results, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to parse output for mkpcli %s: %w", strings.Join(args, " "), err)
	}

	Sort(versions)
	return GetOnlySince(versions, i.Version), nil
}

func (i *Input) DownloadAsset(folder string) error {
	if i.Params.SkipDownload {
		return nil
	}

	filePath := path.Join(folder, i.Params.Filename)
	args := []string{"download", "--product", i.Source.ProductSlug, "--product-version", i.Version.VersionNumber, "--filename", filePath}

	if i.Params.AcceptEula {
		args = append(args, "--accept-eula")
	}
	if i.Params.Filter != "" {
		args = append(args, "--filter", i.Params.Filter)
	}

	_, err := i.run(args)
	return err
}

func (i *Input) GetInputVersion() *Version {
	return i.Version
}

func (i *Input) GetInputSlug() string {
	return i.Source.ProductSlug
}
