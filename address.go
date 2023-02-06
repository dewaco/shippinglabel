package shippinglabel

type AddressTypeCode string

const (
	AddressTypeShipping AddressTypeCode = "SHIPPING"
	AddressTypeReturn   AddressTypeCode = "RETURN"
)

type Address struct {
	ID           int             `json:"id,omitempty"`
	Company      string          `json:"company,omitempty"`   // Name 1
	FirstName    string          `json:"firstName,omitempty"` // Name 2
	LastName     string          `json:"lastName,omitempty"`  // Name 3
	Street       string          `json:"street,omitempty"`
	StreetNumber string          `json:"streetNumber,omitempty"`
	PostalCode   string          `json:"postalCode,omitempty"`
	City         string          `json:"city,omitempty"`
	Country      string          `json:"country,omitempty"`
	State        string          `json:"state,omitempty"`
	Mail         string          `json:"mail,omitempty"`
	Phone        string          `json:"phone,omitempty"`
	VATNumber    string          `json:"vatNumber,omitempty"`
	AddressType  AddressTypeCode `json:"addressType,omitempty"`
}
