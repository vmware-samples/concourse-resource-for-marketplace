// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var outCmd = &cobra.Command{
	Use: "out",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("out called")
	},
}

func init() {
	outCmd.SetOut(outCmd.OutOrStdout())
}
