// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexType
type SignerInfoType struct {
	PartnerID  *PartnerIDType            `xml:"PartnerID"`
	UserID     *UserIDType               `xml:"UserID"`
	Name       *NameType                 `xml:"Name,omitempty"`
	Timestamp  *TimestampType            `xml:"Timestamp"`
	Permission *SignerInfoTypePermission `xml:"Permission"`

	Any []*w3c.Any
}

func NewSignerInfoType() *SignerInfoType {
	return new(SignerInfoType)
}

func (me *SignerInfoType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *SignerInfoType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *SignerInfoType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *SignerInfoType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *SignerInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *SignerInfoType) AddName() *NameType {
	me.Name = new(NameType)
	return me.Name
}

func (me *SignerInfoType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *SignerInfoType) AddTimestamp() *TimestampType {
	me.Timestamp = new(TimestampType)
	return me.Timestamp
}

func (me *SignerInfoType) SetPermission(value *SignerInfoTypePermission) {
	me.Permission = value
}

func (me *SignerInfoType) AddPermission() *SignerInfoTypePermission {
	me.Permission = new(SignerInfoTypePermission)
	return me.Permission
}
