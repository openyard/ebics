// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package s002

import w3c "github.com/openyard/ebics/3.0/w3c"
import ds "github.com/openyard/ebics/3.0/xmldsig"

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

func (me *OrderSignatureDataType) SetSignatureValue(value *w3c.Base64Binary) {
	me.SignatureValue = value
}

func (me *OrderSignatureDataType) SetPartnerID(value *PartnerIDType) {
	me.PartnerID = value
}

func (me *OrderSignatureDataType) SetUserID(value *UserIDType) {
	me.UserID = value
}

func (me *OrderSignatureDataType) SetX509Data(value *ds.X509Data) {
	me.X509Data = value
}
