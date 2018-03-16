// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type TransactionPhaseType w3c.Token

const (
	TransactionPhaseType_INITIALISATION TransactionPhaseType = "Initialisation"
	TransactionPhaseType_TRANSFER       TransactionPhaseType = "Transfer"
	TransactionPhaseType_RECEIPT        TransactionPhaseType = "Receipt"
)
