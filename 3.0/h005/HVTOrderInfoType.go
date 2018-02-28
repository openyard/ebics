// Generated with goxc v0.1.7 - rev 2ae97d253f5eaa17bab360dad75945920dfceef4
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

// complex type
type HVTOrderInfoType struct {
	MsgName       MessageType                   `xml:"MsgName,omitempty"`
	AccountInfo   HVTAccountInfoType            `xml:"AccountInfo"`
	ExecutionDate HVTOrderInfoTypeExecutionDate `xml:"ExecutionDate,omitempty"`
	Amount        HVTOrderInfoTypeAmount        `xml:"Amount"`
	Description   HVTOrderInfoTypeDescription   `xml:"Description,omitempty"`

	Any []*w3c.Any
}
