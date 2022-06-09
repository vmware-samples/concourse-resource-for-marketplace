// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	// In params
	AcceptEula   bool   `json:"accept_eula"`
	Filter       string `json:"filter"`
	Filename     string `json:"filename"`
	SkipDownload bool   `json:"skip_download"`

	// Out params
	ProductVersion     string `json:"product_version"`
	ProductVersionFile string `json:"product_version_file"`
	AssetType          string `json:"asset_type"`
	File               string `json:"file"`
	Chart              string `json:"chart"`
	Instructions       string `json:"instructions"`
	InstructionsFile   string `json:"instructions_file"`
	ImageTag           string `json:"image_tag"`
	ImageTagFile       string `json:"image_tag_file"`
	ImageTagType       string `json:"image_tag_type"`

	// for mkpcli attach <asset type>
	CreateVersion bool `json:"create_version"`

	// for mkpcli product set
	//OSLFile string `json:"osl_file"`
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

func StringOrFile(stringValue, filePath string) (string, error) {
	if stringValue != "" {
		return stringValue, nil
	} else {
		fileContents, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
		}
		return string(fileContents), nil
	}
}
