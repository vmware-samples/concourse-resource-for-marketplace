// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"

	"github.com/spf13/cobra"
)

func ValidateCheckInput(cmd *cobra.Command, args []string) error {
	input, err := pkg.ParseFromReader(os.Stdin)
	if err != nil {
		return err
	}
	err = input.ValidateSource()
	if err != nil {
		return err
	}
	MarketplaceCLI = input
	return nil
}

var CheckCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check for versions",
	PreRunE: ValidateCheckInput,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		versions, err := MarketplaceCLI.GetVersions()
		if err != nil {
			return fmt.Errorf("failed to get the list of versions: %w", err)
		}

		output, _ := json.Marshal(versions)
		if err != nil {
			return fmt.Errorf("failed to encode the list of versions: %w", err)
		}

		cmd.Println(string(output))
		return nil
	},
}

func init() {
	CheckCmd.SetOut(CheckCmd.OutOrStdout())
}
