package goetsy

type GenericEnvelope struct {
	CommonEnvelope
	Results []interface{} `json:"results"`
}

type CommonEnvelope struct {
	Count      int32                  `json:"count"`
	Type       string                 `json:"type"`
	Pagination Pagination             `json:"pagination"`
	Params     map[string]interface{} `json:"params"`
}

type Pagination struct {
	EffectiveLimit  int32 `json:"effective_limit"`
	EffectiveOffset int32 `json:"effective_offset"`
	NextOffset      int32 `json:"next_offset"`
	EffectivePage   int32 `json:"effective_page"`
	NextPage        int32 `json:"next_page"`
}

type LedgerEntry struct {
	LedgerEntryId  int64  `json:"ledger_entry_id"`
	LedgerId       int32  `json:"ledger_id"`
	Sequence       int32  `json:"sequence"`
	CreditAmount   int32  `json:"credit_amount"`
	DebitAmount    int32  `json:"debit_amount"`
	EntryType      string `json:"entry_type"`
	ReferenceId    int64  `json:"reference_id"`
	RunningBalance int64  `json:"running_balance"`
	CreateDate     int32  `json:"create_date"`
}

type LedgerEntryEnvelope struct {
	CommonEnvelope
	Results []LedgerEntry `json:"results"`
}

const (
	LedgerEntryPayment = "PAYMENT"
	LedgerEntryRefund = "REFUND"
)
