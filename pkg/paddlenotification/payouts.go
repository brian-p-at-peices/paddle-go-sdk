// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// PayoutCreated represents the payout.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type PayoutCreated struct {
	GenericNotificationEvent
	Data PayoutNotification `json:"data"`
}

// PayoutPaid represents the payout.paid event.
// See https://developer.paddle.com/webhooks/overview for more information.
type PayoutPaid struct {
	GenericNotificationEvent
	Data PayoutNotification `json:"data"`
}

// PayoutNotification: New or changed entity.
type PayoutNotification struct {
	// ID: ID for this payout.
	ID string `json:"id,omitempty"`
	// Status: Status of this payout.
	Status Status `json:"status,omitempty"`
	// Amount: Amount paid, or scheduled to be paid, for this payout.
	Amount string `json:"amount,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code for this payout.
	CurrencyCode CurrencyCodePayouts `json:"currency_code,omitempty"`
}
