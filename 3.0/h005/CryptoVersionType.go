// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

// UnionType
type CryptoVersionType struct {
	EncryptionVersionType
	SignatureVersionType
	AuthenticationVersionType
}

func NewCryptoVersionType() *CryptoVersionType {
	return new(CryptoVersionType)
}
