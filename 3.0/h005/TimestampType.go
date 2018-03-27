// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// SimpleType
type TimestampType w3c.DateTime

func NewTimestampType(value *w3c.DateTime) *TimestampType {
	me := (*TimestampType)(value)
	return me
}
