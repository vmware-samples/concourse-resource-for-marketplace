// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"

	"github.com/spf13/cobra"
)

func ValidateOutInput(cmd *cobra.Command, args []string) error {
	input, err := pkg.ParseFromReader(os.Stdin)
	if err != nil {
		return err
	}

	if input.Params.ProductVersion == "" && input.Params.ProductVersionFile == "" {
		return errors.New("no product version specified. Please set either params.product_version or params.product_version_file")
	}

	if input.Params.ProductVersion != "" && input.Params.ProductVersionFile != "" {
		return errors.New("both params.product_version and params.product_version_file cannot be specified at the same time. Please choose one or the other")
	}

	validAssetTypes := []string{"chart", "image", "vm"}
	if input.Params.AssetType == "" {
		return fmt.Errorf("no asset type specified. Please set params.asset_type to one of %s", strings.Join(validAssetTypes, ", "))
	}
	validAssetType := false
	for _, assetType := range validAssetTypes {
		if input.Params.AssetType == assetType {
			validAssetType = true
			break
		}
	}
	if !validAssetType {
		return fmt.Errorf("asset type \"%s\" is not valid. Please set params.asset_type to one of %s", input.Params.AssetType, strings.Join(validAssetTypes, ", "))
	}

	MarketplaceCLI = input
	return nil
}

var outCmd = &cobra.Command{
	Use:     "out",
	PreRunE: ValidateOutInput,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	outCmd.SetOut(outCmd.OutOrStdout())
}
