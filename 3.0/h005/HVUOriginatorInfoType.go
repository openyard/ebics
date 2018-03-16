// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *HVUOriginatorInfoType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *HVUOriginatorInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *HVUOriginatorInfoType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}
