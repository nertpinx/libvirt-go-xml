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

type StorageVolumeSize struct {
	Unit  string `xml:"unit,attr,omitempty",json:"unit,attr,omitempty",yaml:"unit,attr,omitempty"`
	Value uint64 `xml:",chardata",json:",chardata",yaml:",chardata"`
}

type StorageVolumeTargetPermissions struct {
	Owner string `xml:"owner,omitempty",json:"owner,omitempty",yaml:"owner,omitempty"`
	Group string `xml:"group,omitempty",json:"group,omitempty",yaml:"group,omitempty"`
	Mode  string `xml:"mode,omitempty",json:"mode,omitempty",yaml:"mode,omitempty"`
	Label string `xml:"label,omitempty",json:"label,omitempty",yaml:"label,omitempty"`
}

type StorageVolumeTargetFeature struct {
	LazyRefcounts *struct{} `xml:"lazy_refcounts",json:"lazy_refcounts",yaml:"lazy_refcounts"`
}

type StorageVolumeTargetFormat struct {
	Type string `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
}

type StorageVolumeTargetTimestamps struct {
	Atime string `xml:"atime",json:"atime",yaml:"atime"`
	Mtime string `xml:"mtime",json:"mtime",yaml:"mtime"`
	Ctime string `xml:"ctime",json:"ctime",yaml:"ctime"`
}

type StorageVolumeTarget struct {
	Path        string                          `xml:"path,omitempty",json:"path,omitempty",yaml:"path,omitempty"`
	Format      *StorageVolumeTargetFormat      `xml:"format",json:"format",yaml:"format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions",json:"permissions",yaml:"permissions"`
	Timestamps  *StorageVolumeTargetTimestamps  `xml:"timestamps",json:"timestamps",yaml:"timestamps"`
	Compat      string                          `xml:"compat,omitempty",json:"compat,omitempty",yaml:"compat,omitempty"`
	NoCOW       *struct{}                       `xml:"nocow",json:"nocow",yaml:"nocow"`
	Features    []StorageVolumeTargetFeature    `xml:"features",json:"features",yaml:"features"`
	Encryption  *StorageEncryption              `xml:"encryption",json:"encryption",yaml:"encryption"`
}

type StorageVolumeBackingStore struct {
	Path        string                          `xml:"path",json:"path",yaml:"path"`
	Format      *StorageVolumeTargetFormat      `xml:"format",json:"format",yaml:"format"`
	Permissions *StorageVolumeTargetPermissions `xml:"permissions",json:"permissions",yaml:"permissions"`
}

type StorageVolume struct {
	XMLName      xml.Name                   `xml:"volume",json:"volume",yaml:"volume"`
	Type         string                     `xml:"type,attr,omitempty",json:"type,attr,omitempty",yaml:"type,attr,omitempty"`
	Name         string                     `xml:"name",json:"name",yaml:"name"`
	Key          string                     `xml:"key,omitempty",json:"key,omitempty",yaml:"key,omitempty"`
	Allocation   *StorageVolumeSize         `xml:"allocation",json:"allocation",yaml:"allocation"`
	Capacity     *StorageVolumeSize         `xml:"capacity",json:"capacity",yaml:"capacity"`
	Physical     *StorageVolumeSize         `xml:"physical",json:"physical",yaml:"physical"`
	Target       *StorageVolumeTarget       `xml:"target",json:"target",yaml:"target"`
	BackingStore *StorageVolumeBackingStore `xml:"backingStore",json:"backingStore",yaml:"backingStore"`
}

func (s *StorageVolume) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *StorageVolume) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
