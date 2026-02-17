package dvp

const (
	TokenTypeNone    = 0
	TokenTypeERC20   = 1
	TokenTypeERC3643 = 2
)

const (
	ServiceName         = "dvp"
	SettlementOpened    = "SettlementOpened"
	SettlementAccepted  = "SettlementAccepted"
	SettlementCanceling = "SettlementCanceling"
	SettlementCanceled  = "SettlementCanceled"
	SettlementClosing   = "SettlementClosing"
	SettlementSettled   = "SettlementSettled"
)

const (
	SettlementStatusNew      = 0
	SettlementStatusOpen     = 1
	SettlementStatusAccepted = 2
	SettlementStatusClosing  = 3
	SettlementStatusSettled  = 4
	SettlementStatusCanceled = 5
)
