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

type NetworkBridge struct {
	Name            string `xml:"name,attr,omitempty",json:"name,attr,omitempty",yaml:"name,attr,omitempty"`
	STP             string `xml:"stp,attr,omitempty",json:"stp,attr,omitempty",yaml:"stp,attr,omitempty"`
	Delay           string `xml:"delay,attr,omitempty",json:"delay,attr,omitempty",yaml:"delay,attr,omitempty"`
	MACTableManager string `xml:"macTableManager,attr,omitempty",json:"macTableManager,attr,omitempty",yaml:"macTableManager,attr,omitempty"`
}

type NetworkVirtualPort struct {
	Params *NetworkVirtualPortParams `xml:"parameters",json:"parameters",yaml:"parameters"`
}

type NetworkVirtualPortParams struct {
	Any          *NetworkVirtualPortParamsAny          `xml:"-",json:"-",yaml:"-"`
	VEPA8021QBG  *NetworkVirtualPortParamsVEPA8021QBG  `xml:"-",json:"-",yaml:"-"`
	VNTag8011QBH *NetworkVirtualPortParamsVNTag8021QBH `xml:"-",json:"-",yaml:"-"`
	OpenVSwitch  *NetworkVirtualPortParamsOpenVSwitch  `xml:"-",json:"-",yaml:"-"`
	MidoNet      *NetworkVirtualPortParamsMidoNet      `xml:"-",json:"-",yaml:"-"`
}

type NetworkVirtualPortParamsAny struct {
	ManagerID     *uint  `xml:"managerid,attr",json:"managerid,attr",yaml:"managerid,attr"`
	TypeID        *uint  `xml:"typeid,attr",json:"typeid,attr",yaml:"typeid,attr"`
	TypeIDVersion *uint  `xml:"typeidversion,attr",json:"typeidversion,attr",yaml:"typeidversion,attr"`
	InstanceID    string `xml:"instanceid,attr,omitempty",json:"instanceid,attr,omitempty",yaml:"instanceid,attr,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty",json:"profileid,attr,omitempty",yaml:"profileid,attr,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty",json:"interfaceid,attr,omitempty",yaml:"interfaceid,attr,omitempty"`
}

type NetworkVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *uint  `xml:"managerid,attr",json:"managerid,attr",yaml:"managerid,attr"`
	TypeID        *uint  `xml:"typeid,attr",json:"typeid,attr",yaml:"typeid,attr"`
	TypeIDVersion *uint  `xml:"typeidversion,attr",json:"typeidversion,attr",yaml:"typeidversion,attr"`
	InstanceID    string `xml:"instanceid,attr,omitempty",json:"instanceid,attr,omitempty",yaml:"instanceid,attr,omitempty"`
}

type NetworkVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty",json:"profileid,attr,omitempty",yaml:"profileid,attr,omitempty"`
}

type NetworkVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty",json:"interfaceid,attr,omitempty",yaml:"interfaceid,attr,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty",json:"profileid,attr,omitempty",yaml:"profileid,attr,omitempty"`
}

type NetworkVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty",json:"interfaceid,attr,omitempty",yaml:"interfaceid,attr,omitempty"`
}

type NetworkDomain struct {
	Name      string `xml:"name,attr,omitempty",json:"name,attr,omitempty",yaml:"name,attr,omitempty"`
	LocalOnly string `xml:"localOnly,attr,omitempty",json:"localOnly,attr,omitempty",yaml:"localOnly,attr,omitempty"`
}

type NetworkForwardNATAddress struct {
	Start string `xml:"start,attr",json:"start,attr",yaml:"start,attr"`
	End   string `xml:"end,attr",json:"end,attr",yaml:"end,attr"`
}

type NetworkForwardNATPort struct {
	Start uint `xml:"start,attr",json:"start,attr",yaml:"start,attr"`
	End   uint `xml:"end,attr",json:"end,attr",yaml:"end,attr"`
}

type NetworkForwardNAT struct {
	Addresses []NetworkForwardNATAddress `xml:"address",json:"address",yaml:"address"`
	Ports     []NetworkForwardNATPort    `xml:"port",json:"port",yaml:"port"`
}

type NetworkForward struct {
	Mode       string                    `xml:"mode,attr,omitempty",json:"mode,attr,omitempty",yaml:"mode,attr,omitempty"`
	Dev        string                    `xml:"dev,attr,omitempty",json:"dev,attr,omitempty",yaml:"dev,attr,omitempty"`
	Managed    string                    `xml:"managed,attr,omitempty",json:"managed,attr,omitempty",yaml:"managed,attr,omitempty"`
	Driver     *NetworkForwardDriver     `xml:"driver",json:"driver",yaml:"driver"`
	PFs        []NetworkForwardPF        `xml:"pf",json:"pf",yaml:"pf"`
	NAT        *NetworkForwardNAT        `xml:"nat",json:"nat",yaml:"nat"`
	Interfaces []NetworkForwardInterface `xml:"interface",json:"interface",yaml:"interface"`
	Addresses  []NetworkForwardAddress   `xml:"address",json:"address",yaml:"address"`
}

