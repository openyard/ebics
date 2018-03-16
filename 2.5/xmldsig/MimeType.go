// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// Attribute
type MimeType struct {
	Value *w3c.String `xml:"MimeType,attr"`
}

func NewMimeType() *MimeType {
	return new(MimeType)
}

func (me *MimeType) SetValue(value *w3c.String) {
	me.Value = value
}
