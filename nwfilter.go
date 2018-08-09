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
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NWFilter struct {
	XMLName  xml.Name `xml:"filter" json:"filter" yaml:"filter"`
	Name     string   `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	UUID     string   `xml:"uuid,omitempty" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Chain    string   `xml:"chain,attr,omitempty" json:"chain,omitempty" yaml:"chain,omitempty"`
	Priority int      `xml:"priority,attr,omitempty" json:"priority,omitempty" yaml:"priority,omitempty"`
	Entries  []NWFilterEntry
}

type NWFilterEntry struct {
	Rule *NWFilterRule
	Ref  *NWFilterRef
}

type NWFilterRef struct {
	Filter     string              `xml:"filter,attr" json:"filter,omitempty" yaml:"filter,omitempty"`
	Parameters []NWFilterParameter `xml:"parameter" json:"parameter,omitempty" yaml:"parameter,omitempty"`
}

type NWFilterParameter struct {
	Name  string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Value string `xml:"value,attr" json:"value,omitempty" yaml:"value,omitempty"`
}

type NWFilterField struct {
	Var  string
	Str  string
	Uint *uint
}

type NWFilterRule struct {
	Action     string `xml:"action,attr,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	Direction  string `xml:"direction,attr,omitempty" json:"direction,omitempty" yaml:"direction,omitempty"`
	Priority   int    `xml:"priority,attr,omitempty" json:"priority,omitempty" yaml:"priority,omitempty"`
	StateMatch string `xml:"statematch,attr,omitempty" json:"statematch,omitempty" yaml:"statematch,omitempty"`

	ARP         *NWFilterRuleARP         `xml:"arp" json:"arp,omitempty" yaml:"arp,omitempty"`
	RARP        *NWFilterRuleRARP        `xml:"rarp" json:"rarp,omitempty" yaml:"rarp,omitempty"`
	MAC         *NWFilterRuleMAC         `xml:"mac" json:"mac,omitempty" yaml:"mac,omitempty"`
	VLAN        *NWFilterRuleVLAN        `xml:"vlan" json:"vlan,omitempty" yaml:"vlan,omitempty"`
	STP         *NWFilterRuleSTP         `xml:"stp" json:"stp,omitempty" yaml:"stp,omitempty"`
	IP          *NWFilterRuleIP          `xml:"ip" json:"ip,omitempty" yaml:"ip,omitempty"`
	IPv6        *NWFilterRuleIPv6        `xml:"ipv6" json:"ipv6" yaml:"ipv6"`
	TCP         *NWFilterRuleTCP         `xml:"tcp" json:"tcp,omitempty" yaml:"tcp,omitempty"`
	UDP         *NWFilterRuleUDP         `xml:"udp" json:"udp,omitempty" yaml:"udp,omitempty"`
	UDPLite     *NWFilterRuleUDPLite     `xml:"udplite" json:"udplite,omitempty" yaml:"udplite,omitempty"`
	ESP         *NWFilterRuleESP         `xml:"esp" json:"esp,omitempty" yaml:"esp,omitempty"`
	AH          *NWFilterRuleAH          `xml:"ah" json:"ah,omitempty" yaml:"ah,omitempty"`
	SCTP        *NWFilterRuleSCTP        `xml:"sctp" json:"sctp,omitempty" yaml:"sctp,omitempty"`
	ICMP        *NWFilterRuleICMP        `xml:"icmp" json:"icmp,omitempty" yaml:"icmp,omitempty"`
	All         *NWFilterRuleAll         `xml:"all" json:"all,omitempty" yaml:"all,omitempty"`
	IGMP        *NWFilterRuleIGMP        `xml:"igmp" json:"igmp,omitempty" yaml:"igmp,omitempty"`
	TCPIPv6     *NWFilterRuleTCPIPv6     `xml:"tcp-ipv6" json:"tcp-ipv6" yaml:"tcp-ipv6"`
	UDPIPv6     *NWFilterRuleUDPIPv6     `xml:"udp-ipv6" json:"udp-ipv6" yaml:"udp-ipv6"`
	UDPLiteIPv6 *NWFilterRuleUDPLiteIPv6 `xml:"udplite-ipv6" json:"udplite-ipv6" yaml:"udplite-ipv6"`
	ESPIPv6     *NWFilterRuleESPIPv6     `xml:"esp-ipv6" json:"esp-ipv6" yaml:"esp-ipv6"`
	AHIPv6      *NWFilterRuleAHIPv6      `xml:"ah-ipv6" json:"ah-ipv6" yaml:"ah-ipv6"`
	SCTPIPv6    *NWFilterRuleSCTPIPv6    `xml:"sctp-ipv6" json:"sctp-ipv6" yaml:"sctp-ipv6"`
	ICMPv6      *NWFilterRuleICMPIPv6    `xml:"icmpv6" json:"icmpv6" yaml:"icmpv6"`
	AllIPv6     *NWFilterRuleAllIPv6     `xml:"all-ipv6" json:"all-ipv6" yaml:"all-ipv6"`
}

type NWFilterRuleCommonMAC struct {
	SrcMACAddr NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcmacaddr,omitempty" yaml:"srcmacaddr,omitempty"`
	SrcMACMask NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcmacmask,omitempty" yaml:"srcmacmask,omitempty"`
	DstMACAddr NWFilterField `xml:"dstmacaddr,attr,omitempty" json:"dstmacaddr,omitempty" yaml:"dstmacaddr,omitempty"`
	DstMACMask NWFilterField `xml:"dstmacmask,attr,omitempty" json:"dstmacmask,omitempty" yaml:"dstmacmask,omitempty"`
}