type NetworkForwardDriver struct {
	Name string `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
}

type NetworkForwardPF struct {
	Dev string `xml:"dev,attr",json:"dev,attr",yaml:"dev,attr"`
}

type NetworkForwardAddress struct {
	PCI *NetworkForwardAddressPCI `xml:"-",json:"-",yaml:"-"`
}

type NetworkForwardAddressPCI struct {
	Domain   *uint `xml:"domain,attr",json:"domain,attr",yaml:"domain,attr"`
	Bus      *uint `xml:"bus,attr",json:"bus,attr",yaml:"bus,attr"`
	Slot     *uint `xml:"slot,attr",json:"slot,attr",yaml:"slot,attr"`
	Function *uint `xml:"function,attr",json:"function,attr",yaml:"function,attr"`
}

type NetworkForwardInterface struct {
	XMLName xml.Name `xml:"interface",json:"interface",yaml:"interface"`
	Dev     string   `xml:"dev,attr,omitempty",json:"dev,attr,omitempty",yaml:"dev,attr,omitempty"`
}

type NetworkMAC struct {
	Address string `xml:"address,attr,omitempty",json:"address,attr,omitempty",yaml:"address,attr,omitempty"`
}

type NetworkDHCPRange struct {
	XMLName xml.Name `xml:"range",json:"range",yaml:"range"`
	Start   string   `xml:"start,attr,omitempty",json:"start,attr,omitempty",yaml:"start,attr,omitempty"`
	End     string   `xml:"end,attr,omitempty",json:"end,attr,omitempty",yaml:"end,attr,omitempty"`
}

type NetworkDHCPHost struct {
	XMLName xml.Name `xml:"host",json:"host",yaml:"host"`
	ID      string   `xml:"id,attr,omitempty",json:"id,attr,omitempty",yaml:"id,attr,omitempty"`
	MAC     string   `xml:"mac,attr,omitempty",json:"mac,attr,omitempty",yaml:"mac,attr,omitempty"`
	Name    string   `xml:"name,attr,omitempty",json:"name,attr,omitempty",yaml:"name,attr,omitempty"`
	IP      string   `xml:"ip,attr,omitempty",json:"ip,attr,omitempty",yaml:"ip,attr,omitempty"`
}

type NetworkBootp struct {
	File   string `xml:"file,attr,omitempty",json:"file,attr,omitempty",yaml:"file,attr,omitempty"`
	Server string `xml:"server,attr,omitempty",json:"server,attr,omitempty",yaml:"server,attr,omitempty"`
}

type NetworkDHCP struct {
	Ranges []NetworkDHCPRange `xml:"range",json:"range",yaml:"range"`
	Hosts  []NetworkDHCPHost  `xml:"host",json:"host",yaml:"host"`
	Bootp  []NetworkBootp     `xml:"bootp",json:"bootp",yaml:"bootp"`
}

type NetworkIP struct {
	Address  string       `xml:"address,attr,omitempty",json:"address,attr,omitempty",yaml:"address,attr,omitempty"`
	Family   string       `xml:"family,attr,omitempty",json:"family,attr,omitempty",yaml:"family,attr,omitempty"`
	Netmask  string       `xml:"netmask,attr,omitempty",json:"netmask,attr,omitempty",yaml:"netmask,attr,omitempty"`
	Prefix   uint         `xml:"prefix,attr,omitempty",json:"prefix,attr,omitempty",yaml:"prefix,attr,omitempty"`
	LocalPtr string       `xml:"localPtr,attr,omitempty",json:"localPtr,attr,omitempty",yaml:"localPtr,attr,omitempty"`
	DHCP     *NetworkDHCP `xml:"dhcp",json:"dhcp",yaml:"dhcp"`
	TFTP     *NetworkTFTP `xml:"tftp",json:"tftp",yaml:"tftp"`
}

type NetworkTFTP struct {
	Root string `xml:"root,attr,omitempty",json:"root,attr,omitempty",yaml:"root,attr,omitempty"`
}

type NetworkRoute struct {
	Family  string `xml:"family,attr,omitempty",json:"family,attr,omitempty",yaml:"family,attr,omitempty"`
	Address string `xml:"address,attr,omitempty",json:"address,attr,omitempty",yaml:"address,attr,omitempty"`
	Netmask string `xml:"netmask,attr,omitempty",json:"netmask,attr,omitempty",yaml:"netmask,attr,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty",json:"prefix,attr,omitempty",yaml:"prefix,attr,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty",json:"gateway,attr,omitempty",yaml:"gateway,attr,omitempty"`
	Metric  string `xml:"metric,attr,omitempty",json:"metric,attr,omitempty",yaml:"metric,attr,omitempty"`
}

type NetworkDNSForwarder struct {
	Domain string `xml:"domain,attr,omitempty",json:"domain,attr,omitempty",yaml:"domain,attr,omitempty"`
	Addr   string `xml:"addr,attr,omitempty",json:"addr,attr,omitempty",yaml:"addr,attr,omitempty"`
}

type NetworkDNSTXT struct {
	XMLName xml.Name `xml:"txt",json:"txt",yaml:"txt"`
	Name    string   `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	Value   string   `xml:"value,attr",json:"value,attr",yaml:"value,attr"`
}

