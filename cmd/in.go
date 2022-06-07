// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"errors"
	"os"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"

	"github.com/spf13/cobra"
)

func ValidateInInput(cmd *cobra.Command, args []string) error {
	var err error
	input, err = pkg.ParseFromReader(os.Stdin)
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
	return nil
}

var inCmd = &cobra.Command{
	Use:     "in",
	Short:   "Get a version",
	PreRunE: ValidateInInput,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		err := input.DownloadAsset()
		if err != nil {
			return err
		}

		//output, _ := json.Marshal(versions)
		//cmd.Println(string(output))
		return nil
	},
}

func init() {
	inCmd.SetOut(inCmd.OutOrStdout())
}
