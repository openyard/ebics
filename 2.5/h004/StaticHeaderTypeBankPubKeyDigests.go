// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

// ComplexType
type StaticHeaderTypeBankPubKeyDigests struct {
	Authentication *StaticHeaderTypeBankPubKeyDigestsAuthentication `xml:"Authentication"`
	Encryption     *StaticHeaderTypeBankPubKeyDigestsEncryption     `xml:"Encryption"`
	Signature      *StaticHeaderTypeBankPubKeyDigestsSignature      `xml:"Signature,omitempty"`
}

func NewStaticHeaderTypeBankPubKeyDigests() *StaticHeaderTypeBankPubKeyDigests {
	return new(StaticHeaderTypeBankPubKeyDigests)
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetAuthentication(value *StaticHeaderTypeBankPubKeyDigestsAuthentication) {
	me.Authentication = value
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetEncryption(value *StaticHeaderTypeBankPubKeyDigestsEncryption) {
	me.Encryption = value
}

func (me *StaticHeaderTypeBankPubKeyDigests) SetSignature(value *StaticHeaderTypeBankPubKeyDigestsSignature) {
	me.Signature = value
}
