// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"os"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
)

var (
	cmdToExecute string
	input        *pkg.Input
)

func Execute() {
	var err error
	if cmdToExecute == "check" {
		err = checkCmd.Execute()
	} else if cmdToExecute == "in" {
		err = inCmd.Execute()
	} else if cmdToExecute == "out" {
		err = outCmd.Execute()
	}
	if err != nil {
		os.Exit(1)
	}
}
