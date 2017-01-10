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
	LedgerEntryRefund  = "REFUND"
)

type Payment struct {
	PaymentId   int64 `json:"payment_id"`
	ShopId      int32 `json:"shop_id"`
	BuyerUserId int32 `json:"buyer_user_id"`

	ReceiptId int64 `json:"receipt_id"`

	AmountGross   int32 `json:"amount_gross"`
	AmountFees    int32 `json:"amount_fees"`
	AmountNet     int32 `json:"amount_net"`
	PostedGross   int32 `json:"posted_gross"`
	PostedFees    int32 `json:"posted_fees"`
	PostedNet     int32 `json:"posted_net"`
	AdjustedGross int32 `json:"adjusted_gross"`
	AdjustedFees  int32 `json:"adjusted_fees"`
	AdjustedNet   int32 `json:"adjusted_net"`

	Currency      string `json:"currency"`
	ShopCurrency  string `json:"shop_currency"`
	BuyerCurrency string `json:"buyer_currency"`

	ShippingUserId    int32 `json:"shipping_user_id"`
	ShippingAddressId int32 `json:"shipping_address_id"`
	BillingAddressId  int32 `json:"billing_address_id"`

	Status string `json:"status"`

	ShippedDate int32 `json:"shipped_date"`
	CreateDate  int32 `json:"create_date"`
	UpdateDate  int32 `json:"update_date"`
}

type Receipt struct {
	AdjustedGrandTotal int32         `json:"adjusted_grandtotal"`
	BuyerEmail         string        `json:"buyer_email"`
	BuyerUserId        int64         `json:"buyer_user_id"`
	City               string        `json:"city"`
	//ChannelBadgeSuffixString
	CountryId          int64  `json:"country_id"`
	CreationTime       int32  `json:"creation_tsz"`
	CurrencyCode       string `json:"currency_code"`
	DiscountAmount     int32  `json:"discount_amt"`
	FirstLine          string `json:"first_line"`
	GrandTotal         int32  `json:"grandtotal"`
	// HasLocalDelivery
	LastModifiedTime   int32  `json:"last_modified_tsz"`
	MessageFromBuyer   string `json:"message_from_buyer"`
	MessageFromPayment string `json:"message_from_payment"`
	MessageFromSeller  string `json:"message_from_seller"`
	Name               string `json:"name"`
	//OrderId
	PaymentEmail       string `json:"payment_email"`
	PaymentMethod      string `json:"payment_method"`
	ReceiptId          int64  `json:"receipt_id"`
	//ReceiptType
	SecondLine         string `json:"second_line"`
	SellerEmail        string `json:"seller_email"`
	SellerUserId       int64  `json:"seller_user_id"`
	//ShowChannelBadge
	//ShippingCarrier
	//ShippingDetails
	//ShippingNote
	//ShippingNotificationDate
	//ShippingTrackingCode
	//Shipments
	State              string `json:"state"`
	Subtotal           int32  `json:"subtotal"`
	TotalPrice         int32  `json:"total_price"`
	TotalShippingCost  int32  `json:"total_shipping_cost"`
	TotalTaxCost       int32  `json:"total_tax_cost"`
	TotalVatCost       int32  `json:"total_vat_cost"`
	//TransparentPriceMessage
	WasPaid            bool   `json:"was_paid"`
	WasShipped         bool   `json:"was_shipped"`
	Zip                string `json:"zip"`

	/**Shipment information associated to this receipt. */

	//Coupon 	private 	transactions_r 	Coupon 	The coupon for the receipt.
}

type PaymentEnvelope struct {
	CommonEnvelope
	Results []Receipt `json:"results"`
}
