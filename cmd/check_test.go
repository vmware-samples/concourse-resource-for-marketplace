// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package cmd_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"vmware-samples/concourse-resource-for-marketplace/m/v2/cmd"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg/pkgfakes"
)

var _ = Describe("check", func() {
	var (
		marketplaceCLI *pkgfakes.FakeMarketplaceCLI
		stdout         *Buffer
	)

	BeforeEach(func() {
		marketplaceCLI = &pkgfakes.FakeMarketplaceCLI{}
		cmd.MarketplaceCLI = marketplaceCLI

		stdout = NewBuffer()
		cmd.CheckCmd.SetOut(stdout)

		marketplaceCLI.GetVersionsReturns([]*pkg.Version{
			{VersionNumber: "1.1.1"},
			{VersionNumber: "2.2.2"},
		}, nil)
	})

	It("prints the list of versions", func() {
		err := cmd.CheckCmd.RunE(cmd.CheckCmd, []string{})
		Expect(err).ToNot(HaveOccurred())

		By("calling the marketplace cli", func() {
			Expect(marketplaceCLI.GetVersionsCallCount()).To(Equal(1))
		})

		By("printing the list of versions", func() {
			Expect(stdout).Should(Say("\"versionnumber\":\"1.1.1\""))
			Expect(stdout).Should(Say("\"versionnumber\":\"2.2.2\""))
		})
	})

	When("getting the list of versions fails", func() {
		BeforeEach(func() {
			marketplaceCLI.GetVersionsReturns(nil, errors.New("get versions failed"))
		})
		It("returns an error", func() {
			err := cmd.CheckCmd.RunE(cmd.CheckCmd, []string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to get the list of versions: get versions failed"))
		})
	})
})
