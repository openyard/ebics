// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package s001

import w3c "github.com/openyard/ebics/2.5/w3c"
import ds "github.com/openyard/ebics/2.5/xmldsig"

// ComplexType
type OrderSignatureDataType struct {
	SignatureVersion *SignatureVersionType `xml:"SignatureVersion"`
	SignatureValue   *w3c.Base64Binary     `xml:"SignatureValue"`
	PartnerID        *PartnerIDType        `xml:"PartnerID"`
	UserID           *UserIDType           `xml:"UserID"`
	X509Data         *ds.X509Data          `xml:"http://www.w3.org/2000/09/xmldsig# X509Data,omitempty"`
}

func NewOrderSignatureDataType() *OrderSignatureDataType {
	return new(OrderSignatureDataType)
}

func (me *OrderSignatureDataType) SetSignatureVersion(value *SignatureVersionType) {
	me.SignatureVersion = value
}

func (me *OrderSignatureDataType) AddSignatureVersion() *SignatureVersionType {
	me.SignatureVersion = new(SignatureVersionType)
	return me.SignatureVersion
}

func (me *OrderSignatureDataType) SetSignatureValue(value *w3c.Base64Binary) {
	me.SignatureValue = value
}

func (me *OrderSignatureDataType) AddSignatureValue() *w3c.Base64Binary {
	me.SignatureValue = new(w3c.Base64Binary)
	return me.SignatureValue
}

func (me *OrderSignatureDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *OrderSignatureDataType) AddPartnerID() *PartnerIDType {
	me.PartnerID = new(PartnerIDType)
	return me.PartnerID
}

func (me *OrderSignatureDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *OrderSignatureDataType) AddUserID() *UserIDType {
	me.UserID = new(UserIDType)
	return me.UserID
}

func (me *OrderSignatureDataType) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
