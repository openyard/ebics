// Generated with goxc v0.1.1 - rev bae2cf01854d664b985cae6986076979716034c7
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type BankCodeRoleType w3c.Token

const (
	BankCodeRoleType_ORIGINATOR    BankCodeRoleType = "Originator"
	BankCodeRoleType_RECIPIENT     BankCodeRoleType = "Recipient"
	BankCodeRoleType_CORRESPONDENT BankCodeRoleType = "Correspondent"
	BankCodeRoleType_OTHER         BankCodeRoleType = "Other"
)
