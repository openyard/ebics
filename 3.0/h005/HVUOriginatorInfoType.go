// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
