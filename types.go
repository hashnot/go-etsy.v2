package goetsy

type CommonEnvelope struct {
	Count int32 `json:"count"`
}

type LedgerEntry struct {
	LedgerEntryId  int64 `json:"ledger_entry_id"`
	LedgerId       int32 `json:"ledger_id"`
	Sequence       int32 `json:"sequence"`
	CreditAmount   int32 `json:"credit_amount"`
	DebitAmount    int32 `json:"debit_amount"`
	EntryType      string `json:"entry_type"`
	ReferenceId    int64 `json:"reference_id"`
	RunningBalance int64 `json:"running_balance"`
	CreateDate     int32 `json:"create_date"`
}

type LedgerEntryEnvelope struct {
	*CommonEnvelope
	Results []LedgerEntry `json:"results"`
}