type NWFilterRuleCommonIP struct {
	SrcMACAddr     NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcmacaddr,omitempty" yaml:"srcmacaddr,omitempty"`
	SrcIPAddr      NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcipaddr,omitempty" yaml:"srcipaddr,omitempty"`
	SrcIPMask      NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcipmask,omitempty" yaml:"srcipmask,omitempty"`
	DstIPAddr      NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstipaddr,omitempty" yaml:"dstipaddr,omitempty"`
	DstIPMask      NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstipmask,omitempty" yaml:"dstipmask,omitempty"`
	SrcIPFrom      NWFilterField `xml:"srcipfrom,attr,omitempty" json:"srcipfrom,omitempty" yaml:"srcipfrom,omitempty"`
	SrcIPTo        NWFilterField `xml:"srcipto,attr,omitempty" json:"srcipto,omitempty" yaml:"srcipto,omitempty"`
	DstIPFrom      NWFilterField `xml:"dstipfrom,attr,omitempty" json:"dstipfrom,omitempty" yaml:"dstipfrom,omitempty"`
	DstIPTo        NWFilterField `xml:"dstipto,attr,omitempty" json:"dstipto,omitempty" yaml:"dstipto,omitempty"`
	DSCP           NWFilterField `xml:"dscp,attr" json:"dscp,omitempty" yaml:"dscp,omitempty"`
	ConnLimitAbove NWFilterField `xml:"connlimit-above,attr" json:"connlimit-above,omitempty" yaml:"connlimit-above,omitempty"`
	State          NWFilterField `xml:"state,attr,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	IPSet          NWFilterField `xml:"ipset,attr,omitempty" json:"ipset,omitempty" yaml:"ipset,omitempty"`
	IPSetFlags     NWFilterField `xml:"ipsetflags,attr,omitempty" json:"ipsetflags,omitempty" yaml:"ipsetflags,omitempty"`
}

type NWFilterRuleCommonPort struct {
	SrcPortStart NWFilterField `xml:"srcportstart,attr" json:"srcportstart,omitempty" yaml:"srcportstart,omitempty"`
	SrcPortEnd   NWFilterField `xml:"srcportend,attr" json:"srcportend,omitempty" yaml:"srcportend,omitempty"`
	DstPortStart NWFilterField `xml:"dstportstart,attr" json:"dstportstart,omitempty" yaml:"dstportstart,omitempty"`
	DstPortEnd   NWFilterField `xml:"dstportend,attr" json:"dstportend,omitempty" yaml:"dstportend,omitempty"`
}

type NWFilterRuleARP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	HWType        NWFilterField `xml:"hwtype,attr" json:"hwtype,omitempty" yaml:"hwtype,omitempty"`
	ProtocolType  NWFilterField `xml:"protocoltype,attr" json:"protocoltype,omitempty" yaml:"protocoltype,omitempty"`
	OpCode        NWFilterField `xml:"opcode,attr,omitempty" json:"opcode,omitempty" yaml:"opcode,omitempty"`
	ARPSrcMACAddr NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpsrcmacaddr,omitempty" yaml:"arpsrcmacaddr,omitempty"`
	ARPDstMACAddr NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpdstmacaddr,omitempty" yaml:"arpdstmacaddr,omitempty"`
	ARPSrcIPAddr  NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpsrcipaddr,omitempty" yaml:"arpsrcipaddr,omitempty"`
	ARPSrcIPMask  NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpsrcipmask,omitempty" yaml:"arpsrcipmask,omitempty"`
	ARPDstIPAddr  NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpdstipaddr,omitempty" yaml:"arpdstipaddr,omitempty"`
	ARPDstIPMask  NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpdstipmask,omitempty" yaml:"arpdstipmask,omitempty"`
	Gratuitous    NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty" yaml:"gratuitous,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleRARP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	HWType        NWFilterField `xml:"hwtype,attr" json:"hwtype,omitempty" yaml:"hwtype,omitempty"`
	ProtocolType  NWFilterField `xml:"protocoltype,attr" json:"protocoltype,omitempty" yaml:"protocoltype,omitempty"`
	OpCode        NWFilterField `xml:"opcode,attr,omitempty" json:"opcode,omitempty" yaml:"opcode,omitempty"`
	ARPSrcMACAddr NWFilterField `xml:"arpsrcmacaddr,attr,omitempty" json:"arpsrcmacaddr,omitempty" yaml:"arpsrcmacaddr,omitempty"`
	ARPDstMACAddr NWFilterField `xml:"arpdstmacaddr,attr,omitempty" json:"arpdstmacaddr,omitempty" yaml:"arpdstmacaddr,omitempty"`
	ARPSrcIPAddr  NWFilterField `xml:"arpsrcipaddr,attr,omitempty" json:"arpsrcipaddr,omitempty" yaml:"arpsrcipaddr,omitempty"`
	ARPSrcIPMask  NWFilterField `xml:"arpsrcipmask,attr,omitempty" json:"arpsrcipmask,omitempty" yaml:"arpsrcipmask,omitempty"`
	ARPDstIPAddr  NWFilterField `xml:"arpdstipaddr,attr,omitempty" json:"arpdstipaddr,omitempty" yaml:"arpdstipaddr,omitempty"`
	ARPDstIPMask  NWFilterField `xml:"arpdstipmask,attr,omitempty" json:"arpdstipmask,omitempty" yaml:"arpdstipmask,omitempty"`
	Gratuitous    NWFilterField `xml:"gratuitous,attr,omitempty" json:"gratuitous,omitempty" yaml:"gratuitous,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleMAC struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	ProtocolID NWFilterField `xml:"protocolid,attr,omitempty" json:"protocolid,omitempty" yaml:"protocolid,omitempty"`
	Comment    string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleVLAN struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	VLANID        NWFilterField `xml:"vlanid,attr,omitempty" json:"vlanid,omitempty" yaml:"vlanid,omitempty"`
	EncapProtocol NWFilterField `xml:"encap-protocol,attr,omitempty" json:"encap-protocol,omitempty" yaml:"encap-protocol,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleSTP struct {
	Match             NWFilterField `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	SrcMACAddr        NWFilterField `xml:"srcmacaddr,attr,omitempty" json:"srcmacaddr,omitempty" yaml:"srcmacaddr,omitempty"`
	SrcMACMask        NWFilterField `xml:"srcmacmask,attr,omitempty" json:"srcmacmask,omitempty" yaml:"srcmacmask,omitempty"`
	Type              NWFilterField `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Flags             NWFilterField `xml:"flags,attr" json:"flags,omitempty" yaml:"flags,omitempty"`
	RootPriority      NWFilterField `xml:"root-priority,attr" json:"root-priority,omitempty" yaml:"root-priority,omitempty"`
	RootPriorityHi    NWFilterField `xml:"root-priority-hi,attr" json:"root-priority-hi,omitempty" yaml:"root-priority-hi,omitempty"`
	RootAddress       NWFilterField `xml:"root-address,attr,omitempty" json:"root-address,omitempty" yaml:"root-address,omitempty"`
	RootAddressMask   NWFilterField `xml:"root-address-mask,attr,omitempty" json:"root-address-mask,omitempty" yaml:"root-address-mask,omitempty"`
	RootCost          NWFilterField `xml:"root-cost,attr" json:"root-cost,omitempty" yaml:"root-cost,omitempty"`
	RootCostHi        NWFilterField `xml:"root-cost-hi,attr" json:"root-cost-hi,omitempty" yaml:"root-cost-hi,omitempty"`
	SenderPriority    NWFilterField `xml:"sender-priority,attr" json:"sender-priority,omitempty" yaml:"sender-priority,omitempty"`
	SenderPriorityHi  NWFilterField `xml:"sender-priority-hi,attr" json:"sender-priority-hi,omitempty" yaml:"sender-priority-hi,omitempty"`
	SenderAddress     NWFilterField `xml:"sender-address,attr,omitempty" json:"sender-address,omitempty" yaml:"sender-address,omitempty"`
	SenderAddressMask NWFilterField `xml:"sender-address-mask,attr,omitempty" json:"sender-address-mask,omitempty" yaml:"sender-address-mask,omitempty"`
	Port              NWFilterField `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
	PortHi            NWFilterField `xml:"port-hi,attr" json:"port-hi,omitempty" yaml:"port-hi,omitempty"`
	Age               NWFilterField `xml:"age,attr" json:"age,omitempty" yaml:"age,omitempty"`
	AgeHi             NWFilterField `xml:"age-hi,attr" json:"age-hi,omitempty" yaml:"age-hi,omitempty"`
	MaxAge            NWFilterField `xml:"max-age,attr" json:"max-age,omitempty" yaml:"max-age,omitempty"`
	MaxAgeHi          NWFilterField `xml:"max-age-hi,attr" json:"max-age-hi,omitempty" yaml:"max-age-hi,omitempty"`
	HelloTime         NWFilterField `xml:"hello-time,attr" json:"hello-time,omitempty" yaml:"hello-time,omitempty"`
	HelloTimeHi       NWFilterField `xml:"hello-time-hi,attr" json:"hello-time-hi,omitempty" yaml:"hello-time-hi,omitempty"`
	ForwardDelay      NWFilterField `xml:"forward-delay,attr" json:"forward-delay,omitempty" yaml:"forward-delay,omitempty"`
	ForwardDelayHi    NWFilterField `xml:"forward-delay-hi,attr" json:"forward-delay-hi,omitempty" yaml:"forward-delay-hi,omitempty"`
	Comment           string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleIP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcipaddr,omitempty" yaml:"srcipaddr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcipmask,omitempty" yaml:"srcipmask,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstipaddr,omitempty" yaml:"dstipaddr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstipmask,omitempty" yaml:"dstipmask,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty" yaml:"protocol,omitempty"`
	NWFilterRuleCommonPort
	DSCP    NWFilterField `xml:"dscp,attr" json:"dscp,omitempty" yaml:"dscp,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonMAC
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty" json:"srcipaddr,omitempty" yaml:"srcipaddr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty" json:"srcipmask,omitempty" yaml:"srcipmask,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty" json:"dstipaddr,omitempty" yaml:"dstipaddr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty" json:"dstipmask,omitempty" yaml:"dstipmask,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty" json:"protocol,omitempty" yaml:"protocol,omitempty"`
	NWFilterRuleCommonPort
	Type    NWFilterField `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	TypeEnd NWFilterField `xml:"typeend,attr" json:"typeend,omitempty" yaml:"typeend,omitempty"`
	Code    NWFilterField `xml:"code,attr" json:"code,omitempty" yaml:"code,omitempty"`
	CodeEnd NWFilterField `xml:"codeend,attr" json:"codeend,omitempty" yaml:"codeend,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleTCP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Option  NWFilterField `xml:"option,attr" json:"option,omitempty" yaml:"option,omitempty"`
	Flags   NWFilterField `xml:"flags,attr,omitempty" json:"flags,omitempty" yaml:"flags,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleUDP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleUDPLite struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleESP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleAH struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleSCTP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleICMP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Type    NWFilterField `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Code    NWFilterField `xml:"code,attr" json:"code,omitempty" yaml:"code,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleAll struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleIGMP struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleTCPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Option  NWFilterField `xml:"option,attr" json:"option,omitempty" yaml:"option,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleUDPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleUDPLiteIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleESPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleAHIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleSCTPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleICMPIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Type    NWFilterField `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	Code    NWFilterField `xml:"code,attr" json:"code,omitempty" yaml:"code,omitempty"`
	Comment string        `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

type NWFilterRuleAllIPv6 struct {
	Match string `xml:"match,attr,omitempty" json:"match,omitempty" yaml:"match,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

func (s *NWFilterField) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s == nil {
		return xml.Attr{}, nil
	}
	if s.Str != "" {
		return xml.Attr{
			Name:  name,
			Value: s.Str,
		}, nil
	} else if s.Var != "" {
		return xml.Attr{
			Name:  name,
			Value: "$" + s.Str,
		}, nil
	} else if s.Uint != nil {
		return xml.Attr{
			Name:  name,
			Value: fmt.Sprintf("0x%x", *s.Uint),
		}, nil
	} else {
		return xml.Attr{}, nil
	}
}

