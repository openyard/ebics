// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h000

import w3c "github.com/openyard/ebics/2.5/w3c"

// SimpleType
type ReportTextType w3c.NormalizedString

func NewReportTextType(value *w3c.NormalizedString) *ReportTextType {
	me := (*ReportTextType)(value)
	return me
}
