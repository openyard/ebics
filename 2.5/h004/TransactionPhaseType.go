// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type TransactionPhaseType w3c.Token

const (
	TransactionPhaseType_INITIALISATION TransactionPhaseType = "Initialisation"
	TransactionPhaseType_TRANSFER       TransactionPhaseType = "Transfer"
	TransactionPhaseType_RECEIPT        TransactionPhaseType = "Receipt"
)