type NetworkDNSHostHostname struct {
	Hostname string `xml:",chardata",json:",chardata",yaml:",chardata"`
}

type NetworkDNSHost struct {
	XMLName   xml.Name                 `xml:"host",json:"host",yaml:"host"`
	IP        string                   `xml:"ip,attr",json:"ip,attr",yaml:"ip,attr"`
	Hostnames []NetworkDNSHostHostname `xml:"hostname",json:"hostname",yaml:"hostname"`
}

type NetworkDNSSRV struct {
	XMLName  xml.Name `xml:"srv",json:"srv",yaml:"srv"`
	Service  string   `xml:"service,attr,omitempty",json:"service,attr,omitempty",yaml:"service,attr,omitempty"`
	Protocol string   `xml:"protocol,attr,omitempty",json:"protocol,attr,omitempty",yaml:"protocol,attr,omitempty"`
	Target   string   `xml:"target,attr,omitempty",json:"target,attr,omitempty",yaml:"target,attr,omitempty"`
	Port     uint     `xml:"port,attr,omitempty",json:"port,attr,omitempty",yaml:"port,attr,omitempty"`
	Priority uint     `xml:"priority,attr,omitempty",json:"priority,attr,omitempty",yaml:"priority,attr,omitempty"`
	Weight   uint     `xml:"weight,attr,omitempty",json:"weight,attr,omitempty",yaml:"weight,attr,omitempty"`
	Domain   string   `xml:"domain,attr,omitempty",json:"domain,attr,omitempty",yaml:"domain,attr,omitempty"`
}

type NetworkDNS struct {
	Enable            string                `xml:"enable,attr,omitempty",json:"enable,attr,omitempty",yaml:"enable,attr,omitempty"`
	ForwardPlainNames string                `xml:"forwardPlainNames,attr,omitempty",json:"forwardPlainNames,attr,omitempty",yaml:"forwardPlainNames,attr,omitempty"`
	Forwarders        []NetworkDNSForwarder `xml:"forwarder",json:"forwarder",yaml:"forwarder"`
	TXTs              []NetworkDNSTXT       `xml:"txt",json:"txt",yaml:"txt"`
	Host              []NetworkDNSHost      `xml:"host",json:"host",yaml:"host"`
	SRVs              []NetworkDNSSRV       `xml:"srv",json:"srv",yaml:"srv"`
}

type NetworkMetadata struct {
	XML string `xml:",innerxml",json:",innerxml",yaml:",innerxml"`
}

type NetworkMTU struct {
	Size uint `xml:"size,attr",json:"size,attr",yaml:"size,attr"`
}

type Network struct {
	XMLName             xml.Name            `xml:"network",json:"network",yaml:"network"`
	IPv6                string              `xml:"ipv6,attr,omitempty",json:"ipv6,attr,omitempty",yaml:"ipv6,attr,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty",json:"trustGuestRxFilters,attr,omitempty",yaml:"trustGuestRxFilters,attr,omitempty"`
	Name                string              `xml:"name,omitempty",json:"name,omitempty",yaml:"name,omitempty"`
	UUID                string              `xml:"uuid,omitempty",json:"uuid,omitempty",yaml:"uuid,omitempty"`
	Metadata            *NetworkMetadata    `xml:"metadata",json:"metadata",yaml:"metadata"`
	Forward             *NetworkForward     `xml:"forward",json:"forward",yaml:"forward"`
	Bridge              *NetworkBridge      `xml:"bridge",json:"bridge",yaml:"bridge"`
	MTU                 *NetworkMTU         `xml:"mtu",json:"mtu",yaml:"mtu"`
	MAC                 *NetworkMAC         `xml:"mac",json:"mac",yaml:"mac"`
	Domain              *NetworkDomain      `xml:"domain",json:"domain",yaml:"domain"`
	DNS                 *NetworkDNS         `xml:"dns",json:"dns",yaml:"dns"`
	VLAN                *NetworkVLAN        `xml:"vlan",json:"vlan",yaml:"vlan"`
	Bandwidth           *NetworkBandwidth   `xml:"bandwidth",json:"bandwidth",yaml:"bandwidth"`
	IPs                 []NetworkIP         `xml:"ip",json:"ip",yaml:"ip"`
	Routes              []NetworkRoute      `xml:"route",json:"route",yaml:"route"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport",json:"virtualport",yaml:"virtualport"`
	PortGroups          []NetworkPortGroup  `xml:"portgroup",json:"portgroup",yaml:"portgroup"`
}

