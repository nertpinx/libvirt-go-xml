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

type StorageEncryptionSecret struct {
	Type string `xml:"type,attr" json:"type,omitempty" yaml:"type,omitempty"`
	UUID string `xml:"uuid,attr" json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type StorageEncryptionCipher struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Size uint64 `xml:"size,attr" json:"size,omitempty" yaml:"size,omitempty"`
	Mode string `xml:"mode,attr" json:"mode,omitempty" yaml:"mode,omitempty"`
	Hash string `xml:"hash,attr" json:"hash,omitempty" yaml:"hash,omitempty"`
}

type StorageEncryptionIvgen struct {
	Name string `xml:"name,attr" json:"name,omitempty" yaml:"name,omitempty"`
	Hash string `xml:"hash,attr" json:"hash,omitempty" yaml:"hash,omitempty"`
}

type StorageEncryption struct {
	Format string                   `xml:"format,attr" json:"format,omitempty" yaml:"format,omitempty"`
	Secret *StorageEncryptionSecret `xml:"secret" json:"secret" yaml:"secret"`
	Cipher *StorageEncryptionCipher `xml:"cipher" json:"cipher" yaml:"cipher"`
	Ivgen  *StorageEncryptionIvgen  `xml:"ivgen" json:"ivgen" yaml:"ivgen"`
}
