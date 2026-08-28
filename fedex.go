package shippinglabel

// FedEx is the one carrier that cannot be used the moment its credentials are stored. Shippinglabel
// is a FedEx Integrator and transacts on the customer's own account, so FedEx requires that account
// to be registered once before it will accept a shipment for it. The registration runs through
// RegisterFedEx and one of the two second factors below.

// FedEx parcel custom fields. Set them in Parcel.CustomFields; only the first parcel of a shipment
// is evaluated, and keys belonging to another carrier are ignored.
const (
	// FedExPackagingType selects the packaging, e.g. FedExPackagingYourPackaging.
	FedExPackagingType = "fedexPackagingType"
	// FedExSignatureOption overrides what the booked services imply, e.g. FedExSignatureIndirect.
	FedExSignatureOption = "fedexSignatureOption"
	// FedExLabelResponseOption switches between an encoded label and a link to it.
	FedExLabelResponseOption = "fedexLabelResponseOption"
)

// Values for FedExPackagingType. Defaults to your own packaging when unset.
const (
	FedExPackagingYourPackaging = "YOUR_PACKAGING"
	FedExPackagingEnvelope      = "FEDEX_ENVELOPE"
	FedExPackagingPak           = "FEDEX_PAK"
	FedExPackagingBox           = "FEDEX_BOX"
	FedExPackagingTube          = "FEDEX_TUBE"
	FedExPackaging10KgBox       = "FEDEX_10KG_BOX"
	FedExPackaging25KgBox       = "FEDEX_25KG_BOX"
)

// Values for FedExSignatureOption. The services SIGNATURE and NAMED_PERSON_ONLY already request a
// direct signature and VISUAL_CHECK_OF_AGE an adult one; these two are only reachable here.
const (
	FedExSignatureDirect   = "DIRECT"
	FedExSignatureIndirect = "INDIRECT"
	FedExSignatureAdult    = "ADULT"
	FedExSignatureNone     = "NO_SIGNATURE_REQUIRED"
)

// Values for FedExLabelResponseOption.
const (
	// FedExLabelResponseLabel returns the label encoded in the response. Default.
	FedExLabelResponseLabel = "LABEL"
	// FedExLabelResponseURL makes FedEx return a link instead; it is downloaded for you either way.
	FedExLabelResponseURL = "URL_ONLY"
)

// FedExPinOption is the channel FedEx delivers the secure code over.
type FedExPinOption string

const (
	FedExPinSMS   FedExPinOption = "SMS"
	FedExPinCall  FedExPinOption = "CALL"
	FedExPinEmail FedExPinOption = "EMAIL"
)

// FedExRegistration starts the registration of a customer's FedEx account.
type FedExRegistration struct {
	// CustomerName is the account holder as FedEx holds it, at most 50 characters.
	CustomerName string `json:"customerName,omitempty"`
	// Address has to match what FedEx holds for the account, otherwise the request is refused.
	Address *Address `json:"address,omitempty"`
	// EULAAccepted confirms the customer accepted the FedEx end user license agreement. Mandatory:
	// FedEx requires the agreement to be shown in full and accepted before credentials are issued.
	EULAAccepted bool `json:"eulaAccepted"`
	// AccountNumber is only needed while the carrier itself has not been saved yet.
	AccountNumber string `json:"accountNumber,omitempty"`
}

// FedExRegistrationOptions tells you which second factor FedEx offers for the account.
type FedExRegistrationOptions struct {
	// Registered is true when FedEx issued the credentials right away and no second factor is due.
	Registered bool `json:"registered"`
	// MFARequired reports that a second factor is still outstanding.
	MFARequired bool `json:"mfaRequired,omitempty"`
	// SecureCode lists the channels available for the code, see FedExPinOption.
	SecureCode []string `json:"secureCode,omitempty"`
	// Invoice reports whether validating with a recent invoice is offered.
	Invoice bool `json:"invoice,omitempty"`
	// Email and PhoneNumber are masked by FedEx and only meant to be shown to the customer.
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// FedExInvoice validates the account with a FedEx invoice instead of a secure code.
type FedExInvoice struct {
	Number int64 `json:"number,omitempty"`
	// Date in format YYYY-MM-DD. FedEx does not accept invoices older than 90 days.
	Date     string  `json:"date,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
}
