// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"vmware-samples/concourse-resource-for-marketplace/m/v2/cmd"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg/pkgfakes"
)

var _ = Describe("in", func() {
	var (
		marketplaceCLI *pkgfakes.FakeMarketplaceCLI
		stdout         *Buffer
		tempDir        string
	)

	BeforeEach(func() {
		marketplaceCLI = &pkgfakes.FakeMarketplaceCLI{}
		cmd.MarketplaceCLI = marketplaceCLI

		stdout = NewBuffer()
		cmd.InCmd.SetOut(stdout)

		var err error
		tempDir, err = ioutil.TempDir("", "test-concourse-resource-for-marketplace-in-tests-*")
		Expect(err).ToNot(HaveOccurred())

		marketplaceCLI.GetProductJSONReturns([]byte("this is the product JSON, really"), nil)
		marketplaceCLI.DownloadAssetReturns(nil)
		marketplaceCLI.GetInputSlugReturns("my-product-slug")
		marketplaceCLI.GetInputVersionReturns(&pkg.Version{VersionNumber: "1.2.3"})
	})

	AfterEach(func() {
		Expect(os.RemoveAll(tempDir)).To(Succeed())
	})

	It("gets the product json and downloads the asset", func() {
		err := cmd.InCmd.RunE(cmd.InCmd, []string{tempDir})
		Expect(err).ToNot(HaveOccurred())

		By("calling the marketplace cli to get the product JSON", func() {
			Expect(marketplaceCLI.GetProductJSONCallCount()).To(Equal(1))
		})

		By("saving the product JSON to the directory", func() {
			contents, err := ioutil.ReadFile(path.Join(tempDir, "product.json"))
			Expect(err).ToNot(HaveOccurred())
			Expect(string(contents)).To(Equal("this is the product JSON, really"))
		})

		By("calling the marketplace cli to download the asset", func() {
			Expect(marketplaceCLI.DownloadAssetCallCount()).To(Equal(1))
			Expect(marketplaceCLI.DownloadAssetArgsForCall(0)).To(Equal(tempDir))
		})

		By("printing the metadata", func() {
			Expect(marketplaceCLI.GetInputSlugCallCount()).To(Equal(1))
			Expect(marketplaceCLI.GetInputVersionCallCount()).To(Equal(2))
			var output *cmd.InOutput
			Expect(json.Unmarshal(stdout.Contents(), &output)).To(Succeed())
			Expect(output.Version.VersionNumber).To(Equal("1.2.3"))
			Expect(output.Metadata).To(HaveLen(2))
			Expect(output.Metadata[0].Name).To(Equal("slug"))
			Expect(output.Metadata[0].Value).To(Equal("my-product-slug"))
			Expect(output.Metadata[1].Name).To(Equal("version"))
			Expect(output.Metadata[1].Value).To(Equal("1.2.3"))
		})
	})

	When("getting the product JSON fails", func() {
		BeforeEach(func() {
			marketplaceCLI.GetProductJSONReturns(nil, errors.New("get product JSON failed"))
		})
		It("returns an error", func() {
			err := cmd.InCmd.RunE(cmd.InCmd, []string{tempDir})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to get product.json: get product JSON failed"))
		})
	})

	When("writing product JSON fails", func() {
		It("returns an error", func() {
			err := cmd.InCmd.RunE(cmd.InCmd, []string{"/dev/null"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to write product.json:"))
		})
	})

	When("downloading the asset fails", func() {
		BeforeEach(func() {
			marketplaceCLI.DownloadAssetReturns(errors.New("download asset failed"))
		})
		It("returns an error", func() {
			err := cmd.InCmd.RunE(cmd.InCmd, []string{tempDir})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to download asset: download asset failed"))
		})
	})
})
