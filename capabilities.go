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
	Sockets int `xml:"sockets,attr" json:"sockets,omitempty" yaml:"sockets,omitempty"`
	Cores   int `xml:"cores,attr" json:"cores,omitempty" yaml:"cores,omitempty"`
	Threads int `xml:"threads,attr" json:"threads,omitempty" yaml:"threads,omitempty"`
}

type CapsHostCPUFeatureFlag struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type CapsHostCPUPageSize struct {
	Size int    `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Unit string `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type CapsHostCPUMicrocode struct {
	Version int `xml:"version,attr" json:"version,omitempty" yaml:"version,omitempty"`
}

type CapsHostCPU struct {
	XMLName      xml.Name                 `xml:"cpu" json:"cpu" yaml:"cpu"`
	Arch         string                   `xml:"arch,omitempty" json:"arch,omitempty" yaml:"arch,omitempty"`
	Model        string                   `xml:"model,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Vendor       string                   `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Topology     *CapsHostCPUTopology     `xml:"topology" json:"topology,omitempty" yaml:"topology,omitempty"`
	FeatureFlags []CapsHostCPUFeatureFlag `xml:"feature" json:"feature,omitempty" yaml:"feature,omitempty"`
	Features     *CapsHostCPUFeatures     `xml:"features" json:"features,omitempty" yaml:"features,omitempty"`
	PageSizes    []CapsHostCPUPageSize    `xml:"pages" json:"pages,omitempty" yaml:"pages,omitempty"`
	Microcode    *CapsHostCPUMicrocode    `xml:"microcode" json:"microcode,omitempty" yaml:"microcode,omitempty"`
}

type CapsHostCPUFeature struct {
}

type CapsHostCPUFeatures struct {
	PAE    *CapsHostCPUFeature `xml:"pae" json:"pae,omitempty" yaml:"pae,omitempty"`
	NonPAE *CapsHostCPUFeature `xml:"nonpae" json:"nonpae,omitempty" yaml:"nonpae,omitempty"`
	SVM    *CapsHostCPUFeature `xml:"svm" json:"svm,omitempty" yaml:"svm,omitempty"`
	VMX    *CapsHostCPUFeature `xml:"vmx" json:"vmx,omitempty" yaml:"vmx,omitempty"`
}

type CapsHostNUMAMemory struct {
	Size uint64 `xml:",chardata" json:"" yaml:""`
	Unit string `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type CapsHostNUMAPageInfo struct {
	Size  int    `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Unit  string `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
	Count uint64 `xml:",chardata" json:"" yaml:""`
}

type CapsHostNUMACPU struct {
	ID       int    `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	SocketID *int   `xml:"socket_id,attr" json:"socket_id,omitempty" yaml:"socket_id,omitempty"`
	CoreID   *int   `xml:"core_id,attr" json:"core_id,omitempty" yaml:"core_id,omitempty"`
	Siblings string `xml:"siblings,attr,omitempty" json:"siblings,omitempty" yaml:"siblings,omitempty"`
}

type CapsHostNUMASibling struct {
	ID    int `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Value int `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type CapsHostNUMACell struct {
	ID        int                    `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Memory    *CapsHostNUMAMemory    `xml:"memory" json:"memory,omitempty" yaml:"memory,omitempty"`
	PageInfo  []CapsHostNUMAPageInfo `xml:"pages" json:"pages,omitempty" yaml:"pages,omitempty"`
	Distances *CapsHostNUMADistances `xml:"distances" json:"distances,omitempty" yaml:"distances,omitempty"`
	CPUS      *CapsHostNUMACPUs      `xml:"cpus" json:"cpus,omitempty" yaml:"cpus,omitempty"`
}

type CapsHostNUMADistances struct {
	Siblings []CapsHostNUMASibling `xml:"sibling" json:"sibling,omitempty" yaml:"sibling,omitempty"`
}

type CapsHostNUMACPUs struct {
	Num  uint              `xml:"num,attr,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
	CPUs []CapsHostNUMACPU `xml:"cpu" json:"cpu,omitempty" yaml:"cpu,omitempty"`
}

type CapsHostNUMATopology struct {
	Cells *CapsHostNUMACells `xml:"cells" json:"cells,omitempty" yaml:"cells,omitempty"`
}

type CapsHostNUMACells struct {
	Num   uint               `xml:"num,attr,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
	Cells []CapsHostNUMACell `xml:"cell" json:"cell,omitempty" yaml:"cell,omitempty"`
}

type CapsHostSecModelLabel struct {
	Type  string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Value string `xml:",chardata" json:"" yaml:""`
}

type CapsHostSecModel struct {
	Name   string                  `xml:"model" json:"model" yaml:"model"`
	DOI    string                  `xml:"doi" json:"doi" yaml:"doi"`
	Labels []CapsHostSecModelLabel `xml:"baselabel" json:"baselabel,omitempty" yaml:"baselabel,omitempty"`
}

type CapsHostMigrationFeatures struct {
	Live          *CapsHostMigrationLive          `xml:"live" json:"live,omitempty" yaml:"live,omitempty"`
	URITransports *CapsHostMigrationURITransports `xml:"uri_transports" json:"uri_transports,omitempty" yaml:"uri_transports,omitempty"`
}

type CapsHostMigrationLive struct {
}

type CapsHostMigrationURITransports struct {
	URI []string `xml:"uri_transport" json:"uri_transport,omitempty" yaml:"uri_transport,omitempty"`
}

type CapsHost struct {
	UUID              string                     `xml:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	CPU               *CapsHostCPU               `xml:"cpu" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	PowerManagement   *CapsHostPowerManagement   `xml:"power_management" json:"power_management,omitempty" yaml:"power_management,omitempty"`
	IOMMU             *CapsHostIOMMU             `xml:"iommu" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	MigrationFeatures *CapsHostMigrationFeatures `xml:"migration_features" json:"migration_features,omitempty" yaml:"migration_features,omitempty"`
	NUMA              *CapsHostNUMATopology      `xml:"topology" json:"topology,omitempty" yaml:"topology,omitempty"`
	Cache             *CapsHostCache             `xml:"cache" json:"cache,omitempty" yaml:"cache,omitempty"`
	SecModel          []CapsHostSecModel         `xml:"secmodel" json:"secmodel,omitempty" yaml:"secmodel,omitempty"`
}

type CapsHostPowerManagement struct {
	SuspendMem    *CapsHostPowerManagementMode `xml:"suspend_mem" json:"suspend_mem,omitempty" yaml:"suspend_mem,omitempty"`
	SuspendDisk   *CapsHostPowerManagementMode `xml:"suspend_disk" json:"suspend_disk,omitempty" yaml:"suspend_disk,omitempty"`
	SuspendHybrid *CapsHostPowerManagementMode `xml:"suspend_hybrid" json:"suspend_hybrid,omitempty" yaml:"suspend_hybrid,omitempty"`
}

type CapsHostPowerManagementMode struct {
}

type CapsHostIOMMU struct {
	Support string `xml:"support,attr" json:"support,omitempty" yaml:"support,omitempty"`
}

type CapsHostCache struct {
	Banks []CapsHostCacheBank `xml:"bank" json:"bank,omitempty" yaml:"bank,omitempty"`
}

type CapsHostCacheBank struct {
	ID      uint                   `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Level   uint                   `xml:"level,attr" json:"level,omitempty" yaml:"level,omitempty"`
	Type    string                 `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Size    uint                   `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Unit    string                 `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
	CPUs    string                 `xml:"cpus,attr" json:"cpus,omitempty" yaml:"cpus,omitempty"`
	Control []CapsHostCacheControl `xml:"control" json:"control,omitempty" yaml:"control,omitempty"`
}

type CapsHostCacheControl struct {
	Granularity uint   `xml:"granularity,attr" json:"granularity,omitempty" yaml:"granularity,omitempty"`
	Min         uint   `xml:"min,attr,omitempty" json:"min,omitempty" yaml:"min,omitempty"`
	Unit        string `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
	Type        string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	MaxAllows   uint   `xml:"maxAllocs,attr" json:"maxAllocs,omitempty" yaml:"maxAllocs,omitempty"`
}

type CapsGuestMachine struct {
	Name      string `xml:",chardata" json:"" yaml:""`
	MaxCPUs   int    `xml:"maxCpus,attr,omitempty" json:"maxCpus,omitempty" yaml:"maxCpus,omitempty"`
	Canonical string `xml:"canonical,attr,omitempty" json:"canonical,omitempty" yaml:"canonical,omitempty"`
}

type CapsGuestDomain struct {
	Type     string             `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Emulator string             `xml:"emulator,omitempty" json:"emulator,omitempty" yaml:"emulator,omitempty"`
	Machines []CapsGuestMachine `xml:"machine" json:"machine,omitempty" yaml:"machine,omitempty"`
}

type CapsGuestArch struct {
	Name     string             `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	WordSize string             `xml:"wordsize" json:"wordsize" yaml:"wordsize"`
	Emulator string             `xml:"emulator" json:"emulator" yaml:"emulator"`
	Loader   string             `xml:"loader,omitempty" json:"loader,omitempty" yaml:"loader,omitempty"`
	Machines []CapsGuestMachine `xml:"machine" json:"machine,omitempty" yaml:"machine,omitempty"`
	Domains  []CapsGuestDomain  `xml:"domain" json:"domain,omitempty" yaml:"domain,omitempty"`
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
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" yaml:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" yaml:"toggle,omitempty"`
}

type CapsGuestFeatureAPIC struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" yaml:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" yaml:"toggle,omitempty"`
}

type CapsGuestFeatureACPI struct {
	Default string `xml:"default,attr,omitempty" json:"default,omitempty" yaml:"default,omitempty"`
	Toggle  string `xml:"toggle,attr,omitempty" json:"toggle,omitempty" yaml:"toggle,omitempty"`
}

type CapsGuestFeatureIA64BE struct {
}

type CapsGuestFeatures struct {
	CPUSelection *CapsGuestFeatureCPUSelection `xml:"cpuselection" json:"cpuselection,omitempty" yaml:"cpuselection,omitempty"`
	DeviceBoot   *CapsGuestFeatureDeviceBoot   `xml:"deviceboot" json:"deviceboot,omitempty" yaml:"deviceboot,omitempty"`
	DiskSnapshot *CapsGuestFeatureDiskSnapshot `xml:"disksnapshot" json:"disksnapshot,omitempty" yaml:"disksnapshot,omitempty"`
	PAE          *CapsGuestFeaturePAE          `xml:"pae" json:"pae,omitempty" yaml:"pae,omitempty"`
	NonPAE       *CapsGuestFeatureNonPAE       `xml:"nonpae" json:"nonpae,omitempty" yaml:"nonpae,omitempty"`
	APIC         *CapsGuestFeatureAPIC         `xml:"apic" json:"apic,omitempty" yaml:"apic,omitempty"`
	ACPI         *CapsGuestFeatureACPI         `xml:"acpi" json:"acpi,omitempty" yaml:"acpi,omitempty"`
	IA64BE       *CapsGuestFeatureIA64BE       `xml:"ia64_be" json:"ia64_be" yaml:"ia64_be"`
}

type CapsGuest struct {
	OSType   string             `xml:"os_type" json:"os_type" yaml:"os_type"`
	Arch     CapsGuestArch      `xml:"arch" json:"arch" yaml:"arch"`
	Features *CapsGuestFeatures `xml:"features" json:"features,omitempty" yaml:"features,omitempty"`
}

type Caps struct {
	XMLName xml.Name    `xml:"capabilities" json:"capabilities" yaml:"capabilities"`
	Host    CapsHost    `xml:"host" json:"host" yaml:"host"`
	Guests  []CapsGuest `xml:"guest" json:"guest,omitempty" yaml:"guest,omitempty"`
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
