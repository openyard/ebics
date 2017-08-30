// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

// complex type
type BTFParamsTyp struct {
	FileName      FileNameStringType    `xml:"fileName,attr"`
	Service       RestrictedServiceType `xml:"Service"`
	SignatureFlag SignatureFlagType     `xml:"SignatureFlag,omitempty"`
	DateRange     DateRangeType         `xml:"DateRange,omitempty"`
	Parameter     Parameter             `xml:"Parameter,omitempty"`
}
