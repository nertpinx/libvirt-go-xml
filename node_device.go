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

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NodeDevice struct {
	XMLName    xml.Name             `xml:"device" json:"device" yaml:"device"`
	Name       string               `xml:"name" json:"name" yaml:"name"`
	Path       string               `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	DevNodes   []NodeDeviceDevNode  `xml:"devnode" json:"devnode" yaml:"devnode"`
	Parent     string               `xml:"parent,omitempty" json:"parent,omitempty" yaml:"parent,omitempty"`
	Driver     *NodeDeviceDriver    `xml:"driver" json:"driver" yaml:"driver"`
	Capability NodeDeviceCapability `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDeviceDevNode struct {
	Type string `xml:"type,attr,omitempty" json:"type,omitempty,omitempty" yaml:"type,omitempty,omitempty"`
	Path string `xml:",chardata" json:",chardata" yaml:",chardata"`
}

type NodeDeviceDriver struct {
	Name string `xml:"name" json:"name" yaml:"name"`
}

type NodeDeviceCapability struct {
	System     *NodeDeviceSystemCapability
	PCI        *NodeDevicePCICapability
	USB        *NodeDeviceUSBCapability
	USBDevice  *NodeDeviceUSBDeviceCapability
	Net        *NodeDeviceNetCapability
	SCSIHost   *NodeDeviceSCSIHostCapability
	SCSITarget *NodeDeviceSCSITargetCapability
	SCSI       *NodeDeviceSCSICapability
	Storage    *NodeDeviceStorageCapability
	DRM        *NodeDeviceDRMCapability
	CCW        *NodeDeviceCCWCapability
	MDev       *NodeDeviceMDevCapability
}

type NodeDeviceIDName struct {
	ID   string `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Name string `xml:",chardata" json:",chardata" yaml:",chardata"`
}

type NodeDevicePCIExpress struct {
	Links []NodeDevicePCIExpressLink `xml:"link" json:"link" yaml:"link"`
}

type NodeDevicePCIExpressLink struct {
	Validity string  `xml:"validity,attr,omitempty" json:"validity,omitempty,omitempty" yaml:"validity,omitempty,omitempty"`
	Speed    float64 `xml:"speed,attr,omitempty" json:"speed,omitempty,omitempty" yaml:"speed,omitempty,omitempty"`
	Port     *uint   `xml:"port,attr" json:"port,omitempty" yaml:"port,omitempty"`
	Width    *uint   `xml:"width,attr" json:"width,omitempty" yaml:"width,omitempty"`
}

type NodeDeviceIOMMUGroup struct {
	Number  int                    `xml:"number,attr" json:"number,omitempty" yaml:"number,omitempty"`
	Address []NodeDevicePCIAddress `xml:"address" json:"address" yaml:"address"`
}

type NodeDeviceNUMA struct {
	Node int `xml:"node,attr" json:"node,omitempty" yaml:"node,omitempty"`
}

type NodeDevicePCICapability struct {
	Domain       *uint                        `xml:"domain" json:"domain" yaml:"domain"`
	Bus          *uint                        `xml:"bus" json:"bus" yaml:"bus"`
	Slot         *uint                        `xml:"slot" json:"slot" yaml:"slot"`
	Function     *uint                        `xml:"function" json:"function" yaml:"function"`
	Product      NodeDeviceIDName             `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Vendor       NodeDeviceIDName             `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	IOMMUGroup   *NodeDeviceIOMMUGroup        `xml:"iommuGroup" json:"iommuGroup" yaml:"iommuGroup"`
	NUMA         *NodeDeviceNUMA              `xml:"numa" json:"numa" yaml:"numa"`
	PCIExpress   *NodeDevicePCIExpress        `xml:"pci-express" json:"pci-express" yaml:"pci-express"`
	Capabilities []NodeDevicePCISubCapability `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDevicePCIAddress struct {
	Domain   *uint `xml:"domain,attr" json:"domain,omitempty" yaml:"domain,omitempty"`
	Bus      *uint `xml:"bus,attr" json:"bus,omitempty" yaml:"bus,omitempty"`
	Slot     *uint `xml:"slot,attr" json:"slot,omitempty" yaml:"slot,omitempty"`
	Function *uint `xml:"function,attr" json:"function,omitempty" yaml:"function,omitempty"`
}

type NodeDevicePCISubCapability struct {
	VirtFunctions *NodeDevicePCIVirtFunctionsCapability
	PhysFunction  *NodeDevicePCIPhysFunctionCapability
	MDevTypes     *NodeDevicePCIMDevTypesCapability
	Bridge        *NodeDevicePCIBridgeCapability
}

type NodeDevicePCIVirtFunctionsCapability struct {
	Address  []NodeDevicePCIAddress `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
	MaxCount int                    `xml:"maxCount,attr,omitempty" json:"maxCount,omitempty,omitempty" yaml:"maxCount,omitempty,omitempty"`
}

type NodeDevicePCIPhysFunctionCapability struct {
	Address NodeDevicePCIAddress `xml:"address,omitempty" json:"address,omitempty" yaml:"address,omitempty"`
}

type NodeDevicePCIMDevTypesCapability struct {
	Types []NodeDevicePCIMDevType `xml:"type" json:"type" yaml:"type"`
}

type NodeDevicePCIMDevType struct {
	ID                 string `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
	Name               string `xml:"name" json:"name" yaml:"name"`
	DeviceAPI          string `xml:"deviceAPI" json:"deviceAPI" yaml:"deviceAPI"`
	AvailableInstances uint   `xml:"availableInstances" json:"availableInstances" yaml:"availableInstances"`
}

type NodeDevicePCIBridgeCapability struct {
}

type NodeDeviceSystemHardware struct {
	Vendor  string `xml:"vendor" json:"vendor" yaml:"vendor"`
	Version string `xml:"version" json:"version" yaml:"version"`
	Serial  string `xml:"serial" json:"serial" yaml:"serial"`
	UUID    string `xml:"uuid" json:"uuid" yaml:"uuid"`
}

type NodeDeviceSystemFirmware struct {
	Vendor      string `xml:"vendor" json:"vendor" yaml:"vendor"`
	Version     string `xml:"version" json:"version" yaml:"version"`
	ReleaseData string `xml:"release_date" json:"release_date" yaml:"release_date"`
}

type NodeDeviceSystemCapability struct {
	Product  string                    `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Hardware *NodeDeviceSystemHardware `xml:"hardware" json:"hardware" yaml:"hardware"`
	Firmware *NodeDeviceSystemFirmware `xml:"firmware" json:"firmware" yaml:"firmware"`
}

type NodeDeviceUSBDeviceCapability struct {
	Bus     int              `xml:"bus" json:"bus" yaml:"bus"`
	Device  int              `xml:"device" json:"device" yaml:"device"`
	Product NodeDeviceIDName `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Vendor  NodeDeviceIDName `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
}

type NodeDeviceUSBCapability struct {
	Number      int    `xml:"number" json:"number" yaml:"number"`
	Class       int    `xml:"class" json:"class" yaml:"class"`
	Subclass    int    `xml:"subclass" json:"subclass" yaml:"subclass"`
	Protocol    int    `xml:"protocol" json:"protocol" yaml:"protocol"`
	Description string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

type NodeDeviceNetOffloadFeatures struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
}

type NodeDeviceNetLink struct {
	State string `xml:"state,attr" json:"state,omitempty" yaml:"state,omitempty"`
	Speed string `xml:"speed,attr,omitempty" json:"speed,omitempty,omitempty" yaml:"speed,omitempty,omitempty"`
}

type NodeDeviceNetSubCapability struct {
	Wireless80211 *NodeDeviceNet80211Capability
	Ethernet80203 *NodeDeviceNet80203Capability
}

type NodeDeviceNet80211Capability struct {
}

type NodeDeviceNet80203Capability struct {
}

type NodeDeviceNetCapability struct {
	Interface  string                         `xml:"interface" json:"interface" yaml:"interface"`
	Address    string                         `xml:"address" json:"address" yaml:"address"`
	Link       *NodeDeviceNetLink             `xml:"link" json:"link" yaml:"link"`
	Features   []NodeDeviceNetOffloadFeatures `xml:"feature,omitempty" json:"feature,omitempty" yaml:"feature,omitempty"`
	Capability []NodeDeviceNetSubCapability   `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDeviceSCSIVPortOpsCapability struct {
	VPorts    int `xml:"vports,omitempty" json:"vports,omitempty" yaml:"vports,omitempty"`
	MaxVPorts int `xml:"maxvports,omitempty" json:"maxvports,omitempty" yaml:"maxvports,omitempty"`
}

type NodeDeviceSCSIFCHostCapability struct {
	WWNN      string `xml:"wwnn,omitempty" json:"wwnn,omitempty" yaml:"wwnn,omitempty"`
	WWPN      string `xml:"wwpn,omitempty" json:"wwpn,omitempty" yaml:"wwpn,omitempty"`
	FabricWWN string `xml:"fabric_wwn,omitempty" json:"fabric_wwn,omitempty" yaml:"fabric_wwn,omitempty"`
}

type NodeDeviceSCSIHostSubCapability struct {
	VPortOps *NodeDeviceSCSIVPortOpsCapability
	FCHost   *NodeDeviceSCSIFCHostCapability
}

type NodeDeviceSCSIHostCapability struct {
	Host       uint                              `xml:"host" json:"host" yaml:"host"`
	UniqueID   *uint                             `xml:"unique_id" json:"unique_id" yaml:"unique_id"`
	Capability []NodeDeviceSCSIHostSubCapability `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDeviceSCSITargetCapability struct {
	Target     string                              `xml:"target" json:"target" yaml:"target"`
	Capability []NodeDeviceSCSITargetSubCapability `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDeviceSCSITargetSubCapability struct {
	FCRemotePort *NodeDeviceSCSIFCRemotePortCapability
}

type NodeDeviceSCSIFCRemotePortCapability struct {
	RPort string `xml:"rport" json:"rport" yaml:"rport"`
	WWPN  string `xml:"wwpn" json:"wwpn" yaml:"wwpn"`
}

type NodeDeviceSCSICapability struct {
	Host   int    `xml:"host" json:"host" yaml:"host"`
	Bus    int    `xml:"bus" json:"bus" yaml:"bus"`
	Target int    `xml:"target" json:"target" yaml:"target"`
	Lun    int    `xml:"lun" json:"lun" yaml:"lun"`
	Type   string `xml:"type" json:"type" yaml:"type"`
}

type NodeDeviceStorageSubCapability struct {
	Removable *NodeDeviceStorageRemovableCapability
}

type NodeDeviceStorageRemovableCapability struct {
	MediaAvailable   *uint  `xml:"media_available" json:"media_available" yaml:"media_available"`
	MediaSize        *uint  `xml:"media_size" json:"media_size" yaml:"media_size"`
	MediaLabel       string `xml:"media_label,omitempty" json:"media_label,omitempty" yaml:"media_label,omitempty"`
	LogicalBlockSize *uint  `xml:"logical_block_size" json:"logical_block_size" yaml:"logical_block_size"`
	NumBlocks        *uint  `xml:"num_blocks" json:"num_blocks" yaml:"num_blocks"`
}

type NodeDeviceStorageCapability struct {
	Block            string                           `xml:"block,omitempty" json:"block,omitempty" yaml:"block,omitempty"`
	Bus              string                           `xml:"bus,omitempty" json:"bus,omitempty" yaml:"bus,omitempty"`
	DriverType       string                           `xml:"drive_type,omitempty" json:"drive_type,omitempty" yaml:"drive_type,omitempty"`
	Model            string                           `xml:"model,omitempty" json:"model,omitempty" yaml:"model,omitempty"`
	Vendor           string                           `xml:"vendor,omitempty" json:"vendor,omitempty" yaml:"vendor,omitempty"`
	Serial           string                           `xml:"serial,omitempty" json:"serial,omitempty" yaml:"serial,omitempty"`
	Size             *uint                            `xml:"size" json:"size" yaml:"size"`
	LogicalBlockSize *uint                            `xml:"logical_block_size" json:"logical_block_size" yaml:"logical_block_size"`
	NumBlocks        *uint                            `xml:"num_blocks" json:"num_blocks" yaml:"num_blocks"`
	Capability       []NodeDeviceStorageSubCapability `xml:"capability" json:"capability" yaml:"capability"`
}

type NodeDeviceDRMCapability struct {
	Type string `xml:"type" json:"type" yaml:"type"`
}

type NodeDeviceCCWCapability struct {
	CSSID *uint `xml:"cssid" json:"cssid" yaml:"cssid"`
	SSID  *uint `xml:"ssid" json:"ssid" yaml:"ssid"`
	DevNo *uint `xml:"devno" json:"devno" yaml:"devno"`
}

type NodeDeviceMDevCapability struct {
	Type       *NodeDeviceMDevCapabilityType `xml:"type" json:"type" yaml:"type"`
	IOMMUGroup *NodeDeviceIOMMUGroup         `xml:"iommuGroup" json:"iommuGroup" yaml:"iommuGroup"`
}

type NodeDeviceMDevCapabilityType struct {
	ID string `xml:"id,attr" json:"id,omitempty" yaml:"id,omitempty"`
}

func (a *NodeDevicePCIAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	marshalUintAttr(&start, "domain", a.Domain, "0x%04x")
	marshalUintAttr(&start, "bus", a.Bus, "0x%02x")
	marshalUintAttr(&start, "slot", a.Slot, "0x%02x")
	marshalUintAttr(&start, "function", a.Function, "0x%x")
	e.EncodeToken(start)
	e.EncodeToken(start.End())
	return nil
}

func (a *NodeDevicePCIAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (c *NodeDeviceCCWCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if c.CSSID != nil {
		cssid := xml.StartElement{
			Name: xml.Name{Local: "cssid"},
		}
		e.EncodeToken(cssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.CSSID)))
		e.EncodeToken(cssid.End())
	}
	if c.SSID != nil {
		ssid := xml.StartElement{
			Name: xml.Name{Local: "ssid"},
		}
		e.EncodeToken(ssid)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%x", *c.SSID)))
		e.EncodeToken(ssid.End())
	}
	if c.DevNo != nil {
		devno := xml.StartElement{
			Name: xml.Name{Local: "devno"},
		}
		e.EncodeToken(devno)
		e.EncodeToken(xml.CharData(fmt.Sprintf("0x%04x", *c.DevNo)))
		e.EncodeToken(devno.End())
	}
	e.EncodeToken(start.End())
	return nil
}

func (c *NodeDeviceCCWCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			cdata, err := d.Token()
			if err != nil {
				return err
			}

			if tok.Name.Local != "cssid" &&
				tok.Name.Local != "ssid" &&
				tok.Name.Local != "devno" {
				continue
			}

			chardata, ok := cdata.(xml.CharData)
			if !ok {
				return fmt.Errorf("Expected text for CCW '%s'", tok.Name.Local)
			}

			valstr := strings.TrimPrefix(string(chardata), "0x")
			val, err := strconv.ParseUint(valstr, 16, 64)
			if err != nil {
				return err
			}

			vali := uint(val)
			if tok.Name.Local == "cssid" {
				c.CSSID = &vali
			} else if tok.Name.Local == "ssid" {
				c.SSID = &vali
			} else if tok.Name.Local == "devno" {
				c.DevNo = &vali
			}
		}
	}
	return nil
}

func (c *NodeDevicePCISubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "virt_functions":
		var virtFuncCaps NodeDevicePCIVirtFunctionsCapability
		if err := d.DecodeElement(&virtFuncCaps, &start); err != nil {
			return err
		}
		c.VirtFunctions = &virtFuncCaps
	case "phys_function":
		var physFuncCaps NodeDevicePCIPhysFunctionCapability
		if err := d.DecodeElement(&physFuncCaps, &start); err != nil {
			return err
		}
		c.PhysFunction = &physFuncCaps
	case "mdev_types":
		var mdevTypeCaps NodeDevicePCIMDevTypesCapability
		if err := d.DecodeElement(&mdevTypeCaps, &start); err != nil {
			return err
		}
		c.MDevTypes = &mdevTypeCaps
	case "pci-bridge":
		var bridgeCaps NodeDevicePCIBridgeCapability
		if err := d.DecodeElement(&bridgeCaps, &start); err != nil {
			return err
		}
		c.Bridge = &bridgeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDevicePCISubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.VirtFunctions != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "virt_functions",
		})
		return e.EncodeElement(c.VirtFunctions, start)
	} else if c.PhysFunction != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "phys_function",
		})
		return e.EncodeElement(c.PhysFunction, start)
	} else if c.MDevTypes != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev_types",
		})
		return e.EncodeElement(c.MDevTypes, start)
	} else if c.Bridge != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci-bridge",
		})
		return e.EncodeElement(c.Bridge, start)
	}
	return nil
}

func (c *NodeDeviceSCSITargetSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "fc_remote_port":
		var fcCaps NodeDeviceSCSIFCRemotePortCapability
		if err := d.DecodeElement(&fcCaps, &start); err != nil {
			return err
		}
		c.FCRemotePort = &fcCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceSCSITargetSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.FCRemotePort != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "fc_remote_port",
		})
		return e.EncodeElement(c.FCRemotePort, start)
	}
	return nil
}

func (c *NodeDeviceSCSIHostSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "fc_host":
		var fcCaps NodeDeviceSCSIFCHostCapability
		if err := d.DecodeElement(&fcCaps, &start); err != nil {
			return err
		}
		c.FCHost = &fcCaps
	case "vport_ops":
		var vportCaps NodeDeviceSCSIVPortOpsCapability
		if err := d.DecodeElement(&vportCaps, &start); err != nil {
			return err
		}
		c.VPortOps = &vportCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceSCSIHostSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.FCHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "fc_host",
		})
		return e.EncodeElement(c.FCHost, start)
	} else if c.VPortOps != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "vport_ops",
		})
		return e.EncodeElement(c.VPortOps, start)
	}
	return nil
}

func (c *NodeDeviceStorageSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "removable":
		var removeCaps NodeDeviceStorageRemovableCapability
		if err := d.DecodeElement(&removeCaps, &start); err != nil {
			return err
		}
		c.Removable = &removeCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceStorageSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Removable != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "removable",
		})
		return e.EncodeElement(c.Removable, start)
	}
	return nil
}

