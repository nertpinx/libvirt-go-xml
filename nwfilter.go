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
	XMLName  xml.Name `xml:"filter",json:"filter",yaml:"filter"`
	Name     string   `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	UUID     string   `xml:"uuid,omitempty",json:"uuid,omitempty",yaml:"uuid,omitempty"`
	Chain    string   `xml:"chain,attr,omitempty",json:"chain,attr,omitempty",yaml:"chain,attr,omitempty"`
	Priority int      `xml:"priority,attr,omitempty",json:"priority,attr,omitempty",yaml:"priority,attr,omitempty"`
	Entries  []NWFilterEntry
}

type NWFilterEntry struct {
	Rule *NWFilterRule
	Ref  *NWFilterRef
}

type NWFilterRef struct {
	Filter     string              `xml:"filter,attr",json:"filter,attr",yaml:"filter,attr"`
	Parameters []NWFilterParameter `xml:"parameter",json:"parameter",yaml:"parameter"`
}

type NWFilterParameter struct {
	Name  string `xml:"name,attr",json:"name,attr",yaml:"name,attr"`
	Value string `xml:"value,attr",json:"value,attr",yaml:"value,attr"`
}

type NWFilterField struct {
	Var  string
	Str  string
	Uint *uint
}

type NWFilterRule struct {
	Action     string `xml:"action,attr,omitempty",json:"action,attr,omitempty",yaml:"action,attr,omitempty"`
	Direction  string `xml:"direction,attr,omitempty",json:"direction,attr,omitempty",yaml:"direction,attr,omitempty"`
	Priority   int    `xml:"priority,attr,omitempty",json:"priority,attr,omitempty",yaml:"priority,attr,omitempty"`
	StateMatch string `xml:"statematch,attr,omitempty",json:"statematch,attr,omitempty",yaml:"statematch,attr,omitempty"`

	ARP         *NWFilterRuleARP         `xml:"arp",json:"arp",yaml:"arp"`
	RARP        *NWFilterRuleRARP        `xml:"rarp",json:"rarp",yaml:"rarp"`
	MAC         *NWFilterRuleMAC         `xml:"mac",json:"mac",yaml:"mac"`
	VLAN        *NWFilterRuleVLAN        `xml:"vlan",json:"vlan",yaml:"vlan"`
	STP         *NWFilterRuleSTP         `xml:"stp",json:"stp",yaml:"stp"`
	IP          *NWFilterRuleIP          `xml:"ip",json:"ip",yaml:"ip"`
	IPv6        *NWFilterRuleIPv6        `xml:"ipv6",json:"ipv6",yaml:"ipv6"`
	TCP         *NWFilterRuleTCP         `xml:"tcp",json:"tcp",yaml:"tcp"`
	UDP         *NWFilterRuleUDP         `xml:"udp",json:"udp",yaml:"udp"`
	UDPLite     *NWFilterRuleUDPLite     `xml:"udplite",json:"udplite",yaml:"udplite"`
	ESP         *NWFilterRuleESP         `xml:"esp",json:"esp",yaml:"esp"`
	AH          *NWFilterRuleAH          `xml:"ah",json:"ah",yaml:"ah"`
	SCTP        *NWFilterRuleSCTP        `xml:"sctp",json:"sctp",yaml:"sctp"`
	ICMP        *NWFilterRuleICMP        `xml:"icmp",json:"icmp",yaml:"icmp"`
	All         *NWFilterRuleAll         `xml:"all",json:"all",yaml:"all"`
	IGMP        *NWFilterRuleIGMP        `xml:"igmp",json:"igmp",yaml:"igmp"`
	TCPIPv6     *NWFilterRuleTCPIPv6     `xml:"tcp-ipv6",json:"tcp-ipv6",yaml:"tcp-ipv6"`
	UDPIPv6     *NWFilterRuleUDPIPv6     `xml:"udp-ipv6",json:"udp-ipv6",yaml:"udp-ipv6"`
	UDPLiteIPv6 *NWFilterRuleUDPLiteIPv6 `xml:"udplite-ipv6",json:"udplite-ipv6",yaml:"udplite-ipv6"`
	ESPIPv6     *NWFilterRuleESPIPv6     `xml:"esp-ipv6",json:"esp-ipv6",yaml:"esp-ipv6"`
	AHIPv6      *NWFilterRuleAHIPv6      `xml:"ah-ipv6",json:"ah-ipv6",yaml:"ah-ipv6"`
	SCTPIPv6    *NWFilterRuleSCTPIPv6    `xml:"sctp-ipv6",json:"sctp-ipv6",yaml:"sctp-ipv6"`
	ICMPv6      *NWFilterRuleICMPIPv6    `xml:"icmpv6",json:"icmpv6",yaml:"icmpv6"`
	AllIPv6     *NWFilterRuleAllIPv6     `xml:"all-ipv6",json:"all-ipv6",yaml:"all-ipv6"`
}

type NWFilterRuleCommonMAC struct {
	SrcMACAddr NWFilterField `xml:"srcmacaddr,attr,omitempty",json:"srcmacaddr,attr,omitempty",yaml:"srcmacaddr,attr,omitempty"`
	SrcMACMask NWFilterField `xml:"srcmacmask,attr,omitempty",json:"srcmacmask,attr,omitempty",yaml:"srcmacmask,attr,omitempty"`
	DstMACAddr NWFilterField `xml:"dstmacaddr,attr,omitempty",json:"dstmacaddr,attr,omitempty",yaml:"dstmacaddr,attr,omitempty"`
	DstMACMask NWFilterField `xml:"dstmacmask,attr,omitempty",json:"dstmacmask,attr,omitempty",yaml:"dstmacmask,attr,omitempty"`
}

type NWFilterRuleCommonIP struct {
	SrcMACAddr     NWFilterField `xml:"srcmacaddr,attr,omitempty",json:"srcmacaddr,attr,omitempty",yaml:"srcmacaddr,attr,omitempty"`
	SrcIPAddr      NWFilterField `xml:"srcipaddr,attr,omitempty",json:"srcipaddr,attr,omitempty",yaml:"srcipaddr,attr,omitempty"`
	SrcIPMask      NWFilterField `xml:"srcipmask,attr,omitempty",json:"srcipmask,attr,omitempty",yaml:"srcipmask,attr,omitempty"`
	DstIPAddr      NWFilterField `xml:"dstipaddr,attr,omitempty",json:"dstipaddr,attr,omitempty",yaml:"dstipaddr,attr,omitempty"`
	DstIPMask      NWFilterField `xml:"dstipmask,attr,omitempty",json:"dstipmask,attr,omitempty",yaml:"dstipmask,attr,omitempty"`
	SrcIPFrom      NWFilterField `xml:"srcipfrom,attr,omitempty",json:"srcipfrom,attr,omitempty",yaml:"srcipfrom,attr,omitempty"`
	SrcIPTo        NWFilterField `xml:"srcipto,attr,omitempty",json:"srcipto,attr,omitempty",yaml:"srcipto,attr,omitempty"`
	DstIPFrom      NWFilterField `xml:"dstipfrom,attr,omitempty",json:"dstipfrom,attr,omitempty",yaml:"dstipfrom,attr,omitempty"`
	DstIPTo        NWFilterField `xml:"dstipto,attr,omitempty",json:"dstipto,attr,omitempty",yaml:"dstipto,attr,omitempty"`
	DSCP           NWFilterField `xml:"dscp,attr",json:"dscp,attr",yaml:"dscp,attr"`
	ConnLimitAbove NWFilterField `xml:"connlimit-above,attr",json:"connlimit-above,attr",yaml:"connlimit-above,attr"`
	State          NWFilterField `xml:"state,attr,omitempty",json:"state,attr,omitempty",yaml:"state,attr,omitempty"`
	IPSet          NWFilterField `xml:"ipset,attr,omitempty",json:"ipset,attr,omitempty",yaml:"ipset,attr,omitempty"`
	IPSetFlags     NWFilterField `xml:"ipsetflags,attr,omitempty",json:"ipsetflags,attr,omitempty",yaml:"ipsetflags,attr,omitempty"`
}

type NWFilterRuleCommonPort struct {
	SrcPortStart NWFilterField `xml:"srcportstart,attr",json:"srcportstart,attr",yaml:"srcportstart,attr"`
	SrcPortEnd   NWFilterField `xml:"srcportend,attr",json:"srcportend,attr",yaml:"srcportend,attr"`
	DstPortStart NWFilterField `xml:"dstportstart,attr",json:"dstportstart,attr",yaml:"dstportstart,attr"`
	DstPortEnd   NWFilterField `xml:"dstportend,attr",json:"dstportend,attr",yaml:"dstportend,attr"`
}

type NWFilterRuleARP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	HWType        NWFilterField `xml:"hwtype,attr",json:"hwtype,attr",yaml:"hwtype,attr"`
	ProtocolType  NWFilterField `xml:"protocoltype,attr",json:"protocoltype,attr",yaml:"protocoltype,attr"`
	OpCode        NWFilterField `xml:"opcode,attr,omitempty",json:"opcode,attr,omitempty",yaml:"opcode,attr,omitempty"`
	ARPSrcMACAddr NWFilterField `xml:"arpsrcmacaddr,attr,omitempty",json:"arpsrcmacaddr,attr,omitempty",yaml:"arpsrcmacaddr,attr,omitempty"`
	ARPDstMACAddr NWFilterField `xml:"arpdstmacaddr,attr,omitempty",json:"arpdstmacaddr,attr,omitempty",yaml:"arpdstmacaddr,attr,omitempty"`
	ARPSrcIPAddr  NWFilterField `xml:"arpsrcipaddr,attr,omitempty",json:"arpsrcipaddr,attr,omitempty",yaml:"arpsrcipaddr,attr,omitempty"`
	ARPSrcIPMask  NWFilterField `xml:"arpsrcipmask,attr,omitempty",json:"arpsrcipmask,attr,omitempty",yaml:"arpsrcipmask,attr,omitempty"`
	ARPDstIPAddr  NWFilterField `xml:"arpdstipaddr,attr,omitempty",json:"arpdstipaddr,attr,omitempty",yaml:"arpdstipaddr,attr,omitempty"`
	ARPDstIPMask  NWFilterField `xml:"arpdstipmask,attr,omitempty",json:"arpdstipmask,attr,omitempty",yaml:"arpdstipmask,attr,omitempty"`
	Gratuitous    NWFilterField `xml:"gratuitous,attr,omitempty",json:"gratuitous,attr,omitempty",yaml:"gratuitous,attr,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleRARP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	HWType        NWFilterField `xml:"hwtype,attr",json:"hwtype,attr",yaml:"hwtype,attr"`
	ProtocolType  NWFilterField `xml:"protocoltype,attr",json:"protocoltype,attr",yaml:"protocoltype,attr"`
	OpCode        NWFilterField `xml:"opcode,attr,omitempty",json:"opcode,attr,omitempty",yaml:"opcode,attr,omitempty"`
	ARPSrcMACAddr NWFilterField `xml:"arpsrcmacaddr,attr,omitempty",json:"arpsrcmacaddr,attr,omitempty",yaml:"arpsrcmacaddr,attr,omitempty"`
	ARPDstMACAddr NWFilterField `xml:"arpdstmacaddr,attr,omitempty",json:"arpdstmacaddr,attr,omitempty",yaml:"arpdstmacaddr,attr,omitempty"`
	ARPSrcIPAddr  NWFilterField `xml:"arpsrcipaddr,attr,omitempty",json:"arpsrcipaddr,attr,omitempty",yaml:"arpsrcipaddr,attr,omitempty"`
	ARPSrcIPMask  NWFilterField `xml:"arpsrcipmask,attr,omitempty",json:"arpsrcipmask,attr,omitempty",yaml:"arpsrcipmask,attr,omitempty"`
	ARPDstIPAddr  NWFilterField `xml:"arpdstipaddr,attr,omitempty",json:"arpdstipaddr,attr,omitempty",yaml:"arpdstipaddr,attr,omitempty"`
	ARPDstIPMask  NWFilterField `xml:"arpdstipmask,attr,omitempty",json:"arpdstipmask,attr,omitempty",yaml:"arpdstipmask,attr,omitempty"`
	Gratuitous    NWFilterField `xml:"gratuitous,attr,omitempty",json:"gratuitous,attr,omitempty",yaml:"gratuitous,attr,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleMAC struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	ProtocolID NWFilterField `xml:"protocolid,attr,omitempty",json:"protocolid,attr,omitempty",yaml:"protocolid,attr,omitempty"`
	Comment    string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleVLAN struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	VLANID        NWFilterField `xml:"vlanid,attr,omitempty",json:"vlanid,attr,omitempty",yaml:"vlanid,attr,omitempty"`
	EncapProtocol NWFilterField `xml:"encap-protocol,attr,omitempty",json:"encap-protocol,attr,omitempty",yaml:"encap-protocol,attr,omitempty"`
	Comment       string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleSTP struct {
	Match             NWFilterField `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	SrcMACAddr        NWFilterField `xml:"srcmacaddr,attr,omitempty",json:"srcmacaddr,attr,omitempty",yaml:"srcmacaddr,attr,omitempty"`
	SrcMACMask        NWFilterField `xml:"srcmacmask,attr,omitempty",json:"srcmacmask,attr,omitempty",yaml:"srcmacmask,attr,omitempty"`
	Type              NWFilterField `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Flags             NWFilterField `xml:"flags,attr",json:"flags,attr",yaml:"flags,attr"`
	RootPriority      NWFilterField `xml:"root-priority,attr",json:"root-priority,attr",yaml:"root-priority,attr"`
	RootPriorityHi    NWFilterField `xml:"root-priority-hi,attr",json:"root-priority-hi,attr",yaml:"root-priority-hi,attr"`
	RootAddress       NWFilterField `xml:"root-address,attr,omitempty",json:"root-address,attr,omitempty",yaml:"root-address,attr,omitempty"`
	RootAddressMask   NWFilterField `xml:"root-address-mask,attr,omitempty",json:"root-address-mask,attr,omitempty",yaml:"root-address-mask,attr,omitempty"`
	RootCost          NWFilterField `xml:"root-cost,attr",json:"root-cost,attr",yaml:"root-cost,attr"`
	RootCostHi        NWFilterField `xml:"root-cost-hi,attr",json:"root-cost-hi,attr",yaml:"root-cost-hi,attr"`
	SenderPriority    NWFilterField `xml:"sender-priority,attr",json:"sender-priority,attr",yaml:"sender-priority,attr"`
	SenderPriorityHi  NWFilterField `xml:"sender-priority-hi,attr",json:"sender-priority-hi,attr",yaml:"sender-priority-hi,attr"`
	SenderAddress     NWFilterField `xml:"sender-address,attr,omitempty",json:"sender-address,attr,omitempty",yaml:"sender-address,attr,omitempty"`
	SenderAddressMask NWFilterField `xml:"sender-address-mask,attr,omitempty",json:"sender-address-mask,attr,omitempty",yaml:"sender-address-mask,attr,omitempty"`
	Port              NWFilterField `xml:"port,attr",json:"port,attr",yaml:"port,attr"`
	PortHi            NWFilterField `xml:"port-hi,attr",json:"port-hi,attr",yaml:"port-hi,attr"`
	Age               NWFilterField `xml:"age,attr",json:"age,attr",yaml:"age,attr"`
	AgeHi             NWFilterField `xml:"age-hi,attr",json:"age-hi,attr",yaml:"age-hi,attr"`
	MaxAge            NWFilterField `xml:"max-age,attr",json:"max-age,attr",yaml:"max-age,attr"`
	MaxAgeHi          NWFilterField `xml:"max-age-hi,attr",json:"max-age-hi,attr",yaml:"max-age-hi,attr"`
	HelloTime         NWFilterField `xml:"hello-time,attr",json:"hello-time,attr",yaml:"hello-time,attr"`
	HelloTimeHi       NWFilterField `xml:"hello-time-hi,attr",json:"hello-time-hi,attr",yaml:"hello-time-hi,attr"`
	ForwardDelay      NWFilterField `xml:"forward-delay,attr",json:"forward-delay,attr",yaml:"forward-delay,attr"`
	ForwardDelayHi    NWFilterField `xml:"forward-delay-hi,attr",json:"forward-delay-hi,attr",yaml:"forward-delay-hi,attr"`
	Comment           string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleIP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty",json:"srcipaddr,attr,omitempty",yaml:"srcipaddr,attr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty",json:"srcipmask,attr,omitempty",yaml:"srcipmask,attr,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty",json:"dstipaddr,attr,omitempty",yaml:"dstipaddr,attr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty",json:"dstipmask,attr,omitempty",yaml:"dstipmask,attr,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty",json:"protocol,attr,omitempty",yaml:"protocol,attr,omitempty"`
	NWFilterRuleCommonPort
	DSCP    NWFilterField `xml:"dscp,attr",json:"dscp,attr",yaml:"dscp,attr"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonMAC
	SrcIPAddr NWFilterField `xml:"srcipaddr,attr,omitempty",json:"srcipaddr,attr,omitempty",yaml:"srcipaddr,attr,omitempty"`
	SrcIPMask NWFilterField `xml:"srcipmask,attr,omitempty",json:"srcipmask,attr,omitempty",yaml:"srcipmask,attr,omitempty"`
	DstIPAddr NWFilterField `xml:"dstipaddr,attr,omitempty",json:"dstipaddr,attr,omitempty",yaml:"dstipaddr,attr,omitempty"`
	DstIPMask NWFilterField `xml:"dstipmask,attr,omitempty",json:"dstipmask,attr,omitempty",yaml:"dstipmask,attr,omitempty"`
	Protocol  NWFilterField `xml:"protocol,attr,omitempty",json:"protocol,attr,omitempty",yaml:"protocol,attr,omitempty"`
	NWFilterRuleCommonPort
	Type    NWFilterField `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	TypeEnd NWFilterField `xml:"typeend,attr",json:"typeend,attr",yaml:"typeend,attr"`
	Code    NWFilterField `xml:"code,attr",json:"code,attr",yaml:"code,attr"`
	CodeEnd NWFilterField `xml:"codeend,attr",json:"codeend,attr",yaml:"codeend,attr"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleTCP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Option  NWFilterField `xml:"option,attr",json:"option,attr",yaml:"option,attr"`
	Flags   NWFilterField `xml:"flags,attr,omitempty",json:"flags,attr,omitempty",yaml:"flags,attr,omitempty"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleUDP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleUDPLite struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleESP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleAH struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleSCTP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleICMP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Type    NWFilterField `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Code    NWFilterField `xml:"code,attr",json:"code,attr",yaml:"code,attr"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleAll struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleIGMP struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleTCPIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Option  NWFilterField `xml:"option,attr",json:"option,attr",yaml:"option,attr"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleUDPIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleUDPLiteIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleESPIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleAHIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleSCTPIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	NWFilterRuleCommonPort
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleICMPIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Type    NWFilterField `xml:"type,attr",json:"type,attr",yaml:"type,attr"`
	Code    NWFilterField `xml:"code,attr",json:"code,attr",yaml:"code,attr"`
	Comment string        `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
}

type NWFilterRuleAllIPv6 struct {
	Match string `xml:"match,attr,omitempty",json:"match,attr,omitempty",yaml:"match,attr,omitempty"`
	NWFilterRuleCommonIP
	Comment string `xml:"comment,attr,omitempty",json:"comment,attr,omitempty",yaml:"comment,attr,omitempty"`
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
