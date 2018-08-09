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
	Name            string `xml:"name,attr,omitempty" json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	STP             string `xml:"stp,attr,omitempty" json:"stp,omitempty,omitempty" yaml:"stp,omitempty,omitempty"`
	Delay           string `xml:"delay,attr,omitempty" json:"delay,omitempty,omitempty" yaml:"delay,omitempty,omitempty"`
	MACTableManager string `xml:"macTableManager,attr,omitempty" json:"macTableManager,omitempty,omitempty" yaml:"macTableManager,omitempty,omitempty"`
}

type NetworkVirtualPort struct {
	Params *NetworkVirtualPortParams `xml:"parameters" json:"parameters" yaml:"parameters"`
}

type NetworkVirtualPortParams struct {
	Any          *NetworkVirtualPortParamsAny          `xml:"-" json:"-" yaml:"-"`
	VEPA8021QBG  *NetworkVirtualPortParamsVEPA8021QBG  `xml:"-" json:"-" yaml:"-"`
	VNTag8011QBH *NetworkVirtualPortParamsVNTag8021QBH `xml:"-" json:"-" yaml:"-"`
	OpenVSwitch  *NetworkVirtualPortParamsOpenVSwitch  `xml:"-" json:"-" yaml:"-"`
	MidoNet      *NetworkVirtualPortParamsMidoNet      `xml:"-" json:"-" yaml:"-"`
}

type NetworkVirtualPortParamsAny struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerid,omitempty" yaml:"managerid,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeid,omitempty" yaml:"typeid,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeidversion,omitempty" yaml:"typeidversion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty,omitempty" yaml:"instanceid,omitempty,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty" json:"profileid,omitempty,omitempty" yaml:"profileid,omitempty,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty,omitempty" yaml:"interfaceid,omitempty,omitempty"`
}

type NetworkVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *uint  `xml:"managerid,attr" json:"managerid,omitempty" yaml:"managerid,omitempty"`
	TypeID        *uint  `xml:"typeid,attr" json:"typeid,omitempty" yaml:"typeid,omitempty"`
	TypeIDVersion *uint  `xml:"typeidversion,attr" json:"typeidversion,omitempty" yaml:"typeidversion,omitempty"`
	InstanceID    string `xml:"instanceid,attr,omitempty" json:"instanceid,omitempty,omitempty" yaml:"instanceid,omitempty,omitempty"`
}

type NetworkVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty" json:"profileid,omitempty,omitempty" yaml:"profileid,omitempty,omitempty"`
}

type NetworkVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty,omitempty" yaml:"interfaceid,omitempty,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty" json:"profileid,omitempty,omitempty" yaml:"profileid,omitempty,omitempty"`
}

type NetworkVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty" json:"interfaceid,omitempty,omitempty" yaml:"interfaceid,omitempty,omitempty"`
}

