// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package xmldsig

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type X509SubjectName w3c.String

func NewX509SubjectName(value *w3c.String) *X509SubjectName {
	me := (*X509SubjectName)(value)
	return me
}
