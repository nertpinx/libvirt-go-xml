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
	"fmt"
	"io"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type DomainControllerPCIHole64 struct {
	Size uint64 `xml:",attr" json:"" yaml:""`
	Unit string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainControllerPCIModel struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainControllerPCITarget struct {
	ChassisNr *uint
	Chassis   *uint
	Port      *uint
	BusNr     *uint
	Index     *uint
	NUMANode  *uint
}

type DomainControllerPCI struct {
	Model  *DomainControllerPCIModel  `xml:"model" json:"model" yaml:"model"`
	Target *DomainControllerPCITarget `xml:"target" json:"target" yaml:"target"`
	Hole64 *DomainControllerPCIHole64 `xml:"pcihole64" json:"pcihole64" yaml:"pcihole64"`
}

type DomainControllerUSBMaster struct {
	StartPort uint `xml:"startport,attr" json:"startport,omitempty" yaml:"startport,omitempty"`
}

type DomainControllerUSB struct {
	Port   *uint                      `xml:"ports,attr" json:"ports,omitempty" yaml:"ports,omitempty"`
	Master *DomainControllerUSBMaster `xml:"master" json:"master" yaml:"master"`
}

type DomainControllerVirtIOSerial struct {
	Ports   *uint `xml:"ports,attr" json:"ports,omitempty" yaml:"ports,omitempty"`
	Vectors *uint `xml:"vectors,attr" json:"vectors,omitempty" yaml:"vectors,omitempty"`
}

type DomainControllerDriver struct {
	Queues     *uint  `xml:"queues,attr" json:"queues,omitempty" yaml:"queues,omitempty"`
	CmdPerLUN  *uint  `xml:"cmd_per_lun,attr" json:"cmd_per_lun,omitempty" yaml:"cmd_per_lun,omitempty"`
	MaxSectors *uint  `xml:"max_sectors,attr" json:"max_sectors,omitempty" yaml:"max_sectors,omitempty"`
	IOEventFD  string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" yaml:"ioeventfd,omitempty"`
	IOThread   uint   `xml:"iothread,attr,omitempty" json:"iothread,omitempty" yaml:"iothread,omitempty"`
	IOMMU      string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS        string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainController struct {
	XMLName      xml.Name                      `xml:"controller" json:"controller" yaml:"controller"`
	Type         string                        `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Index        *uint                         `xml:"index,attr" json:"index,omitempty" yaml:"index,omitempty"`
	Model        string                        `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Driver       *DomainControllerDriver       `xml:"driver" json:"driver" yaml:"driver"`
	PCI          *DomainControllerPCI          `xml:"-" json:"-" yaml:"-"`
	USB          *DomainControllerUSB          `xml:"-" json:"-" yaml:"-"`
	VirtIOSerial *DomainControllerVirtIOSerial `xml:"-" json:"-" yaml:"-"`
	Alias        *DomainAlias                  `xml:"alias" json:"alias" yaml:"alias"`
	Address      *DomainAddress                `xml:"address" json:"address" yaml:"address"`
}

type DomainDiskSecret struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty" yaml:"usage,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DomainDiskAuth struct {
	Username string            `xml:"username,attr,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Secret   *DomainDiskSecret `xml:"secret" json:"secret" yaml:"secret"`
}

type DomainDiskSourceHost struct {
	Transport string `xml:"transport,attr,omitempty" json:"transport,omitempty" yaml:"transport,omitempty"`
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Port      string `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Socket    string `xml:"socket,attr,omitempty" json:"socket,omitempty" yaml:"socket,omitempty"`
}

type DomainDiskReservationsSource DomainChardevSource

type DomainDiskReservations struct {
	Enabled string                        `xml:"enabled,attr,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Managed string                        `xml:"managed,attr,omitempty" json:"managed,omitempty" yaml:"managed,omitempty"`
	Source  *DomainDiskReservationsSource `xml:"source" json:"source" yaml:"source"`
}

type DomainDiskSource struct {
	File          *DomainDiskSourceFile    `xml:"-" json:"-" yaml:"-"`
	Block         *DomainDiskSourceBlock   `xml:"-" json:"-" yaml:"-"`
	Dir           *DomainDiskSourceDir     `xml:"-" json:"-" yaml:"-"`
	Network       *DomainDiskSourceNetwork `xml:"-" json:"-" yaml:"-"`
	Volume        *DomainDiskSourceVolume  `xml:"-" json:"-" yaml:"-"`
	StartupPolicy string                   `xml:"startupPolicy,attr,omitempty" json:"startupPolicy,omitempty" yaml:"startupPolicy,omitempty"`
	Encryption    *DomainDiskEncryption    `xml:"encryption" json:"encryption" yaml:"encryption"`
	Reservations  *DomainDiskReservations  `xml:"reservations" json:"reservations" yaml:"reservations"`
}

type DomainDiskSourceFile struct {
	File     string                 `xml:"file,attr,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainDiskSourceBlock struct {
	Dev      string                 `xml:"dev,attr,omitempty" json:"dev,omitempty" yaml:"dev,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainDiskSourceDir struct {
	Dir string `xml:"dir,attr,omitempty" json:"dir,omitempty" yaml:"dir,omitempty"`
}

type DomainDiskSourceNetwork struct {
	Protocol  string                            `xml:"protocol,attr,omitempty" json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Name      string                            `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	TLS       string                            `xml:"tls,attr,omitempty" json:"tls,omitempty" yaml:"tls,omitempty"`
	Hosts     []DomainDiskSourceHost            `xml:"host" json:"host" yaml:"host"`
	Initiator *DomainDiskSourceNetworkInitiator `xml:"initiator" json:"initiator" yaml:"initiator"`
	Snapshot  *DomainDiskSourceNetworkSnapshot  `xml:"snapshot" json:"snapshot" yaml:"snapshot"`
	Config    *DomainDiskSourceNetworkConfig    `xml:"config" json:"config" yaml:"config"`
	Auth      *DomainDiskAuth                   `xml:"auth" json:"auth" yaml:"auth"`
}

type DomainDiskSourceNetworkInitiator struct {
	IQN *DomainDiskSourceNetworkIQN `xml:"iqn" json:"iqn" yaml:"iqn"`
}

type DomainDiskSourceNetworkIQN struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainDiskSourceNetworkSnapshot struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainDiskSourceNetworkConfig struct {
	File string `xml:"file,attr" json:"file,omitempty" yaml:"file,omitempty"`
}

type DomainDiskSourceVolume struct {
	Pool     string                 `xml:"pool,attr,omitempty" json:"pool,omitempty" yaml:"pool,omitempty"`
	Volume   string                 `xml:"volume,attr,omitempty" json:"volume,omitempty" yaml:"volume,omitempty"`
	Mode     string                 `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainDiskDriver struct {
	Name         string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type         string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Cache        string `xml:"cache,attr,omitempty" json:"cache,omitempty" yaml:"cache,omitempty"`
	ErrorPolicy  string `xml:"error_policy,attr,omitempty" json:"error_policy,omitempty" yaml:"error_policy,omitempty"`
	RErrorPolicy string `xml:"rerror_policy,attr,omitempty" json:"rerror_policy,omitempty" yaml:"rerror_policy,omitempty"`
	IO           string `xml:"io,attr,omitempty" json:"io,omitempty" yaml:"io,omitempty"`
	IOEventFD    string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" yaml:"ioeventfd,omitempty"`
	EventIDX     string `xml:"event_idx,attr,omitempty" json:"event_idx,omitempty" yaml:"event_idx,omitempty"`
	CopyOnRead   string `xml:"copy_on_read,attr,omitempty" json:"copy_on_read,omitempty" yaml:"copy_on_read,omitempty"`
	Discard      string `xml:"discard,attr,omitempty" json:"discard,omitempty" yaml:"discard,omitempty"`
	IOThread     *uint  `xml:"iothread,attr" json:"iothread,omitempty" yaml:"iothread,omitempty"`
	DetectZeros  string `xml:"detect_zeroes,attr,omitempty" json:"detect_zeroes,omitempty" yaml:"detect_zeroes,omitempty"`
	Queues       *uint  `xml:"queues,attr" json:"queues,omitempty" yaml:"queues,omitempty"`
	IOMMU        string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS          string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainDiskTarget struct {
	Dev       string `xml:"dev,attr,omitempty" json:"dev,omitempty" yaml:"dev,omitempty"`
	Bus       string `xml:"bus,attr,omitempty" json:"bus,omitempty" yaml:"bus,omitempty"`
	Tray      string `xml:"tray,attr,omitempty" json:"tray,omitempty" yaml:"tray,omitempty"`
	Removable string `xml:"removable,attr,omitempty" json:"removable,omitempty" yaml:"removable,omitempty"`
}

type DomainDiskEncryption struct {
	Format string            `xml:"format,attr,omitempty" json:"format,omitempty" yaml:"format,omitempty"`
	Secret *DomainDiskSecret `xml:"secret" json:"secret" yaml:"secret"`
}

type DomainDiskReadOnly struct {
}

type DomainDiskShareable struct {
}

type DomainDiskTransient struct {
}

type DomainDiskIOTune struct {
	TotalBytesSec          uint64 `xml:"total_bytes_sec,omitempty" json:"total_bytes_sec,omitempty" yaml:"total_bytes_sec,omitempty"`
	ReadBytesSec           uint64 `xml:"read_bytes_sec,omitempty" json:"read_bytes_sec,omitempty" yaml:"read_bytes_sec,omitempty"`
	WriteBytesSec          uint64 `xml:"write_bytes_sec,omitempty" json:"write_bytes_sec,omitempty" yaml:"write_bytes_sec,omitempty"`
	TotalIopsSec           uint64 `xml:"total_iops_sec,omitempty" json:"total_iops_sec,omitempty" yaml:"total_iops_sec,omitempty"`
	ReadIopsSec            uint64 `xml:"read_iops_sec,omitempty" json:"read_iops_sec,omitempty" yaml:"read_iops_sec,omitempty"`
	WriteIopsSec           uint64 `xml:"write_iops_sec,omitempty" json:"write_iops_sec,omitempty" yaml:"write_iops_sec,omitempty"`
	TotalBytesSecMax       uint64 `xml:"total_bytes_sec_max,omitempty" json:"total_bytes_sec_max,omitempty" yaml:"total_bytes_sec_max,omitempty"`
	ReadBytesSecMax        uint64 `xml:"read_bytes_sec_max,omitempty" json:"read_bytes_sec_max,omitempty" yaml:"read_bytes_sec_max,omitempty"`
	WriteBytesSecMax       uint64 `xml:"write_bytes_sec_max,omitempty" json:"write_bytes_sec_max,omitempty" yaml:"write_bytes_sec_max,omitempty"`
	TotalIopsSecMax        uint64 `xml:"total_iops_sec_max,omitempty" json:"total_iops_sec_max,omitempty" yaml:"total_iops_sec_max,omitempty"`
	ReadIopsSecMax         uint64 `xml:"read_iops_sec_max,omitempty" json:"read_iops_sec_max,omitempty" yaml:"read_iops_sec_max,omitempty"`
	WriteIopsSecMax        uint64 `xml:"write_iops_sec_max,omitempty" json:"write_iops_sec_max,omitempty" yaml:"write_iops_sec_max,omitempty"`
	TotalBytesSecMaxLength uint64 `xml:"total_bytes_sec_max_length,omitempty" json:"total_bytes_sec_max_length,omitempty" yaml:"total_bytes_sec_max_length,omitempty"`
	ReadBytesSecMaxLength  uint64 `xml:"read_bytes_sec_max_length,omitempty" json:"read_bytes_sec_max_length,omitempty" yaml:"read_bytes_sec_max_length,omitempty"`
	WriteBytesSecMaxLength uint64 `xml:"write_bytes_sec_max_length,omitempty" json:"write_bytes_sec_max_length,omitempty" yaml:"write_bytes_sec_max_length,omitempty"`
	TotalIopsSecMaxLength  uint64 `xml:"total_iops_sec_max_length,omitempty" json:"total_iops_sec_max_length,omitempty" yaml:"total_iops_sec_max_length,omitempty"`
	ReadIopsSecMaxLength   uint64 `xml:"read_iops_sec_max_length,omitempty" json:"read_iops_sec_max_length,omitempty" yaml:"read_iops_sec_max_length,omitempty"`
	WriteIopsSecMaxLength  uint64 `xml:"write_iops_sec_max_length,omitempty" json:"write_iops_sec_max_length,omitempty" yaml:"write_iops_sec_max_length,omitempty"`
	SizeIopsSec            uint64 `xml:"size_iops_sec,omitempty" json:"size_iops_sec,omitempty" yaml:"size_iops_sec,omitempty"`
	GroupName              string `xml:"group_name,omitempty" json:"group_name,omitempty" yaml:"group_name,omitempty"`
}

type DomainDiskGeometry struct {
	Cylinders uint   `xml:"cyls,attr" json:"cyls,omitempty" yaml:"cyls,omitempty"`
	Headers   uint   `xml:"heads,attr" json:"heads,omitempty" yaml:"heads,omitempty"`
	Sectors   uint   `xml:"secs,attr" json:"secs,omitempty" yaml:"secs,omitempty"`
	Trans     string `xml:"trans,attr,omitempty" json:"trans,omitempty" yaml:"trans,omitempty"`
}

type DomainDiskBlockIO struct {
	LogicalBlockSize  uint `xml:"logical_block_size,attr,omitempty" json:"logical_block_size,omitempty" yaml:"logical_block_size,omitempty"`
	PhysicalBlockSize uint `xml:"physical_block_size,attr,omitempty" json:"physical_block_size,omitempty" yaml:"physical_block_size,omitempty"`
}

type DomainDiskFormat struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainDiskBackingStore struct {
	Index        uint                    `xml:"index,attr,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Format       *DomainDiskFormat       `xml:"format" json:"format" yaml:"format"`
	Source       *DomainDiskSource       `xml:"source" json:"source" yaml:"source"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore" yaml:"backingStore"`
}

type DomainDiskMirror struct {
	Job    string            `xml:"job,attr,omitempty" json:"job,omitempty" yaml:"job,omitempty"`
	Ready  string            `xml:"ready,attr,omitempty" json:"ready,omitempty" yaml:"ready,omitempty"`
	Format *DomainDiskFormat `xml:"format" json:"format" yaml:"format"`
	Source *DomainDiskSource `xml:"source" json:"source" yaml:"source"`
}

type DomainDisk struct {
	XMLName      xml.Name                `xml:"disk" json:"disk" yaml:"disk"`
	Device       string                  `xml:"device,attr,omitempty" json:"device,omitempty" yaml:"device,omitempty"`
	RawIO        string                  `xml:"rawio,attr,omitempty" json:"rawio,omitempty" yaml:"rawio,omitempty"`
	SGIO         string                  `xml:"sgio,attr,omitempty" json:"sgio,omitempty" yaml:"sgio,omitempty"`
	Snapshot     string                  `xml:"snapshot,attr,omitempty" json:"snapshot,omitempty" yaml:"snapshot,omitempty"`
	Driver       *DomainDiskDriver       `xml:"driver" json:"driver" yaml:"driver"`
	Auth         *DomainDiskAuth         `xml:"auth" json:"auth" yaml:"auth"`
	Source       *DomainDiskSource       `xml:"source" json:"source" yaml:"source"`
	BackingStore *DomainDiskBackingStore `xml:"backingStore" json:"backingStore" yaml:"backingStore"`
	Geometry     *DomainDiskGeometry     `xml:"geometry" json:"geometry" yaml:"geometry"`
	BlockIO      *DomainDiskBlockIO      `xml:"blockio" json:"blockio" yaml:"blockio"`
	Mirror       *DomainDiskMirror       `xml:"mirror" json:"mirror" yaml:"mirror"`
	Target       *DomainDiskTarget       `xml:"target" json:"target" yaml:"target"`
	IOTune       *DomainDiskIOTune       `xml:"iotune" json:"iotune" yaml:"iotune"`
	ReadOnly     *DomainDiskReadOnly     `xml:"readonly" json:"readonly" yaml:"readonly"`
	Shareable    *DomainDiskShareable    `xml:"shareable" json:"shareable" yaml:"shareable"`
	Transient    *DomainDiskTransient    `xml:"transient" json:"transient" yaml:"transient"`
	Serial       string                  `xml:"serial,omitempty" json:"serial,omitempty" yaml:"serial,omitempty"`
	WWN          string                  `xml:"wwn,omitempty" json:"wwn,omitempty" yaml:"wwn,omitempty"`
	Vendor       string                  `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Product      string                  `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Encryption   *DomainDiskEncryption   `xml:"encryption" json:"encryption" yaml:"encryption"`
	Boot         *DomainDeviceBoot       `xml:"boot" json:"boot" yaml:"boot"`
	Alias        *DomainAlias            `xml:"alias" json:"alias" yaml:"alias"`
	Address      *DomainAddress          `xml:"address" json:"address" yaml:"address"`
}

type DomainFilesystemDriver struct {
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Format   string `xml:"format,attr,omitempty" json:"format,omitempty" yaml:"format,omitempty"`
	Name     string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	WRPolicy string `xml:"wrpolicy,attr,omitempty" json:"wrpolicy,omitempty" yaml:"wrpolicy,omitempty"`
	IOMMU    string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS      string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainFilesystemSource struct {
	Mount    *DomainFilesystemSourceMount    `xml:"-" json:"-" yaml:"-"`
	Block    *DomainFilesystemSourceBlock    `xml:"-" json:"-" yaml:"-"`
	File     *DomainFilesystemSourceFile     `xml:"-" json:"-" yaml:"-"`
	Template *DomainFilesystemSourceTemplate `xml:"-" json:"-" yaml:"-"`
	RAM      *DomainFilesystemSourceRAM      `xml:"-" json:"-" yaml:"-"`
	Bind     *DomainFilesystemSourceBind     `xml:"-" json:"-" yaml:"-"`
	Volume   *DomainFilesystemSourceVolume   `xml:"-" json:"-" yaml:"-"`
}

type DomainFilesystemSourceMount struct {
	Dir string `xml:"dir,attr" json:"dir,omitempty" yaml:"dir,omitempty"`
}

type DomainFilesystemSourceBlock struct {
	Dev string `xml:"dev,attr" json:"dev,omitempty" yaml:"dev,omitempty"`
}

type DomainFilesystemSourceFile struct {
	File string `xml:"file,attr" json:"file,omitempty" yaml:"file,omitempty"`
}

type DomainFilesystemSourceTemplate struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainFilesystemSourceRAM struct {
	Usage uint   `xml:"usage,attr" json:"usage,omitempty" yaml:"usage,omitempty"`
	Units string `xml:"units,attr,omitempty" json:"units,omitempty" yaml:"units,omitempty"`
}

type DomainFilesystemSourceBind struct {
	Dir string `xml:"dir,attr" json:"dir,omitempty" yaml:"dir,omitempty"`
}

type DomainFilesystemSourceVolume struct {
	Pool   string `xml:"pool,attr" json:"pool,omitempty" yaml:"pool,omitempty"`
	Volume string `xml:"volume,attr" json:"volume,omitempty" yaml:"volume,omitempty"`
}

type DomainFilesystemTarget struct {
	Dir string `xml:"dir,attr" json:"dir,omitempty" yaml:"dir,omitempty"`
}

type DomainFilesystemReadOnly struct {
}

type DomainFilesystemSpaceHardLimit struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainFilesystemSpaceSoftLimit struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainFilesystem struct {
	XMLName        xml.Name                        `xml:"filesystem" json:"filesystem" yaml:"filesystem"`
	AccessMode     string                          `xml:"accessmode,attr,omitempty" json:"accessmode,omitempty" yaml:"accessmode,omitempty"`
	Driver         *DomainFilesystemDriver         `xml:"driver" json:"driver" yaml:"driver"`
	Source         *DomainFilesystemSource         `xml:"source" json:"source" yaml:"source"`
	Target         *DomainFilesystemTarget         `xml:"target" json:"target" yaml:"target"`
	ReadOnly       *DomainFilesystemReadOnly       `xml:"readonly" json:"readonly" yaml:"readonly"`
	SpaceHardLimit *DomainFilesystemSpaceHardLimit `xml:"space_hard_limit" json:"space_hard_limit" yaml:"space_hard_limit"`
	SpaceSoftLimit *DomainFilesystemSpaceSoftLimit `xml:"space_soft_limit" json:"space_soft_limit" yaml:"space_soft_limit"`
	Alias          *DomainAlias                    `xml:"alias" json:"alias" yaml:"alias"`
	Address        *DomainAddress                  `xml:"address" json:"address" yaml:"address"`
}

type DomainInterfaceMAC struct {
	Address string `xml:"address,attr" json:"address,omitempty" yaml:"address,omitempty"`
}

type DomainInterfaceModel struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainInterfaceSource struct {
	User      *DomainInterfaceSourceUser     `xml:"-" json:"-" yaml:"-"`
	Ethernet  *DomainInterfaceSourceEthernet `xml:"-" json:"-" yaml:"-"`
	VHostUser *DomainChardevSource           `xml:"-" json:"-" yaml:"-"`
	Server    *DomainInterfaceSourceServer   `xml:"-" json:"-" yaml:"-"`
	Client    *DomainInterfaceSourceClient   `xml:"-" json:"-" yaml:"-"`
	MCast     *DomainInterfaceSourceMCast    `xml:"-" json:"-" yaml:"-"`
	Network   *DomainInterfaceSourceNetwork  `xml:"-" json:"-" yaml:"-"`
	Bridge    *DomainInterfaceSourceBridge   `xml:"-" json:"-" yaml:"-"`
	Internal  *DomainInterfaceSourceInternal `xml:"-" json:"-" yaml:"-"`
	Direct    *DomainInterfaceSourceDirect   `xml:"-" json:"-" yaml:"-"`
	Hostdev   *DomainInterfaceSourceHostdev  `xml:"-" json:"-" yaml:"-"`
	UDP       *DomainInterfaceSourceUDP      `xml:"-" json:"-" yaml:"-"`
}

type DomainInterfaceSourceUser struct {
}

type DomainInterfaceSourceEthernet struct {
	IP    []DomainInterfaceIP    `xml:"ip" json:"ip" yaml:"ip"`
	Route []DomainInterfaceRoute `xml:"route" json:"route" yaml:"route"`
}

type DomainInterfaceSourceServer struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local" yaml:"local"`
}

type DomainInterfaceSourceClient struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local" yaml:"local"`
}

type DomainInterfaceSourceMCast struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local" yaml:"local"`
}

type DomainInterfaceSourceNetwork struct {
	Network   string `xml:"network,attr,omitempty" json:"network,omitempty" yaml:"network,omitempty"`
	PortGroup string `xml:"portgroup,attr,omitempty" json:"portgroup,omitempty" yaml:"portgroup,omitempty"`
}

type DomainInterfaceSourceBridge struct {
	Bridge string `xml:"bridge,attr" json:"bridge,omitempty" yaml:"bridge,omitempty"`
}

type DomainInterfaceSourceInternal struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainInterfaceSourceDirect struct {
	Dev  string `xml:"dev,attr,omitempty" json:"dev,omitempty" yaml:"dev,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainInterfaceSourceHostdev struct {
	PCI *DomainHostdevSubsysPCISource `xml:"-" json:"-" yaml:"-"`
	USB *DomainHostdevSubsysUSBSource `xml:"-" json:"-" yaml:"-"`
}

type DomainInterfaceSourceUDP struct {
	Address string                      `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local" json:"local" yaml:"local"`
}

type DomainInterfaceSourceLocal struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    uint   `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainInterfaceTarget struct {
	Dev string `xml:"dev,attr" json:"dev,omitempty" yaml:"dev,omitempty"`
}

type DomainInterfaceLink struct {
	State string `xml:"state,attr" json:"state,omitempty" yaml:"state,omitempty"`
}

type DomainDeviceBoot struct {
	Order    uint   `xml:"order,attr" json:"order,omitempty" yaml:"order,omitempty"`
	LoadParm string `xml:"loadparm,attr,omitempty" json:"loadparm,omitempty" yaml:"loadparm,omitempty"`
}

type DomainInterfaceScript struct {
	Path string `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
}

type DomainInterfaceDriver struct {
	Name        string                      `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	TXMode      string                      `xml:"txmode,attr,omitempty" json:"txmode,omitempty" yaml:"txmode,omitempty"`
	IOEventFD   string                      `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" yaml:"ioeventfd,omitempty"`
	EventIDX    string                      `xml:"event_idx,attr,omitempty" json:"event_idx,omitempty" yaml:"event_idx,omitempty"`
	Queues      uint                        `xml:"queues,attr,omitempty" json:"queues,omitempty" yaml:"queues,omitempty"`
	RXQueueSize uint                        `xml:"rx_queue_size,attr,omitempty" json:"rx_queue_size,omitempty" yaml:"rx_queue_size,omitempty"`
	TXQueueSize uint                        `xml:"tx_queue_size,attr,omitempty" json:"tx_queue_size,omitempty" yaml:"tx_queue_size,omitempty"`
	IOMMU       string                      `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS         string                      `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
	Host        *DomainInterfaceDriverHost  `xml:"host" json:"host" yaml:"host"`
	Guest       *DomainInterfaceDriverGuest `xml:"guest" json:"guest" yaml:"guest"`
}

type DomainInterfaceDriverHost struct {
	CSum     string `xml:"csum,attr,omitempty" json:"csum,omitempty" yaml:"csum,omitempty"`
	GSO      string `xml:"gso,attr,omitempty" json:"gso,omitempty" yaml:"gso,omitempty"`
	TSO4     string `xml:"tso4,attr,omitempty" json:"tso4,omitempty" yaml:"tso4,omitempty"`
	TSO6     string `xml:"tso6,attr,omitempty" json:"tso6,omitempty" yaml:"tso6,omitempty"`
	ECN      string `xml:"ecn,attr,omitempty" json:"ecn,omitempty" yaml:"ecn,omitempty"`
	UFO      string `xml:"ufo,attr,omitempty" json:"ufo,omitempty" yaml:"ufo,omitempty"`
	MrgRXBuf string `xml:"mrg_rxbuf,attr,omitempty" json:"mrg_rxbuf,omitempty" yaml:"mrg_rxbuf,omitempty"`
}

type DomainInterfaceDriverGuest struct {
	CSum string `xml:"csum,attr,omitempty" json:"csum,omitempty" yaml:"csum,omitempty"`
	TSO4 string `xml:"tso4,attr,omitempty" json:"tso4,omitempty" yaml:"tso4,omitempty"`
	TSO6 string `xml:"tso6,attr,omitempty" json:"tso6,omitempty" yaml:"tso6,omitempty"`
	ECN  string `xml:"ecn,attr,omitempty" json:"ecn,omitempty" yaml:"ecn,omitempty"`
	UFO  string `xml:"ufo,attr,omitempty" json:"ufo,omitempty" yaml:"ufo,omitempty"`
}

type DomainInterfaceVirtualPort struct {
	Params *DomainInterfaceVirtualPortParams `xml:"parameters" json:"parameters" yaml:"parameters"`
}

type DomainInterfaceVirtualPortParams struct {
	Any          *DomainInterfaceVirtualPortParamsAny          `xml:"-" json:"-" yaml:"-"`
	VEPA8021QBG  *DomainInterfaceVirtualPortParamsVEPA8021QBG  `xml:"-" json:"-" yaml:"-"`
	VNTag8011QBH *DomainInterfaceVirtualPortParamsVNTag8021QBH `xml:"-" json:"-" yaml:"-"`
	OpenVSwitch  *DomainInterfaceVirtualPortParamsOpenVSwitch  `xml:"-" json:"-" yaml:"-"`
	MidoNet      *DomainInterfaceVirtualPortParamsMidoNet      `xml:"-" json:"-" yaml:"-"`
}

type DomainInterfaceVirtualPortParamsAny struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerid,omitempty" yaml:"managerid,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeid,omitempty" yaml:"typeid,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeidversion,omitempty" yaml:"typeidversion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty" yaml:"instanceid,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileid,omitempty" yaml:"profileid,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty" yaml:"interfaceid,omitempty"`
}

type DomainInterfaceVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerid,omitempty" yaml:"managerid,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeid,omitempty" yaml:"typeid,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeidversion,omitempty" yaml:"typeidversion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty" yaml:"instanceid,omitempty"`
}

type DomainInterfaceVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileid,omitempty" yaml:"profileid,omitempty"`
}

type DomainInterfaceVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty" yaml:"interfaceid,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileid,omitempty" yaml:"profileid,omitempty"`
}

type DomainInterfaceVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty" yaml:"interfaceid,omitempty"`
}

type DomainInterfaceBandwidthParams struct {
	Average *int `xml:"average,attr" json:"average,omitempty" yaml:"average,omitempty"`
	Peak    *int `xml:"peak,attr" json:"peak,omitempty" yaml:"peak,omitempty"`
	Burst   *int `xml:"burst,attr" json:"burst,omitempty" yaml:"burst,omitempty"`
	Floor   *int `xml:"floor,attr" json:"floor,omitempty" yaml:"floor,omitempty"`
}

type DomainInterfaceBandwidth struct {
	Inbound  *DomainInterfaceBandwidthParams `xml:"inbound" json:"inbound" yaml:"inbound"`
	Outbound *DomainInterfaceBandwidthParams `xml:"outbound" json:"outbound" yaml:"outbound"`
}

type DomainInterfaceVLan struct {
	Trunk string                   `xml:"trunk,attr,omitempty" json:"trunk,omitempty" yaml:"trunk,omitempty"`
	Tags  []DomainInterfaceVLanTag `xml:"tag" json:"tag" yaml:"tag"`
}

type DomainInterfaceVLanTag struct {
	ID         uint   `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty" yaml:"nativeMode,omitempty"`
}

type DomainInterfaceGuest struct {
	Dev    string `xml:"dev,attr,omitempty" json:"dev,omitempty" yaml:"dev,omitempty"`
	Actual string `xml:"actual,attr,omitempty" json:"actual,omitempty" yaml:"actual,omitempty"`
}

type DomainInterfaceFilterRef struct {
	Filter     string                       `xml:"filter,attr" json:"filter,omitempty" yaml:"filter,omitempty"`
	Parameters []DomainInterfaceFilterParam `xml:"parameter" json:"parameter" yaml:"parameter"`
}

type DomainInterfaceFilterParam struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value string `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainInterfaceBackend struct {
	Tap   string `xml:"tap,attr,omitempty" json:"tap,omitempty" yaml:"tap,omitempty"`
	VHost string `xml:"vhost,attr,omitempty" json:"vhost,omitempty" yaml:"vhost,omitempty"`
}

type DomainInterfaceTune struct {
	SndBuf uint `xml:"sndbuf" json:"sndbuf" yaml:"sndbuf"`
}

type DomainInterfaceMTU struct {
	Size uint `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
}

type DomainInterfaceCoalesce struct {
	RX *DomainInterfaceCoalesceRX `xml:"rx" json:"rx" yaml:"rx"`
}

type DomainInterfaceCoalesceRX struct {
	Frames *DomainInterfaceCoalesceRXFrames `xml:"frames" json:"frames" yaml:"frames"`
}

type DomainInterfaceCoalesceRXFrames struct {
	Max *uint `xml:"max,attr" json:"max,omitempty" yaml:"max,omitempty"`
}

type DomainROM struct {
	Bar     string `xml:"bar,attr,omitempty" json:"bar,omitempty" yaml:"bar,omitempty"`
	File    string `xml:"file,attr,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Enabled string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type DomainInterfaceIP struct {
	Address string `xml:"address,attr" json:"address,omitempty" yaml:"address,omitempty"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" yaml:"family,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Peer    string `xml:"peer,attr,omitempty" json:"peer,omitempty" yaml:"peer,omitempty"`
}

type DomainInterfaceRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" yaml:"family,omitempty"`
	Address string `xml:"address,attr" json:"address,omitempty" yaml:"address,omitempty"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty" yaml:"netmask,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Gateway string `xml:"gateway,attr" json:"gateway,omitempty" yaml:"gateway,omitempty"`
	Metric  uint   `xml:"metric,attr,omitempty" json:"metric,omitempty" yaml:"metric,omitempty"`
}

type DomainInterface struct {
	XMLName             xml.Name                    `xml:"interface" json:"interface" yaml:"interface"`
	Managed             string                      `xml:"managed,attr,omitempty" json:"managed,omitempty" yaml:"managed,omitempty"`
	TrustGuestRXFilters string                      `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty" yaml:"trustGuestRxFilters,omitempty"`
	MAC                 *DomainInterfaceMAC         `xml:"mac" json:"mac" yaml:"mac"`
	Source              *DomainInterfaceSource      `xml:"source" json:"source" yaml:"source"`
	Boot                *DomainDeviceBoot           `xml:"boot" json:"boot" yaml:"boot"`
	VLan                *DomainInterfaceVLan        `xml:"vlan" json:"vlan" yaml:"vlan"`
	VirtualPort         *DomainInterfaceVirtualPort `xml:"virtualport" json:"virtualport" yaml:"virtualport"`
	IP                  []DomainInterfaceIP         `xml:"ip" json:"ip" yaml:"ip"`
	Route               []DomainInterfaceRoute      `xml:"route" json:"route" yaml:"route"`
	Script              *DomainInterfaceScript      `xml:"script" json:"script" yaml:"script"`
	Target              *DomainInterfaceTarget      `xml:"target" json:"target" yaml:"target"`
	Guest               *DomainInterfaceGuest       `xml:"guest" json:"guest" yaml:"guest"`
	Model               *DomainInterfaceModel       `xml:"model" json:"model" yaml:"model"`
	Driver              *DomainInterfaceDriver      `xml:"driver" json:"driver" yaml:"driver"`
	Backend             *DomainInterfaceBackend     `xml:"backend" json:"backend" yaml:"backend"`
	FilterRef           *DomainInterfaceFilterRef   `xml:"filterref" json:"filterref" yaml:"filterref"`
	Tune                *DomainInterfaceTune        `xml:"tune" json:"tune" yaml:"tune"`
	Link                *DomainInterfaceLink        `xml:"link" json:"link" yaml:"link"`
	MTU                 *DomainInterfaceMTU         `xml:"mtu" json:"mtu" yaml:"mtu"`
	Bandwidth           *DomainInterfaceBandwidth   `xml:"bandwidth" json:"bandwidth" yaml:"bandwidth"`
	Coalesce            *DomainInterfaceCoalesce    `xml:"coalesce" json:"coalesce" yaml:"coalesce"`
	ROM                 *DomainROM                  `xml:"rom" json:"rom" yaml:"rom"`
	Alias               *DomainAlias                `xml:"alias" json:"alias" yaml:"alias"`
	Address             *DomainAddress              `xml:"address" json:"address" yaml:"address"`
}

type DomainChardevSource struct {
	Null      *DomainChardevSourceNull      `xml:"-" json:"-" yaml:"-"`
	VC        *DomainChardevSourceVC        `xml:"-" json:"-" yaml:"-"`
	Pty       *DomainChardevSourcePty       `xml:"-" json:"-" yaml:"-"`
	Dev       *DomainChardevSourceDev       `xml:"-" json:"-" yaml:"-"`
	File      *DomainChardevSourceFile      `xml:"-" json:"-" yaml:"-"`
	Pipe      *DomainChardevSourcePipe      `xml:"-" json:"-" yaml:"-"`
	StdIO     *DomainChardevSourceStdIO     `xml:"-" json:"-" yaml:"-"`
	UDP       *DomainChardevSourceUDP       `xml:"-" json:"-" yaml:"-"`
	TCP       *DomainChardevSourceTCP       `xml:"-" json:"-" yaml:"-"`
	UNIX      *DomainChardevSourceUNIX      `xml:"-" json:"-" yaml:"-"`
	SpiceVMC  *DomainChardevSourceSpiceVMC  `xml:"-" json:"-" yaml:"-"`
	SpicePort *DomainChardevSourceSpicePort `xml:"-" json:"-" yaml:"-"`
	NMDM      *DomainChardevSourceNMDM      `xml:"-" json:"-" yaml:"-"`
}

type DomainChardevSourceNull struct {
}

type DomainChardevSourceVC struct {
}

type DomainChardevSourcePty struct {
	Path     string                 `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainChardevSourceDev struct {
	Path     string                 `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainChardevSourceFile struct {
	Path     string                 `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	Append   string                 `xml:"append,attr,omitempty" json:"append,omitempty" yaml:"append,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainChardevSourcePipe struct {
	Path     string                 `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	SecLabel []DomainDeviceSecLabel `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainChardevSourceStdIO struct {
}

type DomainChardevSourceUDP struct {
	BindHost       string `xml:"-" json:"-" yaml:"-"`
	BindService    string `xml:"-" json:"-" yaml:"-"`
	ConnectHost    string `xml:"-" json:"-" yaml:"-"`
	ConnectService string `xml:"-" json:"-" yaml:"-"`
}

type DomainChardevSourceReconnect struct {
	Enabled string `xml:"enabled,attr" json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Timeout *uint  `xml:"timeout,attr" json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type DomainChardevSourceTCP struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Host      string                        `xml:"host,attr,omitempty" json:"host,omitempty" yaml:"host,omitempty"`
	Service   string                        `xml:"service,attr,omitempty" json:"service,omitempty" yaml:"service,omitempty"`
	TLS       string                        `xml:"tls,attr,omitempty" json:"tls,omitempty" yaml:"tls,omitempty"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect" yaml:"reconnect"`
}

type DomainChardevSourceUNIX struct {
	Mode      string                        `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Path      string                        `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	Reconnect *DomainChardevSourceReconnect `xml:"reconnect" json:"reconnect" yaml:"reconnect"`
	SecLabel  []DomainDeviceSecLabel        `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
}

type DomainChardevSourceSpiceVMC struct {
}

type DomainChardevSourceSpicePort struct {
	Channel string `xml:"channel,attr" json:"channel,omitempty" yaml:"channel,omitempty"`
}

type DomainChardevSourceNMDM struct {
	Master string `xml:"master,attr" json:"master,omitempty" yaml:"master,omitempty"`
	Slave  string `xml:"slave,attr" json:"slave,omitempty" yaml:"slave,omitempty"`
}

type DomainChardevTarget struct {
	Type  string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"` // is guest agent connected?
	Port  *uint  `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainConsoleTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Port *uint  `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainSerialTarget struct {
	Type  string                   `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Port  *uint                    `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
	Model *DomainSerialTargetModel `xml:"model" json:"model" yaml:"model"`
}

type DomainSerialTargetModel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainParallelTarget struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Port *uint  `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainChannelTarget struct {
	VirtIO   *DomainChannelTargetVirtIO   `xml:"-" json:"-" yaml:"-"`
	Xen      *DomainChannelTargetXen      `xml:"-" json:"-" yaml:"-"`
	GuestFWD *DomainChannelTargetGuestFWD `xml:"-" json:"-" yaml:"-"`
}

type DomainChannelTargetVirtIO struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"` // is guest agent connected?
}

type DomainChannelTargetXen struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"` // is guest agent connected?
}

type DomainChannelTargetGuestFWD struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Port    string `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainAlias struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainAddressPCI struct {
	Domain        *uint  `xml:"domain,attr" json:"domain,omitempty" yaml:"domain,omitempty"`
	Bus           *uint  `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Slot          *uint  `xml:"slot,attr" json:"slot,omitempty" yaml:"slot,omitempty"`
	Function      *uint  `xml:"function,attr" json:"function,omitempty" yaml:"function,omitempty"`
	MultiFunction string `xml:"multifunction,attr,omitempty" json:"multifunction,omitempty" yaml:"multifunction,omitempty"`
}

type DomainAddressUSB struct {
	Bus    *uint  `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Port   string `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Device *uint  `xml:"device,attr" json:"device,omitempty" yaml:"device,omitempty"`
}

