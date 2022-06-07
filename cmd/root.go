// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"os"

	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
)

var (
	cmdToExecute   string
	MarketplaceCLI pkg.MarketplaceCLI
)

func Execute() {
	var err error
	if cmdToExecute == "check" {
		err = CheckCmd.Execute()
	} else if cmdToExecute == "in" {
		err = InCmd.Execute()
	} else if cmdToExecute == "out" {
		err = outCmd.Execute()
	}
	if err != nil {
		os.Exit(1)
	}
}
