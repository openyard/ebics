// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

// complex type
type FDLOrderParamsType struct {
	DateRange  FDLOrderParamsTypeDateRange `xml:"DateRange,omitempty"`
	FileFormat FileFormatType              `xml:"FileFormat"`
	Parameter  Parameter                   `xml:"Parameter,omitempty"`
}
