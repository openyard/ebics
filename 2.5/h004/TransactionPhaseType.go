// Generated with goxc vgoxc-0.1.10 - rev e8baacfe36e4067177cedfe1884d18a3ba2f1d75
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type TransactionPhaseType w3c.Token

const (
	TransactionPhaseType_INITIALISATION TransactionPhaseType = "Initialisation"
	TransactionPhaseType_TRANSFER       TransactionPhaseType = "Transfer"
	TransactionPhaseType_RECEIPT        TransactionPhaseType = "Receipt"
)
