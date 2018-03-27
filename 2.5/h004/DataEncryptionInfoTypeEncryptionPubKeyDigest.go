// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

// ComplexElement
type DataEncryptionInfoTypeEncryptionPubKeyDigest struct {
	Value   *PubKeyDigestType      `xml:",chardata"`
	Version *EncryptionVersionType `xml:"Version,attr"`
}

func NewDataEncryptionInfoTypeEncryptionPubKeyDigest() *DataEncryptionInfoTypeEncryptionPubKeyDigest {
	return new(DataEncryptionInfoTypeEncryptionPubKeyDigest)
}

func (me *DataEncryptionInfoTypeEncryptionPubKeyDigest) SetVersion(value *EncryptionVersionType) {
	me.Version = value
}
