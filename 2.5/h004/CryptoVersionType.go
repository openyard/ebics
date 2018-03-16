// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

// UnionType
type CryptoVersionType struct {
	EncryptionVersionType
	SignatureVersionType
	AuthenticationVersionType
}

func NewCryptoVersionType() *CryptoVersionType {
	return new(CryptoVersionType)
}