func (c *NodeDeviceNetSubCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "80211":
		var wlanCaps NodeDeviceNet80211Capability
		if err := d.DecodeElement(&wlanCaps, &start); err != nil {
			return err
		}
		c.Wireless80211 = &wlanCaps
	case "80203":
		var ethCaps NodeDeviceNet80203Capability
		if err := d.DecodeElement(&ethCaps, &start); err != nil {
			return err
		}
		c.Ethernet80203 = &ethCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceNetSubCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.Wireless80211 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "80211",
		})
		return e.EncodeElement(c.Wireless80211, start)
	} else if c.Ethernet80203 != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "80203",
		})
		return e.EncodeElement(c.Ethernet80203, start)
	}
	return nil
}

func (c *NodeDeviceCapability) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	typ, ok := getAttr(start.Attr, "type")
	if !ok {
		return fmt.Errorf("Missing node device capability type")
	}

	switch typ {
	case "pci":
		var pciCaps NodeDevicePCICapability
		if err := d.DecodeElement(&pciCaps, &start); err != nil {
			return err
		}
		c.PCI = &pciCaps
	case "system":
		var systemCaps NodeDeviceSystemCapability
		if err := d.DecodeElement(&systemCaps, &start); err != nil {
			return err
		}
		c.System = &systemCaps
	case "usb_device":
		var usbdevCaps NodeDeviceUSBDeviceCapability
		if err := d.DecodeElement(&usbdevCaps, &start); err != nil {
			return err
		}
		c.USBDevice = &usbdevCaps
	case "usb":
		var usbCaps NodeDeviceUSBCapability
		if err := d.DecodeElement(&usbCaps, &start); err != nil {
			return err
		}
		c.USB = &usbCaps
	case "net":
		var netCaps NodeDeviceNetCapability
		if err := d.DecodeElement(&netCaps, &start); err != nil {
			return err
		}
		c.Net = &netCaps
	case "scsi_host":
		var scsiHostCaps NodeDeviceSCSIHostCapability
		if err := d.DecodeElement(&scsiHostCaps, &start); err != nil {
			return err
		}
		c.SCSIHost = &scsiHostCaps
	case "scsi_target":
		var scsiTargetCaps NodeDeviceSCSITargetCapability
		if err := d.DecodeElement(&scsiTargetCaps, &start); err != nil {
			return err
		}
		c.SCSITarget = &scsiTargetCaps
	case "scsi":
		var scsiCaps NodeDeviceSCSICapability
		if err := d.DecodeElement(&scsiCaps, &start); err != nil {
			return err
		}
		c.SCSI = &scsiCaps
	case "storage":
		var storageCaps NodeDeviceStorageCapability
		if err := d.DecodeElement(&storageCaps, &start); err != nil {
			return err
		}
		c.Storage = &storageCaps
	case "drm":
		var drmCaps NodeDeviceDRMCapability
		if err := d.DecodeElement(&drmCaps, &start); err != nil {
			return err
		}
		c.DRM = &drmCaps
	case "ccw":
		var ccwCaps NodeDeviceCCWCapability
		if err := d.DecodeElement(&ccwCaps, &start); err != nil {
			return err
		}
		c.CCW = &ccwCaps
	case "mdev":
		var mdevCaps NodeDeviceMDevCapability
		if err := d.DecodeElement(&mdevCaps, &start); err != nil {
			return err
		}
		c.MDev = &mdevCaps
	}
	d.Skip()
	return nil
}

