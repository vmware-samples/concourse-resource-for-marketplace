// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: BSD-2-Clause

package pkg

import (
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
)

type Version struct {
	VersionNumber string `json:"versionnumber"`
}
type Versions []*Version

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	return v[i].LessThan(*v[j])
}

func (a Version) LessThan(b Version) bool {
	semverA, errA := semver.NewVersion(a.VersionNumber)
	semverB, errB := semver.NewVersion(b.VersionNumber)

	if errA != nil || errB != nil {
		return strings.Compare(a.VersionNumber, b.VersionNumber) < 0
	}

	return semverA.LessThan(*semverB)
}

func Sort(versions []*Version) {
	sort.Sort(Versions(versions))
}

func GetOnlySince(versions []*Version, since *Version) []*Version {
	if since == nil || since.VersionNumber == "" {
		return versions
	}

	var filtered []*Version
	for _, version := range versions {
		if version.VersionNumber == since.VersionNumber {
			filtered = nil
		}
		filtered = append(filtered, version)
	}
	return filtered
}
