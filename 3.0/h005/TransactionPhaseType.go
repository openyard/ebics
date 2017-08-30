// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

type TransactionPhaseType w3c.Token

const (
	TransactionPhaseType_INITIALISATION TransactionPhaseType = "Initialisation"
	TransactionPhaseType_TRANSFER       TransactionPhaseType = "Transfer"
	TransactionPhaseType_RECEIPT        TransactionPhaseType = "Receipt"
)
