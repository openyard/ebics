// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// ComplexElement
type HVTOrderInfoTypeExecutionDate struct {
	Value *w3c.Date `xml:",chardata"`
}

func NewHVTOrderInfoTypeExecutionDate() *HVTOrderInfoTypeExecutionDate {
	return new(HVTOrderInfoTypeExecutionDate)
}
