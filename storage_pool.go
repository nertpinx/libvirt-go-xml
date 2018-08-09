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
 * Copyright (C) 2017 Red Hat, Inc.
 *
 */

package libvirtxml

import "encoding/xml"

type StoragePoolSize struct {
	Unit  string `xml:"unit,attr,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Value uint64 `xml:",attr" json:"" yaml:""`
}

type StoragePoolTargetPermissions struct {
	Owner string `xml:"owner,omitempty" json:"owner,omitempty" yaml:"owner,omitempty"`
	Group string `xml:"group,omitempty" json:"group,omitempty" yaml:"group,omitempty"`
	Mode  string `xml:"mode,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Label string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
}

type StoragePoolTargetTimestamps struct {
	Atime string `xml:"atime" json:"atime" yaml:"atime"`
	Mtime string `xml:"mtime" json:"mtime" yaml:"mtime"`
	Ctime string `xml:"ctime" json:"ctime" yaml:"ctime"`
}

type StoragePoolTarget struct {
	Path        string                        `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	Permissions *StoragePoolTargetPermissions `xml:"permissions" json:"permissions,omitempty" yaml:"permissions,omitempty"`
	Timestamps  *StoragePoolTargetTimestamps  `xml:"timestamps" json:"timestamps,omitempty" yaml:"timestamps,omitempty"`
	Encryption  *StorageEncryption            `xml:"encryption" json:"encryption,omitempty" yaml:"encryption,omitempty"`
}

type StoragePoolSourceFormat struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
}
type StoragePoolSourceHost struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Port string `xml:"port,attr,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
}

type StoragePoolSourceDevice struct {
	Path          string                              `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
	PartSeparator string                              `xml:"part_separator,attr,omitempty" json:"part_separator,omitempty" yaml:"part_separator,omitempty"`
	FreeExtents   []StoragePoolSourceDeviceFreeExtent `xml:"freeExtent" json:"freeExtent" yaml:"freeExtent"`
}

type StoragePoolSourceDeviceFreeExtent struct {
	Start uint64 `xml:"start,attr" json:"start,omitempty" yaml:"start,omitempty"`
	End   uint64 `xml:"end,attr" json:"end,omitempty" yaml:"end,omitempty"`
}

type StoragePoolSourceAuthSecret struct {
	Usage string `xml:"usage,attr,omitempty" json:"usage,omitempty" yaml:"usage,omitempty"`
	UUID  string `xml:"uuid,attr,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type StoragePoolSourceAuth struct {
	Type     string                       `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Username string                       `xml:"username,attr" json:"username,omitempty" yaml:"username,omitempty"`
	Secret   *StoragePoolSourceAuthSecret `xml:"secret" json:"secret,omitempty" yaml:"secret,omitempty"`
}

type StoragePoolSourceVendor struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type StoragePoolSourceProduct struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type StoragePoolPCIAddress struct {
	Domain   *uint `xml:"domain,attr" json:"domain,omitempty,omitempty" yaml:"domain,omitempty,omitempty"`
	Bus      *uint `xml:"bus,attr" json:"bus,omitempty,omitempty" yaml:"bus,omitempty,omitempty"`
	Slot     *uint `xml:"slot,attr" json:"slot,omitempty,omitempty" yaml:"slot,omitempty,omitempty"`
	Function *uint `xml:"function,attr" json:"function,omitempty,omitempty" yaml:"function,omitempty,omitempty"`
}

type StoragePoolSourceAdapterParentAddr struct {
	UniqueID uint64                 `xml:"unique_id,attr" json:"unique_id,omitempty" yaml:"unique_id,omitempty"`
	Address  *StoragePoolPCIAddress `xml:"address" json:"address,omitempty" yaml:"address,omitempty"`
}

type StoragePoolSourceAdapter struct {
	Type       string                              `xml:"type,attr,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Name       string                              `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Parent     string                              `xml:"parent,attr,omitempty" json:"parent,omitempty" yaml:"parent,omitempty"`
	Managed    string                              `xml:"managed,attr,omitempty" json:"managed,omitempty" yaml:"managed,omitempty"`
	WWNN       string                              `xml:"wwnn,attr,omitempty" json:"wwnn,omitempty" yaml:"wwnn,omitempty"`
	WWPN       string                              `xml:"wwpn,attr,omitempty" json:"wwpn,omitempty" yaml:"wwpn,omitempty"`
	ParentAddr *StoragePoolSourceAdapterParentAddr `xml:"parentaddr" json:"parentaddr,omitempty" yaml:"parentaddr,omitempty"`
}

type StoragePoolSourceDir struct {
	Path string `xml:"path,attr" json:"path,omitempty" yaml:"path,omitempty"`
}

type StoragePoolSourceInitiator struct {
	IQN StoragePoolSourceInitiatorIQN `xml:"iqn" json:"iqn" yaml:"iqn"`
}

type StoragePoolSourceInitiatorIQN struct {
	Name string `xml:"name,attr,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

type StoragePoolSource struct {
	Name      string                      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Dir       *StoragePoolSourceDir       `xml:"dir" json:"dir,omitempty" yaml:"dir,omitempty"`
	Host      []StoragePoolSourceHost     `xml:"host" json:"host" yaml:"host"`
	Device    []StoragePoolSourceDevice   `xml:"device" json:"device" yaml:"device"`
	Auth      *StoragePoolSourceAuth      `xml:"auth" json:"auth,omitempty" yaml:"auth,omitempty"`
	Vendor    *StoragePoolSourceVendor    `xml:"vendor" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Product   *StoragePoolSourceProduct   `xml:"product" json:"product,omitempty" yaml:"product,omitempty"`
	Format    *StoragePoolSourceFormat    `xml:"format" json:"format,omitempty" yaml:"format,omitempty"`
	Adapter   *StoragePoolSourceAdapter   `xml:"adapter" json:"adapter,omitempty" yaml:"adapter,omitempty"`
	Initiator *StoragePoolSourceInitiator `xml:"initiator" json:"initiator,omitempty" yaml:"initiator,omitempty"`
}

type StoragePool struct {
	XMLName    xml.Name           `xml:"pool" json:"pool" yaml:"pool"`
	Type       string             `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Name       string             `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	UUID       string             `xml:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Allocation *StoragePoolSize   `xml:"allocation" json:"allocation,omitempty" yaml:"allocation,omitempty"`
	Capacity   *StoragePoolSize   `xml:"capacity" json:"capacity,omitempty" yaml:"capacity,omitempty"`
	Available  *StoragePoolSize   `xml:"available" json:"available,omitempty" yaml:"available,omitempty"`
	Target     *StoragePoolTarget `xml:"target" json:"target,omitempty" yaml:"target,omitempty"`
	Source     *StoragePoolSource `xml:"source" json:"source,omitempty" yaml:"source,omitempty"`
}

func (a *StoragePoolPCIAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *StoragePoolPCIAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		}
	}
	d.Skip()
	return nil
}

func (s *StoragePool) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StoragePool) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
