// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package s001

// ComplexType
type SignaturePubKeyInfoType struct {
	PubKeyInfoType
	SignatureVersion *SignatureVersionType `xml:"SignatureVersion"`
}

func NewSignaturePubKeyInfoType() *PubKeyInfoType {
	return new(PubKeyInfoType)
}

func (me *SignaturePubKeyInfoType) SetSignatureVersion(value *SignatureVersionType) {
	me.SignatureVersion = value
}

func (me *SignaturePubKeyInfoType) AddSignatureVersion() *SignatureVersionType {
	me.SignatureVersion = new(SignatureVersionType)
	return me.SignatureVersion
}
