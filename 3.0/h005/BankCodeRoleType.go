// Generated with goxc v0.1.12 - rev a028cbdc83780b377b6e304cf9d98aa764b4028c
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

type BankCodeRoleType w3c.Token

const (
	BankCodeRoleType_ORIGINATOR    BankCodeRoleType = "Originator"
	BankCodeRoleType_RECIPIENT     BankCodeRoleType = "Recipient"
	BankCodeRoleType_CORRESPONDENT BankCodeRoleType = "Correspondent"
	BankCodeRoleType_OTHER         BankCodeRoleType = "Other"
)