type NetworkDomain struct {
	Name      string `xml:"name,attr,omitempty" json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	LocalOnly string `xml:"localOnly,attr,omitempty" json:"localOnly,omitempty,omitempty" yaml:"localOnly,omitempty,omitempty"`
}

type NetworkForwardNATAddress struct {
	Start string `xml:"start,attr" json:"start,omitempty" yaml:"start,omitempty"`
	End   string `xml:"end,attr" json:"end,omitempty" yaml:"end,omitempty"`
}

type NetworkForwardNATPort struct {
	Start uint `xml:"start,attr" json:"start,omitempty" yaml:"start,omitempty"`
	End   uint `xml:"end,attr" json:"end,omitempty" yaml:"end,omitempty"`
}

type NetworkForwardNAT struct {
	Addresses []NetworkForwardNATAddress `xml:"address" json:"address" yaml:"address"`
	Ports     []NetworkForwardNATPort    `xml:"port" json:"port" yaml:"port"`
}

type NetworkForward struct {
	Mode       string                    `xml:"mode,attr,omitempty" json:"mode,omitempty,omitempty" yaml:"mode,omitempty,omitempty"`
	Dev        string                    `xml:"dev,attr,omitempty" json:"dev,omitempty,omitempty" yaml:"dev,omitempty,omitempty"`
	Managed    string                    `xml:"managed,attr,omitempty" json:"managed,omitempty,omitempty" yaml:"managed,omitempty,omitempty"`
	Driver     *NetworkForwardDriver     `xml:"driver" json:"driver" yaml:"driver"`
	PFs        []NetworkForwardPF        `xml:"pf" json:"pf" yaml:"pf"`
	NAT        *NetworkForwardNAT        `xml:"nat" json:"nat" yaml:"nat"`
	Interfaces []NetworkForwardInterface `xml:"interface" json:"interface" yaml:"interface"`
	Addresses  []NetworkForwardAddress   `xml:"address" json:"address" yaml:"address"`
}

type NetworkForwardDriver struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type NetworkForwardPF struct {
	Dev string `xml:"dev,attr" json:"dev,omitempty" yaml:"dev,omitempty"`
}

type NetworkForwardAddress struct {
	PCI *NetworkForwardAddressPCI `xml:"-" json:"-" yaml:"-"`
}

type NetworkForwardAddressPCI struct {
	Domain   *uint `xml:"domain,attr" json:"domain,omitempty" yaml:"domain,omitempty"`
	Bus      *uint `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Slot     *uint `xml:"slot,attr" json:"slot,omitempty" yaml:"slot,omitempty"`
	Function *uint `xml:"function,attr" json:"function,omitempty" yaml:"function,omitempty"`
}

type NetworkForwardInterface struct {
	XMLName xml.Name `xml:"interface" json:"interface" yaml:"interface"`
	Dev     string   `xml:"dev,attr,omitempty" json:"dev,omitempty,omitempty" yaml:"dev,omitempty,omitempty"`
}

type NetworkMAC struct {
	Address string `xml:"address,attr,omitempty" json:"address,omitempty,omitempty" yaml:"address,omitempty,omitempty"`
}

type NetworkDHCPRange struct {
	XMLName xml.Name `xml:"range" json:"range" yaml:"range"`
	Start   string   `xml:"start,attr,omitempty" json:"start,omitempty,omitempty" yaml:"start,omitempty,omitempty"`
	End     string   `xml:"end,attr,omitempty" json:"end,omitempty,omitempty" yaml:"end,omitempty,omitempty"`
}

