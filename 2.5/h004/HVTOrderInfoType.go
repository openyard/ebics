// Generated with goxc v0.1.2 - rev bbe25b23ba83e8f8464e681ca3e514eee51fd2ba
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVTOrderInfoType struct {
	OrderFormat   OrderFormatType               `xml:"OrderFormat,omitempty"`
	AccountInfo   HVTAccountInfoType            `xml:"AccountInfo"`
	ExecutionDate HVTOrderInfoTypeExecutionDate `xml:"ExecutionDate,omitempty"`
	Amount        HVTOrderInfoTypeAmount        `xml:"Amount"`
	Description   HVTOrderInfoTypeDescription   `xml:"Description,omitempty"`

	Any []*w3c.Any
}