func (c *NodeDeviceCapability) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c.PCI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "pci",
		})
		return e.EncodeElement(c.PCI, start)
	} else if c.System != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "system",
		})
		return e.EncodeElement(c.System, start)
	} else if c.USB != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb",
		})
		return e.EncodeElement(c.USB, start)
	} else if c.USBDevice != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "usb_device",
		})
		return e.EncodeElement(c.USBDevice, start)
	} else if c.Net != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "net",
		})
		return e.EncodeElement(c.Net, start)
	} else if c.SCSI != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi",
		})
		return e.EncodeElement(c.SCSI, start)
	} else if c.SCSIHost != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_host",
		})
		return e.EncodeElement(c.SCSIHost, start)
	} else if c.SCSITarget != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "scsi_target",
		})
		return e.EncodeElement(c.SCSITarget, start)
	} else if c.Storage != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "storage",
		})
		return e.EncodeElement(c.Storage, start)
	} else if c.DRM != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "drm",
		})
		return e.EncodeElement(c.DRM, start)
	} else if c.CCW != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "ccw",
		})
		return e.EncodeElement(c.CCW, start)
	} else if c.MDev != nil {
		start.Attr = append(start.Attr, xml.Attr{
			xml.Name{Local: "type"}, "mdev",
		})
		return e.EncodeElement(c.MDev, start)
	}
	return nil
}

func (c *NodeDevice) Unmarshal(doc string) error {
	return xml.Unmarshal([]byte(doc), c)
}

func (c *NodeDevice) Marshal() (string, error) {
	doc, err := xml.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(doc), nil
}
