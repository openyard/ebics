// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type HVUOriginatorInfoType struct {
	PartnerID *PartnerIDType `xml:"PartnerID"`
	UserID    *UserIDType    `xml:"UserID"`
	Name      *NameType      `xml:"Name,omitempty"`
	Timestamp *TimestampType `xml:"Timestamp"`

	Any []*w3c.Any
}

func NewHVUOriginatorInfoType() *HVUOriginatorInfoType {
	return new(HVUOriginatorInfoType)
}

func (me *HVUOriginatorInfoType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *HVUOriginatorInfoType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *HVUOriginatorInfoType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HVUOriginatorInfoType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *HVUOriginatorInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *HVUOriginatorInfoType) AddName() *NameType {
	me.Name = new(NameType)
	return me.Name
}

func (me *HVUOriginatorInfoType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *HVUOriginatorInfoType) AddTimestamp() *TimestampType {
	me.Timestamp = new(TimestampType)
	return me.Timestamp
}
