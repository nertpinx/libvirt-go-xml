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
 * Copyright (C) 2017 Lian Duan <blazeblue@gmail.com>
 *
 */

package libvirtxml

import (
	"encoding/xml"
)

type Interface struct {
	XMLName  xml.Name            `xml:"interface" json:"interface" yaml:"interface"`
	Name     string              `xml:"name,attr,omitempty" json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	Start    *InterfaceStart     `xml:"start" json:"start" yaml:"start"`
	MTU      *InterfaceMTU       `xml:"mtu" json:"mtu" yaml:"mtu"`
	Protocol []InterfaceProtocol `xml:"protocol" json:"protocol" yaml:"protocol"`
	Link     *InterfaceLink      `xml:"link" json:"link" yaml:"link"`
	MAC      *InterfaceMAC       `xml:"mac" json:"mac" yaml:"mac"`
	Bond     *InterfaceBond      `xml:"bond" json:"bond" yaml:"bond"`
	Bridge   *InterfaceBridge    `xml:"bridge" json:"bridge" yaml:"bridge"`
	VLAN     *InterfaceVLAN      `xml:"vlan" json:"vlan" yaml:"vlan"`
}

type InterfaceStart struct {
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
}

type InterfaceMTU struct {
	Size uint `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
}

type InterfaceProtocol struct {
	Family   string             `xml:"family,attr,omitempty" json:"family,omitempty,omitempty" yaml:"family,omitempty,omitempty"`
	AutoConf *InterfaceAutoConf `xml:"autoconf" json:"autoconf" yaml:"autoconf"`
	DHCP     *InterfaceDHCP     `xml:"dhcp" json:"dhcp" yaml:"dhcp"`
	IPs      []InterfaceIP      `xml:"ip" json:"ip" yaml:"ip"`
	Route    []InterfaceRoute   `xml:"route" json:"route" yaml:"route"`
}

type InterfaceAutoConf struct {
}

type InterfaceDHCP struct {
	PeerDNS string `xml:"peerdns,attr,omitempty" json:"peerdns,omitempty,omitempty" yaml:"peerdns,omitempty,omitempty"`
}

type InterfaceIP struct {
	Address string `xml:"address,attr" json:"address,omitempty" yaml:"address,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty,omitempty" yaml:"prefix,omitempty,omitempty"`
}

type InterfaceRoute struct {
	Gateway string `xml:"gateway,attr" json:"gateway,omitempty" yaml:"gateway,omitempty"`
}

type InterfaceLink struct {
	Speed uint   `xml:"speed,attr,omitempty" json:"speed,omitempty,omitempty" yaml:"speed,omitempty,omitempty"`
	State string `xml:"state,attr,omitempty" json:"state,omitempty,omitempty" yaml:"state,omitempty,omitempty"`
}

type InterfaceMAC struct {
	Address string `xml:"address,attr" json:"address,omitempty" yaml:"address,omitempty"`
}

type InterfaceBond struct {
	Mode       string               `xml:"mode,attr,omitempty" json:"mode,omitempty,omitempty" yaml:"mode,omitempty,omitempty"`
	ARPMon     *InterfaceBondARPMon `xml:"arpmon" json:"arpmon" yaml:"arpmon"`
	MIIMon     *InterfaceBondMIIMon `xml:"miimon" json:"miimon" yaml:"miimon"`
	Interfaces []Interface          `xml:"interface" json:"interface" yaml:"interface"`
}

type InterfaceBondARPMon struct {
	Interval uint   `xml:"interval,attr,omitempty" json:"interval,omitempty,omitempty" yaml:"interval,omitempty,omitempty"`
	Target   string `xml:"target,attr,omitempty" json:"target,omitempty,omitempty" yaml:"target,omitempty,omitempty"`
	Validate string `xml:"validate,attr,omitempty" json:"validate,omitempty,omitempty" yaml:"validate,omitempty,omitempty"`
}

type InterfaceBondMIIMon struct {
	Freq    uint   `xml:"freq,attr,omitempty" json:"freq,omitempty,omitempty" yaml:"freq,omitempty,omitempty"`
	UpDelay uint   `xml:"updelay,attr,omitempty" json:"updelay,omitempty,omitempty" yaml:"updelay,omitempty,omitempty"`
	Carrier string `xml:"carrier,attr,omitempty" json:"carrier,omitempty,omitempty" yaml:"carrier,omitempty,omitempty"`
}

type InterfaceBridge struct {
	STP        string      `xml:"stp,attr,omitempty" json:"stp,omitempty,omitempty" yaml:"stp,omitempty,omitempty"`
	Delay      *float64    `xml:"delay,attr" json:"delay,omitempty" yaml:"delay,omitempty"`
	Interfaces []Interface `xml:"interface" json:"interface" yaml:"interface"`
}

type InterfaceVLAN struct {
	Tag       *uint      `xml:"tag,attr" json:"tag,omitempty" yaml:"tag,omitempty"`
	Interface *Interface `xml:"interface" json:"interface" yaml:"interface"`
}

type interfaceDup Interface

func (s *Interface) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "interface"

	typ := "ethernet"
	if s.Bond != nil {
		typ = "bond"
	} else if s.Bridge != nil {
		typ = "bridge"
	} else if s.VLAN != nil {
		typ = "vlan"
	}

	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "type"},
		Value: typ,
	})

	i := interfaceDup(*s)

	return e.EncodeElement(i, start)
}

func (s *Interface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Interface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
