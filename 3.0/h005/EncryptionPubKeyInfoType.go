// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexType
type EncryptionPubKeyInfoType struct {
	PubKeyInfoType
	EncryptionVersion *EncryptionVersionType `xml:"EncryptionVersion"`
}

func NewEncryptionPubKeyInfoType() *PubKeyInfoType {
	return new(PubKeyInfoType)
}

func (me *EncryptionPubKeyInfoType) SetEncryptionVersion(value *EncryptionVersionType) {
	me.EncryptionVersion = value
}

func (me *EncryptionPubKeyInfoType) AddEncryptionVersion() *EncryptionVersionType {
	me.EncryptionVersion = new(EncryptionVersionType)
	return me.EncryptionVersion
}
