// Generated with goxc v0.1.3 - rev 0e63342ac0a4d5f52582ea0065a462e700069839
package h005

import w3c "github.com/openyard/ebics/3.0/w3c"

type BankCodeRoleType w3c.Token

const (
	BankCodeRoleType_ORIGINATOR    BankCodeRoleType = "Originator"
	BankCodeRoleType_RECIPIENT     BankCodeRoleType = "Recipient"
	BankCodeRoleType_CORRESPONDENT BankCodeRoleType = "Correspondent"
	BankCodeRoleType_OTHER         BankCodeRoleType = "Other"
)