type NetworkDHCPHost struct {
	XMLName xml.Name `xml:"host" json:"host" yaml:"host"`
	ID      string   `xml:"id,attr,omitempty" json:"id,omitempty,omitempty" yaml:"id,omitempty,omitempty"`
	MAC     string   `xml:"mac,attr,omitempty" json:"mac,omitempty,omitempty" yaml:"mac,omitempty,omitempty"`
	Name    string   `xml:"name,attr,omitempty" json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	IP      string   `xml:"ip,attr,omitempty" json:"ip,omitempty,omitempty" yaml:"ip,omitempty,omitempty"`
}

type NetworkBootp struct {
	File   string `xml:"file,attr,omitempty" json:"file,omitempty,omitempty" yaml:"file,omitempty,omitempty"`
	Server string `xml:"server,attr,omitempty" json:"server,omitempty,omitempty" yaml:"server,omitempty,omitempty"`
}

type NetworkDHCP struct {
	Ranges []NetworkDHCPRange `xml:"range" json:"range" yaml:"range"`
	Hosts  []NetworkDHCPHost  `xml:"host" json:"host" yaml:"host"`
	Bootp  []NetworkBootp     `xml:"bootp" json:"bootp" yaml:"bootp"`
}

type NetworkIP struct {
	Address  string       `xml:"address,attr,omitempty" json:"address,omitempty,omitempty" yaml:"address,omitempty,omitempty"`
	Family   string       `xml:"family,attr,omitempty" json:"family,omitempty,omitempty" yaml:"family,omitempty,omitempty"`
	Netmask  string       `xml:"netmask,attr,omitempty" json:"netmask,omitempty,omitempty" yaml:"netmask,omitempty,omitempty"`
	Prefix   uint         `xml:"prefix,attr,omitempty" json:"prefix,omitempty,omitempty" yaml:"prefix,omitempty,omitempty"`
	LocalPtr string       `xml:"localPtr,attr,omitempty" json:"localPtr,omitempty,omitempty" yaml:"localPtr,omitempty,omitempty"`
	DHCP     *NetworkDHCP `xml:"dhcp" json:"dhcp" yaml:"dhcp"`
	TFTP     *NetworkTFTP `xml:"tftp" json:"tftp" yaml:"tftp"`
}

type NetworkTFTP struct {
	Root string `xml:"root,attr,omitempty" json:"root,omitempty,omitempty" yaml:"root,omitempty,omitempty"`
}

type NetworkRoute struct {
	Family  string `xml:"family,attr,omitempty" json:"family,omitempty,omitempty" yaml:"family,omitempty,omitempty"`
	Address string `xml:"address,attr,omitempty" json:"address,omitempty,omitempty" yaml:"address,omitempty,omitempty"`
	Netmask string `xml:"netmask,attr,omitempty" json:"netmask,omitempty,omitempty" yaml:"netmask,omitempty,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty" json:"prefix,omitempty,omitempty" yaml:"prefix,omitempty,omitempty"`
	Gateway string `xml:"gateway,attr,omitempty" json:"gateway,omitempty,omitempty" yaml:"gateway,omitempty,omitempty"`
	Metric  string `xml:"metric,attr,omitempty" json:"metric,omitempty,omitempty" yaml:"metric,omitempty,omitempty"`
}

type NetworkDNSForwarder struct {
	Domain string `xml:"domain,attr,omitempty" json:"domain,omitempty,omitempty" yaml:"domain,omitempty,omitempty"`
	Addr   string `xml:"addr,attr,omitempty" json:"addr,omitempty,omitempty" yaml:"addr,omitempty,omitempty"`
}

type NetworkDNSTXT struct {
	XMLName xml.Name `xml:"txt" json:"txt" yaml:"txt"`
	Name    string   `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value   string   `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type NetworkDNSHostHostname struct {
	Hostname string `xml:",attr" json:"" yaml:""`
}

type NetworkDNSHost struct {
	XMLName   xml.Name                 `xml:"host" json:"host" yaml:"host"`
	IP        string                   `xml:"ip,attr" json:"ip,omitempty" yaml:"ip,omitempty"`
	Hostnames []NetworkDNSHostHostname `xml:"hostname" json:"hostname" yaml:"hostname"`
}

type NetworkDNSSRV struct {
	XMLName  xml.Name `xml:"srv" json:"srv" yaml:"srv"`
	Service  string   `xml:"service,attr,omitempty" json:"service,omitempty,omitempty" yaml:"service,omitempty,omitempty"`
	Protocol string   `xml:"protocol,attr,omitempty" json:"protocol,omitempty,omitempty" yaml:"protocol,omitempty,omitempty"`
	Target   string   `xml:"target,attr,omitempty" json:"target,omitempty,omitempty" yaml:"target,omitempty,omitempty"`
	Port     uint     `xml:"port,attr,omitempty" json:"port,omitempty,omitempty" yaml:"port,omitempty,omitempty"`
	Priority uint     `xml:"priority,attr,omitempty" json:"priority,omitempty,omitempty" yaml:"priority,omitempty,omitempty"`
	Weight   uint     `xml:"weight,attr,omitempty" json:"weight,omitempty,omitempty" yaml:"weight,omitempty,omitempty"`
	Domain   string   `xml:"domain,attr,omitempty" json:"domain,omitempty,omitempty" yaml:"domain,omitempty,omitempty"`
}

type NetworkDNS struct {
	Enable            string                `xml:"enable,attr,omitempty" json:"enable,omitempty,omitempty" yaml:"enable,omitempty,omitempty"`
	ForwardPlainNames string                `xml:"forwardPlainNames,attr,omitempty" json:"forwardPlainNames,omitempty,omitempty" yaml:"forwardPlainNames,omitempty,omitempty"`
	Forwarders        []NetworkDNSForwarder `xml:"forwarder" json:"forwarder" yaml:"forwarder"`
	TXTs              []NetworkDNSTXT       `xml:"txt" json:"txt" yaml:"txt"`
	Host              []NetworkDNSHost      `xml:"host" json:"host" yaml:"host"`
	SRVs              []NetworkDNSSRV       `xml:"srv" json:"srv" yaml:"srv"`
}

type NetworkMetadata struct {
	XML string `xml:",innerxml" json:",innerxml" yaml:",innerxml"`
}

type NetworkMTU struct {
	Size uint `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
}

