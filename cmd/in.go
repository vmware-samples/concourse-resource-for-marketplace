// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"

	"github.com/spf13/cobra"
)

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InOutput struct {
	Version  *pkg.Version `json:"version"`
	Metadata []*Metadata  `json:"metadata"`
}

func ValidateInInput(cmd *cobra.Command, args []string) error {
	input, err := pkg.ParseFromReader(os.Stdin)
	if err != nil {
		return err
	}

	err = input.ValidateSource()
	if err != nil {
		return err
	}

	if input.Version == nil || input.Version.VersionNumber == "" {
		return errors.New("version cannot be empty. Please set version.versionnumber")
	}

	if !input.Params.SkipDownload {
		if input.Params.Filename == "" {
			return errors.New("file name can not be empty. Please set params.filename")
		}
	}

	MarketplaceCLI = input
	return nil
}

var InCmd = &cobra.Command{
	Use:     "in",
	Short:   "Get a version",
	PreRunE: ValidateInInput,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		productJson, err := MarketplaceCLI.GetProductJSON()
		if err != nil {
			return fmt.Errorf("failed to get product.json: %w", err)
		}

		err = ioutil.WriteFile(path.Join(args[0], "product.json"), productJson, 0644)
		if err != nil {
			return fmt.Errorf("failed to write product.json: %w", err)
		}

		err = ioutil.WriteFile(path.Join(args[0], "version"), []byte(MarketplaceCLI.GetInputVersion().VersionNumber), 0644)
		if err != nil {
			return fmt.Errorf("failed to write version: %w", err)
		}

		err = MarketplaceCLI.DownloadAsset(args[0])
		if err != nil {
			return fmt.Errorf("failed to download asset: %w", err)
		}

		output := &InOutput{
			Version: MarketplaceCLI.GetInputVersion(),
			Metadata: []*Metadata{
				{Name: "slug", Value: MarketplaceCLI.GetInputSlug()},
				{Name: "version", Value: MarketplaceCLI.GetInputVersion().VersionNumber},
			},
		}
		encodedOutput, _ := json.Marshal(output)
		cmd.Println(string(encodedOutput))

		return nil
	},
}

func init() {
	InCmd.SetOut(InCmd.OutOrStdout())
}
