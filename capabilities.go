/*
 * This file is part of the libvirt-go-xml project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirtxml

import (
	"encoding/xml"
)

type CapsHostCPUTopology struct {
	Sockets int `xml:"sockets,attr",json:"sockets,attr",yaml:"sockets,attr"`
	Cores   int `xml:"cores,attr",json:"cores,attr",yaml:"cores,attr"`
	Threads int `xml:"threads,attr",json:"threads,attr",yaml:"threads,attr"`
}

type CapsHostCPUFeatureFlag struct {
	Name string `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
}

type CapsHostCPUPageSize struct {
	Size int    `xml:"size,attr",json:"size,attr",yaml:"size,attr"`
	Unit string `xml:"unit,attr",json:"unit,attr",yaml:"unit,attr"`
}

type CapsHostCPUMicrocode struct {
	Version int `xml:"version,attr",json:"version,attr",yaml:"version,attr"`
}

type CapsHostCPU struct {
	XMLName      xml.Name                 `xml:"cpu",json:"cpu",yaml:"cpu"`
	Arch         string                   `xml:"arch,omitempty",json:"arch,omitempty",yaml:"arch,omitempty"`
	Model        string                   `xml:"model,omitempty",json:"model,omitempty",yaml:"model,omitempty"`
	Vendor       string                   `xml:"vendor,omitempty",json:"vendor,omitempty",yaml:"vendor,omitempty"`
	Topology     *CapsHostCPUTopology     `xml:"topology",json:"topology",yaml:"topology"`
	FeatureFlags []CapsHostCPUFeatureFlag `xml:"feature",json:"feature",yaml:"feature"`
	Features     *CapsHostCPUFeatures     `xml:"features",json:"features",yaml:"features"`
	PageSizes    []CapsHostCPUPageSize    `xml:"pages",json:"pages",yaml:"pages"`
	Microcode    *CapsHostCPUMicrocode    `xml:"microcode",json:"microcode",yaml:"microcode"`
}

type CapsHostCPUFeature struct {
}

type CapsHostCPUFeatures struct {
	PAE    *CapsHostCPUFeature `xml:"pae",json:"pae",yaml:"pae"`
	NonPAE *CapsHostCPUFeature `xml:"nonpae",json:"nonpae",yaml:"nonpae"`
	SVM    *CapsHostCPUFeature `xml:"svm",json:"svm",yaml:"svm"`
	VMX    *CapsHostCPUFeature `xml:"vmx",json:"vmx",yaml:"vmx"`
}

type CapsHostNUMAMemory struct {
	Size uint64 `xml:",chardata",json:",chardata",yaml:",chardata"`
	Unit string `xml:"unit,attr",json:"unit,attr",yaml:"unit,attr"`
}

type CapsHostNUMAPageInfo struct {
	Size  int    `xml:"size,attr",json:"size,attr",yaml:"size,attr"`
	Unit  string `xml:"unit,attr",json:"unit,attr",yaml:"unit,attr"`
	Count uint64 `xml:",chardata",json:",chardata",yaml:",chardata"`
}

type CapsHostNUMACPU struct {
	ID       int    `xml:"id,attr",json:"id,attr",yaml:"id,attr"`
	SocketID *int   `xml:"socket_id,attr",json:"socket_id,attr",yaml:"socket_id,attr"`
	CoreID   *int   `xml:"core_id,attr",json:"core_id,attr",yaml:"core_id,attr"`
	Siblings string `xml:"siblings,attr,omitempty",json:"siblings,attr,omitempty",yaml:"siblings,attr,omitempty"`
}

type CapsHostNUMASibling struct {
	ID    int `xml:"id,attr",json:"id,attr",yaml:"id,attr"`
	Value int `xml:"value,attr",json:"value,attr",yaml:"value,attr"`
}

type CapsHostNUMACell struct {
	ID        int                    `xml:"id,attr",json:"id,attr",yaml:"id,attr"`
	Memory    *CapsHostNUMAMemory    `xml:"memory",json:"memory",yaml:"memory"`
	PageInfo  []CapsHostNUMAPageInfo `xml:"pages",json:"pages",yaml:"pages"`
	Distances *CapsHostNUMADistances `xml:"distances",json:"distances",yaml:"distances"`
	CPUS      *CapsHostNUMACPUs      `xml:"cpus",json:"cpus",yaml:"cpus"`
}

type CapsHostNUMADistances struct {
	Siblings []CapsHostNUMASibling `xml:"sibling",json:"sibling",yaml:"sibling"`
}

type CapsHostNUMACPUs struct {
	Num  uint              `xml:"num,attr,omitempty",json:"num,attr,omitempty",yaml:"num,attr,omitempty"`
	CPUs []CapsHostNUMACPU `xml:"cpu",json:"cpu",yaml:"cpu"`
}

type CapsHostNUMATopology struct {
	Cells *CapsHostNUMACells `xml:"cells",json:"cells",yaml:"cells"`
}

type CapsHostNUMACells struct {
	Num   uint               `xml:"num,attr,omitempty",json:"num,attr,omitempty",yaml:"num,attr,omitempty"`
	Cells []CapsHostNUMACell `xml:"cell",json:"cell",yaml:"cell"`
}

type CapsHostSecModelLabel struct {
	Type  string `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Value string `xml:",chardata",json:",chardata",yaml:",chardata"`
}

type CapsHostSecModel struct {
	Name   string                  `xml:"model",json:"model",yaml:"model"`
	DOI    string                  `xml:"doi",json:"doi",yaml:"doi"`
	Labels []CapsHostSecModelLabel `xml:"baselabel",json:"baselabel",yaml:"baselabel"`
}

type CapsHostMigrationFeatures struct {
	Live          *CapsHostMigrationLive          `xml:"live",json:"live",yaml:"live"`
	URITransports *CapsHostMigrationURITransports `xml:"uri_transports",json:"uri_transports",yaml:"uri_transports"`
}

type CapsHostMigrationLive struct {
}

type CapsHostMigrationURITransports struct {
	URI []string `xml:"uri_transport",json:"uri_transport",yaml:"uri_transport"`
}

type CapsHost struct {
	UUID              string                     `xml:"uuid,omitempty",json:"uuid,omitempty",yaml:"uuid,omitempty"`
	CPU               *CapsHostCPU               `xml:"cpu",json:"cpu",yaml:"cpu"`
	PowerManagement   *CapsHostPowerManagement   `xml:"power_management",json:"power_management",yaml:"power_management"`
	IOMMU             *CapsHostIOMMU             `xml:"iommu",json:"iommu",yaml:"iommu"`
	MigrationFeatures *CapsHostMigrationFeatures `xml:"migration_features",json:"migration_features",yaml:"migration_features"`
	NUMA              *CapsHostNUMATopology      `xml:"topology",json:"topology",yaml:"topology"`
	Cache             *CapsHostCache             `xml:"cache",json:"cache",yaml:"cache"`
	SecModel          []CapsHostSecModel         `xml:"secmodel",json:"secmodel",yaml:"secmodel"`
}

type CapsHostPowerManagement struct {
	SuspendMem    *CapsHostPowerManagementMode `xml:"suspend_mem",json:"suspend_mem",yaml:"suspend_mem"`
	SuspendDisk   *CapsHostPowerManagementMode `xml:"suspend_disk",json:"suspend_disk",yaml:"suspend_disk"`
	SuspendHybrid *CapsHostPowerManagementMode `xml:"suspend_hybrid",json:"suspend_hybrid",yaml:"suspend_hybrid"`
}

type CapsHostPowerManagementMode struct {
}

type CapsHostIOMMU struct {
	Support string `xml:"support,attr",json:"support,attr",yaml:"support,attr"`
}

type CapsHostCache struct {
	Banks []CapsHostCacheBank `xml:"bank",json:"bank",yaml:"bank"`
}

type CapsHostCacheBank struct {
	ID      uint                   `xml:"id,attr",json:"id,attr",yaml:"id,attr"`
	Level   uint                   `xml:"level,attr",json:"level,attr",yaml:"level,attr"`
	Type    string                 `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Size    uint                   `xml:"size,attr",json:"size,attr",yaml:"size,attr"`
	Unit    string                 `xml:"unit,attr",json:"unit,attr",yaml:"unit,attr"`
	CPUs    string                 `xml:"cpus,attr",json:"cpus,attr",yaml:"cpus,attr"`
	Control []CapsHostCacheControl `xml:"control",json:"control",yaml:"control"`
}

type CapsHostCacheControl struct {
	Granularity uint   `xml:"granularity,attr",json:"granularity,attr",yaml:"granularity,attr"`
	Min         uint   `xml:"min,attr,omitempty",json:"min,attr,omitempty",yaml:"min,attr,omitempty"`
	Unit        string `xml:"unit,attr",json:"unit,attr",yaml:"unit,attr"`
	Type        string `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	MaxAllows   uint   `xml:"maxAllocs,attr",json:"maxAllocs,attr",yaml:"maxAllocs,attr"`
}

type CapsGuestMachine struct {
	Name      string `xml:",chardata",json:",chardata",yaml:",chardata"`
	MaxCPUs   int    `xml:"maxCpus,attr,omitempty",json:"maxCpus,attr,omitempty",yaml:"maxCpus,attr,omitempty"`
	Canonical string `xml:"canonical,attr,omitempty",json:"canonical,attr,omitempty",yaml:"canonical,attr,omitempty"`
}

type CapsGuestDomain struct {
	Type     string             `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Emulator string             `xml:"emulator,omitempty",json:"emulator,omitempty",yaml:"emulator,omitempty"`
	Machines []CapsGuestMachine `xml:"machine",json:"machine",yaml:"machine"`
}

type CapsGuestArch struct {
	Name     string             `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	WordSize string             `xml:"wordsize",json:"wordsize",yaml:"wordsize"`
	Emulator string             `xml:"emulator",json:"emulator",yaml:"emulator"`
	Loader   string             `xml:"loader,omitempty",json:"loader,omitempty",yaml:"loader,omitempty"`
	Machines []CapsGuestMachine `xml:"machine",json:"machine",yaml:"machine"`
	Domains  []CapsGuestDomain  `xml:"domain",json:"domain",yaml:"domain"`
}

type CapsGuestFeatureCPUSelection struct {
}

type CapsGuestFeatureDeviceBoot struct {
}

type CapsGuestFeaturePAE struct {
}

type CapsGuestFeatureNonPAE struct {
}

type CapsGuestFeatureDiskSnapshot struct {
	Default string `xml:"default,attr,omitempty",json:"default,attr,omitempty",yaml:"default,attr,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty",json:"toggle,attr,omitempty",yaml:"toggle,attr,omitempty"`
}

type CapsGuestFeatureAPIC struct {
	Default string `xml:"default,attr,omitempty",json:"default,attr,omitempty",yaml:"default,attr,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty",json:"toggle,attr,omitempty",yaml:"toggle,attr,omitempty"`
}

type CapsGuestFeatureACPI struct {
	Default string `xml:"default,attr,omitempty",json:"default,attr,omitempty",yaml:"default,attr,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty",json:"toggle,attr,omitempty",yaml:"toggle,attr,omitempty"`
}

type CapsGuestFeatureIA64BE struct {
}

type CapsGuestFeatures struct {
	CPUSelection *CapsGuestFeatureCPUSelection `xml:"cpuselection",json:"cpuselection",yaml:"cpuselection"`
	DeviceBoot   *CapsGuestFeatureDeviceBoot   `xml:"deviceboot",json:"deviceboot",yaml:"deviceboot"`
	DiskSnapshot *CapsGuestFeatureDiskSnapshot `xml:"disksnapshot",json:"disksnapshot",yaml:"disksnapshot"`
	PAE          *CapsGuestFeaturePAE          `xml:"pae",json:"pae",yaml:"pae"`
	NonPAE       *CapsGuestFeatureNonPAE       `xml:"nonpae",json:"nonpae",yaml:"nonpae"`
	APIC         *CapsGuestFeatureAPIC         `xml:"apic",json:"apic",yaml:"apic"`
	ACPI         *CapsGuestFeatureACPI         `xml:"acpi",json:"acpi",yaml:"acpi"`
	IA64BE       *CapsGuestFeatureIA64BE       `xml:"ia64_be",json:"ia64_be",yaml:"ia64_be"`
}

type CapsGuest struct {
	OSType   string             `xml:"os_type",json:"os_type",yaml:"os_type"`
	Arch     CapsGuestArch      `xml:"arch",json:"arch",yaml:"arch"`
	Features *CapsGuestFeatures `xml:"features",json:"features",yaml:"features"`
}

type Caps struct {
	XMLName xml.Name    `xml:"capabilities",json:"capabilities",yaml:"capabilities"`
	Host    CapsHost    `xml:"host",json:"host",yaml:"host"`
	Guests  []CapsGuest `xml:"guest",json:"guest",yaml:"guest"`
}

func (c *CapsHostCPU) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *CapsHostCPU) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (c *Caps) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *Caps) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
