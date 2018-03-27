// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// ComplexType
type DataEncryptionInfoType struct {
	EncryptionPubKeyDigest *DataEncryptionInfoTypeEncryptionPubKeyDigest `xml:"EncryptionPubKeyDigest"`
	TransactionKey         *SymmetricKeyType                             `xml:"TransactionKey"`

	Any []*w3c.Any
}

func NewDataEncryptionInfoType() *DataEncryptionInfoType {
	return new(DataEncryptionInfoType)
}

func (me *DataEncryptionInfoType) SetEncryptionPubKeyDigest(value *DataEncryptionInfoTypeEncryptionPubKeyDigest) {
	me.EncryptionPubKeyDigest = value
}

func (me *DataEncryptionInfoType) AddEncryptionPubKeyDigest() *DataEncryptionInfoTypeEncryptionPubKeyDigest {
	me.EncryptionPubKeyDigest = new(DataEncryptionInfoTypeEncryptionPubKeyDigest)
	return me.EncryptionPubKeyDigest
}

func (me *DataEncryptionInfoType) SetTransactionKey(value *SymmetricKeyType) {
	me.TransactionKey = value
}

func (me *DataEncryptionInfoType) AddTransactionKey() *SymmetricKeyType {
	me.TransactionKey = new(SymmetricKeyType)
	return me.TransactionKey
}
