// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *SignerInfoType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *SignerInfoType) SetName(value *NameType) {
	me.Name = value
}

func (me *SignerInfoType) SetTimestamp(value *TimestampType) {
	me.Timestamp = value
}

func (me *SignerInfoType) SetPermission(value *SignerInfoTypePermission) {
	me.Permission = value
}
