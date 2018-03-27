// Generated with goxc v0.1.13 - rev f5cc87998c35abe9b532e49b5672e8667bcbd00c
package h004

import w3c "github.com/openyard/ebics/2.5/w3c"

type BankCodeRoleType w3c.Token

const (
	BankCodeRoleType_ORIGINATOR    BankCodeRoleType = "Originator"
	BankCodeRoleType_RECIPIENT     BankCodeRoleType = "Recipient"
	BankCodeRoleType_CORRESPONDENT BankCodeRoleType = "Correspondent"
	BankCodeRoleType_OTHER         BankCodeRoleType = "Other"
)
