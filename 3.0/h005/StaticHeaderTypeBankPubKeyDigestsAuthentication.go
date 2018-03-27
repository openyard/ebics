// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// ComplexElement
type StaticHeaderTypeBankPubKeyDigestsAuthentication struct {
	Value   *PubKeyDigestType          `xml:",chardata"`
	Version *AuthenticationVersionType `xml:"Version,attr"`
}

func NewStaticHeaderTypeBankPubKeyDigestsAuthentication() *StaticHeaderTypeBankPubKeyDigestsAuthentication {
	return new(StaticHeaderTypeBankPubKeyDigestsAuthentication)
}

func (me *StaticHeaderTypeBankPubKeyDigestsAuthentication) SetVersion(value *AuthenticationVersionType) {
	me.Version = value
}
