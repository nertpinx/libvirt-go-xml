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

type DomainCaps struct {
	XMLName   xml.Name             `xml:"domainCapabilities",json:"domainCapabilities",yaml:"domainCapabilities"`
	Path      string               `xml:"path",json:"path",yaml:"path"`
	Domain    string               `xml:"domain",json:"domain",yaml:"domain"`
	Machine   string               `xml:"machine,omitempty",json:"machine,omitempty",yaml:"machine,omitempty"`
	Arch      string               `xml:"arch",json:"arch",yaml:"arch"`
	VCPU      *DomainCapsVCPU      `xml:"vcpu",json:"vcpu",yaml:"vcpu"`
	IOThreads *DomainCapsIOThreads `xml:"iothreads",json:"iothreads",yaml:"iothreads"`
	OS        DomainCapsOS         `xml:"os",json:"os",yaml:"os"`
	CPU       *DomainCapsCPU       `xml:"cpu",json:"cpu",yaml:"cpu"`
	Devices   *DomainCapsDevices   `xml:"devices",json:"devices",yaml:"devices"`
	Features  *DomainCapsFeatures  `xml:"features",json:"features",yaml:"features"`
}

type DomainCapsVCPU struct {
	Max uint `xml:"max,attr",json:"max,attr",yaml:"max,attr"`
}

type DomainCapsOS struct {
	Supported string              `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	Loader    *DomainCapsOSLoader `xml:"loader",json:"loader",yaml:"loader"`
}

type DomainCapsOSLoader struct {
	Supported string           `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	Values    []string         `xml:"value",json:"value",yaml:"value"`
	Enums     []DomainCapsEnum `xml:"enum",json:"enum",yaml:"enum"`
}

type DomainCapsIOThreads struct {
	Supported string `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
}

type DomainCapsCPU struct {
	Modes []DomainCapsCPUMode `xml:"mode",json:"mode",yaml:"mode"`
}

type DomainCapsCPUMode struct {
	Name      string                 `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	Supported string                 `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	Models    []DomainCapsCPUModel   `xml:"model",json:"model",yaml:"model"`
	Vendor    string                 `xml:"vendor,omitempty",json:"vendor,omitempty",yaml:"vendor,omitempty"`
	Features  []DomainCapsCPUFeature `xml:"feature",json:"feature",yaml:"feature"`
}

type DomainCapsCPUModel struct {
	Name     string `xml:",chardata",json:",chardata",yaml:",chardata"`
	Usable   string `xml:"usable,attr,omitempty",json:"usable,attr,omitempty",yaml:"usable,attr,omitempty"`
	Fallback string `xml:"fallback,attr,omitempty",json:"fallback,attr,omitempty",yaml:"fallback,attr,omitempty"`
}

type DomainCapsCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty",json:"policy,attr,omitempty",yaml:"policy,attr,omitempty"`
	Name   string `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
}

type DomainCapsEnum struct {
	Name   string   `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	Values []string `xml:"value",json:"value",yaml:"value"`
}

type DomainCapsDevices struct {
	Disk     *DomainCapsDevice `xml:"disk",json:"disk",yaml:"disk"`
	Graphics *DomainCapsDevice `xml:"graphics",json:"graphics",yaml:"graphics"`
	Video    *DomainCapsDevice `xml:"video",json:"video",yaml:"video"`
	HostDev  *DomainCapsDevice `xml:"hostdev",json:"hostdev",yaml:"hostdev"`
}

type DomainCapsDevice struct {
	Supported string           `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	Enums     []DomainCapsEnum `xml:"enum",json:"enum",yaml:"enum"`
}

type DomainCapsFeatures struct {
	GIC        *DomainCapsFeatureGIC        `xml:"gic",json:"gic",yaml:"gic"`
	VMCoreInfo *DomainCapsFeatureVMCoreInfo `xml:"vmcoreinfo",json:"vmcoreinfo",yaml:"vmcoreinfo"`
	GenID      *DomainCapsFeatureGenID      `xml:"genid",json:"genid",yaml:"genid"`
	SEV        *DomainCapsFeatureSEV        `xml:"sev",json:"sev",yaml:"sev"`
}

type DomainCapsFeatureGIC struct {
	Supported string           `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	Enums     []DomainCapsEnum `xml:"enum",json:"enum",yaml:"enum"`
}

type DomainCapsFeatureVMCoreInfo struct {
	Supported string `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
}

type DomainCapsFeatureGenID struct {
	Supported string `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
}

type DomainCapsFeatureSEV struct {
	Supported       string `xml:"supported,attr",json:"supported,attr",yaml:"supported,attr"`
	CBitPos         uint   `xml:"cbitpos,omitempty",json:"cbitpos,omitempty",yaml:"cbitpos,omitempty"`
	ReducedPhysBits uint   `xml:"reducedPhysBits,omitempty",json:"reducedPhysBits,omitempty",yaml:"reducedPhysBits,omitempty"`
}

func (c *DomainCaps) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *DomainCaps) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