func (s *NWFilterField) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}
	if attr.Value[0] == '$' {
		s.Var = attr.Value[1:]
	}
	if strings.HasPrefix(attr.Value, "0x") {
		val, err := strconv.ParseUint(attr.Value[2:], 16, 64)
		if err != nil {
			return err
		}
		uval := uint(val)
		s.Uint = &uval
	}
	s.Str = attr.Value
	return nil
}

func (a *NWFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "filter"
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "name"},
		Value: a.Name,
	})
	if a.Chain != "" {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "chain"},
			Value: a.Chain,
		})
	}
	if a.Priority != 0 {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "priority"},
			Value: fmt.Sprintf("%d", a.Priority),
		})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	if a.UUID != "" {
		uuid := xml.StartElement{
			Name: xml.Name{Local: "uuid"},
		}
		e.EncodeToken(uuid)
		e.EncodeToken(xml.CharData(a.UUID))
		e.EncodeToken(uuid.End())
	}

	for _, entry := range a.Entries {
		if entry.Rule != nil {
			rule := xml.StartElement{
				Name: xml.Name{Local: "rule"},
			}
			e.EncodeElement(entry.Rule, rule)
		} else if entry.Ref != nil {
			ref := xml.StartElement{
				Name: xml.Name{Local: "filterref"},
			}
			e.EncodeElement(entry.Ref, ref)
		}
	}

	err = e.EncodeToken(start.End())
	if err != nil {
		return err
	}
	return nil
}

