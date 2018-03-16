// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type DataDigestType struct {
	DigestType
	SignatureVersion *SignatureVersionType `xml:"SignatureVersion,attr,omitempty"`
}

func NewDataDigestType() *DigestType {
	return new(DigestType)
}

func (me *DataDigestType) SetSignatureVersion(value *SignatureVersionType) {
	me.SignatureVersion = value
}
