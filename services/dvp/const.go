package dvp

const (
	TokenTypeNone = iota
	TokenTypeERC20
	TokenTypeERC3643
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
	SettlementStatusNew = iota
	SettlementStatusOpen
	SettlementStatusAccepted
	SettlementStatusClosing
	SettlementStatusSettled
	SettlementStatusCanceled
)
