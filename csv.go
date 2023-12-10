package shippinglabel

import "time"

type EncodingCode string

const (
	EncodingUTF8     EncodingCode = "UTF-8"
	EncodingISO88591 EncodingCode = "ISO-8859-1"
	EncodingUTF8BOM  EncodingCode = "UTF-8-BOM"
)

type CSVProfile struct {
	ID         int          `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	Delimiter  string       `json:"delimiter,omitempty"`
	DateFormat string       `json:"dateFormat,omitempty"`
	Encoding   EncodingCode `json:"encoding,omitempty"`
	Mapping    []*Mapping   `json:"mapping,omitempty"`
	IsDefault  bool         `json:"isDefault,omitempty"`
	Created    *time.Time   `json:"created,omitempty"`
}

type Mapping struct {
	HeaderField string `json:"headerField,omitempty"`
	ValueID     string `json:"valueId,omitempty"`
}
