// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"encoding/json"
	"os"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"

	"github.com/spf13/cobra"
)

func ValidateCheckInput(cmd *cobra.Command, args []string) error {
	var err error
	input, err = pkg.ParseFromReader(os.Stdin)
	if err != nil {
		return err
	}
	return input.ValidateSource()
}

var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check for versions",
	PreRunE: ValidateCheckInput,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		versions, err := input.GetVersions()
		if err != nil {
			return err
		}

		output, _ := json.Marshal(versions)
		cmd.Println(string(output))
		return nil
	},
}

func init() {
	checkCmd.SetOut(checkCmd.OutOrStdout())
}
