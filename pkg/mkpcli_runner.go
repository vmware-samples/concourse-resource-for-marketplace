// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"os/exec"
)

//go:generate counterfeiter . Command
type Command interface {
	Output() ([]byte, error)
	Run() error
}

//go:generate counterfeiter . MkpcliRunner
type MkpcliRunner func(input *Input, args ...string) Command

var MakeMkpcliCommand = CommandRunner

func CommandRunner(input *Input, args ...string) Command {
	command := exec.Command("mkpcli", args...)
	command.Env = append(command.Env, "CSP_API_TOKEN="+input.Source.CSPAPIToken)
	if input.Source.MarketplaceEnv != "" {
		command.Env = append(command.Env, "MARKETPLACE_ENV="+input.Source.MarketplaceEnv)
	}
	return command
}
