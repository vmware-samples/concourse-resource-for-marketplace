// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg_test

import (
	"encoding/json"
	"errors"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg"
	"vmware-samples/concourse-resource-for-marketplace/m/v2/pkg/pkgfakes"
)

type MarketplaceCLIVersions struct {
	Number  string `json:"versionnumber"`
	Details string `json:"versiondetails"`
	Status  string `json:"status,omitempty"`
}

var _ = Describe("Marketplace CLI", func() {
	var (
		input              *pkg.Input
		mkpcliCommandMaker *pkgfakes.FakeMkpcliRunner
		mkpcliCommand      *pkgfakes.FakeCommand
	)

	BeforeEach(func() {
		mkpcliCommand = &pkgfakes.FakeCommand{}

		mkpcliCommandMaker = &pkgfakes.FakeMkpcliRunner{}
		mkpcliCommandMaker.Returns(mkpcliCommand)
		pkg.MakeMkpcliCommand = mkpcliCommandMaker.Spy
	})

	Describe("GetProductJSON", func() {
		BeforeEach(func() {
			input = &pkg.Input{
				Source: &pkg.SourceInput{
					CSPAPIToken: "my-csp-api-token",
					ProductSlug: "my-product",
				},
				Version: &pkg.Version{VersionNumber: "1.2.3"},
			}

			mkpcliCommand.OutputReturns([]byte("{\"slug\": \"my-product\"}"), nil)
		})

		It("returns the product JSON", func() {
			productData, err := input.GetProductJSON()
			Expect(err).ToNot(HaveOccurred())

			Expect(string(productData)).To(Equal("{\"slug\": \"my-product\"}"))

			By("constructing the right mkpcli command", func() {
				Expect(mkpcliCommandMaker.CallCount()).To(Equal(1))
				inputArg, cliArgs := mkpcliCommandMaker.ArgsForCall(0)
				Expect(inputArg).To(Equal(input))
				Expect(cliArgs).To(ConsistOf("product", "get", "--output", "json", "--product", "my-product", "--product-version", "1.2.3"))
			})

			By("invoking the mkpcli", func() {
				Expect(mkpcliCommand.OutputCallCount()).To(Equal(1))
			})
		})

		When("the command fails", func() {
			BeforeEach(func() {
				exitErr := &exec.ExitError{
					Stderr: []byte("a totally expected error occurred"),
				}
				mkpcliCommand.OutputReturns(nil, exitErr)
			})
			It("returns an error", func() {
				_, err := input.GetProductJSON()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to run mkpcli product get --output json --product my-product --product-version 1.2.3:\na totally expected error occurred\n<nil>"))
			})
			When("the error is not an ExitError", func() {
				BeforeEach(func() {
					mkpcliCommand.OutputReturns(nil, errors.New("a totally generic error occurred"))
				})
				It("returns a less helpful error", func() {
					_, err := input.GetProductJSON()
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("failed to run mkpcli product get --output json --product my-product --product-version 1.2.3: a totally generic error occurred"))
				})
			})
		})
	})

	Describe("GetVersions", func() {
		BeforeEach(func() {
			input = &pkg.Input{
				Source: &pkg.SourceInput{
					CSPAPIToken: "my-csp-api-token",
					ProductSlug: "my-product",
				},
			}

			versions := []*MarketplaceCLIVersions{
				{
					Number:  "1.1.1",
					Status:  "PENDING",
					Details: "1.1.1 details",
				},
				{
					Number:  "3.3.3",
					Status:  "PENDING",
					Details: "3.3.3 details",
				},
				{
					Number:  "2.2.2",
					Status:  "PENDING",
					Details: "2.2.2 details",
				},
			}
			versionDetails, err := json.Marshal(versions)
			Expect(err).ToNot(HaveOccurred())
			mkpcliCommand.OutputReturns(versionDetails, nil)
		})

		It("gets the sorted list of versions", func() {
			versions, err := input.GetVersions()
			Expect(err).ToNot(HaveOccurred())

			Expect(versions).To(HaveLen(3))
			Expect(versions[0].VersionNumber).To(Equal("1.1.1"))
			Expect(versions[1].VersionNumber).To(Equal("2.2.2"))
			Expect(versions[2].VersionNumber).To(Equal("3.3.3"))

			By("constructing the right mkpcli command", func() {
				Expect(mkpcliCommandMaker.CallCount()).To(Equal(1))
				inputArg, cliArgs := mkpcliCommandMaker.ArgsForCall(0)
				Expect(inputArg).To(Equal(input))
				Expect(cliArgs).To(ConsistOf("product", "list-versions", "--output", "json", "--product", "my-product"))
			})

			By("invoking the mkpcli", func() {
				Expect(mkpcliCommand.OutputCallCount()).To(Equal(1))
			})
		})

		When("an existing version is specified", func() {
			BeforeEach(func() {
				input.Version = &pkg.Version{
					VersionNumber: "2.2.2",
				}
			})
			It("only returns the versions since that one", func() {
				versions, err := input.GetVersions()
				Expect(err).ToNot(HaveOccurred())

				Expect(versions).To(HaveLen(2))
				Expect(versions[0].VersionNumber).To(Equal("2.2.2"))
				Expect(versions[1].VersionNumber).To(Equal("3.3.3"))
			})

			When("the specified version does not exist", func() {
				BeforeEach(func() {
					input.Version = &pkg.Version{
						VersionNumber: "9.9.9",
					}
				})
				It("returns all versions", func() {
					versions, err := input.GetVersions()
					Expect(err).ToNot(HaveOccurred())

					Expect(versions).To(HaveLen(3))
					Expect(versions[0].VersionNumber).To(Equal("1.1.1"))
					Expect(versions[1].VersionNumber).To(Equal("2.2.2"))
					Expect(versions[2].VersionNumber).To(Equal("3.3.3"))
				})
			})
		})

		When("the command fails", func() {
			BeforeEach(func() {
				exitErr := &exec.ExitError{
					Stderr: []byte("a totally expected error occurred"),
				}
				mkpcliCommand.OutputReturns(nil, exitErr)
			})
			It("returns an error", func() {
				_, err := input.GetVersions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to run mkpcli product list-versions --output json --product my-product:\na totally expected error occurred\n<nil>"))
			})
			When("the error is not an ExitError", func() {
				BeforeEach(func() {
					mkpcliCommand.OutputReturns(nil, errors.New("a totally generic error occurred"))
				})
				It("returns a less helpful error", func() {
					_, err := input.GetVersions()
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("failed to run mkpcli product list-versions --output json --product my-product: a totally generic error occurred"))
				})
			})
		})

		When("the output is not parseable", func() {
			BeforeEach(func() {
				mkpcliCommand.OutputReturns([]byte("this is not parseable JSON"), nil)
			})
			It("returns an error", func() {
				_, err := input.GetVersions()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to parse output for mkpcli product list-versions --output json --product my-product: invalid character 'h' in literal true (expecting 'r')"))
			})
		})
	})

	Describe("DownloadAsset", func() {
		BeforeEach(func() {
			input = &pkg.Input{
				Source: &pkg.SourceInput{
					CSPAPIToken: "my-csp-api-token",
					ProductSlug: "my-product",
				},
				Version: &pkg.Version{VersionNumber: "1.2.3"},
				Params: &pkg.ParamsInput{
					Filename: "kubernetes-1.22.8.ova",
				},
			}
			mkpcliCommand.OutputReturns(nil, nil)
		})

		It("downloads the asset", func() {
			err := input.DownloadAsset("/path/to/download/folder")
			Expect(err).ToNot(HaveOccurred())

			By("constructing the right mkpcli command", func() {
				Expect(mkpcliCommandMaker.CallCount()).To(Equal(1))
				inputArg, cliArgs := mkpcliCommandMaker.ArgsForCall(0)
				Expect(inputArg).To(Equal(input))
				Expect(cliArgs).To(ConsistOf("download", "--product", "my-product", "--product-version", "1.2.3", "--filename", "/path/to/download/folder/kubernetes-1.22.8.ova"))
			})

			By("invoking the mkpcli", func() {
				Expect(mkpcliCommand.OutputCallCount()).To(Equal(1))
			})
		})

		When("the filter is also given", func() {
			BeforeEach(func() {
				input.Params.Filter = "kubernetes"
			})
			It("downloads the asset with the filter", func() {
				err := input.DownloadAsset("/path/to/download/folder")
				Expect(err).ToNot(HaveOccurred())

				By("constructing the right mkpcli command", func() {
					Expect(mkpcliCommandMaker.CallCount()).To(Equal(1))
					inputArg, cliArgs := mkpcliCommandMaker.ArgsForCall(0)
					Expect(inputArg).To(Equal(input))
					Expect(cliArgs).To(ConsistOf("download", "--product", "my-product", "--product-version", "1.2.3", "--filename", "/path/to/download/folder/kubernetes-1.22.8.ova", "--filter", "kubernetes"))
				})
			})
		})

		When("the command fails", func() {
			BeforeEach(func() {
				exitErr := &exec.ExitError{
					Stderr: []byte("a totally expected error occurred"),
				}
				mkpcliCommand.OutputReturns(nil, exitErr)
			})
			It("returns an error", func() {
				err := input.DownloadAsset("/path/to/download/folder")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to run mkpcli download --product my-product --product-version 1.2.3 --filename /path/to/download/folder/kubernetes-1.22.8.ova:\na totally expected error occurred\n<nil>"))
			})
			When("the error is not an ExitError", func() {
				BeforeEach(func() {
					mkpcliCommand.OutputReturns(nil, errors.New("a totally generic error occurred"))
				})
				It("returns a less helpful error", func() {
					err := input.DownloadAsset("/path/to/download/folder")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("failed to run mkpcli download --product my-product --product-version 1.2.3 --filename /path/to/download/folder/kubernetes-1.22.8.ova: a totally generic error occurred"))
				})
			})
		})
	})
})
