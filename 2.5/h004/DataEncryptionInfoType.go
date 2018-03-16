// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
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

func (me *DataEncryptionInfoType) SetTransactionKey(value *SymmetricKeyType) {
	me.TransactionKey = value
}