func (a *NWFilter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	name, ok := getAttr(start.Attr, "name")
	if !ok {
		return fmt.Errorf("Missing filter name")
	}
	a.Name = name
	a.Chain, _ = getAttr(start.Attr, "chain")
	prio, ok := getAttr(start.Attr, "priority")
	if ok {
		val, err := strconv.ParseInt(prio, 10, 64)
		if err != nil {
			return err
		}
		a.Priority = int(val)
	}

	for {
		tok, err := d.Token()
		if err == io.EOF {
			break
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			{
				if tok.Name.Local == "uuid" {
					txt, err := d.Token()
					if err != nil {
						return err
					}

					txt2, ok := txt.(xml.CharData)
					if !ok {
						return fmt.Errorf("Expected UUID string")
					}
					a.UUID = string(txt2)
				} else if tok.Name.Local == "rule" {
					entry := NWFilterEntry{
						Rule: &NWFilterRule{},
					}

					d.DecodeElement(entry.Rule, &tok)

					a.Entries = append(a.Entries, entry)
				} else if tok.Name.Local == "filterref" {
					entry := NWFilterEntry{
						Ref: &NWFilterRef{},
					}

					d.DecodeElement(entry.Ref, &tok)

					a.Entries = append(a.Entries, entry)
				}
			}
		}

	}
	return nil
}

func (s *NWFilter) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), s)
}

func (s *NWFilter) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
