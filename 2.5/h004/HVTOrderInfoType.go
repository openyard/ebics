// Generated with goxc vgoxc-0.1.9 - rev 260439f4ef82f3f152002242cdec0bb97e6118c3
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

// complex type
type HVTOrderInfoType struct {
	OrderFormat   *OrderFormatType               `xml:"OrderFormat,omitempty"`
	AccountInfo   *HVTAccountInfoType            `xml:"AccountInfo"`
	ExecutionDate *HVTOrderInfoTypeExecutionDate `xml:"ExecutionDate,omitempty"`
	Amount        *HVTOrderInfoTypeAmount        `xml:"Amount"`
	Description   *HVTOrderInfoTypeDescription   `xml:"Description,omitempty"`

	Any []*w3c.Any
}
