// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Input struct {
	Source  *SourceInput `json:"source"`
	Version *Version     `json:"version"`
	Params  *ParamsInput `json:"params"`
}

type SourceInput struct {
	CSPAPIToken    string `json:"csp_api_token"`
	MarketplaceEnv string `json:"marketplace_env"`
	ProductSlug    string `json:"product_slug"`
}

type ParamsInput struct {
	AcceptEula   bool   `json:"accept_eula"`
	Filter       string `json:"filter"`
	Filename     string `json:"filename"`
	SkipDownload bool   `json:"skip_download"`
}

func ParseFromReader(inputReader io.Reader) (*Input, error) {
	var input *Input
	decoder := json.NewDecoder(inputReader)
	err := decoder.Decode(&input)
	if err != nil {
		return nil, fmt.Errorf("failed to decode input params: %w", err)
	}
	return input, nil
}

func (i *Input) ValidateSource() error {
	if i.Source == nil {
		return errors.New("source was not defined")
	}
	if i.Source.CSPAPIToken == "" {
		return errors.New("CSP API token must be defined. Please set source.csp_api_token")
	}
	if i.Source.ProductSlug == "" {
		return errors.New("marketplace product slug must be defined. Please set source.product_slug")
	}
	return nil
}