type Network struct {
	XMLName             xml.Name            `xml:"network" json:"network" yaml:"network"`
	IPv6                string              `xml:"ipv6,attr,omitempty" json:"ipv6,omitempty,omitempty" yaml:"ipv6,omitempty,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty,omitempty" yaml:"trustGuestRxFilters,omitempty,omitempty"`
	Name                string              `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	UUID                string              `xml:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Metadata            *NetworkMetadata    `xml:"metadata" json:"metadata" yaml:"metadata"`
	Forward             *NetworkForward     `xml:"forward" json:"forward" yaml:"forward"`
	Bridge              *NetworkBridge      `xml:"bridge" json:"bridge" yaml:"bridge"`
	MTU                 *NetworkMTU         `xml:"mtu" json:"mtu" yaml:"mtu"`
	MAC                 *NetworkMAC         `xml:"mac" json:"mac" yaml:"mac"`
	Domain              *NetworkDomain      `xml:"domain" json:"domain" yaml:"domain"`
	DNS                 *NetworkDNS         `xml:"dns" json:"dns" yaml:"dns"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan" yaml:"vlan"`
	Bandwidth           *NetworkBandwidth   `xml:"bandwidth" json:"bandwidth" yaml:"bandwidth"`
	IPs                 []NetworkIP         `xml:"ip" json:"ip" yaml:"ip"`
	Routes              []NetworkRoute      `xml:"route" json:"route" yaml:"route"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualport" yaml:"virtualport"`
	PortGroups          []NetworkPortGroup  `xml:"portgroup" json:"portgroup" yaml:"portgroup"`
}

type NetworkPortGroup struct {
	XMLName             xml.Name            `xml:"portgroup" json:"portgroup" yaml:"portgroup"`
	Name                string              `xml:"name,attr,omitempty" json:"name,omitempty,omitempty" yaml:"name,omitempty,omitempty"`
	Default             string              `xml:"default,attr,omitempty" json:"default,omitempty,omitempty" yaml:"default,omitempty,omitempty"`
	TrustGuestRxFilters string              `xml:"trustGuestRxFilters,attr,omitempty" json:"trustGuestRxFilters,omitempty,omitempty" yaml:"trustGuestRxFilters,omitempty,omitempty"`
	VLAN                *NetworkVLAN        `xml:"vlan" json:"vlan" yaml:"vlan"`
	VirtualPort         *NetworkVirtualPort `xml:"virtualport" json:"virtualport" yaml:"virtualport"`
}

type NetworkVLAN struct {
	Trunk string           `xml:"trunk,attr,omitempty" json:"trunk,omitempty,omitempty" yaml:"trunk,omitempty,omitempty"`
	Tags  []NetworkVLANTag `xml:"tag" json:"tag" yaml:"tag"`
}

type NetworkVLANTag struct {
	ID         uint   `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	NativeMode string `xml:"nativeMode,attr,omitempty" json:"nativeMode,omitempty,omitempty" yaml:"nativeMode,omitempty,omitempty"`
}

type NetworkBandwidthParams struct {
	Average *uint `xml:"average,attr" json:"average,omitempty" yaml:"average,omitempty"`
	Peak    *uint `xml:"peak,attr" json:"peak,omitempty" yaml:"peak,omitempty"`
	Burst   *uint `xml:"burst,attr" json:"burst,omitempty" yaml:"burst,omitempty"`
	Floor   *uint `xml:"floor,attr" json:"floor,omitempty" yaml:"floor,omitempty"`
}

type NetworkBandwidth struct {
	Inbound  *NetworkBandwidthParams `xml:"inbound" json:"inbound" yaml:"inbound"`
	Outbound *NetworkBandwidthParams `xml:"outbound" json:"outbound" yaml:"outbound"`
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
