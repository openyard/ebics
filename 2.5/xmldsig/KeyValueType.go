// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type KeyValueType struct {
	DSAKeyValue DSAKeyValue `xml:"DSAKeyValue"`
	RSAKeyValue RSAKeyValue `xml:"RSAKeyValue"`

	Any []*w3c.Any
}