type NetworkPortGroup struct {
	XMLName             xml.Name            `xml:"portgroup",json:"portgroup",yaml:"portgroup"`
	Name                string              `xml:"name,attr,omitempty",json:"name,attr,omitempty",yaml:"name,attr,omitempty"`
	Default             string              `xml:"default,attr,omitempty",json:"default,attr,omitempty",yaml:"default,attr,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty",json:"trustGuestRxFilters,attr,omitempty",yaml:"trustGuestRxFilters,attr,omitempty"`
	VLAN                *NetworkVLAN        `xml:"vlan",json:"vlan",yaml:"vlan"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport",json:"virtualport",yaml:"virtualport"`
}

type NetworkVLAN struct {
	Trunk string           `xml:"trunk,attr,omitempty",json:"trunk,attr,omitempty",yaml:"trunk,attr,omitempty"`
	Tags  []NetworkVLANTag `xml:"tag",json:"tag",yaml:"tag"`
}

type NetworkVLANTag struct {
	ID         uint   `xml:"id,attr",json:"id,attr",yaml:"id,attr"`
	NativeMode string `xml:"nativeMode,attr,omitempty",json:"nativeMode,attr,omitempty",yaml:"nativeMode,attr,omitempty"`
}

type NetworkBandwidthParams struct {
	Average *uint `xml:"average,attr",json:"average,attr",yaml:"average,attr"`
	Peak    *uint `xml:"peak,attr",json:"peak,attr",yaml:"peak,attr"`
	Burst   *uint `xml:"burst,attr",json:"burst,attr",yaml:"burst,attr"`
	Floor   *uint `xml:"floor,attr",json:"floor,attr",yaml:"floor,attr"`
}

type NetworkBandwidth struct {
	Inbound  *NetworkBandwidthParams `xml:"inbound",json:"inbound",yaml:"inbound"`
	Outbound *NetworkBandwidthParams `xml:"outbound",json:"outbound",yaml:"outbound"`
}

func (a *NetworkVirtualPortParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

func (a *NetworkVirtualPortParams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

type networkVirtualPort NetworkVirtualPort

func (a *NetworkVirtualPort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	vp := networkVirtualPort(*a)
	return e.EncodeElement(&vp, start)
}

func (a *NetworkVirtualPort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	a.Params = &NetworkVirtualPortParams{}
	if !ok {
		var any NetworkVirtualPortParamsAny
		a.Params.Any = &any
	} else if typ == "802.1Qbg" {
		var vepa NetworkVirtualPortParamsVEPA8021QBG
		a.Params.VEPA8021QBG = &vepa
	} else if typ == "802.1Qbh" {
		var vntag NetworkVirtualPortParamsVNTag8021QBH
		a.Params.VNTag8011QBH = &vntag
	} else if typ == "openvswitch" {
		var ovs NetworkVirtualPortParamsOpenVSwitch
		a.Params.OpenVSwitch = &ovs
	} else if typ == "midonet" {
		var mido NetworkVirtualPortParamsMidoNet
		a.Params.MidoNet = &mido
	}

	vp := networkVirtualPort(*a)
	err := d.DecodeElement(&vp, &start)
	if err != nil {
		return err
	}
	*a = NetworkVirtualPort(vp)
	return nil
}

func (a *NetworkForwardAddressPCI) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *NetworkForwardAddressPCI) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (a *NetworkForwardAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(a.PCI, start)
	} else {
		return nil
	}
}

func (a *NetworkForwardAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	if typ == "pci" {
		a.PCI = &NetworkForwardAddressPCI{}
		return d.DecodeElement(a.PCI, &start)
	}

	return nil
}

func (s *NetworkDHCPHost) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPHost) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSHost) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSHost) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkPortGroup) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkPortGroup) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSTXT) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSTXT) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDNSSRV) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDNSSRV) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkDHCPRange) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkDHCPRange) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *NetworkForwardInterface) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NetworkForwardInterface) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}

func (s *Network) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *Network) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
