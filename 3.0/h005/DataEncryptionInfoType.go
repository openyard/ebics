// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

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