type DomainAddressDrive struct {
	Controller *uint `xml:"controller,attr" json:"controller,omitempty" yaml:"controller,omitempty"`
	Bus        *uint `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Target     *uint `xml:"target,attr" json:"target,omitempty" yaml:"target,omitempty"`
	Unit       *uint `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainAddressDIMM struct {
	Slot *uint   `xml:"slot,attr" json:"slot,omitempty" yaml:"slot,omitempty"`
	Base *uint64 `xml:"base,attr" json:"base,omitempty" yaml:"base,omitempty"`
}

type DomainAddressISA struct {
	IOBase *uint `xml:"iobase,attr" json:"iobase,omitempty" yaml:"iobase,omitempty"`
	IRQ    *uint `xml:"irq,attr" json:"irq,omitempty" yaml:"irq,omitempty"`
}

type DomainAddressVirtioMMIO struct {
}

type DomainAddressCCW struct {
	CSSID *uint `xml:"cssid,attr" json:"cssid,omitempty" yaml:"cssid,omitempty"`
	SSID  *uint `xml:"ssid,attr" json:"ssid,omitempty" yaml:"ssid,omitempty"`
	DevNo *uint `xml:"devno,attr" json:"devno,omitempty" yaml:"devno,omitempty"`
}

type DomainAddressVirtioSerial struct {
	Controller *uint `xml:"controller,attr" json:"controller,omitempty" yaml:"controller,omitempty"`
	Bus        *uint `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Port       *uint `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
}

type DomainAddressSpaprVIO struct {
	Reg *uint64 `xml:"reg,attr" json:"reg,omitempty" yaml:"reg,omitempty"`
}

type DomainAddressCCID struct {
	Controller *uint `xml:"controller,attr" json:"controller,omitempty" yaml:"controller,omitempty"`
	Slot       *uint `xml:"slot,attr" json:"slot,omitempty" yaml:"slot,omitempty"`
}

type DomainAddressVirtioS390 struct {
}

type DomainAddress struct {
	PCI          *DomainAddressPCI
	Drive        *DomainAddressDrive
	VirtioSerial *DomainAddressVirtioSerial
	CCID         *DomainAddressCCID
	USB          *DomainAddressUSB
	SpaprVIO     *DomainAddressSpaprVIO
	VirtioS390   *DomainAddressVirtioS390
	CCW          *DomainAddressCCW
	VirtioMMIO   *DomainAddressVirtioMMIO
	ISA          *DomainAddressISA
	DIMM         *DomainAddressDIMM
}

type DomainChardevLog struct {
	File   string `xml:"file,attr" json:"file,omitempty" yaml:"file,omitempty"`
	Append string `xml:"append,attr,omitempty" json:"append,omitempty" yaml:"append,omitempty"`
}

type DomainConsole struct {
	XMLName  xml.Name               `xml:"console" json:"console" yaml:"console"`
	TTY      string                 `xml:"tty,attr,omitempty" json:"tty,omitempty" yaml:"tty,omitempty"`
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Target   *DomainConsoleTarget   `xml:"target" json:"target" yaml:"target"`
	Log      *DomainChardevLog      `xml:"log" json:"log" yaml:"log"`
	Alias    *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address  *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainSerial struct {
	XMLName  xml.Name               `xml:"serial" json:"serial" yaml:"serial"`
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Target   *DomainSerialTarget    `xml:"target" json:"target" yaml:"target"`
	Log      *DomainChardevLog      `xml:"log" json:"log" yaml:"log"`
	Alias    *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address  *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainParallel struct {
	XMLName  xml.Name               `xml:"parallel" json:"parallel" yaml:"parallel"`
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Target   *DomainParallelTarget  `xml:"target" json:"target" yaml:"target"`
	Log      *DomainChardevLog      `xml:"log" json:"log" yaml:"log"`
	Alias    *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address  *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainChardevProtocol struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainChannel struct {
	XMLName  xml.Name               `xml:"channel" json:"channel" yaml:"channel"`
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Target   *DomainChannelTarget   `xml:"target" json:"target" yaml:"target"`
	Log      *DomainChardevLog      `xml:"log" json:"log" yaml:"log"`
	Alias    *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address  *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainRedirDev struct {
	XMLName  xml.Name               `xml:"redirdev" json:"redirdev" yaml:"redirdev"`
	Bus      string                 `xml:"bus,attr,omitempty" json:"bus,omitempty" yaml:"bus,omitempty"`
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Boot     *DomainDeviceBoot      `xml:"boot" json:"boot" yaml:"boot"`
	Alias    *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address  *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainRedirFilter struct {
	USB []DomainRedirFilterUSB `xml:"usbdev" json:"usbdev" yaml:"usbdev"`
}

type DomainRedirFilterUSB struct {
	Class   *uint  `xml:"class,attr" json:"class,omitempty" yaml:"class,omitempty"`
	Vendor  *uint  `xml:"vendor,attr" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Product *uint  `xml:"product,attr" json:"product,omitempty" yaml:"product,omitempty"`
	Version string `xml:"version,attr,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
	Allow   string `xml:"allow,attr" json:"allow,omitempty" yaml:"allow,omitempty"`
}

type DomainInput struct {
	XMLName xml.Name           `xml:"input" json:"input" yaml:"input"`
	Type    string             `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Bus     string             `xml:"bus,attr,omitempty" json:"bus,omitempty" yaml:"bus,omitempty"`
	Driver  *DomainInputDriver `xml:"driver" json:"driver" yaml:"driver"`
	Source  *DomainInputSource `xml:"source" json:"source" yaml:"source"`
	Alias   *DomainAlias       `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress     `xml:"address" json:"address" yaml:"address"`
}

type DomainInputDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS   string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainInputSource struct {
	EVDev string `xml:"evdev,attr" json:"evdev,omitempty" yaml:"evdev,omitempty"`
}

type DomainGraphicListenerAddress struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
}

type DomainGraphicListenerNetwork struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Network string `xml:"network,attr,omitempty" json:"network,omitempty" yaml:"network,omitempty"`
}

type DomainGraphicListenerSocket struct {
	Socket string `xml:"socket,attr,omitempty" json:"socket,omitempty" yaml:"socket,omitempty"`
}

type DomainGraphicListener struct {
	Address *DomainGraphicListenerAddress `xml:"-" json:"-" yaml:"-"`
	Network *DomainGraphicListenerNetwork `xml:"-" json:"-" yaml:"-"`
	Socket  *DomainGraphicListenerSocket  `xml:"-" json:"-" yaml:"-"`
}

type DomainGraphicChannel struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainGraphicFileTransfer struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty" yaml:"enable,omitempty"`
}

type DomainGraphicsSDLGL struct {
	Enable string `xml:"enable,attr,omitempty" json:"enable,omitempty" yaml:"enable,omitempty"`
}

type DomainGraphicSDL struct {
	Display    string               `xml:"display,attr,omitempty" json:"display,omitempty" yaml:"display,omitempty"`
	XAuth      string               `xml:"xauth,attr,omitempty" json:"xauth,omitempty" yaml:"xauth,omitempty"`
	FullScreen string               `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty" yaml:"fullscreen,omitempty"`
	GL         *DomainGraphicsSDLGL `xml:"gl" json:"gl" yaml:"gl"`
}

type DomainGraphicVNC struct {
	Socket        string                  `xml:"socket,attr,omitempty" json:"socket,omitempty" yaml:"socket,omitempty"`
	Port          int                     `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	AutoPort      string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty" yaml:"autoport,omitempty"`
	WebSocket     int                     `xml:"websocket,attr,omitempty" json:"websocket,omitempty" yaml:"websocket,omitempty"`
	Keymap        string                  `xml:"keymap,attr,omitempty" json:"keymap,omitempty" yaml:"keymap,omitempty"`
	SharePolicy   string                  `xml:"sharePolicy,attr,omitempty" json:"sharePolicy,omitempty" yaml:"sharePolicy,omitempty"`
	Passwd        string                  `xml:"passwd,attr,omitempty" json:"passwd,omitempty" yaml:"passwd,omitempty"`
	PasswdValidTo string                  `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty" yaml:"passwdValidTo,omitempty"`
	Connected     string                  `xml:"connected,attr,omitempty" json:"connected,omitempty" yaml:"connected,omitempty"`
	Listen        string                  `xml:"listen,attr,omitempty" json:"listen,omitempty" yaml:"listen,omitempty"`
	Listeners     []DomainGraphicListener `xml:"listen" json:"listen" yaml:"listen"`
}

type DomainGraphicRDP struct {
	Port        int                     `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	AutoPort    string                  `xml:"autoport,attr,omitempty" json:"autoport,omitempty" yaml:"autoport,omitempty"`
	ReplaceUser string                  `xml:"replaceUser,attr,omitempty" json:"replaceUser,omitempty" yaml:"replaceUser,omitempty"`
	MultiUser   string                  `xml:"multiUser,attr,omitempty" json:"multiUser,omitempty" yaml:"multiUser,omitempty"`
	Listen      string                  `xml:"listen,attr,omitempty" json:"listen,omitempty" yaml:"listen,omitempty"`
	Listeners   []DomainGraphicListener `xml:"listen" json:"listen" yaml:"listen"`
}

type DomainGraphicDesktop struct {
	Display    string `xml:"display,attr,omitempty" json:"display,omitempty" yaml:"display,omitempty"`
	FullScreen string `xml:"fullscreen,attr,omitempty" json:"fullscreen,omitempty" yaml:"fullscreen,omitempty"`
}

type DomainGraphicSpiceChannel struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainGraphicSpiceImage struct {
	Compression string `xml:"compression,attr" json:"compression,omitempty" yaml:"compression,omitempty"`
}

type DomainGraphicSpiceJPEG struct {
	Compression string `xml:"compression,attr" json:"compression,omitempty" yaml:"compression,omitempty"`
}

type DomainGraphicSpiceZLib struct {
	Compression string `xml:"compression,attr" json:"compression,omitempty" yaml:"compression,omitempty"`
}

type DomainGraphicSpicePlayback struct {
	Compression string `xml:"compression,attr" json:"compression,omitempty" yaml:"compression,omitempty"`
}

type DomainGraphicSpiceStreaming struct {
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainGraphicSpiceMouse struct {
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainGraphicSpiceClipBoard struct {
	CopyPaste string `xml:"copypaste,attr" json:"copypaste,omitempty" yaml:"copypaste,omitempty"`
}

type DomainGraphicSpiceFileTransfer struct {
	Enable string `xml:"enable,attr" json:"enable,omitempty" yaml:"enable,omitempty"`
}

type DomainGraphicSpiceGL struct {
	Enable     string `xml:"enable,attr,omitempty" json:"enable,omitempty" yaml:"enable,omitempty"`
	RenderNode string `xml:"rendernode,attr,omitempty" json:"rendernode,omitempty" yaml:"rendernode,omitempty"`
}

type DomainGraphicSpice struct {
	Port          int                             `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	TLSPort       int                             `xml:"tlsPort,attr,omitempty" json:"tlsPort,omitempty" yaml:"tlsPort,omitempty"`
	AutoPort      string                          `xml:"autoport,attr,omitempty" json:"autoport,omitempty" yaml:"autoport,omitempty"`
	Listen        string                          `xml:"listen,attr,omitempty" json:"listen,omitempty" yaml:"listen,omitempty"`
	Keymap        string                          `xml:"keymap,attr,omitempty" json:"keymap,omitempty" yaml:"keymap,omitempty"`
	DefaultMode   string                          `xml:"defaultMode,attr,omitempty" json:"defaultMode,omitempty" yaml:"defaultMode,omitempty"`
	Passwd        string                          `xml:"passwd,attr,omitempty" json:"passwd,omitempty" yaml:"passwd,omitempty"`
	PasswdValidTo string                          `xml:"passwdValidTo,attr,omitempty" json:"passwdValidTo,omitempty" yaml:"passwdValidTo,omitempty"`
	Connected     string                          `xml:"connected,attr,omitempty" json:"connected,omitempty" yaml:"connected,omitempty"`
	Listeners     []DomainGraphicListener         `xml:"listen" json:"listen" yaml:"listen"`
	Channel       []DomainGraphicSpiceChannel     `xml:"channel" json:"channel" yaml:"channel"`
	Image         *DomainGraphicSpiceImage        `xml:"image" json:"image" yaml:"image"`
	JPEG          *DomainGraphicSpiceJPEG         `xml:"jpeg" json:"jpeg" yaml:"jpeg"`
	ZLib          *DomainGraphicSpiceZLib         `xml:"zlib" json:"zlib" yaml:"zlib"`
	Playback      *DomainGraphicSpicePlayback     `xml:"playback" json:"playback" yaml:"playback"`
	Streaming     *DomainGraphicSpiceStreaming    `xml:"streaming" json:"streaming" yaml:"streaming"`
	Mouse         *DomainGraphicSpiceMouse        `xml:"mouse" json:"mouse" yaml:"mouse"`
	ClipBoard     *DomainGraphicSpiceClipBoard    `xml:"clipboard" json:"clipboard" yaml:"clipboard"`
	FileTransfer  *DomainGraphicSpiceFileTransfer `xml:"filetransfer" json:"filetransfer" yaml:"filetransfer"`
	GL            *DomainGraphicSpiceGL           `xml:"gl" json:"gl" yaml:"gl"`
}

type DomainGraphicEGLHeadless struct {
}

type DomainGraphic struct {
	XMLName     xml.Name                  `xml:"graphics" json:"graphics" yaml:"graphics"`
	SDL         *DomainGraphicSDL         `xml:"-" json:"-" yaml:"-"`
	VNC         *DomainGraphicVNC         `xml:"-" json:"-" yaml:"-"`
	RDP         *DomainGraphicRDP         `xml:"-" json:"-" yaml:"-"`
	Desktop     *DomainGraphicDesktop     `xml:"-" json:"-" yaml:"-"`
	Spice       *DomainGraphicSpice       `xml:"-" json:"-" yaml:"-"`
	EGLHeadless *DomainGraphicEGLHeadless `xml:"-" json:"-" yaml:"-"`
}

type DomainVideoAccel struct {
	Accel3D string `xml:"accel3d,attr,omitempty" json:"accel3d,omitempty" yaml:"accel3d,omitempty"`
	Accel2D string `xml:"accel2d,attr,omitempty" json:"accel2d,omitempty" yaml:"accel2d,omitempty"`
}

type DomainVideoModel struct {
	Type    string            `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Heads   uint              `xml:"heads,attr,omitempty" json:"heads,omitempty" yaml:"heads,omitempty"`
	Ram     uint              `xml:"ram,attr,omitempty" json:"ram,omitempty" yaml:"ram,omitempty"`
	VRam    uint              `xml:"vram,attr,omitempty" json:"vram,omitempty" yaml:"vram,omitempty"`
	VRam64  uint              `xml:"vram64,attr,omitempty" json:"vram64,omitempty" yaml:"vram64,omitempty"`
	VGAMem  uint              `xml:"vgamem,attr,omitempty" json:"vgamem,omitempty" yaml:"vgamem,omitempty"`
	Primary string            `xml:"primary,attr,omitempty" json:"primary,omitempty" yaml:"primary,omitempty"`
	Accel   *DomainVideoAccel `xml:"acceleration" json:"acceleration" yaml:"acceleration"`
}

type DomainVideo struct {
	XMLName xml.Name           `xml:"video" json:"video" yaml:"video"`
	Model   DomainVideoModel   `xml:"model" json:"model" yaml:"model"`
	Driver  *DomainVideoDriver `xml:"driver" json:"driver" yaml:"driver"`
	Alias   *DomainAlias       `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress     `xml:"address" json:"address" yaml:"address"`
}

type DomainVideoDriver struct {
	VGAConf string `xml:"vgaconf,attr,omitempty" json:"vgaconf,omitempty" yaml:"vgaconf,omitempty"`
	IOMMU   string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS     string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainMemBalloonStats struct {
	Period uint `xml:"period,attr" json:"period,omitempty" yaml:"period,omitempty"`
}

type DomainMemBalloon struct {
	XMLName     xml.Name                `xml:"memballoon" json:"memballoon" yaml:"memballoon"`
	Model       string                  `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	AutoDeflate string                  `xml:"autodeflate,attr,omitempty" json:"autodeflate,omitempty" yaml:"autodeflate,omitempty"`
	Driver      *DomainMemBalloonDriver `xml:"driver" json:"driver" yaml:"driver"`
	Stats       *DomainMemBalloonStats  `xml:"stats" json:"stats" yaml:"stats"`
	Alias       *DomainAlias            `xml:"alias" json:"alias" yaml:"alias"`
	Address     *DomainAddress          `xml:"address" json:"address" yaml:"address"`
}

type DomainVSockCID struct {
	Auto    string `xml:"auto,attr,omitempty" json:"auto,omitempty" yaml:"auto,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
}

type DomainVSock struct {
	XMLName xml.Name        `xml:"vsock" json:"vsock" yaml:"vsock"`
	Model   string          `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	CID     *DomainVSockCID `xml:"cid" json:"cid" yaml:"cid"`
	Alias   *DomainAlias    `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress  `xml:"address" json:"address" yaml:"address"`
}

type DomainMemBalloonDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS   string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainPanic struct {
	XMLName xml.Name       `xml:"panic" json:"panic" yaml:"panic"`
	Model   string         `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Alias   *DomainAlias   `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress `xml:"address" json:"address" yaml:"address"`
}

type DomainSoundCodec struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainSound struct {
	XMLName xml.Name           `xml:"sound" json:"sound" yaml:"sound"`
	Model   string             `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	Codec   []DomainSoundCodec `xml:"codec" json:"codec" yaml:"codec"`
	Alias   *DomainAlias       `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress     `xml:"address" json:"address" yaml:"address"`
}

type DomainRNGRate struct {
	Bytes  uint `xml:"bytes,attr" json:"bytes,omitempty" yaml:"bytes,omitempty"`
	Period uint `xml:"period,attr,omitempty" json:"period,omitempty" yaml:"period,omitempty"`
}

type DomainRNGBackend struct {
	Random *DomainRNGBackendRandom `xml:"-" json:"-" yaml:"-"`
	EGD    *DomainRNGBackendEGD    `xml:"-" json:"-" yaml:"-"`
}

type DomainRNGBackendEGD struct {
	Source   *DomainChardevSource   `xml:"source" json:"source" yaml:"source"`
	Protocol *DomainChardevProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
}

type DomainRNGBackendRandom struct {
	Device string `xml:",attr" json:"" yaml:""`
}

type DomainRNG struct {
	XMLName xml.Name          `xml:"rng" json:"rng" yaml:"rng"`
	Model   string            `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	Driver  *DomainRNGDriver  `xml:"driver" json:"driver" yaml:"driver"`
	Rate    *DomainRNGRate    `xml:"rate" json:"rate" yaml:"rate"`
	Backend *DomainRNGBackend `xml:"backend" json:"backend" yaml:"backend"`
	Alias   *DomainAlias      `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress    `xml:"address" json:"address" yaml:"address"`
}

type DomainRNGDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty" json:"iommu,omitempty" yaml:"iommu,omitempty"`
	ATS   string `xml:"ats,attr,omitempty" json:"ats,omitempty" yaml:"ats,omitempty"`
}

type DomainHostdevSubsysUSB struct {
	Source *DomainHostdevSubsysUSBSource `xml:"source" json:"source" yaml:"source"`
}

type DomainHostdevSubsysUSBSource struct {
	Address *DomainAddressUSB `xml:"address" json:"address" yaml:"address"`
}

type DomainHostdevSubsysSCSI struct {
	SGIO      string                         `xml:"sgio,attr,omitempty" json:"sgio,omitempty" yaml:"sgio,omitempty"`
	RawIO     string                         `xml:"rawio,attr,omitempty" json:"rawio,omitempty" yaml:"rawio,omitempty"`
	Source    *DomainHostdevSubsysSCSISource `xml:"source" json:"source" yaml:"source"`
	ReadOnly  *DomainDiskReadOnly            `xml:"readonly" json:"readonly" yaml:"readonly"`
	Shareable *DomainDiskShareable           `xml:"shareable" json:"shareable" yaml:"shareable"`
}

type DomainHostdevSubsysSCSISource struct {
	Host  *DomainHostdevSubsysSCSISourceHost  `xml:"-" json:"-" yaml:"-"`
	ISCSI *DomainHostdevSubsysSCSISourceISCSI `xml:"-" json:"-" yaml:"-"`
}

type DomainHostdevSubsysSCSIAdapter struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainHostdevSubsysSCSISourceHost struct {
	Adapter *DomainHostdevSubsysSCSIAdapter `xml:"adapter" json:"adapter" yaml:"adapter"`
	Address *DomainAddressDrive             `xml:"address" json:"address" yaml:"address"`
}

type DomainHostdevSubsysSCSISourceISCSI struct {
	Name string                 `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Host []DomainDiskSourceHost `xml:"host" json:"host" yaml:"host"`
	Auth *DomainDiskAuth        `xml:"auth" json:"auth" yaml:"auth"`
}

type DomainHostdevSubsysSCSIHost struct {
	Source *DomainHostdevSubsysSCSIHostSource `xml:"source" json:"source" yaml:"source"`
}

type DomainHostdevSubsysSCSIHostSource struct {
	Protocol string `xml:"protocol,attr,omitempty" json:"protocol,omitempty" yaml:"protocol,omitempty"`
	WWPN     string `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty" yaml:"wwpn,omitempty"`
}

type DomainHostdevSubsysPCISource struct {
	Address *DomainAddressPCI `xml:"address" json:"address" yaml:"address"`
}

type DomainHostdevSubsysPCIDriver struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainHostdevSubsysPCI struct {
	Driver *DomainHostdevSubsysPCIDriver `xml:"driver" json:"driver" yaml:"driver"`
	Source *DomainHostdevSubsysPCISource `xml:"source" json:"source" yaml:"source"`
}

type DomainAddressMDev struct {
	UUID string `xml:"uuid,attr" json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DomainHostdevSubsysMDevSource struct {
	Address *DomainAddressMDev `xml:"address" json:"address" yaml:"address"`
}

type DomainHostdevSubsysMDev struct {
	Model   string                         `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Display string                         `xml:"display,attr,omitempty" json:"display,omitempty" yaml:"display,omitempty"`
	Source  *DomainHostdevSubsysMDevSource `xml:"source" json:"source" yaml:"source"`
}

type DomainHostdevCapsStorage struct {
	Source *DomainHostdevCapsStorageSource `xml:"source" json:"source" yaml:"source"`
}

type DomainHostdevCapsStorageSource struct {
	Block string `xml:"block" json:"block" yaml:"block"`
}

type DomainHostdevCapsMisc struct {
	Source *DomainHostdevCapsMiscSource `xml:"source" json:"source" yaml:"source"`
}

type DomainHostdevCapsMiscSource struct {
	Char string `xml:"char" json:"char" yaml:"char"`
}

type DomainIP struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" yaml:"family,omitempty"`
	Prefix  *uint  `xml:"prefix,attr" json:"prefix,omitempty" yaml:"prefix,omitempty"`
}

type DomainRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty" yaml:"family,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty" yaml:"gateway,omitempty"`
}

type DomainHostdevCapsNet struct {
	Source *DomainHostdevCapsNetSource `xml:"source" json:"source" yaml:"source"`
	IP     []DomainIP                  `xml:"ip" json:"ip" yaml:"ip"`
	Route  []DomainRoute               `xml:"route" json:"route" yaml:"route"`
}

type DomainHostdevCapsNetSource struct {
	Interface string `xml:"interface" json:"interface" yaml:"interface"`
}

type DomainHostdev struct {
	Managed        string                       `xml:"managed,attr,omitempty" json:"managed,omitempty" yaml:"managed,omitempty"`
	SubsysUSB      *DomainHostdevSubsysUSB      `xml:"-" json:"-" yaml:"-"`
	SubsysSCSI     *DomainHostdevSubsysSCSI     `xml:"-" json:"-" yaml:"-"`
	SubsysSCSIHost *DomainHostdevSubsysSCSIHost `xml:"-" json:"-" yaml:"-"`
	SubsysPCI      *DomainHostdevSubsysPCI      `xml:"-" json:"-" yaml:"-"`
	SubsysMDev     *DomainHostdevSubsysMDev     `xml:"-" json:"-" yaml:"-"`
	CapsStorage    *DomainHostdevCapsStorage    `xml:"-" json:"-" yaml:"-"`
	CapsMisc       *DomainHostdevCapsMisc       `xml:"-" json:"-" yaml:"-"`
	CapsNet        *DomainHostdevCapsNet        `xml:"-" json:"-" yaml:"-"`
	Boot           *DomainDeviceBoot            `xml:"boot" json:"boot" yaml:"boot"`
	ROM            *DomainROM                   `xml:"rom" json:"rom" yaml:"rom"`
	Alias          *DomainAlias                 `xml:"alias" json:"alias" yaml:"alias"`
	Address        *DomainAddress               `xml:"address" json:"address" yaml:"address"`
}

type DomainMemorydevSource struct {
	NodeMask string                         `xml:"nodemask,omitempty" json:"nodemask,omitempty" yaml:"nodemask,omitempty"`
	PageSize *DomainMemorydevSourcePagesize `xml:"pagesize" json:"pagesize" yaml:"pagesize"`
	Path     string                         `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
}

type DomainMemorydevSourcePagesize struct {
	Value uint64 `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainMemorydevTargetNode struct {
	Value uint `xml:",attr" json:"" yaml:""`
}

type DomainMemorydevTargetSize struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainMemorydevTargetLabel struct {
	Size *DomainMemorydevTargetSize `xml:"size" json:"size" yaml:"size"`
}

type DomainMemorydevTarget struct {
	Size  *DomainMemorydevTargetSize  `xml:"size" json:"size" yaml:"size"`
	Node  *DomainMemorydevTargetNode  `xml:"node" json:"node" yaml:"node"`
	Label *DomainMemorydevTargetLabel `xml:"label" json:"label" yaml:"label"`
}

type DomainMemorydev struct {
	XMLName xml.Name               `xml:"memory" json:"memory" yaml:"memory"`
	Model   string                 `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	Access  string                 `xml:"access,attr,omitempty" json:"access,omitempty" yaml:"access,omitempty"`
	Discard string                 `xml:"discard,attr,omitempty" json:"discard,omitempty" yaml:"discard,omitempty"`
	Source  *DomainMemorydevSource `xml:"source" json:"source" yaml:"source"`
	Target  *DomainMemorydevTarget `xml:"target" json:"target" yaml:"target"`
	Alias   *DomainAlias           `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress         `xml:"address" json:"address" yaml:"address"`
}

type DomainWatchdog struct {
	XMLName xml.Name       `xml:"watchdog" json:"watchdog" yaml:"watchdog"`
	Model   string         `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	Action  string         `xml:"action,attr,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	Alias   *DomainAlias   `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress `xml:"address" json:"address" yaml:"address"`
}

type DomainHub struct {
	Type    string         `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Alias   *DomainAlias   `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress `xml:"address" json:"address" yaml:"address"`
}

type DomainIOMMU struct {
	Model  string             `xml:"model,attr" json:"model,omitempty" yaml:"model,omitempty"`
	Driver *DomainIOMMUDriver `xml:"driver" json:"driver" yaml:"driver"`
}

type DomainIOMMUDriver struct {
	IntRemap    string `xml:"intremap,attr,omitempty" json:"intremap,omitempty" yaml:"intremap,omitempty"`
	CachingMode string `xml:"caching_mode,attr,omitempty" json:"caching_mode,omitempty" yaml:"caching_mode,omitempty"`
	EIM         string `xml:"eim,attr,omitempty" json:"eim,omitempty" yaml:"eim,omitempty"`
	IOTLB       string `xml:"iotlb,attr,omitempty" json:"iotlb,omitempty" yaml:"iotlb,omitempty"`
}

type DomainNVRAM struct {
	Alias   *DomainAlias   `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress `xml:"address" json:"address" yaml:"address"`
}

type DomainLease struct {
	Lockspace string             `xml:"lockspace" json:"lockspace" yaml:"lockspace"`
	Key       string             `xml:"key" json:"key" yaml:"key"`
	Target    *DomainLeaseTarget `xml:"target" json:"target" yaml:"target"`
}

type DomainLeaseTarget struct {
	Path   string `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	Offset uint64 `xml:"offset,attr,omitempty" json:"offset,omitempty" yaml:"offset,omitempty"`
}

type DomainSmartcard struct {
	XMLName     xml.Name                  `xml:"smartcard" json:"smartcard" yaml:"smartcard"`
	Passthrough *DomainChardevSource      `xml:"source" json:"source" yaml:"source"`
	Protocol    *DomainChardevProtocol    `xml:"protocol" json:"protocol" yaml:"protocol"`
	Host        *DomainSmartcardHost      `xml:"-" json:"-" yaml:"-"`
	HostCerts   []DomainSmartcardHostCert `xml:"certificate" json:"certificate" yaml:"certificate"`
	Database    string                    `xml:"database,omitempty" json:"database,omitempty" yaml:"database,omitempty"`
	Alias       *DomainAlias              `xml:"alias" json:"alias" yaml:"alias"`
	Address     *DomainAddress            `xml:"address" json:"address" yaml:"address"`
}

type DomainSmartcardHost struct {
}

type DomainSmartcardHostCert struct {
	File string `xml:",attr" json:"" yaml:""`
}

type DomainTPM struct {
	XMLName xml.Name          `xml:"tpm" json:"tpm" yaml:"tpm"`
	Model   string            `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Backend *DomainTPMBackend `xml:"backend" json:"backend" yaml:"backend"`
	Alias   *DomainAlias      `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress    `xml:"address" json:"address" yaml:"address"`
}

type DomainTPMBackend struct {
	Passthrough *DomainTPMBackendPassthrough `xml:"-" json:"-" yaml:"-"`
	Emulator    *DomainTPMBackendEmulator    `xml:"-" json:"-" yaml:"-"`
}

type DomainTPMBackendPassthrough struct {
	Device *DomainTPMBackendDevice `xml:"device" json:"device" yaml:"device"`
}

type DomainTPMBackendEmulator struct {
	Version string `xml:"version,attr,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
}

type DomainTPMBackendDevice struct {
	Path string `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
}

type DomainShmem struct {
	XMLName xml.Name           `xml:"shmem" json:"shmem" yaml:"shmem"`
	Name    string             `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Size    *DomainShmemSize   `xml:"size" json:"size" yaml:"size"`
	Model   *DomainShmemModel  `xml:"model" json:"model" yaml:"model"`
	Server  *DomainShmemServer `xml:"server" json:"server" yaml:"server"`
	MSI     *DomainShmemMSI    `xml:"msi" json:"msi" yaml:"msi"`
	Alias   *DomainAlias       `xml:"alias" json:"alias" yaml:"alias"`
	Address *DomainAddress     `xml:"address" json:"address" yaml:"address"`
}

type DomainShmemSize struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainShmemModel struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainShmemServer struct {
	Path string `xml:"path,attr,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
}

type DomainShmemMSI struct {
	Enabled   string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Vectors   uint   `xml:"vectors,attr,omitempty" json:"vectors,omitempty" yaml:"vectors,omitempty"`
	IOEventFD string `xml:"ioeventfd,attr,omitempty" json:"ioeventfd,omitempty" yaml:"ioeventfd,omitempty"`
}

type DomainDeviceList struct {
	Emulator     string              `xml:"emulator,omitempty" json:"emulator,omitempty" yaml:"emulator,omitempty"`
	Disks        []DomainDisk        `xml:"disk" json:"disk" yaml:"disk"`
	Controllers  []DomainController  `xml:"controller" json:"controller" yaml:"controller"`
	Leases       []DomainLease       `xml:"lease" json:"lease" yaml:"lease"`
	Filesystems  []DomainFilesystem  `xml:"filesystem" json:"filesystem" yaml:"filesystem"`
	Interfaces   []DomainInterface   `xml:"interface" json:"interface" yaml:"interface"`
	Smartcards   []DomainSmartcard   `xml:"smartcard" json:"smartcard" yaml:"smartcard"`
	Serials      []DomainSerial      `xml:"serial" json:"serial" yaml:"serial"`
	Parallels    []DomainParallel    `xml:"parallel" json:"parallel" yaml:"parallel"`
	Consoles     []DomainConsole     `xml:"console" json:"console" yaml:"console"`
	Channels     []DomainChannel     `xml:"channel" json:"channel" yaml:"channel"`
	Inputs       []DomainInput       `xml:"input" json:"input" yaml:"input"`
	TPMs         []DomainTPM         `xml:"tpm" json:"tpm" yaml:"tpm"`
	Graphics     []DomainGraphic     `xml:"graphics" json:"graphics" yaml:"graphics"`
	Sounds       []DomainSound       `xml:"sound" json:"sound" yaml:"sound"`
	Videos       []DomainVideo       `xml:"video" json:"video" yaml:"video"`
	Hostdevs     []DomainHostdev     `xml:"hostdev" json:"hostdev" yaml:"hostdev"`
	RedirDevs    []DomainRedirDev    `xml:"redirdev" json:"redirdev" yaml:"redirdev"`
	RedirFilters []DomainRedirFilter `xml:"redirfilter" json:"redirfilter" yaml:"redirfilter"`
	Hubs         []DomainHub         `xml:"hub" json:"hub" yaml:"hub"`
	Watchdog     *DomainWatchdog     `xml:"watchdog" json:"watchdog" yaml:"watchdog"`
	MemBalloon   *DomainMemBalloon   `xml:"memballoon" json:"memballoon" yaml:"memballoon"`
	RNGs         []DomainRNG         `xml:"rng" json:"rng" yaml:"rng"`
	NVRAM        *DomainNVRAM        `xml:"nvram" json:"nvram" yaml:"nvram"`
	Panics       []DomainPanic       `xml:"panic" json:"panic" yaml:"panic"`
	Shmems       []DomainShmem       `xml:"shmem" json:"shmem" yaml:"shmem"`
	Memorydevs   []DomainMemorydev   `xml:"memory" json:"memory" yaml:"memory"`
	IOMMU        *DomainIOMMU        `xml:"iommu" json:"iommu" yaml:"iommu"`
	VSock        *DomainVSock        `xml:"vsock" json:"vsock" yaml:"vsock"`
}

type DomainMemory struct {
	Value    uint   `xml:",attr" json:"" yaml:""`
	Unit     string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	DumpCore string `xml:"dumpCore,attr,omitempty" json:"dumpCore,omitempty" yaml:"dumpCore,omitempty"`
}

type DomainCurrentMemory struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainMaxMemory struct {
	Value uint   `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Slots uint   `xml:"slots,attr,omitempty" json:"slots,omitempty" yaml:"slots,omitempty"`
}

type DomainMemoryHugepage struct {
	Size    uint   `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Unit    string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Nodeset string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty" yaml:"nodeset,omitempty"`
}

type DomainMemoryHugepages struct {
	Hugepages []DomainMemoryHugepage `xml:"page" json:"page" yaml:"page"`
}

type DomainMemoryNosharepages struct {
}

type DomainMemoryLocked struct {
}

type DomainMemorySource struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainMemoryAccess struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainMemoryAllocation struct {
	Mode string `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainMemoryDiscard struct {
}

type DomainMemoryBacking struct {
	MemoryHugePages    *DomainMemoryHugepages    `xml:"hugepages" json:"hugepages" yaml:"hugepages"`
	MemoryNosharepages *DomainMemoryNosharepages `xml:"nosharepages" json:"nosharepages" yaml:"nosharepages"`
	MemoryLocked       *DomainMemoryLocked       `xml:"locked" json:"locked" yaml:"locked"`
	MemorySource       *DomainMemorySource       `xml:"source" json:"source" yaml:"source"`
	MemoryAccess       *DomainMemoryAccess       `xml:"access" json:"access" yaml:"access"`
	MemoryAllocation   *DomainMemoryAllocation   `xml:"allocation" json:"allocation" yaml:"allocation"`
	MemoryDiscard      *DomainMemoryDiscard      `xml:"discard" json:"discard" yaml:"discard"`
}

type DomainOSType struct {
	Arch    string `xml:"arch,attr,omitempty" json:"arch,omitempty" yaml:"arch,omitempty"`
	Machine string `xml:"machine,attr,omitempty" json:"machine,omitempty" yaml:"machine,omitempty"`
	Type    string `xml:",attr" json:"" yaml:""`
}

type DomainSMBios struct {
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainNVRam struct {
	NVRam    string `xml:",attr" json:"" yaml:""`
	Template string `xml:"template,attr,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
}

type DomainBootDevice struct {
	Dev string `xml:"dev,attr" json:"dev,omitempty" yaml:"dev,omitempty"`
}

type DomainBootMenu struct {
	Enable  string `xml:"enable,attr,omitempty" json:"enable,omitempty" yaml:"enable,omitempty"`
	Timeout string `xml:"timeout,attr,omitempty" json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type DomainSysInfoBIOS struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoSystem struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoBaseBoard struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoProcessor struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoMemory struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoChassis struct {
	Entry []DomainSysInfoEntry `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfoOEMStrings struct {
	Entry []string `xml:"entry" json:"entry" yaml:"entry"`
}

type DomainSysInfo struct {
	Type       string                   `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	BIOS       *DomainSysInfoBIOS       `xml:"bios" json:"bios" yaml:"bios"`
	System     *DomainSysInfoSystem     `xml:"system" json:"system" yaml:"system"`
	BaseBoard  []DomainSysInfoBaseBoard `xml:"baseBoard" json:"baseBoard" yaml:"baseBoard"`
	Chassis    *DomainSysInfoChassis    `xml:"chassis" json:"chassis" yaml:"chassis"`
	Processor  []DomainSysInfoProcessor `xml:"processor" json:"processor" yaml:"processor"`
	Memory     []DomainSysInfoMemory    `xml:"memory" json:"memory" yaml:"memory"`
	OEMStrings *DomainSysInfoOEMStrings `xml:"oemStrings" json:"oemStrings" yaml:"oemStrings"`
}

type DomainSysInfoEntry struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value string `xml:",attr" json:"" yaml:""`
}

type DomainBIOS struct {
	UseSerial     string `xml:"useserial,attr,omitempty" json:"useserial,omitempty" yaml:"useserial,omitempty"`
	RebootTimeout *int   `xml:"rebootTimeout,attr" json:"rebootTimeout,omitempty" yaml:"rebootTimeout,omitempty"`
}

type DomainLoader struct {
	Path     string `xml:",attr" json:"" yaml:""`
	Readonly string `xml:"readonly,attr,omitempty" json:"readonly,omitempty" yaml:"readonly,omitempty"`
	Secure   string `xml:"secure,attr,omitempty" json:"secure,omitempty" yaml:"secure,omitempty"`
	Type     string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

type DomainACPI struct {
	Tables []DomainACPITable `xml:"table" json:"table" yaml:"table"`
}

type DomainACPITable struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Path string `xml:",attr" json:"" yaml:""`
}

type DomainOSInitEnv struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value string `xml:",attr" json:"" yaml:""`
}

type DomainOS struct {
	Type        *DomainOSType      `xml:"type" json:"type" yaml:"type"`
	Init        string             `xml:"init,omitempty" json:"init,omitempty" yaml:"init,omitempty"`
	InitArgs    []string           `xml:"initarg" json:"initarg" yaml:"initarg"`
	InitEnv     []DomainOSInitEnv  `xml:"initenv" json:"initenv" yaml:"initenv"`
	InitDir     string             `xml:"initdir,omitempty" json:"initdir,omitempty" yaml:"initdir,omitempty"`
	InitUser    string             `xml:"inituser,omitempty" json:"inituser,omitempty" yaml:"inituser,omitempty"`
	InitGroup   string             `xml:"initgroup,omitempty" json:"initgroup,omitempty" yaml:"initgroup,omitempty"`
	Loader      *DomainLoader      `xml:"loader" json:"loader" yaml:"loader"`
	NVRam       *DomainNVRam       `xml:"nvram" json:"nvram" yaml:"nvram"`
	Kernel      string             `xml:"kernel,omitempty" json:"kernel,omitempty" yaml:"kernel,omitempty"`
	Initrd      string             `xml:"initrd,omitempty" json:"initrd,omitempty" yaml:"initrd,omitempty"`
	Cmdline     string             `xml:"cmdline,omitempty" json:"cmdline,omitempty" yaml:"cmdline,omitempty"`
	DTB         string             `xml:"dtb,omitempty" json:"dtb,omitempty" yaml:"dtb,omitempty"`
	ACPI        *DomainACPI        `xml:"acpi" json:"acpi" yaml:"acpi"`
	BootDevices []DomainBootDevice `xml:"boot" json:"boot" yaml:"boot"`
	BootMenu    *DomainBootMenu    `xml:"bootmenu" json:"bootmenu" yaml:"bootmenu"`
	BIOS        *DomainBIOS        `xml:"bios" json:"bios" yaml:"bios"`
	SMBios      *DomainSMBios      `xml:"smbios" json:"smbios" yaml:"smbios"`
}

type DomainResource struct {
	Partition string `xml:"partition,omitempty" json:"partition,omitempty" yaml:"partition,omitempty"`
}

type DomainVCPU struct {
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty" yaml:"placement,omitempty"`
	CPUSet    string `xml:"cpuset,attr,omitempty" json:"cpuset,omitempty" yaml:"cpuset,omitempty"`
	Current   string `xml:"current,attr,omitempty" json:"current,omitempty" yaml:"current,omitempty"`
	Value     int    `xml:",attr" json:"" yaml:""`
}

type DomainVCPUsVCPU struct {
	Id           *uint  `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Enabled      string `xml:"enabled,attr,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty" json:"hotpluggable,omitempty" yaml:"hotpluggable,omitempty"`
	Order        *uint  `xml:"order,attr" json:"order,omitempty" yaml:"order,omitempty"`
}

type DomainVCPUs struct {
	VCPU []DomainVCPUsVCPU `xml:"vcpu" json:"vcpu" yaml:"vcpu"`
}

type DomainCPUModel struct {
	Fallback string `xml:"fallback,attr,omitempty" json:"fallback,omitempty" yaml:"fallback,omitempty"`
	Value    string `xml:",attr" json:"" yaml:""`
	VendorID string `xml:"vendor_id,attr,omitempty" json:"vendor_id,omitempty" yaml:"vendor_id,omitempty"`
}

type DomainCPUTopology struct {
	Sockets int `xml:"sockets,attr,omitempty" json:"sockets,omitempty" yaml:"sockets,omitempty"`
	Cores   int `xml:"cores,attr,omitempty" json:"cores,omitempty" yaml:"cores,omitempty"`
	Threads int `xml:"threads,attr,omitempty" json:"threads,omitempty" yaml:"threads,omitempty"`
}

type DomainCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty" json:"policy,omitempty" yaml:"policy,omitempty"`
	Name   string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type DomainCPUCache struct {
	Level uint   `xml:"level,attr,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Mode  string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type DomainCPU struct {
	XMLName  xml.Name           `xml:"cpu" json:"cpu" yaml:"cpu"`
	Match    string             `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	Mode     string             `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Check    string             `xml:"check,attr,omitempty" json:"check,omitempty" yaml:"check,omitempty"`
	Model    *DomainCPUModel    `xml:"model" json:"model" yaml:"model"`
	Vendor   string             `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Topology *DomainCPUTopology `xml:"topology" json:"topology" yaml:"topology"`
	Cache    *DomainCPUCache    `xml:"cache" json:"cache" yaml:"cache"`
	Features []DomainCPUFeature `xml:"feature" json:"feature" yaml:"feature"`
	Numa     *DomainNuma        `xml:"numa" json:"numa" yaml:"numa"`
}

type DomainNuma struct {
	Cell []DomainCell `xml:"cell" json:"cell" yaml:"cell"`
}

type DomainCell struct {
	ID        *uint                `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	CPUs      string               `xml:"cpus,attr" json:"cpus,omitempty" yaml:"cpus,omitempty"`
	Memory    string               `xml:"memory,attr" json:"memory,omitempty" yaml:"memory,omitempty"`
	Unit      string               `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	MemAccess string               `xml:"memAccess,attr,omitempty" json:"memAccess,omitempty" yaml:"memAccess,omitempty"`
	Discard   string               `xml:"discard,attr,omitempty" json:"discard,omitempty" yaml:"discard,omitempty"`
	Distances *DomainCellDistances `xml:"distances" json:"distances" yaml:"distances"`
}

type DomainCellDistances struct {
	Siblings []DomainCellSibling `xml:"sibling" json:"sibling" yaml:"sibling"`
}

type DomainCellSibling struct {
	ID    uint `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Value uint `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty" json:"offset,omitempty" yaml:"offset,omitempty"`
	Basis      string        `xml:"basis,attr,omitempty" json:"basis,omitempty" yaml:"basis,omitempty"`
	Adjustment string        `xml:"adjustment,attr,omitempty" json:"adjustment,omitempty" yaml:"adjustment,omitempty"`
	TimeZone   string        `xml:"timezone,attr,omitempty" json:"timezone,omitempty" yaml:"timezone,omitempty"`
	Timer      []DomainTimer `xml:"timer" json:"timer" yaml:"timer"`
}

type DomainTimer struct {
	Name       string              `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Track      string              `xml:"track,attr,omitempty" json:"track,omitempty" yaml:"track,omitempty"`
	TickPolicy string              `xml:"tickpolicy,attr,omitempty" json:"tickpolicy,omitempty" yaml:"tickpolicy,omitempty"`
	CatchUp    *DomainTimerCatchUp `xml:"catchup" json:"catchup" yaml:"catchup"`
	Frequency  uint32              `xml:"frequency,attr,omitempty" json:"frequency,omitempty" yaml:"frequency,omitempty"`
	Mode       string              `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Present    string              `xml:"present,attr,omitempty" json:"present,omitempty" yaml:"present,omitempty"`
}

type DomainTimerCatchUp struct {
	Threshold uint `xml:"threshold,attr,omitempty" json:"threshold,omitempty" yaml:"threshold,omitempty"`
	Slew      uint `xml:"slew,attr,omitempty" json:"slew,omitempty" yaml:"slew,omitempty"`
	Limit     uint `xml:"limit,attr,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

type DomainFeature struct {
}

type DomainFeatureState struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
}

type DomainFeatureAPIC struct {
	EOI string `xml:"eoi,attr,omitempty" json:"eoi,omitempty" yaml:"eoi,omitempty"`
}

type DomainFeatureHyperVVendorId struct {
	DomainFeatureState
	Value string `xml:"value,attr,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainFeatureHyperVSpinlocks struct {
	DomainFeatureState
	Retries uint `xml:"retries,attr,omitempty" json:"retries,omitempty" yaml:"retries,omitempty"`
}

type DomainFeatureHyperV struct {
	DomainFeature
	Relaxed   *DomainFeatureState           `xml:"relaxed" json:"relaxed" yaml:"relaxed"`
	VAPIC     *DomainFeatureState           `xml:"vapic" json:"vapic" yaml:"vapic"`
	Spinlocks *DomainFeatureHyperVSpinlocks `xml:"spinlocks" json:"spinlocks" yaml:"spinlocks"`
	VPIndex   *DomainFeatureState           `xml:"vpindex" json:"vpindex" yaml:"vpindex"`
	Runtime   *DomainFeatureState           `xml:"runtime" json:"runtime" yaml:"runtime"`
	Synic     *DomainFeatureState           `xml:"synic" json:"synic" yaml:"synic"`
	STimer    *DomainFeatureState           `xml:"stimer" json:"stimer" yaml:"stimer"`
	Reset     *DomainFeatureState           `xml:"reset" json:"reset" yaml:"reset"`
	VendorId  *DomainFeatureHyperVVendorId  `xml:"vendor_id" json:"vendor_id" yaml:"vendor_id"`
}

type DomainFeatureKVM struct {
	Hidden *DomainFeatureState `xml:"hidden" json:"hidden" yaml:"hidden"`
}

type DomainFeatureGIC struct {
	Version string `xml:"version,attr,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
}

type DomainFeatureIOAPIC struct {
	Driver string `xml:"driver,attr,omitempty" json:"driver,omitempty" yaml:"driver,omitempty"`
}

type DomainFeatureHPT struct {
	Resizing    string                    `xml:"resizing,attr,omitempty" json:"resizing,omitempty" yaml:"resizing,omitempty"`
	MaxPageSize *DomainFeatureHPTPageSize `xml:"maxpagesize" json:"maxpagesize" yaml:"maxpagesize"`
}

type DomainFeatureHPTPageSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Value string `xml:",attr" json:"" yaml:""`
}

type DomainFeatureSMM struct {
	State string                `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	TSeg  *DomainFeatureSMMTSeg `xml:"tseg" json:"tseg" yaml:"tseg"`
}

type DomainFeatureSMMTSeg struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Value uint   `xml:",attr" json:"" yaml:""`
}

type DomainFeatureCapability struct {
	State string `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
}

type DomainLaunchSecurity struct {
	SEV *DomainLaunchSecuritySEV `xml:"-" json:"-" yaml:"-"`
}

type DomainLaunchSecuritySEV struct {
	CBitPos         *uint  `xml:"cbitpos" json:"cbitpos" yaml:"cbitpos"`
	ReducedPhysBits *uint  `xml:"reducedPhysBits" json:"reducedPhysBits" yaml:"reducedPhysBits"`
	Policy          *uint  `xml:"policy" json:"policy" yaml:"policy"`
	DHCert          string `xml:"dhCert" json:"dhCert" yaml:"dhCert"`
	Session         string `xml:"sesion" json:"sesion" yaml:"sesion"`
}

type DomainFeatureCapabilities struct {
	Policy         string                   `xml:"policy,attr,omitempty" json:"policy,omitempty" yaml:"policy,omitempty"`
	AuditControl   *DomainFeatureCapability `xml:"audit_control" json:"audit_control" yaml:"audit_control"`
	AuditWrite     *DomainFeatureCapability `xml:"audit_write" json:"audit_write" yaml:"audit_write"`
	BlockSuspend   *DomainFeatureCapability `xml:"block_suspend" json:"block_suspend" yaml:"block_suspend"`
	Chown          *DomainFeatureCapability `xml:"chown" json:"chown" yaml:"chown"`
	DACOverride    *DomainFeatureCapability `xml:"dac_override" json:"dac_override" yaml:"dac_override"`
	DACReadSearch  *DomainFeatureCapability `xml:"dac_read_Search" json:"dac_read_Search" yaml:"dac_read_Search"`
	FOwner         *DomainFeatureCapability `xml:"fowner" json:"fowner" yaml:"fowner"`
	FSetID         *DomainFeatureCapability `xml:"fsetid" json:"fsetid" yaml:"fsetid"`
	IPCLock        *DomainFeatureCapability `xml:"ipc_lock" json:"ipc_lock" yaml:"ipc_lock"`
	IPCOwner       *DomainFeatureCapability `xml:"ipc_owner" json:"ipc_owner" yaml:"ipc_owner"`
	Kill           *DomainFeatureCapability `xml:"kill" json:"kill" yaml:"kill"`
	Lease          *DomainFeatureCapability `xml:"lease" json:"lease" yaml:"lease"`
	LinuxImmutable *DomainFeatureCapability `xml:"linux_immutable" json:"linux_immutable" yaml:"linux_immutable"`
	MACAdmin       *DomainFeatureCapability `xml:"mac_admin" json:"mac_admin" yaml:"mac_admin"`
	MACOverride    *DomainFeatureCapability `xml:"mac_override" json:"mac_override" yaml:"mac_override"`
	MkNod          *DomainFeatureCapability `xml:"mknod" json:"mknod" yaml:"mknod"`
	NetAdmin       *DomainFeatureCapability `xml:"net_admin" json:"net_admin" yaml:"net_admin"`
	NetBindService *DomainFeatureCapability `xml:"net_bind_service" json:"net_bind_service" yaml:"net_bind_service"`
	NetBroadcast   *DomainFeatureCapability `xml:"net_broadcast" json:"net_broadcast" yaml:"net_broadcast"`
	NetRaw         *DomainFeatureCapability `xml:"net_raw" json:"net_raw" yaml:"net_raw"`
	SetGID         *DomainFeatureCapability `xml:"setgid" json:"setgid" yaml:"setgid"`
	SetFCap        *DomainFeatureCapability `xml:"setfcap" json:"setfcap" yaml:"setfcap"`
	SetPCap        *DomainFeatureCapability `xml:"setpcap" json:"setpcap" yaml:"setpcap"`
	SetUID         *DomainFeatureCapability `xml:"setuid" json:"setuid" yaml:"setuid"`
	SysAdmin       *DomainFeatureCapability `xml:"sys_admin" json:"sys_admin" yaml:"sys_admin"`
	SysBoot        *DomainFeatureCapability `xml:"sys_boot" json:"sys_boot" yaml:"sys_boot"`
	SysChRoot      *DomainFeatureCapability `xml:"sys_chroot" json:"sys_chroot" yaml:"sys_chroot"`
	SysModule      *DomainFeatureCapability `xml:"sys_module" json:"sys_module" yaml:"sys_module"`
	SysNice        *DomainFeatureCapability `xml:"sys_nice" json:"sys_nice" yaml:"sys_nice"`
	SysPAcct       *DomainFeatureCapability `xml:"sys_pacct" json:"sys_pacct" yaml:"sys_pacct"`
	SysPTrace      *DomainFeatureCapability `xml:"sys_ptrace" json:"sys_ptrace" yaml:"sys_ptrace"`
	SysRawIO       *DomainFeatureCapability `xml:"sys_rawio" json:"sys_rawio" yaml:"sys_rawio"`
	SysResource    *DomainFeatureCapability `xml:"sys_resource" json:"sys_resource" yaml:"sys_resource"`
	SysTime        *DomainFeatureCapability `xml:"sys_time" json:"sys_time" yaml:"sys_time"`
	SysTTYCnofig   *DomainFeatureCapability `xml:"sys_tty_config" json:"sys_tty_config" yaml:"sys_tty_config"`
	SysLog         *DomainFeatureCapability `xml:"syslog" json:"syslog" yaml:"syslog"`
	WakeAlarm      *DomainFeatureCapability `xml:"wake_alarm" json:"wake_alarm" yaml:"wake_alarm"`
}

type DomainFeatureList struct {
	PAE          *DomainFeature             `xml:"pae" json:"pae" yaml:"pae"`
	ACPI         *DomainFeature             `xml:"acpi" json:"acpi" yaml:"acpi"`
	APIC         *DomainFeatureAPIC         `xml:"apic" json:"apic" yaml:"apic"`
	HAP          *DomainFeatureState        `xml:"hap" json:"hap" yaml:"hap"`
	Viridian     *DomainFeature             `xml:"viridian" json:"viridian" yaml:"viridian"`
	PrivNet      *DomainFeature             `xml:"privnet" json:"privnet" yaml:"privnet"`
	HyperV       *DomainFeatureHyperV       `xml:"hyperv" json:"hyperv" yaml:"hyperv"`
	KVM          *DomainFeatureKVM          `xml:"kvm" json:"kvm" yaml:"kvm"`
	PVSpinlock   *DomainFeatureState        `xml:"pvspinlock" json:"pvspinlock" yaml:"pvspinlock"`
	PMU          *DomainFeatureState        `xml:"pmu" json:"pmu" yaml:"pmu"`
	VMPort       *DomainFeatureState        `xml:"vmport" json:"vmport" yaml:"vmport"`
	GIC          *DomainFeatureGIC          `xml:"gic" json:"gic" yaml:"gic"`
	SMM          *DomainFeatureSMM          `xml:"smm" json:"smm" yaml:"smm"`
	IOAPIC       *DomainFeatureIOAPIC       `xml:"ioapic" json:"ioapic" yaml:"ioapic"`
	HPT          *DomainFeatureHPT          `xml:"hpt" json:"hpt" yaml:"hpt"`
	HTM          *DomainFeatureState        `xml:"htm" json:"htm" yaml:"htm"`
	Capabilities *DomainFeatureCapabilities `xml:"capabilities" json:"capabilities" yaml:"capabilities"`
	VMCoreInfo   *DomainFeatureState        `xml:"vmcoreinfo" json:"vmcoreinfo" yaml:"vmcoreinfo"`
}

type DomainCPUTuneShares struct {
	Value uint `xml:",attr" json:"" yaml:""`
}

type DomainCPUTunePeriod struct {
	Value uint64 `xml:",attr" json:"" yaml:""`
}

type DomainCPUTuneQuota struct {
	Value int64 `xml:",attr" json:"" yaml:""`
}

type DomainCPUTuneVCPUPin struct {
	VCPU   uint   `xml:"vcpu,attr" json:"vcpu,omitempty" yaml:"vcpu,omitempty"`
	CPUSet string `xml:"cpuset,attr" json:"cpuset,omitempty" yaml:"cpuset,omitempty"`
}

type DomainCPUTuneEmulatorPin struct {
	CPUSet string `xml:"cpuset,attr" json:"cpuset,omitempty" yaml:"cpuset,omitempty"`
}

type DomainCPUTuneIOThreadPin struct {
	IOThread uint   `xml:"iothread,attr" json:"iothread,omitempty" yaml:"iothread,omitempty"`
	CPUSet   string `xml:"cpuset,attr" json:"cpuset,omitempty" yaml:"cpuset,omitempty"`
}

type DomainCPUTuneVCPUSched struct {
	VCPUs     string `xml:"vcpus,attr" json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty" yaml:"scheduler,omitempty"`
	Priority  *int   `xml:"priority,attr" json:"priority,omitempty" yaml:"priority,omitempty"`
}

type DomainCPUTuneIOThreadSched struct {
	IOThreads string `xml:"iothreads,attr" json:"iothreads,omitempty" yaml:"iothreads,omitempty"`
	Scheduler string `xml:"scheduler,attr,omitempty" json:"scheduler,omitempty" yaml:"scheduler,omitempty"`
	Priority  *int   `xml:"priority,attr" json:"priority,omitempty" yaml:"priority,omitempty"`
}

type DomainCPUCacheTune struct {
	VCPUs string                    `xml:"vcpus,attr,omitempty" json:"vcpus,omitempty" yaml:"vcpus,omitempty"`
	Cache []DomainCPUCacheTuneCache `xml:"cache" json:"cache" yaml:"cache"`
}

type DomainCPUCacheTuneCache struct {
	ID    uint   `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Level uint   `xml:"level,attr" json:"level,omitempty" yaml:"level,omitempty"`
	Type  string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Size  uint   `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Unit  string `xml:"unit,attr" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainCPUTune struct {
	Shares         *DomainCPUTuneShares         `xml:"shares" json:"shares" yaml:"shares"`
	Period         *DomainCPUTunePeriod         `xml:"period" json:"period" yaml:"period"`
	Quota          *DomainCPUTuneQuota          `xml:"quota" json:"quota" yaml:"quota"`
	GlobalPeriod   *DomainCPUTunePeriod         `xml:"global_period" json:"global_period" yaml:"global_period"`
	GlobalQuota    *DomainCPUTuneQuota          `xml:"global_quota" json:"global_quota" yaml:"global_quota"`
	EmulatorPeriod *DomainCPUTunePeriod         `xml:"emulator_period" json:"emulator_period" yaml:"emulator_period"`
	EmulatorQuota  *DomainCPUTuneQuota          `xml:"emulator_quota" json:"emulator_quota" yaml:"emulator_quota"`
	IOThreadPeriod *DomainCPUTunePeriod         `xml:"iothread_period" json:"iothread_period" yaml:"iothread_period"`
	IOThreadQuota  *DomainCPUTuneQuota          `xml:"iothread_quota" json:"iothread_quota" yaml:"iothread_quota"`
	VCPUPin        []DomainCPUTuneVCPUPin       `xml:"vcpupin" json:"vcpupin" yaml:"vcpupin"`
	EmulatorPin    *DomainCPUTuneEmulatorPin    `xml:"emulatorpin" json:"emulatorpin" yaml:"emulatorpin"`
	IOThreadPin    []DomainCPUTuneIOThreadPin   `xml:"iothreadpin" json:"iothreadpin" yaml:"iothreadpin"`
	VCPUSched      []DomainCPUTuneVCPUSched     `xml:"vcpusched" json:"vcpusched" yaml:"vcpusched"`
	IOThreadSched  []DomainCPUTuneIOThreadSched `xml:"iothreadsched" json:"iothreadsched" yaml:"iothreadsched"`
	CacheTune      []DomainCPUCacheTune         `xml:"cachetune" json:"cachetune" yaml:"cachetune"`
}

type DomainQEMUCommandlineArg struct {
	Value string `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainQEMUCommandlineEnv struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainQEMUCommandline struct {
	XMLName xml.Name                   `xml:"http://libvirt.org/schemas/domain/qemu/1.0 commandline" json:"http://libvirt.org/schemas/domain/qemu/1.0 commandline" yaml:"http://libvirt.org/schemas/domain/qemu/1.0 commandline"`
	Args    []DomainQEMUCommandlineArg `xml:"arg" json:"arg" yaml:"arg"`
	Envs    []DomainQEMUCommandlineEnv `xml:"env" json:"env" yaml:"env"`
}

type DomainLXCNamespace struct {
	XMLName  xml.Name               `xml:"http://libvirt.org/schemas/domain/lxc/1.0 namespace" json:"http://libvirt.org/schemas/domain/lxc/1.0 namespace" yaml:"http://libvirt.org/schemas/domain/lxc/1.0 namespace"`
	ShareNet *DomainLXCNamespaceMap `xml:"sharenet" json:"sharenet" yaml:"sharenet"`
	ShareIPC *DomainLXCNamespaceMap `xml:"shareipc" json:"shareipc" yaml:"shareipc"`
	ShareUTS *DomainLXCNamespaceMap `xml:"shareuts" json:"shareuts" yaml:"shareuts"`
}

type DomainLXCNamespaceMap struct {
	Type  string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Value string `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type DomainBlockIOTune struct {
	Weight uint                      `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Device []DomainBlockIOTuneDevice `xml:"device" json:"device" yaml:"device"`
}

type DomainBlockIOTuneDevice struct {
	Path          string `xml:"path" json:"path" yaml:"path"`
	Weight        uint   `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	ReadIopsSec   uint   `xml:"read_iops_sec,omitempty" json:"read_iops_sec,omitempty" yaml:"read_iops_sec,omitempty"`
	WriteIopsSec  uint   `xml:"write_iops_sec,omitempty" json:"write_iops_sec,omitempty" yaml:"write_iops_sec,omitempty"`
	ReadBytesSec  uint   `xml:"read_bytes_sec,omitempty" json:"read_bytes_sec,omitempty" yaml:"read_bytes_sec,omitempty"`
	WriteBytesSec uint   `xml:"write_bytes_sec,omitempty" json:"write_bytes_sec,omitempty" yaml:"write_bytes_sec,omitempty"`
}

type DomainPM struct {
	SuspendToMem  *DomainPMPolicy `xml:"suspend-to-mem" json:"suspend-to-mem" yaml:"suspend-to-mem"`
	SuspendToDisk *DomainPMPolicy `xml:"suspend-to-disk" json:"suspend-to-disk" yaml:"suspend-to-disk"`
}

type DomainPMPolicy struct {
	Enabled string `xml:"enabled,attr" json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type DomainSecLabel struct {
	Type       string `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Model      string `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Relabel    string `xml:"relabel,attr,omitempty" json:"relabel,omitempty" yaml:"relabel,omitempty"`
	Label      string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	ImageLabel string `xml:"imagelabel,omitempty" json:"imagelabel,omitempty" yaml:"imagelabel,omitempty"`
	BaseLabel  string `xml:"baselabel,omitempty" json:"baselabel,omitempty" yaml:"baselabel,omitempty"`
}

type DomainDeviceSecLabel struct {
	Model     string `xml:"model,attr,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	LabelSkip string `xml:"labelskip,attr,omitempty" json:"labelskip,omitempty" yaml:"labelskip,omitempty"`
	Relabel   string `xml:"relabel,attr,omitempty" json:"relabel,omitempty" yaml:"relabel,omitempty"`
	Label     string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
}

type DomainNUMATune struct {
	Memory   *DomainNUMATuneMemory   `xml:"memory" json:"memory" yaml:"memory"`
	MemNodes []DomainNUMATuneMemNode `xml:"memnode" json:"memnode" yaml:"memnode"`
}

type DomainNUMATuneMemory struct {
	Mode      string `xml:"mode,attr,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Nodeset   string `xml:"nodeset,attr,omitempty" json:"nodeset,omitempty" yaml:"nodeset,omitempty"`
	Placement string `xml:"placement,attr,omitempty" json:"placement,omitempty" yaml:"placement,omitempty"`
}

type DomainNUMATuneMemNode struct {
	CellID  uint   `xml:"cellid,attr" json:"cellid,omitempty" yaml:"cellid,omitempty"`
	Mode    string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
	Nodeset string `xml:"nodeset,attr" json:"nodeset,omitempty" yaml:"nodeset,omitempty"`
}

type DomainIOThreadIDs struct {
	IOThreads []DomainIOThread `xml:"iothread" json:"iothread" yaml:"iothread"`
}

type DomainIOThread struct {
	ID uint `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
}

type DomainKeyWrap struct {
	Ciphers []DomainKeyWrapCipher `xml:"cipher" json:"cipher" yaml:"cipher"`
}

type DomainKeyWrapCipher struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	State string `xml:"state,attr" json:"state,omitempty" yaml:"state,omitempty"`
}

type DomainIDMap struct {
	UIDs []DomainIDMapRange `xml:"uid" json:"uid" yaml:"uid"`
	GIDs []DomainIDMapRange `xml:"gid" json:"gid" yaml:"gid"`
}

type DomainIDMapRange struct {
	Start  uint `xml:"start,attr" json:"start,omitempty" yaml:"start,omitempty"`
	Target uint `xml:"target,attr" json:"target,omitempty" yaml:"target,omitempty"`
	Count  uint `xml:"count,attr" json:"count,omitempty" yaml:"count,omitempty"`
}

type DomainMemoryTuneLimit struct {
	Value uint64 `xml:",attr" json:"" yaml:""`
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
}

type DomainMemoryTune struct {
	HardLimit     *DomainMemoryTuneLimit `xml:"hard_limit" json:"hard_limit" yaml:"hard_limit"`
	SoftLimit     *DomainMemoryTuneLimit `xml:"soft_limit" json:"soft_limit" yaml:"soft_limit"`
	MinGuarantee  *DomainMemoryTuneLimit `xml:"min_guarantee" json:"min_guarantee" yaml:"min_guarantee"`
	SwapHardLimit *DomainMemoryTuneLimit `xml:"swap_hard_limit" json:"swap_hard_limit" yaml:"swap_hard_limit"`
}

type DomainMetadata struct {
	XML string `xml:",innerxml" json:",innerxml" yaml:",innerxml"`
}

type DomainVMWareDataCenterPath struct {
	XMLName xml.Name `xml:"http://libvirt.org/schemas/domain/vmware/1.0 datacenterpath" json:"http://libvirt.org/schemas/domain/vmware/1.0 datacenterpath" yaml:"http://libvirt.org/schemas/domain/vmware/1.0 datacenterpath"`
	Value   string   `xml:",attr" json:"" yaml:""`
}

type DomainPerf struct {
	Events []DomainPerfEvent `xml:"event" json:"event" yaml:"event"`
}

type DomainPerfEvent struct {
	Name    string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Enabled string `xml:"enabled,attr" json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type DomainGenID struct {
	Value string `xml:",attr" json:"" yaml:""`
}

// NB, try to keep the order of fields in this struct
// matching the order of XML elements that libvirt
// will generate when dumping XML.
type Domain struct {
	XMLName              xml.Name             `xml:"domain" json:"domain" yaml:"domain"`
	Type                 string               `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ID                   *int                 `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Name                 string               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	UUID                 string               `xml:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	GenID                *DomainGenID         `xml:"genid" json:"genid" yaml:"genid"`
	Title                string               `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Description          string               `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Metadata             *DomainMetadata      `xml:"metadata" json:"metadata" yaml:"metadata"`
	MaximumMemory        *DomainMaxMemory     `xml:"maxMemory" json:"maxMemory" yaml:"maxMemory"`
	Memory               *DomainMemory        `xml:"memory" json:"memory" yaml:"memory"`
	CurrentMemory        *DomainCurrentMemory `xml:"currentMemory" json:"currentMemory" yaml:"currentMemory"`
	BlockIOTune          *DomainBlockIOTune   `xml:"blkiotune" json:"blkiotune" yaml:"blkiotune"`
	MemoryTune           *DomainMemoryTune    `xml:"memtune" json:"memtune" yaml:"memtune"`
	MemoryBacking        *DomainMemoryBacking `xml:"memoryBacking" json:"memoryBacking" yaml:"memoryBacking"`
	VCPU                 *DomainVCPU          `xml:"vcpu" json:"vcpu" yaml:"vcpu"`
	VCPUs                *DomainVCPUs         `xml:"vcpus" json:"vcpus" yaml:"vcpus"`
	IOThreads            uint                 `xml:"iothreads,omitempty" json:"iothreads,omitempty" yaml:"iothreads,omitempty"`
	IOThreadIDs          *DomainIOThreadIDs   `xml:"iothreadids" json:"iothreadids" yaml:"iothreadids"`
	CPUTune              *DomainCPUTune       `xml:"cputune" json:"cputune" yaml:"cputune"`
	NUMATune             *DomainNUMATune      `xml:"numatune" json:"numatune" yaml:"numatune"`
	Resource             *DomainResource      `xml:"resource" json:"resource" yaml:"resource"`
	SysInfo              *DomainSysInfo       `xml:"sysinfo" json:"sysinfo" yaml:"sysinfo"`
	Bootloader           string               `xml:"bootloader,omitempty" json:"bootloader,omitempty" yaml:"bootloader,omitempty"`
	BootloaderArgs       string               `xml:"bootloader_args,omitempty" json:"bootloader_args,omitempty" yaml:"bootloader_args,omitempty"`
	OS                   *DomainOS            `xml:"os" json:"os" yaml:"os"`
	IDMap                *DomainIDMap         `xml:"idmap" json:"idmap" yaml:"idmap"`
	Features             *DomainFeatureList   `xml:"features" json:"features" yaml:"features"`
	CPU                  *DomainCPU           `xml:"cpu" json:"cpu" yaml:"cpu"`
	Clock                *DomainClock         `xml:"clock" json:"clock" yaml:"clock"`
	OnPoweroff           string               `xml:"on_poweroff,omitempty" json:"on_poweroff,omitempty" yaml:"on_poweroff,omitempty"`
	OnReboot             string               `xml:"on_reboot,omitempty" json:"on_reboot,omitempty" yaml:"on_reboot,omitempty"`
	OnCrash              string               `xml:"on_crash,omitempty" json:"on_crash,omitempty" yaml:"on_crash,omitempty"`
	PM                   *DomainPM            `xml:"pm" json:"pm" yaml:"pm"`
	Perf                 *DomainPerf          `xml:"perf" json:"perf" yaml:"perf"`
	Devices              *DomainDeviceList    `xml:"devices" json:"devices" yaml:"devices"`
	SecLabel             []DomainSecLabel     `xml:"seclabel" json:"seclabel" yaml:"seclabel"`
	QEMUCommandline      *DomainQEMUCommandline
	LXCNamespace         *DomainLXCNamespace
	VMWareDataCenterPath *DomainVMWareDataCenterPath
	KeyWrap              *DomainKeyWrap        `xml:"keywrap" json:"keywrap" yaml:"keywrap"`
	LaunchSecurity       *DomainLaunchSecurity `xml:"launchSecurity" json:"launchSecurity" yaml:"launchSecurity"`
}

func (d *Domain) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *Domain) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *Domain) MarshalYAML() (string, error) {
	doc, err := yaml.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainController DomainController

type domainControllerPCI struct {
	DomainControllerPCI
	domainController
}

type domainControllerUSB struct {
	DomainControllerUSB
	domainController
}

type domainControllerVirtIOSerial struct {
	DomainControllerVirtIOSerial
	domainController
}

func (a *DomainControllerPCITarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "chassisNr", a.ChassisNr, "%d")
	marshalUintAttr(&start, "chassis", a.Chassis, "%d")
	marshalUintAttr(&start, "port", a.Port, "%d")
	marshalUintAttr(&start, "busNr", a.BusNr, "%d")
	marshalUintAttr(&start, "index", a.Index, "%d")
	e.EncodeToken(start)
	if a.NUMANode != nil {
		node := xml.StartElement{
			Name: xml.Name{Local: "node"},
		}
		e.EncodeToken(node)
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.NUMANode)))
		e.EncodeToken(node.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainControllerPCITarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "chassisNr" {
			if err := unmarshalUintAttr(attr.Value, &a.ChassisNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "chassis" {
			if err := unmarshalUintAttr(attr.Value, &a.Chassis, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshalUintAttr(attr.Value, &a.Port, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "busNr" {
			if err := unmarshalUintAttr(attr.Value, &a.BusNr, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "index" {
			if err := unmarshalUintAttr(attr.Value, &a.Index, 10); err != nil {
				return err
			}
		}
	}
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "node" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					val, err := strconv.ParseUint(string(data), 10, 64)
					if err != nil {
						return err
					}
					vali := uint(val)
					a.NUMANode = &vali
				}
			}
		}
	}
	return nil
}

func (a *DomainController) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "controller"
	if a.Type == "pci" {
		pci := domainControllerPCI{}
		pci.domainController = domainController(*a)
		if a.PCI != nil {
			pci.DomainControllerPCI = *a.PCI
		}
		return e.EncodeElement(pci, start)
	} else if a.Type == "usb" {
		usb := domainControllerUSB{}
		usb.domainController = domainController(*a)
		if a.USB != nil {
			usb.DomainControllerUSB = *a.USB
		}
		return e.EncodeElement(usb, start)
	} else if a.Type == "virtio-serial" {
		vioserial := domainControllerVirtIOSerial{}
		vioserial.domainController = domainController(*a)
		if a.VirtIOSerial != nil {
			vioserial.DomainControllerVirtIOSerial = *a.VirtIOSerial
		}
		return e.EncodeElement(vioserial, start)
	} else {
		gen := domainController(*a)
		return e.EncodeElement(gen, start)
	}
}

func getAttr(attrs []xml.Attr, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value, true
		}
	}
	return "", false
}

func (a *DomainController) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if typ == "pci" {
		var pci domainControllerPCI
		err := d.DecodeElement(&pci, &start)
		if err != nil {
			return err
		}
		*a = DomainController(pci.domainController)
		a.PCI = &pci.DomainControllerPCI
		return nil
	} else if typ == "usb" {
		var usb domainControllerUSB
		err := d.DecodeElement(&usb, &start)
		if err != nil {
			return err
		}
		*a = DomainController(usb.domainController)
		a.USB = &usb.DomainControllerUSB
		return nil
	} else if typ == "virtio-serial" {
		var vioserial domainControllerVirtIOSerial
		err := d.DecodeElement(&vioserial, &start)
		if err != nil {
			return err
		}
		*a = DomainController(vioserial.domainController)
		a.VirtIOSerial = &vioserial.DomainControllerVirtIOSerial
		return nil
	} else {
		var gen domainController
		err := d.DecodeElement(&gen, &start)
		if err != nil {
			return err
		}
		*a = DomainController(gen)
		return nil
	}
}

func (d *DomainGraphic) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainGraphic) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainController) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainController) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainDiskReservationsSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "source"
	src := DomainChardevSource(*a)
	typ := getChardevSourceType(&src)
	if typ != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, typ,
		})
	}
	return e.EncodeElement(&src, start)
}

func (a *DomainDiskReservationsSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "unix"
	}
	src := createChardevSource(typ)
	err := d.DecodeElement(&src, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskReservationsSource(*src)
	return nil
}

type domainDiskSource DomainDiskSource

type domainDiskSourceFile struct {
	DomainDiskSourceFile
	domainDiskSource
}

type domainDiskSourceBlock struct {
	DomainDiskSourceBlock
	domainDiskSource
}

type domainDiskSourceDir struct {
	DomainDiskSourceDir
	domainDiskSource
}

type domainDiskSourceNetwork struct {
	DomainDiskSourceNetwork
	domainDiskSource
}

type domainDiskSourceVolume struct {
	DomainDiskSourceVolume
	domainDiskSource
}

func (a *DomainDiskSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.File != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.File.File == "" {
			return nil
		}
		file := domainDiskSourceFile{
			*a.File, domainDiskSource(*a),
		}
		return e.EncodeElement(&file, start)
	} else if a.Block != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.Block.Dev == "" {
			return nil
		}
		block := domainDiskSourceBlock{
			*a.Block, domainDiskSource(*a),
		}
		return e.EncodeElement(&block, start)
	} else if a.Dir != nil {
		dir := domainDiskSourceDir{
			*a.Dir, domainDiskSource(*a),
		}
		return e.EncodeElement(&dir, start)
	} else if a.Network != nil {
		network := domainDiskSourceNetwork{
			*a.Network, domainDiskSource(*a),
		}
		return e.EncodeElement(&network, start)
	} else if a.Volume != nil {
		if a.StartupPolicy == "" && a.Encryption == nil && a.Volume.Pool == "" && a.Volume.Volume == "" {
			return nil
		}
		volume := domainDiskSourceVolume{
			*a.Volume, domainDiskSource(*a),
		}
		return e.EncodeElement(&volume, start)
	}
	return nil
}

func (a *DomainDiskSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.File != nil {
		file := domainDiskSourceFile{
			*a.File, domainDiskSource(*a),
		}
		err := d.DecodeElement(&file, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(file.domainDiskSource)
		a.File = &file.DomainDiskSourceFile
	} else if a.Block != nil {
		block := domainDiskSourceBlock{
			*a.Block, domainDiskSource(*a),
		}
		err := d.DecodeElement(&block, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(block.domainDiskSource)
		a.Block = &block.DomainDiskSourceBlock
	} else if a.Dir != nil {
		dir := domainDiskSourceDir{
			*a.Dir, domainDiskSource(*a),
		}
		err := d.DecodeElement(&dir, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(dir.domainDiskSource)
		a.Dir = &dir.DomainDiskSourceDir
	} else if a.Network != nil {
		network := domainDiskSourceNetwork{
			*a.Network, domainDiskSource(*a),
		}
		err := d.DecodeElement(&network, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(network.domainDiskSource)
		a.Network = &network.DomainDiskSourceNetwork
	} else if a.Volume != nil {
		volume := domainDiskSourceVolume{
			*a.Volume, domainDiskSource(*a),
		}
		err := d.DecodeElement(&volume, &start)
		if err != nil {
			return err
		}
		*a = DomainDiskSource(volume.domainDiskSource)
		a.Volume = &volume.DomainDiskSourceVolume
	}
	return nil
}

type domainDiskBackingStore DomainDiskBackingStore

func (a *DomainDiskBackingStore) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backingStore"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	disk := domainDiskBackingStore(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDiskBackingStore) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	}
	disk := domainDiskBackingStore(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskBackingStore(disk)
	if !ok && a.Source.File.File == "" {
		a.Source.File = nil
	}
	return nil
}

type domainDiskMirror DomainDiskMirror

func (a *DomainDiskMirror) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "mirror"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
			if a.Source.File.File != "" {
				start.Attr = append(start.Attr, xml.Attr{
					xml.Name{Local: "file"}, a.Source.File.File,
				})
			}
			if a.Format != nil && a.Format.Type != "" {
				start.Attr = append(start.Attr, xml.Attr{
					xml.Name{Local: "format"}, a.Format.Type,
				})
			}
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	disk := domainDiskMirror(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDiskMirror) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	}
	disk := domainDiskMirror(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDiskMirror(disk)
	if !ok {
		if a.Source.File.File == "" {
			file, ok := getAttr(start.Attr, "file")
			if ok {
				a.Source.File.File = file
			} else {
				a.Source.File = nil
			}
		}
		if a.Format == nil {
			format, ok := getAttr(start.Attr, "format")
			if ok {
				a.Format = &DomainDiskFormat{
					Type: format,
				}
			}
		}
	}
	return nil
}

type domainDisk DomainDisk

func (a *DomainDisk) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "disk"
	if a.Source != nil {
		if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.Dir != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "dir",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	disk := domainDisk(*a)
	return e.EncodeElement(disk, start)
}

func (a *DomainDisk) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "file"
	}
	a.Source = &DomainDiskSource{}
	if typ == "file" {
		a.Source.File = &DomainDiskSourceFile{}
	} else if typ == "block" {
		a.Source.Block = &DomainDiskSourceBlock{}
	} else if typ == "network" {
		a.Source.Network = &DomainDiskSourceNetwork{}
	} else if typ == "dir" {
		a.Source.Dir = &DomainDiskSourceDir{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainDiskSourceVolume{}
	}
	disk := domainDisk(*a)
	err := d.DecodeElement(&disk, &start)
	if err != nil {
		return err
	}
	*a = DomainDisk(disk)
	return nil
}

func (d *DomainDisk) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainDisk) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainFilesystemSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Mount != nil {
		return e.EncodeElement(a.Mount, start)
	} else if a.Block != nil {
		return e.EncodeElement(a.Block, start)
	} else if a.File != nil {
		return e.EncodeElement(a.File, start)
	} else if a.Template != nil {
		return e.EncodeElement(a.Template, start)
	} else if a.RAM != nil {
		return e.EncodeElement(a.RAM, start)
	} else if a.Bind != nil {
		return e.EncodeElement(a.Bind, start)
	} else if a.Volume != nil {
		return e.EncodeElement(a.Volume, start)
	}
	return nil
}

func (a *DomainFilesystemSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Mount != nil {
		return d.DecodeElement(a.Mount, &start)
	} else if a.Block != nil {
		return d.DecodeElement(a.Block, &start)
	} else if a.File != nil {
		return d.DecodeElement(a.File, &start)
	} else if a.Template != nil {
		return d.DecodeElement(a.Template, &start)
	} else if a.RAM != nil {
		return d.DecodeElement(a.RAM, &start)
	} else if a.Bind != nil {
		return d.DecodeElement(a.Bind, &start)
	} else if a.Volume != nil {
		return d.DecodeElement(a.Volume, &start)
	}
	return nil
}

type domainFilesystem DomainFilesystem

func (a *DomainFilesystem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "filesystem"
	if a.Source != nil {
		if a.Source.Mount != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "mount",
			})
		} else if a.Source.Block != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "block",
			})
		} else if a.Source.File != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "file",
			})
		} else if a.Source.Template != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "template",
			})
		} else if a.Source.RAM != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "ram",
			})
		} else if a.Source.Bind != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "bind",
			})
		} else if a.Source.Volume != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "volume",
			})
		}
	}
	fs := domainFilesystem(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainFilesystem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "mount"
	}
	a.Source = &DomainFilesystemSource{}
	if typ == "mount" {
		a.Source.Mount = &DomainFilesystemSourceMount{}
	} else if typ == "block" {
		a.Source.Block = &DomainFilesystemSourceBlock{}
	} else if typ == "file" {
		a.Source.File = &DomainFilesystemSourceFile{}
	} else if typ == "template" {
		a.Source.Template = &DomainFilesystemSourceTemplate{}
	} else if typ == "ram" {
		a.Source.RAM = &DomainFilesystemSourceRAM{}
	} else if typ == "bind" {
		a.Source.Bind = &DomainFilesystemSourceBind{}
	} else if typ == "volume" {
		a.Source.Volume = &DomainFilesystemSourceVolume{}
	}
	fs := domainFilesystem(*a)
	err := d.DecodeElement(&fs, &start)
	if err != nil {
		return err
	}
	*a = DomainFilesystem(fs)
	return nil
}

func (d *DomainFilesystem) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainFilesystem) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainInterfaceVirtualPortParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "parameters"
	if a.Any != nil {
		return e.EncodeElement(a.Any, start)
	} else if a.VEPA8021QBG != nil {
		return e.EncodeElement(a.VEPA8021QBG, start)
	} else if a.VNTag8011QBH != nil {
		return e.EncodeElement(a.VNTag8011QBH, start)
	} else if a.OpenVSwitch != nil {
		return e.EncodeElement(a.OpenVSwitch, start)
	} else if a.MidoNet != nil {
		return e.EncodeElement(a.MidoNet, start)
	}
	return nil
}

func (a *DomainInterfaceVirtualPortParams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Any != nil {
		return d.DecodeElement(a.Any, &start)
	} else if a.VEPA8021QBG != nil {
		return d.DecodeElement(a.VEPA8021QBG, &start)
	} else if a.VNTag8011QBH != nil {
		return d.DecodeElement(a.VNTag8011QBH, &start)
	} else if a.OpenVSwitch != nil {
		return d.DecodeElement(a.OpenVSwitch, &start)
	} else if a.MidoNet != nil {
		return d.DecodeElement(a.MidoNet, &start)
	}
	return nil
}

type domainInterfaceVirtualPort DomainInterfaceVirtualPort

func (a *DomainInterfaceVirtualPort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "virtualport"
	if a.Params != nil {
		if a.Params.Any != nil {
			/* no type attr wanted */
		} else if a.Params.VEPA8021QBG != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbg",
			})
		} else if a.Params.VNTag8011QBH != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "802.1Qbh",
			})
		} else if a.Params.OpenVSwitch != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "openvswitch",
			})
		} else if a.Params.MidoNet != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "midonet",
			})
		}
	}
	vp := domainInterfaceVirtualPort(*a)
	return e.EncodeElement(&vp, start)
}

func (a *DomainInterfaceVirtualPort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	a.Params = &DomainInterfaceVirtualPortParams{}
	if !ok {
		var any DomainInterfaceVirtualPortParamsAny
		a.Params.Any = &any
	} else if typ == "802.1Qbg" {
		var vepa DomainInterfaceVirtualPortParamsVEPA8021QBG
		a.Params.VEPA8021QBG = &vepa
	} else if typ == "802.1Qbh" {
		var vntag DomainInterfaceVirtualPortParamsVNTag8021QBH
		a.Params.VNTag8011QBH = &vntag
	} else if typ == "openvswitch" {
		var ovs DomainInterfaceVirtualPortParamsOpenVSwitch
		a.Params.OpenVSwitch = &ovs
	} else if typ == "midonet" {
		var mido DomainInterfaceVirtualPortParamsMidoNet
		a.Params.MidoNet = &mido
	}

	vp := domainInterfaceVirtualPort(*a)
	err := d.DecodeElement(&vp, &start)
	if err != nil {
		return err
	}
	*a = DomainInterfaceVirtualPort(vp)
	return nil
}

func (a *DomainInterfaceSourceHostdev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if a.PCI != nil {
		addr := xml.StartElement{
			Name: xml.Name{Local: "address"},
		}
		addr.Attr = append(addr.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		e.EncodeElement(a.PCI.Address, addr)
	} else if a.USB != nil {
		addr := xml.StartElement{
			Name: xml.Name{Local: "address"},
		}
		addr.Attr = append(addr.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		e.EncodeElement(a.USB.Address, addr)
	}
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainInterfaceSourceHostdev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "address" {
				typ, ok := getAttr(tok.Attr, "type")
				if !ok {
					return fmt.Errorf("Missing hostdev address type attribute")
				}

				if typ == "pci" {
					a.PCI = &DomainHostdevSubsysPCISource{
						&DomainAddressPCI{},
					}
					err := d.DecodeElement(a.PCI.Address, &tok)
					if err != nil {
						return err
					}
				} else if typ == "usb" {
					a.USB = &DomainHostdevSubsysUSBSource{
						&DomainAddressUSB{},
					}
					err := d.DecodeElement(a.USB, &tok)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainInterfaceSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.User != nil {
		/* We don't want an empty <source></source> for User mode */
		//return e.EncodeElement(a.User, start)
		return nil
	} else if a.Ethernet != nil {
		if len(a.Ethernet.IP) > 0 && len(a.Ethernet.Route) > 0 {
			return e.EncodeElement(a.Ethernet, start)
		}
		return nil
	} else if a.VHostUser != nil {
		typ := getChardevSourceType(a.VHostUser)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
		return e.EncodeElement(a.VHostUser, start)
	} else if a.Server != nil {
		return e.EncodeElement(a.Server, start)
	} else if a.Client != nil {
		return e.EncodeElement(a.Client, start)
	} else if a.MCast != nil {
		return e.EncodeElement(a.MCast, start)
	} else if a.Network != nil {
		return e.EncodeElement(a.Network, start)
	} else if a.Bridge != nil {
		return e.EncodeElement(a.Bridge, start)
	} else if a.Internal != nil {
		return e.EncodeElement(a.Internal, start)
	} else if a.Direct != nil {
		return e.EncodeElement(a.Direct, start)
	} else if a.Hostdev != nil {
		return e.EncodeElement(a.Hostdev, start)
	} else if a.UDP != nil {
		return e.EncodeElement(a.UDP, start)
	}
	return nil
}

func (a *DomainInterfaceSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.User != nil {
		return d.DecodeElement(a.User, &start)
	} else if a.Ethernet != nil {
		return d.DecodeElement(a.Ethernet, &start)
	} else if a.VHostUser != nil {
		typ, ok := getAttr(start.Attr, "type")
		if !ok {
			typ = "pty"
		}
		a.VHostUser = createChardevSource(typ)
		return d.DecodeElement(a.VHostUser, &start)
	} else if a.Server != nil {
		return d.DecodeElement(a.Server, &start)
	} else if a.Client != nil {
		return d.DecodeElement(a.Client, &start)
	} else if a.MCast != nil {
		return d.DecodeElement(a.MCast, &start)
	} else if a.Network != nil {
		return d.DecodeElement(a.Network, &start)
	} else if a.Bridge != nil {
		return d.DecodeElement(a.Bridge, &start)
	} else if a.Internal != nil {
		return d.DecodeElement(a.Internal, &start)
	} else if a.Direct != nil {
		return d.DecodeElement(a.Direct, &start)
	} else if a.Hostdev != nil {
		return d.DecodeElement(a.Hostdev, &start)
	} else if a.UDP != nil {
		return d.DecodeElement(a.UDP, &start)
	}
	return nil
}

type domainInterface DomainInterface

func (a *DomainInterface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "interface"
	if a.Source != nil {
		if a.Source.User != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "user",
			})
		} else if a.Source.Ethernet != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "ethernet",
			})
		} else if a.Source.VHostUser != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "vhostuser",
			})
		} else if a.Source.Server != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "server",
			})
		} else if a.Source.Client != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "client",
			})
		} else if a.Source.MCast != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "mcast",
			})
		} else if a.Source.Network != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "network",
			})
		} else if a.Source.Bridge != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "bridge",
			})
		} else if a.Source.Internal != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "internal",
			})
		} else if a.Source.Direct != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "direct",
			})
		} else if a.Source.Hostdev != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "hostdev",
			})
		} else if a.Source.UDP != nil {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, "udp",
			})
		}
	}
	fs := domainInterface(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainInterface) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing interface type attribute")
	}
	a.Source = &DomainInterfaceSource{}
	if typ == "user" {
		a.Source.User = &DomainInterfaceSourceUser{}
	} else if typ == "ethernet" {
		a.Source.Ethernet = &DomainInterfaceSourceEthernet{}
	} else if typ == "vhostuser" {
		a.Source.VHostUser = &DomainChardevSource{}
	} else if typ == "server" {
		a.Source.Server = &DomainInterfaceSourceServer{}
	} else if typ == "client" {
		a.Source.Client = &DomainInterfaceSourceClient{}
	} else if typ == "mcast" {
		a.Source.MCast = &DomainInterfaceSourceMCast{}
	} else if typ == "network" {
		a.Source.Network = &DomainInterfaceSourceNetwork{}
	} else if typ == "bridge" {
		a.Source.Bridge = &DomainInterfaceSourceBridge{}
	} else if typ == "internal" {
		a.Source.Internal = &DomainInterfaceSourceInternal{}
	} else if typ == "direct" {
		a.Source.Direct = &DomainInterfaceSourceDirect{}
	} else if typ == "hostdev" {
		a.Source.Hostdev = &DomainInterfaceSourceHostdev{}
	} else if typ == "udp" {
		a.Source.UDP = &DomainInterfaceSourceUDP{}
	}
	fs := domainInterface(*a)
	err := d.DecodeElement(&fs, &start)
	if err != nil {
		return err
	}
	*a = DomainInterface(fs)
	return nil
}

func (d *DomainInterface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInterface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainSmartcard DomainSmartcard

func (a *DomainSmartcard) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "smartcard"
	if a.Passthrough != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "passthrough",
		})
		typ := getChardevSourceType(a.Passthrough)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	} else if a.Host != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "host",
		})
	} else if len(a.HostCerts) != 0 {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "host-certificates",
		})
	}
	smartcard := domainSmartcard(*a)
	return e.EncodeElement(smartcard, start)
}

func (a *DomainSmartcard) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mode, ok := getAttr(start.Attr, "mode")
	if !ok {
		return fmt.Errorf("Missing mode on smartcard device")
	}
	if mode == "host" {
		a.Host = &DomainSmartcardHost{}
	} else if mode == "passthrough" {
		typ, ok := getAttr(start.Attr, "type")
		if !ok {
			typ = "pty"
		}
		a.Passthrough = createChardevSource(typ)
	}
	smartcard := domainSmartcard(*a)
	err := d.DecodeElement(&smartcard, &start)
	if err != nil {
		return err
	}
	*a = DomainSmartcard(smartcard)
	return nil
}

func (d *DomainSmartcard) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSmartcard) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainTPMBackend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backend"
	if a.Passthrough != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "passthrough",
		})
		err := e.EncodeElement(a.Passthrough, start)
		if err != nil {
			return err
		}
	} else if a.Emulator != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "emulator",
		})
		err := e.EncodeElement(a.Emulator, start)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *DomainTPMBackend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing TPM backend type")
	}
	if typ == "passthrough" {
		a.Passthrough = &DomainTPMBackendPassthrough{}
		err := d.DecodeElement(a.Passthrough, &start)
		if err != nil {
			return err
		}
	} else if typ == "emulator" {
		a.Emulator = &DomainTPMBackendEmulator{}
		err := d.DecodeElement(a.Emulator, &start)
		if err != nil {
			return err
		}
	} else {
		d.Skip()
	}
	return nil
}

func (d *DomainTPM) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainTPM) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainShmem) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainShmem) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func getChardevSourceType(s *DomainChardevSource) string {
	if s.Null != nil {
		return "null"
	} else if s.VC != nil {
		return "vc"
	} else if s.Pty != nil {
		return "pty"
	} else if s.Dev != nil {
		return "dev"
	} else if s.File != nil {
		return "file"
	} else if s.Pipe != nil {
		return "pipe"
	} else if s.StdIO != nil {
		return "stdio"
	} else if s.UDP != nil {
		return "udp"
	} else if s.TCP != nil {
		return "tcp"
	} else if s.UNIX != nil {
		return "unix"
	} else if s.SpiceVMC != nil {
		return "spicevmc"
	} else if s.SpicePort != nil {
		return "spiceport"
	} else if s.NMDM != nil {
		return "nmdm"
	}
	return ""
}

func createChardevSource(typ string) *DomainChardevSource {
	switch typ {
	case "null":
		return &DomainChardevSource{
			Null: &DomainChardevSourceNull{},
		}
	case "vc":
		return &DomainChardevSource{
			VC: &DomainChardevSourceVC{},
		}
	case "pty":
		return &DomainChardevSource{
			Pty: &DomainChardevSourcePty{},
		}
	case "dev":
		return &DomainChardevSource{
			Dev: &DomainChardevSourceDev{},
		}
	case "file":
		return &DomainChardevSource{
			File: &DomainChardevSourceFile{},
		}
	case "pipe":
		return &DomainChardevSource{
			Pipe: &DomainChardevSourcePipe{},
		}
	case "stdio":
		return &DomainChardevSource{
			StdIO: &DomainChardevSourceStdIO{},
		}
	case "udp":
		return &DomainChardevSource{
			UDP: &DomainChardevSourceUDP{},
		}
	case "tcp":
		return &DomainChardevSource{
			TCP: &DomainChardevSourceTCP{},
		}
	case "unix":
		return &DomainChardevSource{
			UNIX: &DomainChardevSourceUNIX{},
		}
	case "spicevmc":
		return &DomainChardevSource{
			SpiceVMC: &DomainChardevSourceSpiceVMC{},
		}
	case "spiceport":
		return &DomainChardevSource{
			SpicePort: &DomainChardevSourceSpicePort{},
		}
	case "nmdm":
		return &DomainChardevSource{
			NMDM: &DomainChardevSourceNMDM{},
		}
	}

	return nil
}

type domainChardevSourceUDPFlat struct {
	Mode    string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
	Host    string `xml:"host,attr,omitempty" json:"host,omitempty" yaml:"host,omitempty"`
	Service string `xml:"service,attr,omitempty" json:"service,omitempty" yaml:"service,omitempty"`
}

func (a *DomainChardevSource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Null != nil {
		return nil
	} else if a.VC != nil {
		return nil
	} else if a.Pty != nil {
		if a.Pty.Path != "" {
			return e.EncodeElement(a.Pty, start)
		}
		return nil
	} else if a.Dev != nil {
		return e.EncodeElement(a.Dev, start)
	} else if a.File != nil {
		return e.EncodeElement(a.File, start)
	} else if a.Pipe != nil {
		return e.EncodeElement(a.Pipe, start)
	} else if a.StdIO != nil {
		return nil
	} else if a.UDP != nil {
		srcs := []domainChardevSourceUDPFlat{
			domainChardevSourceUDPFlat{
				Mode:    "bind",
				Host:    a.UDP.BindHost,
				Service: a.UDP.BindService,
			},
			domainChardevSourceUDPFlat{
				Mode:    "connect",
				Host:    a.UDP.ConnectHost,
				Service: a.UDP.ConnectService,
			},
		}
		if srcs[0].Host != "" || srcs[0].Service != "" {
			err := e.EncodeElement(&srcs[0], start)
			if err != nil {
				return err
			}
		}
		if srcs[1].Host != "" || srcs[1].Service != "" {
			err := e.EncodeElement(&srcs[1], start)
			if err != nil {
				return err
			}
		}
	} else if a.TCP != nil {
		return e.EncodeElement(a.TCP, start)
	} else if a.UNIX != nil {
		if a.UNIX.Path == "" {
			return nil
		}
		return e.EncodeElement(a.UNIX, start)
	} else if a.SpiceVMC != nil {
		return nil
	} else if a.SpicePort != nil {
		return e.EncodeElement(a.SpicePort, start)
	} else if a.NMDM != nil {
		return e.EncodeElement(a.NMDM, start)
	}
	return nil
}

func (a *DomainChardevSource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if a.Null != nil {
		d.Skip()
		return nil
	} else if a.VC != nil {
		d.Skip()
		return nil
	} else if a.Pty != nil {
		return d.DecodeElement(a.Pty, &start)
	} else if a.Dev != nil {
		return d.DecodeElement(a.Dev, &start)
	} else if a.File != nil {
		return d.DecodeElement(a.File, &start)
	} else if a.Pipe != nil {
		return d.DecodeElement(a.Pipe, &start)
	} else if a.StdIO != nil {
		d.Skip()
		return nil
	} else if a.UDP != nil {
		src := domainChardevSourceUDPFlat{}
		err := d.DecodeElement(&src, &start)
		if src.Mode == "connect" {
			a.UDP.ConnectHost = src.Host
			a.UDP.ConnectService = src.Service
		} else {
			a.UDP.BindHost = src.Host
			a.UDP.BindService = src.Service
		}
		return err
	} else if a.TCP != nil {
		return d.DecodeElement(a.TCP, &start)
	} else if a.UNIX != nil {
		return d.DecodeElement(a.UNIX, &start)
	} else if a.SpiceVMC != nil {
		d.Skip()
		return nil
	} else if a.SpicePort != nil {
		return d.DecodeElement(a.SpicePort, &start)
	} else if a.NMDM != nil {
		return d.DecodeElement(a.NMDM, &start)
	}
	return nil
}

type domainConsole DomainConsole

func (a *DomainConsole) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "console"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainConsole(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainConsole) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainConsole(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainConsole(con)
	return nil
}

func (d *DomainConsole) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainConsole) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainSerial DomainSerial

func (a *DomainSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "serial"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	s := domainSerial(*a)
	return e.EncodeElement(s, start)
}

func (a *DomainSerial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainSerial(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainSerial(con)
	return nil
}

func (d *DomainSerial) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSerial) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainParallel DomainParallel

func (a *DomainParallel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "parallel"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	s := domainParallel(*a)
	return e.EncodeElement(s, start)
}

func (a *DomainParallel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainParallel(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainParallel(con)
	return nil
}

func (d *DomainParallel) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainParallel) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainInput) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainInput) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainVideo) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainVideo) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainChannelTarget DomainChannelTarget

func (a *DomainChannelTarget) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.VirtIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio",
		})
		return e.EncodeElement(a.VirtIO, start)
	} else if a.Xen != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "xen",
		})
		return e.EncodeElement(a.Xen, start)
	} else if a.GuestFWD != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "guestfwd",
		})
		return e.EncodeElement(a.GuestFWD, start)
	}
	return nil
}

func (a *DomainChannelTarget) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing channel target type")
	}
	if typ == "virtio" {
		a.VirtIO = &DomainChannelTargetVirtIO{}
		return d.DecodeElement(a.VirtIO, &start)
	} else if typ == "xen" {
		a.Xen = &DomainChannelTargetXen{}
		return d.DecodeElement(a.Xen, &start)
	} else if typ == "guestfwd" {
		a.GuestFWD = &DomainChannelTargetGuestFWD{}
		return d.DecodeElement(a.GuestFWD, &start)
	}
	d.Skip()
	return nil
}

type domainChannel DomainChannel

func (a *DomainChannel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "channel"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainChannel(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainChannel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainChannel(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainChannel(con)
	return nil
}

func (d *DomainChannel) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainChannel) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainRedirFilterUSB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "class", a.Class, "0x%02x")
	marshalUintAttr(&start, "vendor", a.Vendor, "0x%04x")
	marshalUintAttr(&start, "product", a.Product, "0x%04x")
	if a.Version != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "version"}, a.Version,
		})
	}
	start.Attr = append(start.Attr, xml.Attr{
		xml.Name{Local: "allow"}, a.Allow,
	})
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainRedirFilterUSB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "class" && attr.Value != "-1" {
			if err := unmarshalUintAttr(attr.Value, &a.Class, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "product" && attr.Value != "-1" {
			if err := unmarshalUintAttr(attr.Value, &a.Product, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "vendor" && attr.Value != "-1" {
			if err := unmarshalUintAttr(attr.Value, &a.Vendor, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "version" && attr.Value != "-1" {
			a.Version = attr.Value
		} else if attr.Name.Local == "allow" {
			a.Allow = attr.Value
		}
	}
	d.Skip()
	return nil
}

type domainRedirDev DomainRedirDev

func (a *DomainRedirDev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "redirdev"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	fs := domainRedirDev(*a)
	return e.EncodeElement(fs, start)
}

func (a *DomainRedirDev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainRedirDev(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainRedirDev(con)
	return nil
}

func (d *DomainRedirDev) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainRedirDev) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainMemBalloon) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemBalloon) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainVSock) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainVSock) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainSound) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainSound) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

type domainRNGBackendEGD DomainRNGBackendEGD

func (a *DomainRNGBackendEGD) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "backend"
	if a.Source != nil {
		typ := getChardevSourceType(a.Source)
		if typ != "" {
			start.Attr = append(start.Attr, xml.Attr{
				xml.Name{Local: "type"}, typ,
			})
		}
	}
	egd := domainRNGBackendEGD(*a)
	return e.EncodeElement(egd, start)
}

func (a *DomainRNGBackendEGD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		typ = "pty"
	}
	a.Source = createChardevSource(typ)
	con := domainRNGBackendEGD(*a)
	err := d.DecodeElement(&con, &start)
	if err != nil {
		return err
	}
	*a = DomainRNGBackendEGD(con)
	return nil
}

func (a *DomainRNGBackend) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Random != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "random",
		})
		return e.EncodeElement(a.Random, start)
	} else if a.EGD != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "model"}, "egd",
		})
		return e.EncodeElement(a.EGD, start)
	}
	return nil
}

func (a *DomainRNGBackend) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	model, ok := getAttr(start.Attr, "model")
	if !ok {
		return nil
	}
	if model == "random" {
		a.Random = &DomainRNGBackendRandom{}
		err := d.DecodeElement(a.Random, &start)
		if err != nil {
			return err
		}
	} else if model == "egd" {
		a.EGD = &DomainRNGBackendEGD{}
		err := d.DecodeElement(a.EGD, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
}

func (d *DomainRNG) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainRNG) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainHostdevSubsysSCSISource) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.Host != nil {
		return e.EncodeElement(a.Host, start)
	} else if a.ISCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "protocol"}, "iscsi",
		})
		return e.EncodeElement(a.ISCSI, start)
	}
	return nil
}

func (a *DomainHostdevSubsysSCSISource) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	proto, ok := getAttr(start.Attr, "protocol")
	if !ok {
		a.Host = &DomainHostdevSubsysSCSISourceHost{}
		err := d.DecodeElement(a.Host, &start)
		if err != nil {
			return err
		}
	}
	if proto == "iscsi" {
		a.ISCSI = &DomainHostdevSubsysSCSISourceISCSI{}
		err := d.DecodeElement(a.ISCSI, &start)
		if err != nil {
			return err
		}
	}
	d.Skip()
	return nil
}

type domainHostdev DomainHostdev

type domainHostdevSubsysSCSI struct {
	DomainHostdevSubsysSCSI
	domainHostdev
}

type domainHostdevSubsysSCSIHost struct {
	DomainHostdevSubsysSCSIHost
	domainHostdev
}

type domainHostdevSubsysUSB struct {
	DomainHostdevSubsysUSB
	domainHostdev
}

type domainHostdevSubsysPCI struct {
	DomainHostdevSubsysPCI
	domainHostdev
}

type domainHostdevSubsysMDev struct {
	DomainHostdevSubsysMDev
	domainHostdev
}

type domainHostdevCapsStorage struct {
	DomainHostdevCapsStorage
	domainHostdev
}

type domainHostdevCapsMisc struct {
	DomainHostdevCapsMisc
	domainHostdev
}

type domainHostdevCapsNet struct {
	DomainHostdevCapsNet
	domainHostdev
}

func (a *DomainHostdev) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "hostdev"
	if a.SubsysSCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		scsi := domainHostdevSubsysSCSI{}
		scsi.domainHostdev = domainHostdev(*a)
		scsi.DomainHostdevSubsysSCSI = *a.SubsysSCSI
		return e.EncodeElement(scsi, start)
	} else if a.SubsysSCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		scsi_host := domainHostdevSubsysSCSIHost{}
		scsi_host.domainHostdev = domainHostdev(*a)
		scsi_host.DomainHostdevSubsysSCSIHost = *a.SubsysSCSIHost
		return e.EncodeElement(scsi_host, start)
	} else if a.SubsysUSB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		usb := domainHostdevSubsysUSB{}
		usb.domainHostdev = domainHostdev(*a)
		usb.DomainHostdevSubsysUSB = *a.SubsysUSB
		return e.EncodeElement(usb, start)
	} else if a.SubsysPCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		pci := domainHostdevSubsysPCI{}
		pci.domainHostdev = domainHostdev(*a)
		pci.DomainHostdevSubsysPCI = *a.SubsysPCI
		return e.EncodeElement(pci, start)
	} else if a.SubsysMDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "subsystem",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		mdev := domainHostdevSubsysMDev{}
		mdev.domainHostdev = domainHostdev(*a)
		mdev.DomainHostdevSubsysMDev = *a.SubsysMDev
		return e.EncodeElement(mdev, start)
	} else if a.CapsStorage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		storage := domainHostdevCapsStorage{}
		storage.domainHostdev = domainHostdev(*a)
		storage.DomainHostdevCapsStorage = *a.CapsStorage
		return e.EncodeElement(storage, start)
	} else if a.CapsMisc != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "misc",
		})
		misc := domainHostdevCapsMisc{}
		misc.domainHostdev = domainHostdev(*a)
		misc.DomainHostdevCapsMisc = *a.CapsMisc
		return e.EncodeElement(misc, start)
	} else if a.CapsNet != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "mode"}, "capabilities",
		})
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		net := domainHostdevCapsNet{}
		net.domainHostdev = domainHostdev(*a)
		net.DomainHostdevCapsNet = *a.CapsNet
		return e.EncodeElement(net, start)
	} else {
		gen := domainHostdev(*a)
		return e.EncodeElement(gen, start)
	}
}

func (a *DomainHostdev) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mode, ok := getAttr(start.Attr, "mode")
	if !ok {
		return fmt.Errorf("Missing 'mode' attribute on domain hostdev")
	}
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain controller")
	}
	if mode == "subsystem" {
		if typ == "scsi" {
			var scsi domainHostdevSubsysSCSI
			err := d.DecodeElement(&scsi, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi.domainHostdev)
			a.SubsysSCSI = &scsi.DomainHostdevSubsysSCSI
			return nil
		} else if typ == "scsi_host" {
			var scsi_host domainHostdevSubsysSCSIHost
			err := d.DecodeElement(&scsi_host, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(scsi_host.domainHostdev)
			a.SubsysSCSIHost = &scsi_host.DomainHostdevSubsysSCSIHost
			return nil
		} else if typ == "usb" {
			var usb domainHostdevSubsysUSB
			err := d.DecodeElement(&usb, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(usb.domainHostdev)
			a.SubsysUSB = &usb.DomainHostdevSubsysUSB
			return nil
		} else if typ == "pci" {
			var pci domainHostdevSubsysPCI
			err := d.DecodeElement(&pci, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(pci.domainHostdev)
			a.SubsysPCI = &pci.DomainHostdevSubsysPCI
			return nil
		} else if typ == "mdev" {
			var mdev domainHostdevSubsysMDev
			err := d.DecodeElement(&mdev, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(mdev.domainHostdev)
			a.SubsysMDev = &mdev.DomainHostdevSubsysMDev
			return nil
		}
	} else if mode == "capabilities" {
		if typ == "storage" {
			var storage domainHostdevCapsStorage
			err := d.DecodeElement(&storage, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(storage.domainHostdev)
			a.CapsStorage = &storage.DomainHostdevCapsStorage
			return nil
		} else if typ == "misc" {
			var misc domainHostdevCapsMisc
			err := d.DecodeElement(&misc, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(misc.domainHostdev)
			a.CapsMisc = &misc.DomainHostdevCapsMisc
			return nil
		} else if typ == "net" {
			var net domainHostdevCapsNet
			err := d.DecodeElement(&net, &start)
			if err != nil {
				return err
			}
			*a = DomainHostdev(net.domainHostdev)
			a.CapsNet = &net.DomainHostdevCapsNet
			return nil
		}
	}
	var gen domainHostdev
	err := d.DecodeElement(&gen, &start)
	if err != nil {
		return err
	}
	*a = DomainHostdev(gen)
	return nil
}

func (d *DomainHostdev) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainHostdev) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainGraphicListener) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "listen"
	if a.Address != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "address",
		})
		return e.EncodeElement(a.Address, start)
	} else if a.Network != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "network",
		})
		return e.EncodeElement(a.Network, start)
	} else if a.Socket != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "socket",
		})
		return e.EncodeElement(a.Socket, start)
	} else {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "none",
		})
		e.EncodeToken(start)
		e.EncodeToken(start.End())
	}
	return nil
}

func (a *DomainGraphicListener) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain graphics listen")
	}
	if typ == "address" {
		var addr DomainGraphicListenerAddress
		err := d.DecodeElement(&addr, &start)
		if err != nil {
			return err
		}
		a.Address = &addr
		return nil
	} else if typ == "network" {
		var net DomainGraphicListenerNetwork
		err := d.DecodeElement(&net, &start)
		if err != nil {
			return err
		}
		a.Network = &net
		return nil
	} else if typ == "socket" {
		var sock DomainGraphicListenerSocket
		err := d.DecodeElement(&sock, &start)
		if err != nil {
			return err
		}
		a.Socket = &sock
		return nil
	} else if typ == "none" {
		d.Skip()
	}
	return nil
}

func (a *DomainGraphic) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "graphics"
	if a.SDL != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "sdl",
		})
		return e.EncodeElement(a.SDL, start)
	} else if a.VNC != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "vnc",
		})
		return e.EncodeElement(a.VNC, start)
	} else if a.RDP != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "rdp",
		})
		return e.EncodeElement(a.RDP, start)
	} else if a.Desktop != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "desktop",
		})
		return e.EncodeElement(a.Desktop, start)
	} else if a.Spice != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spice",
		})
		return e.EncodeElement(a.Spice, start)
	} else if a.EGLHeadless != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "egl-headless",
		})
		return e.EncodeElement(a.EGLHeadless, start)
	}
	return nil
}

func (a *DomainGraphic) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing 'type' attribute on domain graphics")
	}
	if typ == "sdl" {
		var sdl DomainGraphicSDL
		err := d.DecodeElement(&sdl, &start)
		if err != nil {
			return err
		}
		a.SDL = &sdl
		return nil
	} else if typ == "vnc" {
		var vnc DomainGraphicVNC
		err := d.DecodeElement(&vnc, &start)
		if err != nil {
			return err
		}
		a.VNC = &vnc
		return nil
	} else if typ == "rdp" {
		var rdp DomainGraphicRDP
		err := d.DecodeElement(&rdp, &start)
		if err != nil {
			return err
		}
		a.RDP = &rdp
		return nil
	} else if typ == "desktop" {
		var desktop DomainGraphicDesktop
		err := d.DecodeElement(&desktop, &start)
		if err != nil {
			return err
		}
		a.Desktop = &desktop
		return nil
	} else if typ == "spice" {
		var spice DomainGraphicSpice
		err := d.DecodeElement(&spice, &start)
		if err != nil {
			return err
		}
		a.Spice = &spice
		return nil
	} else if typ == "egl-headless" {
		var egl DomainGraphicEGLHeadless
		err := d.DecodeElement(&egl, &start)
		if err != nil {
			return err
		}
		a.EGLHeadless = &egl
		return nil
	}
	return nil
}

func (d *DomainMemorydev) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainMemorydev) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (d *DomainWatchdog) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainWatchdog) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func marshalUintAttr(start *xml.StartElement, name string, val *uint, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func marshalUint64Attr(start *xml.StartElement, name string, val *uint64, format string) {
	if val != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: name}, fmt.Sprintf(format, *val),
		})
	}
}

func (a *DomainAddressPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	if a.MultiFunction != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "multifunction"}, a.MultiFunction,
		})
	}
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressUSB) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	if a.Port != "" {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "port"}, a.Port,
		})
	}
	marshalUintAttr(&start, "device", a.Device, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDrive) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	marshalUintAttr(&start, "target", a.Target, "%d")
	marshalUintAttr(&start, "unit", a.Unit, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressDIMM) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "slot", a.Slot, "%d")
	marshalUint64Attr(&start, "base", a.Base, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressISA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "iobase", a.IOBase, "0x%x")
	marshalUintAttr(&start, "irq", a.IRQ, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioMMIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCW) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "cssid", a.CSSID, "0x%x")
	marshalUintAttr(&start, "ssid", a.SSID, "0x%x")
	marshalUintAttr(&start, "devno", a.DevNo, "0x%04x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioSerial) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "bus", a.Bus, "%d")
	marshalUintAttr(&start, "port", a.Port, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressSpaprVIO) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUint64Attr(&start, "reg", a.Reg, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressCCID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "controller", a.Controller, "%d")
	marshalUintAttr(&start, "slot", a.Slot, "%d")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddressVirtioS390) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *DomainAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return e.EncodeElement(a.USB, start)
	} else if a.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(a.PCI, start)
	} else if a.Drive != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drive",
		})
		return e.EncodeElement(a.Drive, start)
	} else if a.DIMM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "dimm",
		})
		return e.EncodeElement(a.DIMM, start)
	} else if a.ISA != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "isa",
		})
		return e.EncodeElement(a.ISA, start)
	} else if a.VirtioMMIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-mmio",
		})
		return e.EncodeElement(a.VirtioMMIO, start)
	} else if a.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return e.EncodeElement(a.CCW, start)
	} else if a.VirtioSerial != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-serial",
		})
		return e.EncodeElement(a.VirtioSerial, start)
	} else if a.SpaprVIO != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "spapr-vio",
		})
		return e.EncodeElement(a.SpaprVIO, start)
	} else if a.CCID != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccid",
		})
		return e.EncodeElement(a.CCID, start)
	} else if a.VirtioS390 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virtio-s390",
		})
		return e.EncodeElement(a.VirtioS390, start)
	} else {
		return nil
	}
}

func unmarshalUint64Attr(valstr string, valptr **uint64, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseUint(valstr, base, 64)
	if err != nil {
		return err
	}
	*valptr = &val
	return nil
}

func unmarshalUintAttr(valstr string, valptr **uint, base int) error {
	if base == 16 {
		valstr = strings.TrimPrefix(valstr, "0x")
	}
	val, err := strconv.ParseUint(valstr, base, 64)
	if err != nil {
		return err
	}
	vali := uint(val)
	*valptr = &vali
	return nil
}

func (a *DomainAddressUSB) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			a.Port = attr.Value
		} else if attr.Name.Local == "device" {
			if err := unmarshalUintAttr(attr.Value, &a.Device, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "domain" {
			if err := unmarshalUintAttr(attr.Value, &a.Domain, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalUintAttr(attr.Value, &a.Slot, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "function" {
			if err := unmarshalUintAttr(attr.Value, &a.Function, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "multifunction" {
			a.MultiFunction = attr.Value
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressDrive) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalUintAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "target" {
			if err := unmarshalUintAttr(attr.Value, &a.Target, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "unit" {
			if err := unmarshalUintAttr(attr.Value, &a.Unit, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressDIMM) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "slot" {
			if err := unmarshalUintAttr(attr.Value, &a.Slot, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "base" {
			if err := unmarshalUint64Attr(attr.Value, &a.Base, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressISA) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "iobase" {
			if err := unmarshalUintAttr(attr.Value, &a.IOBase, 16); err != nil {
				return err
			}
		} else if attr.Name.Local == "irq" {
			if err := unmarshalUintAttr(attr.Value, &a.IRQ, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioMMIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddressCCW) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "cssid" {
			if err := unmarshalUintAttr(attr.Value, &a.CSSID, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "ssid" {
			if err := unmarshalUintAttr(attr.Value, &a.SSID, 0); err != nil {
				return err
			}
		} else if attr.Name.Local == "devno" {
			if err := unmarshalUintAttr(attr.Value, &a.DevNo, 0); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioSerial) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalUintAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "bus" {
			if err := unmarshalUintAttr(attr.Value, &a.Bus, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "port" {
			if err := unmarshalUintAttr(attr.Value, &a.Port, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressSpaprVIO) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "reg" {
			if err := unmarshalUint64Attr(attr.Value, &a.Reg, 16); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressCCID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "controller" {
			if err := unmarshalUintAttr(attr.Value, &a.Controller, 10); err != nil {
				return err
			}
		} else if attr.Name.Local == "slot" {
			if err := unmarshalUintAttr(attr.Value, &a.Slot, 10); err != nil {
				return err
			}
		}
	}
	d.Skip()
	return nil
}

func (a *DomainAddressVirtioS390) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	d.Skip()
	return nil
}

func (a *DomainAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
			break
		}
	}
	if typ == "" {
		d.Skip()
		return nil
	}

	if typ == "usb" {
		a.USB = &DomainAddressUSB{}
		return d.DecodeElement(a.USB, &start)
	} else if typ == "pci" {
		a.PCI = &DomainAddressPCI{}
		return d.DecodeElement(a.PCI, &start)
	} else if typ == "drive" {
		a.Drive = &DomainAddressDrive{}
		return d.DecodeElement(a.Drive, &start)
	} else if typ == "dimm" {
		a.DIMM = &DomainAddressDIMM{}
		return d.DecodeElement(a.DIMM, &start)
	} else if typ == "isa" {
		a.ISA = &DomainAddressISA{}
		return d.DecodeElement(a.ISA, &start)
	} else if typ == "virtio-mmio" {
		a.VirtioMMIO = &DomainAddressVirtioMMIO{}
		return d.DecodeElement(a.VirtioMMIO, &start)
	} else if typ == "ccw" {
		a.CCW = &DomainAddressCCW{}
		return d.DecodeElement(a.CCW, &start)
	} else if typ == "virtio-serial" {
		a.VirtioSerial = &DomainAddressVirtioSerial{}
		return d.DecodeElement(a.VirtioSerial, &start)
	} else if typ == "spapr-vio" {
		a.SpaprVIO = &DomainAddressSpaprVIO{}
		return d.DecodeElement(a.SpaprVIO, &start)
	} else if typ == "ccid" {
		a.CCID = &DomainAddressCCID{}
		return d.DecodeElement(a.CCID, &start)
	} else if typ == "spapr-vio" {
		a.VirtioS390 = &DomainAddressVirtioS390{}
		return d.DecodeElement(a.VirtioS390, &start)
	}

	return nil
}

func (d *DomainCPU) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), d)
}

func (d *DomainCPU) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (a *DomainLaunchSecuritySEV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	cbitpos := xml.StartElement{
		Name: xml.Name{Local: "cbitpos"},
	}
	e.EncodeToken(cbitpos)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.CBitPos)))
	e.EncodeToken(cbitpos.End())

	reducedPhysBits := xml.StartElement{
		Name: xml.Name{Local: "reducedPhysBits"},
	}
	e.EncodeToken(reducedPhysBits)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *a.ReducedPhysBits)))
	e.EncodeToken(reducedPhysBits.End())

	if a.Policy != nil {
		policy := xml.StartElement{
			Name: xml.Name{Local: "policy"},
		}
		e.EncodeToken(policy)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *a.Policy)))
		e.EncodeToken(policy.End())
	}

	dhcert := xml.StartElement{
		Name: xml.Name{Local: "dhCert"},
	}
	e.EncodeToken(dhcert)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%s", a.DHCert)))
	e.EncodeToken(dhcert.End())

	session := xml.StartElement{
		Name: xml.Name{Local: "session"},
	}
	e.EncodeToken(session)
	e.EncodeToken(xml.CharData(fmt.Sprintf("%s", a.Session)))
	e.EncodeToken(session.End())

	e.EncodeToken(start.End())

	return nil
}

func (a *DomainLaunchSecuritySEV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "policy" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalUintAttr(string(data), &a.Policy, 16); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "cbitpos" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalUintAttr(string(data), &a.CBitPos, 10); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "reducedPhysBits" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					if err := unmarshalUintAttr(string(data), &a.ReducedPhysBits, 10); err != nil {
						return err
					}
				}
			} else if tok.Name.Local == "dhCert" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					a.DHCert = string(data)
				}
			} else if tok.Name.Local == "session" {
				data, err := d.Token()
				if err != nil {
					return err
				}
				switch data := data.(type) {
				case xml.CharData:
					a.Session = string(data)
				}
			}
		}
	}
	return nil
}

func (a *DomainLaunchSecurity) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if a.SEV != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "sev",
		})
		return e.EncodeElement(a.SEV, start)
	} else {
		return nil
	}

}

func (a *DomainLaunchSecurity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typ string
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			typ = attr.Value
		}
	}

	if typ == "" {
		d.Skip()
		return nil
	}

	if typ == "sev" {
		a.SEV = &DomainLaunchSecuritySEV{}
		return d.DecodeElement(a.SEV, &start)
	}

	return nil
}
